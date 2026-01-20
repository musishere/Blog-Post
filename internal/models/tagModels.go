package models

import (
	"time"

	"github.com/google/uuid"
)

type Tag struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid.generate_v4();primaryKey" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	Slug      string    `gorm:"unique;not null" json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Posts []Post `gorm:"many2many" json:"post_tags"`
}
