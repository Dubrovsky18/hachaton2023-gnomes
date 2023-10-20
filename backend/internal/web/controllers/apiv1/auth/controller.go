package auth

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/services"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/web/controllers/apiv1"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	apiv1.BaseController
	studentService *services.StudentService
}

func NewController(service *services.StudentService) *Controller {
	return &Controller{
		studentService: service,
	}
}

// DefineRoutes defines the routes for the message API
func (ctrl *Controller) DefineRoutes(r *fiber.App) {
	apiVer := r.Group("/api/v1/")
	auth := apiVer.Group("/auth")
	oauth2 := apiVer.Group("/oauth2")

	{
		auth.Post("/login/:role", ctrl.loginAuth)
	}

	{
		oauth2.Post("/login/:role", ctrl.loginOAuth2)
	}

}
