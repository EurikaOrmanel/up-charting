package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Artist struct {
	Base
	Stagename string           `json:"stagename"`
	Firstname string           `json:"firstname"`
	Lastname  string           `json:"lastname"`
	Platforms []ArtistPlatform `json:"platforms"`
}

func (adM *Artist) BeforeCreate(tx *gorm.DB) error {
	adM.ID = uuid.New()

	return nil
}
