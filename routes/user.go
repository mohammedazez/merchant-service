package routes

import (
	"merchant-service/auth"
	"merchant-service/config"
	"merchant-service/handler"
	"merchant-service/layer/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connection()
	authService             = auth.NewService()
	userService             = user.NewService(userRepository)
	userRepository          = user.NewRepository(DB)
	userHandler             = handler.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	r.POST("users/register", userHandler.RegisterUserHandler) // done
	r.POST("users/login", userHandler.LoginUserHandler)       // done
}
