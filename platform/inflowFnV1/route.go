package inflowV1

import "github.com/gofiber/fiber/v2"


func RegisterInflowFnV1(api fiber.Router) {
	api.Post("/:action/run",actionHandlers) // run method of IAction
	api.Get("/:action/load",getActionApplication) // get dialog form of action command
	api.All("/:action/set",settingsHandler) // settings method of IAction
	api.Get("/:action/help/:name",describeHandler) // doc based on Descibe arg on registerAction
}