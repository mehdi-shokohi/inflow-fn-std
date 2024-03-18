package v1Handler

import (
	"reflect"

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
	s := reflect.TypeOf(fn.Argument)
	for i := 0; i < s.NumField(); i++ {
		if v := s.Field(i).Tag.Get("intent"); v == c.Query("for") {
				o:=reflect.New(s.Field(i).Type).Interface()
				return inflowV1.Send(c, fiber.StatusOK, o)
		}
	}
	

	return inflowV1.Send(c, fiber.StatusOK, fn.Argument)

}

func helpObjHandler(c *fiber.Ctx) error {
	return nil
}

func helpMdHandler(c *fiber.Ctx) error {
	return nil
}
