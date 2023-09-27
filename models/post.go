package models

import (
	"github.com/google/uuid"
)

type Post struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;column:post_id" json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`

	// Relationship: each post belongs to a user
	Author User `gorm:"foreignKey:user_id" json:"user"`

	// Relationship: a post can have many comments
	Comments []Comment `gorm:"foreignKey:post_id" json:"comments"`
}
