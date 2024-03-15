package platform

import (
	"github.com/gofiber/fiber/v2"
	v1Handler "github.com/mehdi-shokohi/inflow-fn-std/platform/v1Handler"
)

func RegisterFunctions(api fiber.Router) {
	apiGroup := api.Group("/fn")
	apiGroup.Get("/protocol", inflowDiscoveryProtocol)
	v1Handler.RegisterInflowFnV1(apiGroup)
}

func inflowDiscoveryProtocol(c *fiber.Ctx) error {
	return c.JSON(struct{ ProtoV string }{ProtoV: "inflowFn.V1"})
}
