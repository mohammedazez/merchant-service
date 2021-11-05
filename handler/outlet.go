package handler

import (
	"merchant-service/auth"
	"merchant-service/entity"
	"merchant-service/helper"
	"merchant-service/layer/outlet"

	"github.com/gin-gonic/gin"
)

type outletHandler struct {
	outletService outlet.Service
	authService   auth.Service
}

func NewOutletHandler(outletService outlet.Service, authService auth.Service) *outletHandler {
	return &outletHandler{outletService, authService}
}

// CREATE NEW Outlet
func (h *outletHandler) CreateOutletHandler(c *gin.Context) {

	// file, err := c.FormFile("Picture")

	userData := c.MustGet("currentUser").(gin.H)
	userID := userData["user_id"].(string)

	if len(userID) == 0 {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user not authorize / not login"})

		c.JSON(401, responseError)
		return
	}
	// path := fmt.Sprintf("images/picture-%d-%s", userData, file.Filename)

	// err = c.SaveUploadedFile(file, path)

	// userID := strconv.Itoa(userData)

	var inputOutlet entity.OutletInput

	if err := c.ShouldBindJSON(&inputOutlet); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	// pathOutletSave := "https://cangkoel.herokuapp.com/" + path

	newOutlet, err := h.outletService.CreateOutlet(inputOutlet, userID)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success create new Outlet", 201, "Status OK", newOutlet)
	c.JSON(201, response)
}
