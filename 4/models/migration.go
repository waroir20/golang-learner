package models

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	log.Info().Msg("Running migration")
	return db.AutoMigrate(
		User{},
		Address{},
	)
}
