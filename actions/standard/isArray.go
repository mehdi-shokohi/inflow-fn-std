package standardActions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/inflowFnV1"

)


func RunIsArrayCommand(c *fiber.Ctx) error {
	body, err := inflowV1.GetBodyAs[map[string]interface{}](c)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	res:=map[string]interface{}{"key":body.InlineParams["key"].(string),"isArray":false}
	if _,ok:=body.Body[body.InlineParams["key"].(string)].([]interface{});ok{
		res["isArray"]=true

	}

	return inflowV1.Send(c, fiber.StatusOK, res)
}
