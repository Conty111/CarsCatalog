package app

import (
	"context"
	"gorm.io/gorm"
	"net/http"

	"github.com/Conty111/CarsCatalog/internal/app/dependencies"
	"github.com/Conty111/CarsCatalog/internal/app/initializers"
	"github.com/rs/zerolog/log"
)

// Application is a main struct for the application that contains general information
type Application struct {
	httpServer *http.Server
	db         *gorm.DB
	Container  *dependencies.Container
}

// InitializeApplication initializes new application
func InitializeApplication() (*Application, error) {
	initializers.InitializeEnvs()

	if err := initializers.InitializeLogs(); err != nil {
		return nil, err
	}

	app, err := BuildApplication()
	if err != nil {
		return nil, err
	}
	err = initializers.InitializeMigrations(app.db)
	if err != nil {
		return nil, err
	}
	return app, nil
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
	log.Info().Msg("gracefully stopping")
	return a.httpServer.Shutdown(context.TODO())
}

func (a *Application) startHTTPServer() {
	go func() {
		log.Info().Str("HTTPServerAddress", a.httpServer.Addr).Msg("started http server")

		// service connections
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panic().Err(err).Msg("HTTP Server stopped")
		}
	}()
}
