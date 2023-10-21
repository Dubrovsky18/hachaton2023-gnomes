package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"golang.org/x/oauth2"
	"net/http"
)

const (
	clientID     = "1028434575027-br2d20a7qbuli837iumif8skmp385dqb.apps.googleusercontent.com"
	clientSecret = "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"
	redirectURL  = "https://www.figma.com/proto/lj8n9p0lfJvJkOS1TrXYQa/Untitled?page-id=0%3A1&type=design&node-id=1-1770&viewport=329%2C168%2C0.14&t=X79cT6CK1AOZuvMq-1&scaling=scale-down&starting-point-node-id=1%3A1770&mode=design"
)

var (
	oauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
)

func (ctrl *Controller) loginOAuth2(c *gin.Context) {
	key := clientSecret  // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30 // 30 days
	isProd := false

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New(clientID, clientSecret, redirectURL, "email", "profile"),
	)

	authURL := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, authURL)
}
func (ctrl *Controller) handleOAuth2Callback(c *gin.Context) {
	code := c.Query("code")

	token, err := oauthConfig.Exchange(c.Request.Context(), code)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	client := oauthConfig.Client(c.Request.Context(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	// Используйте полученный токен для получения информации о пользователе
	// ...

	c.String(http.StatusOK, "Successfully authenticated")
}
