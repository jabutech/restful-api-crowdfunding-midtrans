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
	// Create new variable with type struct RegisterUserInput
	var input user.RegisterUserInput

	// Get data body json from request user and save to variable input
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

	// send response json with status 200, and argument response fromat
	c.JSON(http.StatusOK, response)
}

// Function Handler Login
func (h *userHandler) Login(c *gin.Context) {
	// Create variable with struct LoginInput
	var input user.LoginInput

	// Get data body request input request and save to variable input
	err := c.ShouldBindJSON(&input)
	// If error validation
	if err != nil {
		// Iteration error with helper format validation error
		errors := helper.FormatValidationError(err)
		// Create new map for handle error
		errorMessage := gin.H{"errors": errors}

		// Create format response with helper
		response := helper.ApiResponse(
			"Login failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage, // handle format error from validation
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// If no error, login proccess with service login with set argument input
	loggedinUser, err := h.userService.Login(input)
	// If error validation
	if err != nil {
		// Create new map for handle error
		errorMessage := gin.H{"errors": err.Error()}

		// Create format response with helper
		response := helper.ApiResponse(
			"Login failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage, // handle format error from validation
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// If no error, create response format with use Helper format user
	formatter := user.FormatUser(loggedinUser, "token")

	// Create format response with helper ApiResponse
	response := helper.ApiResponse(
		"Successfuly Loggedin",
		http.StatusOK,
		"success",
		formatter,
	)

	// send response json with status 200, and argument response fromat
	c.JSON(http.StatusOK, response)
}

// Handle for check email
func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// Create var input with type struct CheckEmailInput
	var input user.CheckEmailInput

	// Get data json body and save to var input
	err := c.ShouldBindJSON(&input)
	// If error validation
	if err != nil {
		// Iteration error with helper format validation error
		errors := helper.FormatValidationError(err)
		// Create new map for handle error
		errorMessage := gin.H{"errors": errors}

		// Create format response with helper
		response := helper.ApiResponse(
			"Email Checking failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage, // handle format error from validation
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// If no error validation, Check email is available on database
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	// If error validation
	if err != nil {
		// Create new map for handle error
		errorMessage := gin.H{"errors": "Server error"}

		// Create format response with helper
		response := helper.ApiResponse(
			"Email checking failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage, // handle format error from validation
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.ApiResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
