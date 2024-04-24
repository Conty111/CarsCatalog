package initializers

import (
	"github.com/Conty111/CarsCatalog/internal/app/build"
)

// InitializeBuildInfo creates new build.Info
func InitializeBuildInfo() *build.Info {
	return build.NewInfo()
}
