package infra

import (
	"merchant-service/auth"
	"merchant-service/infra/mysql"

	"merchant-service/handler"
	"merchant-service/service"
	"merchant-service/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = mysql.Connection()

	authService    = auth.NewService()
	userService    = service.NewUserService(userRepository)
	userRepository = storage.NewDao(DB)
	userHandler    = handler.NewUserHandler(userService, authService)

	productRepository = storage.NewProductDao(DB)
	productService    = service.NewProductService(productRepository)
	productHandler    = handler.NewProductHandler(productService, authService)
)

func RegisterApi(r *gin.Engine) {
	api := r.Group("/api")
	{
		// user
		api.POST("/users/register", userHandler.RegisterUserHandler)
		api.POST("/users/login", userHandler.LoginUserHandler)
		api.GET("/users", Middleware(userService, authService), userHandler.ShowAllUserHandler)
		api.GET("/users/:user_id", Middleware(userService, authService), userHandler.GetUserByIDHandler)
		api.PUT("/users/:user_id", Middleware(userService, authService), userHandler.UpdateUserByIDHandler)
		api.DELETE("/users/:user_id", Middleware(userService, authService), userHandler.DeleteUserByIDHandler)
		api.POST("/users/outlet", Middleware(userService, authService), userHandler.CreateOutletUserHandler)
		api.GET("/users/outlet", userHandler.ShowAllOutletUserHandler)

		// product
		api.POST("/product", Middleware(userService, authService), productHandler.CreateProductHandler)
		api.POST("/product/display-image", Middleware(userService, authService), productHandler.CreateDisplayImageProduct)
		api.GET("/product/:outlet_id", productHandler.GetProductOutletByIDHandler)
		api.GET("/product", productHandler.ShowAllProductHandler)
		api.GET("/product/detail/:product_id", productHandler.GetProductByIDHandler)
		api.PUT("/product/:product_id", Middleware(userService, authService), productHandler.UpdateProductByIDHandler)
		api.DELETE("/product/:product_id", Middleware(userService, authService), productHandler.DeleteProductByIDHandler)
	}

}
