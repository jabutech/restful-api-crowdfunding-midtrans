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

// Function for create new data campaign transaction
func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	// Get data payload json
	err := c.ShouldBindJSON(&input)

	// If error validation
	if err != nil {
		// Iteration error with helper format validation error
		errors := helper.FormatValidationError(err)
		// Create new map for handle error
		errorMessage := gin.H{"errors": errors}

		// Create format response with helper
		response := helper.ApiResponse(
			"Failed to create transaction",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage, // handle format error from validation
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Get data current user is logged in
	currentUser := c.MustGet("currentUser").(user.User)
	// Insert current user to property inputData.User
	input.User = currentUser

	// Create transaction
	newTransaction, err := h.service.CreateTransaction(input)

	// If error validation
	if err != nil {
		// Create format response with helper
		response := helper.ApiResponse(
			"Failed to create transaction",
			http.StatusBadRequest,
			"error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create format response
	response := helper.ApiResponse(
		"Success to create transaction",
		http.StatusOK,
		"success",
		transaction.FormatTransaction(newTransaction),
	)
	c.JSON(http.StatusOK, response)
}
