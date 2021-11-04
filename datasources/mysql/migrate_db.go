package datasources

import (
	"database/sql"
	"fmt"
	"log"

	env "merchant-service/configs"

	"merchant-service/schema"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client   *sql.DB
	Database *gorm.DB
)

func Init_db() {
	dbUsername := env.EnvFile.CONFIGDB_USER
	dbPassword := env.EnvFile.CONFIGDB_PASSWORD
	dbHost := env.EnvFile.CONFIGDB_HOST
	dbSchema := env.EnvFile.CONFIGDB_DBNAME

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		dbUsername, dbPassword, dbHost, dbSchema,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Printf("%#v", err)
	}
	log.Println("database successfully configured")

	autoCreate := env.EnvFile.DB_AUTO_CREATE
	if autoCreate {
		db, err := gorm.Open("mysql", dataSourceName)
		if err != nil {
			log.Println("MySQL:", err)
		}

		fmt.Println("Dropping and recreating all tables...")
		Database = db

		schema.AutoMigrate(db)
		fmt.Println("All tables recreated successfully...")
	}

}
