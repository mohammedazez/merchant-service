package handler

import (
	"merchant-service/auth"
	"merchant-service/domain/dto"
	"merchant-service/service"
	"merchant-service/utils/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ProductService
	authService    auth.Service
}

func NewProductHandler(productService service.ProductService, authService auth.Service) *productHandler {
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

	var inputProduct dto.ProductInput

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

// SHOW ALL Product
func (h *productHandler) ShowAllProductHandler(c *gin.Context) {
	product, err := h.productService.ShowAllProduct()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all Product", 200, "status OK", product)
	c.JSON(200, response)
}

// UPDATE Product BY ID
func (h *productHandler) UpdateProductByIDHandler(c *gin.Context) {
	id := c.Params.ByName("product_id")

	userData := c.MustGet("currentUser").(gin.H)
	userID := userData["user_id"].(string)

	if len(userID) == 0 {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "you are not user, not authorize"})

		c.JSON(401, responseError)
		return
	}

	var updateProductInput dto.UpdateProductInput

	if err := c.ShouldBindJSON(&updateProductInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	product, err := h.productService.UpdateProductByID(id, updateProductInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update Product by ID", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

// DELETE Product BY ID
func (h *productHandler) DeleteProductByIDHandler(c *gin.Context) {
	id := c.Params.ByName("product_id")

	userData := c.MustGet("currentUser").(gin.H)
	userID := userData["user_id"].(string)

	if len(userID) == 0 {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "you are not admin, not authorize"})

		c.JSON(401, responseError)
		return
	}

	product, err := h.productService.DeleteProductByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request delete Product", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success delete Product by ID", 200, "success", product)
	c.JSON(200, response)
}

// FIND Outlet BY ID
func (h *productHandler) GetProductOutletByIDHandler(c *gin.Context) {
	id := c.Params.ByName("outlet_id")

	outlet, err := h.productService.FindOutletUserByID(id)
	if err != nil {
		responseError := helper.APIResponse("input params error", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get outlet by ID", 200, "success", outlet)
	c.JSON(200, response)
}
