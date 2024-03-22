package v1Handler

import (
	"github.com/gofiber/fiber/v2"
	inflowV1 "github.com/mehdi-shokohi/inflow-fn-std/platform/inflowFnV1"
)

func RegisterInflowFnV1(api fiber.Router) {
	api.Use(safeRecovery)
	api.Post("/:action/run", actionHandlers)       // run method of IAction
	api.Get("/:action/load", getActionApplication) // get dialog form of action command
	api.Get("", describeHandler)                   // doc based on Descibe arg on registerAction
	api.Get("/help/md/:name", helpMdHandler)
	api.Get("/help/o/:name", helpObjHandler)

}

func safeRecovery(c *fiber.Ctx) error {
	defer func() error {
		if v := recover(); v != nil {

			body := make(map[string]interface{})
			if err:=c.BodyParser(&body);err!=nil{
				return inflowV1.Send(c, fiber.StatusOK, string(c.Body()))
			}
			return inflowV1.Send(c, fiber.StatusOK, body["_data"])
		}
		return nil
	}()
	return c.Next()
}

