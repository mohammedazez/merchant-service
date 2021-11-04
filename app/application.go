package app

import (
	"merchant-service/configs"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

func StartApplication() {

	// mapUrls()

	port := configs.EnvFile.APPS_PORT
	router.Run(port)
}
