package main

import (
	"github.com/bounteous/artixweb-packages/router"
	"github.com/bounteous/artixweb-packages/tools"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

func main() {
	var settings fiber.Settings

	// Prefork feature is disabled due to a bug related
	// to the args flags, since they are not being detected
	settings.Prefork = false
	settings.ServerHeader = tools.Config.API.Header

	app := fiber.New(&settings)

	app.Use(cors.New())

	apiPackages := app.Group("/package")
	apiPackages.Get("/find/:packageName", router.List.Package.Find)

	app.Listen(tools.Config.API.Port)
}
