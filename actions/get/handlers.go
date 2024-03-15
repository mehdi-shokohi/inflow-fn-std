package standardActions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/std"
)

type StandardFunc struct{}

func (a *StandardFunc) Run(c *fiber.Ctx) error {
	body, err := std.GetBodyAs[IsoMappingArgs](c)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}


	return std.Send(c, fiber.StatusOK, body)
}

func (a *StandardFunc) Arguments(c *fiber.Ctx) error {
	// Return Form Application Model
	return std.Send(c, fiber.StatusOK, IsoMappingArgs{})
}
func (a *StandardFunc) Settings(c *fiber.Ctx) error {
	// Global Setting for Fn
	if c.Method() == fiber.MethodGet {

		return std.Send(c, fiber.StatusOK, "show settings fields and model")
	}
	if c.Method() == fiber.MethodPut {
		return std.Send(c, fiber.StatusOK, "new value")

	}
	return nil
}
