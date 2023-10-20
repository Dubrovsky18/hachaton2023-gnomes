package pkg

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type ErrorResponse struct {
	Error       string `json:"error"`
	Message     string `json:"message"`
	ServiceName string `json:"serviceName"`
}

type StatusResponse struct {
	Message     string `json:"message"`
	ServiceName string `json:"serviceName"`
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

func NewErrorResponse(c *fiber.Ctx, statusCode int, message string) {
	log.Fatalf("Error: %s", message)
	c.Status(statusCode).JSON(ErrorResponse{message, "Error", "Client"})
}

func NewJsonResponse(c *fiber.Ctx, statusCode int, message, serviceName string, mapMessage map[string]string) error {
	newMessage := Message{Message: message, Response: mapMessage, ServiceName: serviceName}
	return c.Status(statusCode).JSON(newMessage)
}

func NewJsonInterfaceResponse(c *fiber.Ctx, statusCode int, message, serviceName string, modelsMessage interface{}) error {
	newMessage := MessageModel{Message: message, Response: modelsMessage, ServiceName: serviceName}
	return c.Status(statusCode).JSON(newMessage)
}

func NewSetContext(c *fiber.Ctx, keyCTX string, valueCTX interface{}) {
	c.Locals(keyCTX, valueCTX)
}
