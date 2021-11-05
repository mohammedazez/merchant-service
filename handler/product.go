package handler

import (
	"merchant-service/auth"
	"merchant-service/entity"
	"merchant-service/helper"
	"merchant-service/layer/product"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService product.Service
	authService    auth.Service
}

func NewProductHandler(productService product.Service, authService auth.Service) *productHandler {
	return &productHandler{productService, authService}
}

// CREATE NEW Product
func (h *productHandler) CreateProductHandler(c *gin.Context) {

	userData := c.MustGet("currentUser").(gin.H)
	userID := userData["user_id"].(string)

	if len(userID) == 0 {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "you are not user, not authorize"})

		c.JSON(401, responseError)
		return
	}

	var inputProduct entity.ProductInput

	if err := c.ShouldBindJSON(&inputProduct); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newProduct, err := h.productService.CreateProduct(inputProduct)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success create new Product", 201, "Status OK", newProduct)
	c.JSON(201, response)
}
