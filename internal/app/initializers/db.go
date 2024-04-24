package initializers

import (
	"github.com/Conty111/CarsCatalog/internal/app/dependencies"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase(container *dependencies.Container) *gorm.DB {
	db, err := gorm.Open(postgres.Open(container.Config.DB.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("error while connecting to database")
	}
	return db
}
