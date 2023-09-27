package models

import (
	"github.com/google/uuid"
)

type Comment struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;column:comment_id" json:"id"`
	Text string    `json:"text"`

	// Relationship: each comment belongs to a user
	Author User `gorm:"foreignKey:user_id" json:"user"`

	// Relationship: each comment belongs to a post
	Post Post `gorm:"foreignKey:post_id" json:"post"`
}
