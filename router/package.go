package router

import (
	"strings"

	"github.com/gofiber/fiber"
)

// Package struct
type Package struct {
	Find func(*fiber.Ctx)
}

// RouteList all
type RouteList struct {
	Package Package
}

// Find func
func find(fiber *fiber.Ctx) {
	info := []string{"Searching for package info ->", fiber.Params("packageName"), "\n"}
	fiber.Send(strings.Join(info, " "))
}

func listRouter() RouteList {
	var router RouteList
	router.Package.Find = find

	return router
}

// Router list
var Router RouteList = listRouter()
