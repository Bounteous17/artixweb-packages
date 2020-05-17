package main

import (
	"github.com/bounteous/artixweb-packages/router"
	"github.com/bounteous/artixweb-packages/tools"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/packages/find/:packageName", router.Find)

	app.Listen(tools.Config.API.Port)
}
