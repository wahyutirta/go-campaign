package handler

import (
	"gocampaign/helper"
	"gocampaign/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)

		errMessage := gin.H{"errors": errors}

		message := "Register Account Failed"
		response := helper.APIResponse(message, http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	formatter := user.FormatUser(newUser, "token")

	if err != nil {
		message := "Register Account Failed"
		response := helper.APIResponse(message, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	message := "Account has been recorded"

	response := helper.APIResponse(message, http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// user memasukan input email dan password
	// input ditangkap handler
	// mapping input user ke struct
	// input struct passing service
	// di service mencari dengan bantuan repository user dengan email
	// mencocokan password
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}
		message := "Login Failed"
		response := helper.APIResponse(message, http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedUser, err := h.userService.Login(input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}
		message := "Login Failed"
		response := helper.APIResponse(message, http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedUser, "token")
	response := helper.APIResponse("Successfully Logging In", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvaibility(c *gin.Context) {
	// input email dari user
	// input email di-mapping ke struct input
	// struct di-passing ke service
	// service akan manggil repository - email sudah ada atau belum
	// repository - db
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}
		message := "Email Checking Failed"
		response := helper.APIResponse(message, http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {

		errMessage := gin.H{"errors": "Server error"}
		message := "Email Checking Failed"
		response := helper.APIResponse(message, http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	var metaMessage string

	if isEmailAvailable {
		metaMessage = "Email is available"
	} else {
		metaMessage = "Email has been registered"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "error", data)
	c.JSON(http.StatusOK, response)
}
