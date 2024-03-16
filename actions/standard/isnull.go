package standardActions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/inflowFnV1"

)


func  RunIsNullCommand(c *fiber.Ctx) error {
	body, err := inflowV1.GetBodyAs[map[string]interface{}](c)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	res:=map[string]interface{}{"key":body.InlineParams["key"].(string),"isNull":false}
	if v,ok:=body.Body[body.InlineParams["key"].(string)];ok && v==nil{
		res["isNull"]=true
	}
	return inflowV1.Send(c, fiber.StatusOK, res)
}
