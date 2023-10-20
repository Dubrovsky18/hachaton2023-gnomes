package auth

import (
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

func (ctrl *Controller) loginOAuth2(c *fiber.Ctx) error {

	conf := &oauth2.Config{
		Endpoint: oauth2.Endpoint{
			AuthURL:  GOOGLE_AUTH_URI,
			TokenURL: GOOGLE_TOKEN_URI,
		},
	}

	authURL := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(authURL)
}
