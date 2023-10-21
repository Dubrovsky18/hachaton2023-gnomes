package render

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
)

const (
	// ContentTypeHeader is a name of header for Content-Type value
	ContentTypeHeader = "Content-Type"
)

// JSONAPIPayload is marshalling function for JSONAPI payload
func JSONAPIPayload(ctx *gin.Context, statusCode int, payload interface{}) {
	ctx.Header(ContentTypeHeader, jsonapi.MediaType)
	ctx.Status(statusCode)

	if err := jsonapi.MarshalPayload(ctx.Writer, payload); err != nil {
		logger.Error("jsonapi.MarshalPayload failed", "error", err)

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

// Errors renders errors in JSONAPI format
func Errors(ctx *gin.Context, statusCode int, errs []*jsonapi.ErrorObject) {
	ctx.Header(ContentTypeHeader, jsonapi.MediaType)
	ctx.Status(statusCode)

	if err := jsonapi.MarshalErrors(ctx.Writer, errs); err != nil {
		logger.Error("jsonapi.MarshalErrors failed", "error", err)

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
