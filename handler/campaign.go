package handler

import (
	"fmt"
	"gocampaign/campaign"
	"gocampaign/helper"
	"gocampaign/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// tangkap parameter ke handler
// handler -> service
// service menentukan repository yang akan dicall
// repository findall findbyid
// db
type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Gets Campaigns Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List Of Campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
	return

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	//api/v1/campaigns/2
	// handler : mapping id yg di url ke struct input => service call formatter
	// service : input struct => tangkap param id di url, panggil rep
	// repository : get campaign by id
	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to bind ID from uri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Campaign detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
	return
}

// tangkap parameter dari user ke input struct
// ambil current user dari jwt
// panggil service, parameter input struct (dan juga buat slug)
// panggil repository untuk simpan data campaign baru

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput
	currentUser := c.MustGet("currentUser").(user.User)
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errMeesage := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed To Bind Campaign Input", http.StatusBadRequest, "error", errMeesage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Failed To Create Campaign Input", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to Create Campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
	return

}

// user masukan input
// handler
// mapping dari input ke input struct
// input dari user, dan juga input yang ada di uri
// service
// repository update data camaign

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to bind ID from uri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputData)

	if err != nil {
		errMeesage := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed To Bind Campaign Input", http.StatusBadRequest, "error", errMeesage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	newCampaign, err := h.service.UpdateCampaign(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed To Update Campaign Input", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to Update Campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
	return

}

func (h *campaignHandler) UploadImage(c *gin.Context) {
	// handler
	// tangkap input dan ubah ke struct input
	// save image campaign ke suatu folder
	// service (kondisi memanggil point 2 di repo1)
	// repository :
	// 	1. create image/save data image ke tabel campaign_image
	// 	2. ubah is_primary true ke false
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed To Verify Campaign Image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	file, err := c.FormFile("file")

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed To Upload Campaign Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to save avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to save avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Success uploads campaign image", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
