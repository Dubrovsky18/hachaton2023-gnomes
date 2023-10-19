package render

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/jsonapi"
)

const (
	// ContentTypeHeader is a name of header for Content-Type value
	ContentTypeHeader = "Content-Type"
)

// JSONAPIPayload is marshalling function for JSONAPI payload
func JSONAPIPayload(ctx *fiber.Ctx, statusCode int, payload interface{}) {
	ctx.Set(ContentTypeHeader, jsonapi.MediaType)
	ctx.Status(statusCode)

	if err := jsonapi.MarshalPayload(ctx.Response().BodyWriter(), payload); err != nil {
		logger.Error("jsonapi.MarshalPayload failed", "error", err)

		err = ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		if err != nil {
			logger.Error("Send string in internal/web/render/jsonapi")
			return
		}
		return
	}
}
