package standardActions

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/inflowFnV1"
)

func RunCastCommand(c *fiber.Ctx) error {
	body, err := inflowV1.GetBodyAs[map[string]interface{}](c)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	for _, v := range body.InlineParams["fields"].([]interface{}) {
		if body.InlineParams["cast_to"].(string) == "int" {

			digit, err := strconv.ParseFloat(body.Body[v.(string)].(string), 64)
			if err != nil {
				continue
			}
			body.Body[v.(string)] = digit
		} else if body.InlineParams["cast_to"].(string) == "string" {

			body.Body[v.(string)] = fmt.Sprintf("%f", body.Body[v.(string)])
		}
	}

	return inflowV1.Send(c, fiber.StatusOK, body)
}
