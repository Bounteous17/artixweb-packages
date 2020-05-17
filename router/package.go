package router

import (
	"github.com/bounteous/artixweb-packages/gitea"

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
	var params gitea.SearchParams
	params.PackageName = fiber.Params("packageName")

	fiber.JSON(gitea.Search(params).Data)
}

func listRouter() RouteList {
	var router RouteList
	router.Package.Find = find

	return router
}

// List list
var List RouteList = listRouter()
