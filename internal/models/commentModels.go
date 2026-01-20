package models

import "github.com/google/uuid"

type Comments struct {
	ID       uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	PostID   uuid.UUIDs `gorm:"type:uuid;not null" json:"post_id"`
	UserID   *uuid.UUID `gorm:"type:uuid" json:"user_id"`
	ParentID uuid.UUID  `gorm:"type:uuid" json:"parent_id"`
}
