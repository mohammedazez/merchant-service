package main

import (
	"merchant-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.UserRoute(r)

	r.Run()
}
