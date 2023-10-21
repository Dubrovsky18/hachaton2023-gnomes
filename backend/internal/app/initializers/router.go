package initializers

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/app/dependencies"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/web/controllers/apiv1"
	apiV1Auth "github.com/Dubrovsky18/hachaton2023-gnomes/internal/web/controllers/apiv1/auth"
	apiV1Status "github.com/Dubrovsky18/hachaton2023-gnomes/internal/web/controllers/apiv1/status"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/web/router"
	"github.com/gin-gonic/gin"
)

// InitializeRouter initializes new gin router
func InitializeRouter(container *dependencies.Container) *gin.Engine {
	r := router.NewRouter()

	ctrls := buildControllers(container)

	for i := range ctrls {
		ctrls[i].DefineRoutes(r)
	}

	return r
}

func buildControllers(container *dependencies.Container) []apiv1.Controller {
	return []apiv1.Controller{
		apiV1Status.NewController(container.BuildInfo),
		apiV1Auth.NewController(&container.Template),
		//apiv1Swagger.NewController(),

	}
}
