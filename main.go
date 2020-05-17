package main

import (
	"github.com/bounteous/artixweb-packages/router"
	"github.com/bounteous/artixweb-packages/tools"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New(&fiber.Settings{
		// Prefork feature is disabled due to a bug related
		// to the args flags, since they are not being detected
		Prefork:      false,
		ServerHeader: tools.Config.API.Header,
	})

	app.Get("/packages/find/:packageName", router.Find)

	app.Listen(tools.Config.API.Port)
}
