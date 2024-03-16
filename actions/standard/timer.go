package standardActions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/inflowFnV1"
)


func RunTimerCommand(c *fiber.Ctx) error {
	body, err := inflowV1.GetBodyAs[map[string]interface{}](c)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	headersInfo, err := inflowV1.GetHeadersAs[map[string]interface{}](c)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	body.Body["headerInfo"] = headersInfo
	return inflowV1.Send(c, fiber.StatusOK, body)
}
