package main

import (
	"log"

	json "github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	conf "github.com/mehdi-shokohi/inflow-fn-std/config"
	_ "github.com/mehdi-shokohi/inflow-fn-std/actions"
	"github.com/mehdi-shokohi/inflow-fn-std/platform"
)

func main() {
	// http server config
	app := fiber.New(fiber.Config{
		Prefork:     false,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New())

	app.Use(logger.New())
	platform.RegisterFunctions(app)
	log.Fatal(app.Listen(conf.GetPort()))
}
