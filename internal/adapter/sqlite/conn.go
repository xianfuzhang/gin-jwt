package sqlite

import (
	// "test/v2/internal/adapter/sqlite/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDB() {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&models.User{})
	DB = db
}
