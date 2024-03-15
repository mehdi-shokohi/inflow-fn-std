package platform

import (
	"github.com/gofiber/fiber/v2"
	inflowV1 "github.com/mehdi-shokohi/inflow-fn-std/platform/inflowFnV1"
)

func RegisterFunctions(api fiber.Router) {
	apiGroup := api.Group("/fn")
	apiGroup.Get("/protocol", inflowDiscoveryProtocol)
	inflowV1.RegisterInflowFnV1(apiGroup)
}

func inflowDiscoveryProtocol(c *fiber.Ctx) error {
	return c.JSON(struct{ ProtoV string }{ProtoV: "inflowFn.V1"})
}
