package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Genre struct {
	Base
	Name string `gorm:"not null" json:"name,omitempty"`
}

func (adM *Genre) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}
