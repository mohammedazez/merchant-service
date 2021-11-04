package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	APPS_PORT         string
	CONFIGDB_HOST     string
	CONFIGDB_USER     string
	CONFIGDB_PASSWORD string
	CONFIGDB_DBNAME   string
	DB_AUTO_CREATE    bool
}

var (
	EnvFile *Env
)

func init() {
	_ = godotenv.Load()

	env := &Env{}
	env.APPS_PORT = os.Getenv("APPS_PORT")
	env.CONFIGDB_HOST = os.Getenv("CONFIGDB_HOST")
	env.CONFIGDB_USER = os.Getenv("CONFIGDB_USER")
	env.CONFIGDB_PASSWORD = os.Getenv("CONFIGDB_PASSWORD")
	env.CONFIGDB_DBNAME = os.Getenv("CONFIGDB_DBNAME")
	env.DB_AUTO_CREATE = os.Getenv("DB_AUTO_CREATE") == "true"

	EnvFile = env
}
