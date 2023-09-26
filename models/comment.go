package models

import (
	"github.com/google/uuid"
)

// models/comment.go
type Comment struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;column:comment_id" json:"id"`
	Content string    `json:"content"`

	// Relationship: each comment belongs to a user
	Author User `gorm:"foreignKey:AuthorID" json:"author"`

	// Relationship: each comment belongs to a post
	Post Post `gorm:"foreignKey:PostID" json:"post"`
}
