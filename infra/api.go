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
		api.GET("/users/:user_id", Middleware(userService, authService), userHandler.GetUserByIDHandler)    // preload
		api.PUT("/users/:user_id", Middleware(userService, authService), userHandler.UpdateUserByIDHandler) // preload
		api.DELETE("/users/:user_id", Middleware(userService, authService), userHandler.DeleteUserByIDHandler)
		api.POST("/users/outlet", Middleware(userService, authService), userHandler.CreateOutletUserHandler)
		api.GET("/users/outlet", Middleware(userService, authService), userHandler.ShowAllOutletUserHandler)

		// product
		api.POST("/product", Middleware(userService, authService), productHandler.CreateProductHandler)
		api.GET("/product/:outlet_id", Middleware(userService, authService), productHandler.GetProductOutletByIDHandler) // preload
		api.GET("/product", Middleware(userService, authService), productHandler.ShowAllProductHandler)
		api.PUT("/product/:product_id", Middleware(userService, authService), productHandler.UpdateProductByIDHandler)
		api.DELETE("/product/:product_id", Middleware(userService, authService), productHandler.DeleteProductByIDHandler)
	}

}
