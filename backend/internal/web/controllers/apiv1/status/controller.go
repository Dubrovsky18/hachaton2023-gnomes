package status

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/app/build"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/web/controllers/apiv1"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/web/render"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var (
	_ apiv1.Controller = (*Controller)(nil)
)

// Controller is a controller implementation for status checks
type Controller struct {
	apiv1.BaseController
	buildInfo *build.Info
}

// NewController creates new status controller instance
func NewController(bi *build.Info) *Controller {
	return &Controller{
		buildInfo: bi,
	}
}

// GetStatus godoc
// @Summary Get Application Status
// @Description get status
// @ID get-status
// @Accept json
// @Produce json
// @Success 200 {object} ResponseDoc
// @Router /api/v1/status [get]
func (ctrl *Controller) GetStatus(ctx *fiber.Ctx) error {
	render.JSONAPIPayload(ctx, http.StatusOK, &Response{
		Status: http.StatusText(http.StatusOK),
		Build:  ctrl.buildInfo,
	})
	return nil
}

// DefineRoutes adds controller routes to the router
func (ctrl *Controller) DefineRoutes(r fiber.Router) {
	r.Get("/api/v1/status", ctrl.GetStatus)
}
