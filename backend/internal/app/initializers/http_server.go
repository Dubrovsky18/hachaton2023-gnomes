package initializers

import (
	"fmt"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/app/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CustomHandler struct {
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("log middleware")
		next.ServeHTTP(w, r)
	})
}

// InitializeHTTPServer create new http.Server instance
func InitializeHTTPServer(app *gin.Engine, cfg *config.HTTPConfig) *http.Server {
	// Создаем экземпляр Fiber-совместимого обработчика

	return &http.Server{
		Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
}
