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
		c.JSON(http.StatusBadRequest, nil)
	}

	newUser, err := h.userService.RegisterUser(input)
	formatter := user.FormatUser(newUser, "token")

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	message := "Account has been recorded"

	response := helper.APIResponse(message, http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
