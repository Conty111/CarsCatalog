package dependencies

import (
	"github.com/Conty111/CarsCatalog/internal/app/build"
	"github.com/Conty111/CarsCatalog/internal/configs"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1/car"
)

// Container is a DI container for application
type Container struct {
	BuildInfo  *build.Info
	Config     *configs.Configuration
	CarService car.Service
}
