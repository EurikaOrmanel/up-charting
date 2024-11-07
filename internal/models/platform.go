package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Platform struct {
	Base
	Name string `gorm:"not null" json:"name"`
	Url  string `json:"url"`
}

func (adM *Platform) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}
