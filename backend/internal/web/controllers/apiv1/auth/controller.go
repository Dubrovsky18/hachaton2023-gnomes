package auth

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/services"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/web/controllers/apiv1"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	apiv1.BaseController
	studentService *services.StudentService
	teacherService *services.TeacherService
	adminService   *services.AdminService
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
		auth.Post("/register", ctrl.register)
	}

	{
		oauth2.Post("/login/:role", ctrl.loginOAuth2)
		oauth2.Get("/callback", ctrl.loginOAuth2)
		oauth2.Get("/register", ctrl.register)

	}

	identity := apiVer.Group("/in")

	{
		student := identity.Group("/student")
		{
			student.Get("/profile")
			student.Put("/profile")
			student.Get("/schedule")
		}

		teacher := identity.Group("/teacher")
		{
			teacher.Get("/profile")
			teacher.Put("/profile")
			teacherSchedule := teacher.Group("/schedule")
			{
				// через JWT token определённое расписание
				teacherSchedule.Get("/")
				teacherSchedule.Post("/wishes")
			}
		}

		admin := identity.Group("/admin")
		{
			admin.Get("/profiles", ctrl.GetProfiles)
			admin.Post("/profiles", ctrl.ChangeUsers)

			adminSchedule := admin.Group("/schedule")
			{
				adminSchedule.Get("/")
				adminSchedule.Post("/generate")
				adminSchedule.Post("/change")
				adminSchedule.Get("/wishes")
			}

			adminCreate := admin.Group("/create")
			{
				adminCreate.Post("/student")
				adminCreate.Post("/teacher")
			}
		}
	}

}
