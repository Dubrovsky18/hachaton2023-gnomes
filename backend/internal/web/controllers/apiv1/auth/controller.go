package auth

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/services"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/web/controllers/apiv1"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	apiv1.BaseController
	services *services.Service
}

func NewController(service *services.Service) *Controller {
	return &Controller{
		services: service,
	}

}

type responseWriter struct {
	gin.ResponseWriter
	statusCode int
	length     int
}

// DefineRoutes defines the routes for the message API
func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	apiVer := r.Group("/api/v1")
	oauth2 := apiVer.Group("/oauth2")

	r.Use(func(c *gin.Context) {
		c.Writer = &responseWriter{c.Writer, c.Writer.Status(), 0}
		c.Next()
	})

	oauth2.GET("/login", ctrl.loginOAuth2)
	oauth2.GET("/callback", ctrl.handleOAuth2Callback)
	//apiVer.GET("/auth/google/callback")

	//
	//{
	//	student := apiVer.Group("/student")
	//	{
	//		student.Get("/profile")
	//		student.Put("/profile")
	//		student.Get("/schedule")
	//	}
	//
	//	teacher := apiVer.Group("/teacher")
	//	{
	//		teacher.Get("/profile")
	//		teacher.Put("/profile")
	//		teacherSchedule := teacher.Group("/schedule")
	//		{
	//			// через JWT token определённое расписание
	//			teacherSchedule.Get("/")
	//			teacherSchedule.Post("/wishes")
	//		}
	//	}

	//admin := apiVer.Group("/admin")
	//{
	//	admin.GET("/profiles", ctrl.GetProfiles)
	//	admin.POST("/profiles", ctrl.ChangeUsers)

	//		adminSchedule := admin.Group("/schedule")
	//		{
	//			adminSchedule.Get("/")
	//			adminSchedule.Post("/generate")
	//			adminSchedule.Post("/change")
	//			adminSchedule.Get("/wishes")
	//		}
	//
	//		adminCreate := admin.Group("/create")
	//		{
	//			adminCreate.Post("/student")
	//			adminCreate.Post("/teacher")
	//		}
	//	}
	//}

}
