package schema

import (
	"github.com/jinzhu/gorm"
)

var (
	Database *gorm.DB
)

func AutoMigrate(database *gorm.DB) {

	Database = database

	// Auto migrate
	database.AutoMigrate(
		Users{},
	)
}
