package handler

import (
	"gocampaign/helper"
	"gocampaign/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

		errMessage := gin.H{"erros": errors}

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
