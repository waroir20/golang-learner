package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model   // Used for PK and some auditing fields
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Country      string
	PostalCode   string
	//Start is needed for GORM
	UserID *uint
	User   *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	//End is needed for GORM
}

// BeforeCreate - Implements the callbacks.BeforeCreateInterface
func (u *Address) BeforeCreate(tx *gorm.DB) error {
	_ = tx //Fake use tx to eliminate IDE check
	// TODO add any other logic you like
	return modelValidator.Struct(u)
}

// BeforeUpdate - Implements the callbacks.BeforeUpdateInterface
func (u *Address) BeforeUpdate(tx *gorm.DB) error {
	_ = tx //Fake use tx to eliminate IDE check
	// TODO add any other logic you like
	return modelValidator.Struct(u)
}
