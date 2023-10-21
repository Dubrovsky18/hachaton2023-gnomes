package pkg

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/web/render"
	"github.com/gin-gonic/gin"
	"log"
)

type ErrorResponse struct {
	Error       string `json:"error"`
	Message     string `json:"message"`
	ServiceName string `json:"serviceName"`
}

type StatusResponse struct {
	Status      string `json:"status,omitempty"`
	Message     string `json:"message,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`
}

type Message struct {
	Message     string            `json:"*message*,omitempty"`
	Response    map[string]string `json:"response,omitempty"`
	ServiceName string            `json:"service_name,omitempty"`
}

type MessageModel struct {
	Message     string      `json:"message,omitempty"`
	Response    interface{} `json:"response,omitempty"`
	ServiceName string      `json:"service_name,omitempty"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Fatalf("Error: %s", message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message, "Error", "Client"})
}

func NewJsonResponse(c *gin.Context, statusCode int, message, serviceName string, mapMessage map[string]string) {
	newMessage := Message{Message: message, Response: mapMessage, ServiceName: serviceName}
	c.JSON(statusCode, newMessage)
}

func NewJsonInterfaceResponse(c *gin.Context, statusCode int, message, serviceName string, modelsMessage interface{}) {
	newMessage := MessageModel{Message: message, Response: modelsMessage, ServiceName: serviceName}
	render.JSONAPIPayload(c, statusCode, newMessage)
}

func NewSetContext(c *gin.Context, keyCTX string, valueCTX any) {
	c.Set(keyCTX, valueCTX)
}

func NewStatusResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, StatusResponse{Status: message})
}
