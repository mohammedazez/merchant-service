package routes

import (
	"merchant-service/handler"
	"merchant-service/layer/outlet"

	"github.com/gin-gonic/gin"
)

var (
	outletRepository = outlet.NewRepository(DB)
	outletService    = outlet.NewService(outletRepository)
	outletHandler    = handler.NewOutletHandler(outletService, authService)
)

func OutletRoute(r *gin.Engine) {
	r.POST("outlet", handler.Middleware(userService, authService), outletHandler.CreateOutletHandler)
	r.GET("outlet", handler.Middleware(userService, authService), outletHandler.ShowAllOutletHandler)
	r.GET("outlet/:outlet_id", handler.Middleware(userService, authService), outletHandler.GetOutletByIDHandler)

}
