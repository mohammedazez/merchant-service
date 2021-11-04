package handler

import (
	"merchant-service/entity"
	"merchant-service/helper"
	"merchant-service/layer/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// CREATE NEW USER OR REGISTER
func (h *userHandler) RegisterUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	// max len password 6
	if err := helper.ValidatePassword(inputUser.Password); err != nil {

		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"error": "error validate password length < 6"})

		c.JSON(400, responseError)
		return
	}

	newUser, err := h.userService.RegisterUser(inputUser)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success create new user", 201, "Status OK", newUser)
	c.JSON(201, response)
}

//LOGIN User
func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var inputLoginUser entity.LoginUserInput

	if err := c.ShouldBindJSON(&inputLoginUser); err != nil {
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	userData, err := h.userService.LoginUser(inputLoginUser)

	if err != nil {
		responseError := helper.APIResponse("input data required", 401, "bad request", gin.H{"errors": err.Error()})

		c.JSON(401, responseError)
		return
	}

	token, err := h.authService.GenerateTokenUser(userData.ID)

	if err != nil {
		responseError := helper.APIResponse("input data required", 500, "bad request", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success login user", 200, "success", gin.H{"token": token, "role": userData.Role, "id": userData.ID})
	c.JSON(200, response)
}
