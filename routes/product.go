package routes

import (
	"merchant-service/handler"
	"merchant-service/layer/product"

	"github.com/gin-gonic/gin"
)

var (
	productRepository = product.NewRepository(DB)
	productService    = product.NewService(productRepository)
	productHandler    = handler.NewProductHandler(productService, authService)
)

func ProductRoute(r *gin.Engine) {
	r.POST("product", handler.Middleware(userService, authService), productHandler.CreateProductHandler)
}
