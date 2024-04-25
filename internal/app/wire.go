//go:build wireinject
// +build wireinject

package app

import (
	"github.com/Conty111/CarsCatalog/internal/app/dependencies"
	"github.com/Conty111/CarsCatalog/internal/app/initializers"
	"github.com/Conty111/CarsCatalog/internal/configs"
	"github.com/Conty111/CarsCatalog/internal/external_api"
	"github.com/Conty111/CarsCatalog/internal/interfaces"
	"github.com/Conty111/CarsCatalog/internal/repositories"
	"github.com/Conty111/CarsCatalog/internal/services"
	"github.com/google/wire"
)

func BuildApplication() (*Application, error) {
	wire.Build(
		initializers.InitializeBuildInfo,
		configs.GetConfig,
		initializers.InitializeDatabase,
		wire.InterfaceValue(new(interfaces.CarManager), new(repositories.CarRepository)),
		//wire.InterfaceValue(new(interfaces.UserManager), new(repositories.UserRepository)),
		wire.InterfaceValue(new(interfaces.UserProvider), new(repositories.UserRepository)),
		external_api.NewClient,
		services.NewCarService,
		wire.Struct(new(dependencies.Container), "*"),
		initializers.InitializeRouter,
		initializers.InitializeHTTPServerConfig,
		initializers.InitializeHTTPServer,
		wire.Struct(new(Application), "*"),
	)

	return &Application{}, nil
}
