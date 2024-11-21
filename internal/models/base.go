package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID  `gorm:"primary_key;not null;type:uuid;unique" json:"id"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func (bM *Base) BeforeCreate(tx *gorm.DB) error {
	bM.ID = uuid.New()
	return nil
}
