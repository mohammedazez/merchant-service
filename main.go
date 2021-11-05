package main

import (
	"merchant-service/handler"
	"merchant-service/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(handler.CORSMiddleware())

	routes.UserRoute(r)
	routes.OutletRoute(r)

	dbPort := os.Getenv("APP_PORT")
	r.Run(dbPort)
}
