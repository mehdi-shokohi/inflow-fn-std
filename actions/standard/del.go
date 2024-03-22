package standardActions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/inflowFnV1"
)


func RunDeleteCommand(c *fiber.Ctx) error {

	body, err := inflowV1.GetBodyAs[map[string]interface{}](c)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}

	for _,v:=range body.InlineParams["delete_key"].([]interface{}){

		delete(body.Body,v.(string))
	}
	return inflowV1.Send(c, fiber.StatusOK, body.Body)
}

