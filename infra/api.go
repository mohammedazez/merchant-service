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

	outletRepository = storage.NewDao(DB)
	outletService    = service.NewOutletService(outletRepository)
	outletHandler    = handler.NewOutletHandler(outletService, authService)

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

		// outlet
		api.POST("/outlet", Middleware(userService, authService), outletHandler.CreateOutletHandler)
		api.GET("/outlet", Middleware(userService, authService), outletHandler.ShowAllOutletHandler)
		api.GET("/outlet/:outlet_id", Middleware(userService, authService), outletHandler.GetOutletByIDHandler)

		// product
		api.POST("/product", Middleware(userService, authService), productHandler.CreateProductHandler)
		api.GET("/product", Middleware(userService, authService), productHandler.ShowAllProductHandler)
		api.GET("/product/:product_id", Middleware(userService, authService), productHandler.GetProductByIDHandler)
		api.PUT("/product/:product_id", Middleware(userService, authService), productHandler.UpdateProductByIDHandler)
		api.DELETE("/product/:product_id", Middleware(userService, authService), productHandler.DeleteProductByIDHandler)

	}

}
