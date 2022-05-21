package handler

import (
	"bwacroudfunding/helper"
	"bwacroudfunding/transaction"
	"bwacroudfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHanlder(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

// Function for get data campaign transaction
func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput
	// Get data from uri
	err := c.ShouldBindUri(&input)
	// If error
	if err != nil {
		// Create format response with helper ApiResponse
		response := helper.ApiResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		// Create response json
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Get data current user is logged in
	currentUser := c.MustGet("currentUser").(user.User)
	// Insert current user to property inputData.User
	input.User = currentUser

	// Find campaign user service
	transactions, err := h.service.GetTransactionByCampaignID(input)
	// If error
	if err != nil {
		// Create format response with helper ApiResponse
		response := helper.ApiResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		// Create response json
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// If no error, create format response
	response := helper.ApiResponse("Campaign's transaction", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransaction(c *gin.Context) {
	// Get current user logged in
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	// Get transaction based on current user logged in
	transactions, err := h.service.GetTransactionByUserID(userID)
	// If error
	if err != nil {
		// Create format response with helper ApiResponse
		response := helper.ApiResponse("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		// Create response json
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// If no error, create format response
	response := helper.ApiResponse("User's transaction", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
