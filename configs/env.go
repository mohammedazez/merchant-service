package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	APPS_PORT string
}

var (
	EnvFile *Env
)

func init() {
	_ = godotenv.Load()

	env := &Env{}
	env.APPS_PORT = os.Getenv("APPS_PORT")

	EnvFile = env
}
