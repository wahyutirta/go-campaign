package handler

import (
	"fmt"
	"gocampaign/auth"
	"gocampaign/helper"
	"gocampaign/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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

	if err != nil {
		message := "Register Account Failed"
		response := helper.APIResponse(message, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)

	if err != nil {
		message := "Generate Token Failed"
		response := helper.APIResponse(message, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

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

	token, err := h.authService.GenerateToken(loggedUser.ID)
	if err != nil {
		fmt.Println(err)
		message := "Generate Login Token Failed"
		response := helper.APIResponse(message, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedUser, token)
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

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar") // nama parameter untuk file
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar", http.StatusInternalServerError, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 2
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to save avatar file", http.StatusInternalServerError, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to update avatar in DB", http.StatusInternalServerError, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true, "user": user}
	response := helper.APIResponse("Success uploads avatar file", http.StatusOK, "success", data)
	c.JSON(http.StatusBadRequest, response)

}
