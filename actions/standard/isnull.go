package standardActions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/inflowFnV1"
)

func RunIsNullCommand(c *fiber.Ctx) error {
	body, err := inflowV1.GetBodyAs[map[string]interface{}](c)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	body.Body["isNull"] = false
	if body.Body[body.InlineParams["key"].(string)]==nil {
		body.Body["isNull"] = true
	}
	return inflowV1.Send(c, fiber.StatusOK, body.Body)
}
