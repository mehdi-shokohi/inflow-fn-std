package standardActions

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/inflowFnV1"
)

type Holder struct {
	Minutes uint64
	Fn      func()
}

var cronHolder map[string]Holder
var s *gocron.Scheduler

func RunPolicyCron(c *fiber.Ctx) error {
	body, err := inflowV1.GetBodyAs[map[string]interface{}](c)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	params := body.InlineParams
	if params["period"].(float64)<=0{
		removeCron(params["docId"].(string))
		return inflowV1.Send(c, fiber.StatusOK, params)

	}
	holder := Holder{Minutes: uint64(params["period"].(float64)), Fn: callService(body.InlineParams)}
	addToCron(body.InlineParams["docId"].(string), holder)
	return inflowV1.Send(c, fiber.StatusOK, params)
}
func removeCron(cId string){
	if cronHolder == nil {
		cronHolder = make(map[string]Holder)
	}

	if s == nil {
		s = gocron.NewScheduler(time.UTC)
	}
	delete(cronHolder,cId)
	s.Clear()
	for _, v := range cronHolder {
		s.Every(int(v.Minutes)).Minute().Do(v.Fn)
		s.StartAsync()
	}
}
func addToCron(cId string, holder Holder) {
	if cronHolder == nil {
		cronHolder = make(map[string]Holder)
	}
	cronHolder[cId] = holder
	if s == nil {
		s = gocron.NewScheduler(time.UTC)
	}
	s.Clear()
	for _, v := range cronHolder {
		s.Every(int(v.Minutes)).Minute().Do(v.Fn)
		s.StartAsync()
	}
}
func callService(params map[string]interface{}) func() {
	url := params["url"].(map[string]interface{})
	return func() {
		data := make(map[string]string)
		data["docId"] = params["docId"].(string)
		data["policyIdentifier"] = params["policyId"].(string)
		headers := map[string][]string{url["access_key"].(string): {url["token"].(string)},"Content-Type":{"application/json"}}
		urlAddress := url["url"].(string)
			body,err:=json.Marshal(data)
			if err!=nil{
				fmt.Println(err)
				return
			}
			fmt.Println(string(body))
		result,err:=Request("POST",urlAddress,body,headers)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(string(result))
		//localhost:2404/api/entry/fractal/gate/6628dc5693aeaa3cbd90b8e2/engine/process

	}
}
