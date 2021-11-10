package mysql

import (
	"fmt"
	"merchant-service/migration"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var (
// 	database   *gorm.DB
// 	autoCreate bool
// )

func Connection() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		return nil
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	// autoCreate = os.Getenv("MIGRATE") == "true"

	// if autoCreate {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&migration.User{})
	db.AutoMigrate(&migration.Outlet{})
	db.AutoMigrate(&migration.Product{})
	db.AutoMigrate(&migration.ImageProduct{})

	return db
	// }

	// return database
}
