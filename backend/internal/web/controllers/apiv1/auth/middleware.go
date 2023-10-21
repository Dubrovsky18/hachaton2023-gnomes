package auth

import (
	"errors"
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Client"
	userCtx             = "userUUID"
)

func getUserHeaderToken(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
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

func getUserUUID(c *gin.Context) (int, error) {
	uuid, ok := c.Get(userCtx)
	if !ok {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, "user is not found")
		return 0, errors.New("user id not found")
	}

	userUUID, ok := uuid.(int)
	if !ok {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}

	return userUUID, nil
}
