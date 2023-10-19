package render

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg/logger"
	"github.com/gobuffalo/envy"
	"github.com/gofiber/fiber/v2"
)

type ResponseError struct {
	Error       string `json:"error"`
	Message     string `json:"message.go"`
	ServiceName string `json:"serviceName"`
}

var ServiceName string = envy.Get("SERVICE_NAME", "Gobase")

func SendError(ctx *fiber.Ctx, statusCode int, err error, message string) {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	ctx.Status(statusCode)
	var response = ResponseError{
		Error:       err.Error(),
		Message:     message,
		ServiceName: ServiceName,
	}

	if err := ctx.JSON(&response); err != nil {
		logger.Error("jsonapi.MarshalErrors failed", "error", err)

		ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		return
	}
}
