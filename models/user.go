package models

import (
	"github.com/google/uuid"
)

const (
	Admin   string = "admin"
	General string = "general"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;column:user_id" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`

	// Relationship: a user can have many posts
	Posts []Post `gorm:"foreignKey:user_id" json:"posts"`

	// Relationship: a user can have many comments
	Comments []Comment `gorm:"foreignKey:user_id" json:"comments"`
}
