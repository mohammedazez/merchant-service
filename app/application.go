package app

import (
	"fmt"
	"merchant-service/infra"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

func StartApplication() {
	router.Use(infra.CORSMiddleware())
	infra.RegisterApi(router)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
