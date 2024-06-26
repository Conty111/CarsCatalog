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
		wire.FieldsOf(new(*configs.Configuration), "HTTPServer"),
		initializers.InitializeDatabase,
		repositories.NewCarRepository,
		repositories.NewUserRepository,
		external_api.NewClient,
		wire.Bind(new(interfaces.CarManager), new(*repositories.CarRepository)),
		//wire.Bind(new(interfaces.UserManager), new(*repositories.UserRepository)),
		wire.Bind(new(interfaces.UserProvider), new(*repositories.UserRepository)),
		wire.Bind(new(external_api.ExternalAPIClient), new(*external_api.Client)),
		services.NewCarService,
		wire.Struct(new(dependencies.Container), "*"),
		initializers.InitializeRouter,
		initializers.InitializeHTTPServer,
		wire.Struct(new(Application), "*"),
	)

	return &Application{}, nil
}
