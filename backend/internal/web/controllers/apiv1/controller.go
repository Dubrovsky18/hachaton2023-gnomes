package apiv1

import (
	"github.com/gofiber/fiber/v2"
)

// Controller is an interface for HTTP controllers
type Controller interface {
	DefineRoutes(fiber.Router)
}
