package auth

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

const (
	key                  = "1028434575027-br2d20a7qbuli837iumif8skmp385dqb.apps.googleusercontent.com"
	sec                  = "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"
	GOOGLE_AUTH_URI      = "https://accounts.google.com/o/oauth2/auth"
	GOOGLE_TOKEN_URI     = "https://accounts.google.com/o/oauth2/token"
	GOOGLE_USER_INFO_URI = "https://www.googleapis.com/oauth2/v1/userinfo"
)

var GOOGLE_SCOPES = []string{
	"https://www.googleapis.com/auth/userinfo.email",
	"https://www.googleapis.com/auth/userinfo.profile",
}

var conf = &oauth2.Config{
	Endpoint: oauth2.Endpoint{
		AuthURL:  GOOGLE_AUTH_URI,
		TokenURL: GOOGLE_TOKEN_URI,
	},
}

func (ctrl *Controller) loginOAuth2(c *fiber.Ctx) error {

	authURL := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(authURL)
}

func (ctrl *Controller) handleOAuth2Callback(c *fiber.Ctx) error {
	code := c.Query("code")

	token, err := conf.Exchange(c.Context(), code)
	if err != nil {
		return err
	}

	client := conf.Client(c.Context(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email string `json:"email"`
		// Другие поля, которые вам интересны
	}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return err
	}

	// Используйте информацию о пользователе по своему усмотрению
	// Например, сохраните email в сессии или базе данных

	return c.SendString("Successfully authenticated")
}
