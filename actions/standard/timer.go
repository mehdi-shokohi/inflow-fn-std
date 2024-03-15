package standardActions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/std"
)


func RunTimerCommand(c *fiber.Ctx) error {
	body, err := std.GetBodyAs[map[string]interface{}](c)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	// db:=std.GetDefaultDb(c.Context(),"std",body)
	// db.Insert()

	return std.Send(c, fiber.StatusOK, body)
}
