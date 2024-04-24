//go:build wireinject
// +build wireinject

package app

import (
	"github.com/Conty111/CarsCatalog/internal/app/dependencies"
	"github.com/Conty111/CarsCatalog/internal/app/initializers"
	"github.com/Conty111/CarsCatalog/internal/configs"
	"github.com/google/wire"
)

func BuildApplication() (*Application, error) {
	wire.Build(
		initializers.InitializeBuildInfo,
		configs.GetConfig,
		wire.Struct(new(dependencies.Container), "*"),
		// initialize API client
		initializers.InitializeDatabase,
		initializers.InitializeRouter,
		initializers.InitializeHTTPServerConfig,
		initializers.InitializeHTTPServer,
		wire.Struct(new(Application), "*"),
	)

	return &Application{}, nil
}
