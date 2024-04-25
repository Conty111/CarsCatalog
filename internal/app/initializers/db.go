package initializers

import (
	"github.com/Conty111/CarsCatalog/internal/configs"
	"github.com/Conty111/CarsCatalog/internal/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase(cfg *configs.Configuration) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DB.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("error while connecting to database")
	}
	return db
}

func InitializeMigrations(db *gorm.DB) error {
	var (
		users models.User
		cars  models.Car
	)
	err := db.AutoMigrate(&users, &cars)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run migrations")
		return err
	}
	return nil
}
