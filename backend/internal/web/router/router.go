package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// NewRouter create new gin router instance
func NewRouter() *fiber.App {
	r := fiber.New()
	r.Use(recover.New(), logger.New())

	return r
}
