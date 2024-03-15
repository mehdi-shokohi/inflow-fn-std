package inflowV1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/std"
)

func actionHandlers(c *fiber.Ctx) error {
	action := c.Params("action")
	fn := std.GetAction(action)
	if fn == nil {
		return std.Send(c, fiber.StatusNotFound, nil)
	}
	return fn.Run(c)

}

func describeHandler(c *fiber.Ctx) error {
	nameParam := c.Params("name")
	if nameParam == "all" {
		//return All Commands doc
		return std.Send(c, fiber.StatusOK, std.GetActions())
	}
	return nil
}

func getActionApplication(c *fiber.Ctx) error {
	action := c.Params("action")
	fn := std.GetAction(action)
	if fn == nil {
		return std.Send(c, fiber.StatusNotFound, nil)
	}
	return fn.Arguments(c)
}

func settingsHandler(c *fiber.Ctx) error {
	action := c.Params("action")
	fn := std.GetAction(action)
	if fn == nil {
		return std.Send(c, fiber.StatusNotFound, nil)
	}
	return fn.Settings(c)
}
