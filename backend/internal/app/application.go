package app

import (
	"context"
	"fmt"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/app/dependencies"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/app/initializers"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/repository"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/repository/postgresql"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/services"
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg/logger"

	"net/http"
)

// Application is a main struct for the application that contains general information
type Application struct {
	httpServer *http.Server
	Container  *dependencies.Container
}

// InitializeApplication initializes new application
func InitializeApplication() (*Application, error) {
	initializers.InitializeEnv()
	appConfig := initializers.InitializeAppConfig()
	info := initializers.InitializeBuildInfo()

	reposPostgres := repository.NewPostgresDB()
	st := postgresql.NewStudentPostgres(reposPostgres)
	te := postgresql.NewTeacherPostgres(reposPostgres)
	ad := postgresql.NewAdminPostgres(reposPostgres)
	sc := postgresql.NewSchedulePostgres(reposPostgres)
	templateServices := services.NewService(st, te, ad, sc)

	container := &dependencies.Container{
		BuildInfo: info,
		Template:  *templateServices,
	}

	router := initializers.InitializeRouter(container)
	server := initializers.InitializeHTTPServer(router, appConfig.HTTP)

	return &Application{
		httpServer: server,
		Container:  container,
	}, nil
}

// Start starts application services
func (a *Application) Start(ctx context.Context, cli bool) {
	if cli {
		return
	}

	a.startHTTPServer()
}

// Stop stops application services
func (a *Application) Stop() (err error) {
	return a.httpServer.Shutdown(context.TODO())
}

func (a *Application) startHTTPServer() {
	go func() {
		logger.Info(fmt.Sprintf("started http server on address: %s", a.httpServer.Addr))

		// service connections
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("HTTP Server stopped", "error", err)
		}
	}()
}
