package initializers

import "github.com/Dubrovsky18/hachaton2023-gnomes/internal/app/build"

// InitializeBuildInfo creates new build.Info
func InitializeBuildInfo() *build.Info {
	return build.NewInfo()
}
