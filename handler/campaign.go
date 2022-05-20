package handler

import (
	"bwacroudfunding/campaign"
	"bwacroudfunding/helper"
	"bwacroudfunding/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

// Instance
func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

// Function for get campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	// Get user id from parameter and convert to inter
	userID, _ := strconv.Atoi(c.Query("user_id"))

	// Get campaign
	campaigns, err := h.service.GetCampaigns(userID)
	// If error
	if err != nil {
		// Create format response with helper ApiResponse
		response := helper.ApiResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		// Create response json
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// If no error, create format api response
	response := helper.ApiResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	// Return json
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// Create new var value with struct GetampaignDetailInput
	var input campaign.GetCampaignDetailInput

	// Get id campaig from uri
	err := c.ShouldBindUri(&input)
	// If error
	if err != nil {
		// Create format response with helper ApiResponse
		response := helper.ApiResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		// Create response json
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Get campain use service
	campaignDetail, err := h.service.GetCampaignByID(input)
	// If error
	if err != nil {
		// Create format response with helper ApiResponse
		response := helper.ApiResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		// Create response json
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// If no error, create format response
	response := helper.ApiResponse("Campaign detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

// Function for create campaign
func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	// Create a var input with value campaign.CreateCampaignInput
	var input campaign.CreateCampaignInput

	// Get data payload
	err := c.ShouldBindJSON(&input)
	// If error validation
	if err != nil {
		// Iteration error with helper format validation error
		errors := helper.FormatValidationError(err)
		// Create new map for handle error
		errorMessage := gin.H{"errors": errors}

		// Create format response with helper
		response := helper.ApiResponse(
			"Failed to create campaign",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage, // handle format error from validation
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Get data current user is logged in
	currentUser := c.MustGet("currentUser").(user.User)
	// Insert current user to property input.User
	input.User = currentUser

	// Create campaign
	newCampaign, err := h.service.CreateCampaign(input)
	// If error validation
	if err != nil {
		// Create format response with helper
		response := helper.ApiResponse(
			"Failed to create campaign",
			http.StatusBadRequest,
			"error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create format response
	response := helper.ApiResponse(
		"Success to create campaign",
		http.StatusOK,
		"error",
		campaign.FormatCampaign(newCampaign),
	)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	// Create new var value with struct GetampaignDetailInput
	var campaignId campaign.GetCampaignDetailInput

	// Get id campaig from uri
	err := c.ShouldBindUri(&campaignId)
	// If error
	if err != nil {
		// Create format response with helper ApiResponse
		response := helper.ApiResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		// Create response json
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create a var inputData with value campaign.CreateCampaignInput
	var inputData campaign.CreateCampaignInput
	// Get data payload
	err = c.ShouldBindJSON(&inputData)
	// If error validation
	if err != nil {
		// Iteration error with helper format validation error
		errors := helper.FormatValidationError(err)
		// Create new map for handle error
		errorMessage := gin.H{"errors": errors}

		// Create format response with helper
		response := helper.ApiResponse(
			"Failed to update campaign",
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
	inputData.User = currentUser

	// Update campaign with service
	updatedCampaign, err := h.service.UpdateCampaign(campaignId, inputData)
	// If error
	if err != nil {
		// Create format response with helper ApiResponse
		response := helper.ApiResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		// Create response json
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create format response with helper ApiResponse
	response := helper.ApiResponse("Success to update campaign", http.StatusOK, "success", campaign.FormatCampaign(updatedCampaign))
	// Return response json
	c.JSON(http.StatusOK, response)
}
