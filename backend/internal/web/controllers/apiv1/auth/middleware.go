package auth

import (
	"errors"
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Client"
	userCtx             = "userUUID"
)

func getUserHeaderToken(c *fiber.Ctx) (string, error) {
	header := c.GetRespHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return headerParts[1], nil

}

func getUserUUID(c *fiber.Ctx) (int, error) {
	uuid, ok := c.Locals(userCtx).(int)
	if !ok {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, "user is not found")
		return 0, errors.New("user id not found")
	}

	return uuid, nil
}
