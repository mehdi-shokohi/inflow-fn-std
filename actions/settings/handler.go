package settingsHandler

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn-std/platform/std"
)

type PersistSettings struct{}
func (a *PersistSettings) Run(c *fiber.Ctx) error {
	body, err := std.GetBodyAs[SettingsArgs](c)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}


	return std.Send(c, fiber.StatusOK, body)
}

func (a *PersistSettings) Arguments(c *fiber.Ctx) error {
	// Return Form Application Model
	return std.Send(c, fiber.StatusOK, SettingsArgs{})
}

func (a *PersistSettings) Settings(c *fiber.Ctx) error {
	// Global Setting for Fn
	body, err := std.GetRAWBodyAs[SettingsArgs](c)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	os.Setenv("MONGO_DB_URI",body.DataBaseUri)
	os.Setenv("KEYDB_URI",body.KeyDbUri)

	return nil
}
