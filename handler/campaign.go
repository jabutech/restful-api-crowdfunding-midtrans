package handler

import (
	"bwacroudfunding/campaign"
	"bwacroudfunding/helper"
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
