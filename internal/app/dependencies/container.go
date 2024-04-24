package dependencies

import (
	"github.com/Conty111/CarsCatalog/internal/app/build"
	"github.com/Conty111/CarsCatalog/internal/configs"
)

// Container is a DI container for application
type Container struct {
	BuildInfo *build.Info
	Config    *configs.Configuration
}
