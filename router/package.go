package router

import (
	"strings"

	"github.com/gofiber/fiber"
)

// Find func
func Find(fiber *fiber.Ctx) {
	info := []string{"Searching for package info ->", fiber.Params("packageName"), "\n"}
	fiber.Send(strings.Join(info, " "))
}
