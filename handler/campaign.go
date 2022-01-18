package handler

import (
	"gocampaign/campaign"
	"gocampaign/helper"
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
	// handler : mapping id yg di url ke struct input => service call formatter
	// service : input struct => tangkap param id di url, panggil rep
	// repository : get campaign by id

}
