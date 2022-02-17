package handler

import (
	"gocampaign/helper"
	"gocampaign/transaction"
	"gocampaign/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

//parameter di uri
// tangkap parameter, mapping ke input struct
// panggil service, input struct sebagai parameter
// service, dengan campaign id, bisa panggil repo
// repo mencari data transaction suatu campaign

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput
	err := c.ShouldBindUri(&input)
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	if err != nil {

		response := helper.APIResponse("Failed to bind campaign's id transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.service.GetTransactionByCampaignID(input)
	if err != nil {

		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Campaign Transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

// get user transaction
// handler
// ambil id user dari jwt/middle ware
// service
// repo => ambil data transactions (preload campaign)

func (h *transactionHandler) GetUserTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	transactions, err := h.service.GetTransactionByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("User's Transactions", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, response)

}
