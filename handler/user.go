package handler

import (
	"bwacroudfunding/helper"
	"bwacroudfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// Handler Register user
func (h *userHandler) RegisterUser(c *gin.Context) {
	// Create new variable with value sturuct RegisterUserInput
	var input user.RegisterUserInput

	// Get data json from request user
	err := c.ShouldBindJSON(&input)

	// If error validation
	if err != nil {
		// Iteration error with helper format validation error
		errors := helper.FormatValidationError(err)
		// Create new map for handle error
		errorMessage := gin.H{"errors": errors}

		// Create format response with helper
		response := helper.ApiResponse(
			"Register account failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage, // handle format error from validation
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// If no error, register user with service Register User
	newUser, err := h.userService.RegisterUser(input)
	// If error
	if err != nil {
		// Create format response with helper
		response := helper.ApiResponse(
			"Register account failed",
			http.StatusBadRequest,
			"error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create response format user with formatter
	formatter := user.FormatUser(newUser, "rahasiatoken")

	// Create format response with helper
	response := helper.ApiResponse(
		"Account has been registered",
		http.StatusOK,
		"success",
		formatter,
	)

	// If no error, send response 200
	c.JSON(http.StatusOK, response)
}
