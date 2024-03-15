package v1Handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/inflowFnV1"
)

func actionHandlers(c *fiber.Ctx) error {
	action := c.Params("action")
	fn := inflowV1.GetAction(action)
	if fn == nil {
		return inflowV1.Send(c, fiber.StatusNotFound, nil)
	}
	return fn.Run(c)

}
func actionHandlers2(c *fiber.Ctx) error {
	action := c.Params("action")
	fn := inflowV1.GetAction(action)
	if fn == nil {
		return inflowV1.Send(c, fiber.StatusNotFound, nil)
	}
	return fn.Run(c)

}
func describeHandler(c *fiber.Ctx) error {

		//return All Commands doc
		return inflowV1.Send(c, fiber.StatusOK, inflowV1.GetActions())
}

func getActionApplication(c *fiber.Ctx) error {
	action := c.Params("action")
	fn := inflowV1.GetAction(action)
	if fn == nil {
		return inflowV1.Send(c, fiber.StatusNotFound, nil)
	}
	return inflowV1.Send(c, fiber.StatusNotFound, fn.Argument)

}


func helpObjHandler(c *fiber.Ctx)error{
	return nil
}

func helpMdHandler(c *fiber.Ctx)error{
	return nil
}
