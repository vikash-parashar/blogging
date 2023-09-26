// config/config.go
package config

import (
	"blogging/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
	var err error
	db, err := gorm.Open("postgres", "host=localhost user=postgres password=admin dbname=blogging port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	DB = db
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	// defer db.Close()
}
