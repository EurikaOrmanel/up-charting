package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	Base
	Fullname  string
	Password  string
	Email     string `gorm:"unique"`
	Phone     string `gorm:"unique"`
	Verified  bool
}

func (adM *Admin) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()
	if len(adM.Email) == 0 && len(adM.Phone) == 0 {
		return fmt.Errorf("email or phone number must be provided.")
	}
	return nil
}
