package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model           // Used for PK and some auditing fields
	FirstName  string    `gorm:"not null" validate:"required"`
	LastName   string    `gorm:"not null"`
	DOB        string    `gorm:"not null"`
	Addresses  []Address `gorm:"constraint:OnDelete:CASCADE;"`
}

// BeforeCreate - Implements the callbacks.BeforeCreateInterface
func (u *User) BeforeCreate(tx *gorm.DB) error {
	_ = tx //Fake use tx to eliminate IDE check
	// TODO add any other logic you like
	return modelValidator.Struct(u)
}

// BeforeUpdate - Implements the callbacks.BeforeUpdateInterface
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	_ = tx //Fake use tx to eliminate IDE check
	// TODO add any other logic you like
	return modelValidator.Struct(u)
}
