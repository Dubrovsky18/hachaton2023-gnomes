package dependencies

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/app/build"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/services"
)

// Container is a DI container for application
type Container struct {
	BuildInfo *build.Info
	Template  services.TemplateService
}
