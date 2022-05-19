package main

import (
	"bwacroudfunding/auth"
	"bwacroudfunding/campaign"
	"bwacroudfunding/handler"
	"bwacroudfunding/helper"
	"bwacroudfunding/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Url connection to mysql
	dsn := "root:root@tcp(127.0.0.1:3306)/bwafunding?charset=utf8mb4&parseTime=True&loc=Local"
	// Open connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Check if error
	if err != nil {
		log.Fatal(err.Error())
	}

	// Repository
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)

	// Service
	authService := auth.NewService()
	userService := user.NewService(userRepository)

	// Handler
	userHandler := handler.NewUserHandler(userService, authService)

	// Create new router
	router := gin.Default()
	// Create new group api version 1
	api := router.Group("/api/v1")

	// Endpoint
	// Endpoint register user
	api.POST("/users", userHandler.RegisterUser)
	// Endpoint login user
	api.POST("/sessions", userHandler.Login)
	// Endpoint email checkers
	api.POST("/email-checkers", userHandler.CheckEmailAvailability)
	// Endpoint avatars
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	// Run router
	router.Run()
}

// Function for auth middleware
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get header with name `Authorization`
		authHeader := c.GetHeader("Authorization")

		// If inside authHeader doesn't have `Bearer`
		if !strings.Contains(authHeader, "Bearer") {
			// Create format response with helper
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			// Stop process and return response
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// If there is, create new variable with empty string value
		tokenString := ""
		// Split authHeader with white space
		arrayToken := strings.Split(authHeader, " ")
		// If length arrayToken is same the 2
		if len(arrayToken) == 2 {
			// Get arrayToken with index 1 / only token jwt
			tokenString = arrayToken[1]
		}

		// Validation token with authService Validation Token
		token, err := authService.ValidateToken(tokenString)
		// If error
		if err != nil {
			// Create format response with helper
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			// Stop process and return response
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Get payload token
		claim, ok := token.Claims.(jwt.MapClaims)
		// If not `ok` and token invalid
		if !ok || !token.Valid {
			// Create format response with helper
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			// Stop process and return response
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Get payload `user_id` and convert to type `float64` and type `int`
		userId := int(claim["user_id"].(float64))

		// Find user on db with service
		user, err := userService.GetUserByID(userId)
		// If error
		if err != nil {
			// Create format response with helper
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			// Stop process and return response
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Set user to context with name `currentUser`
		c.Set("currentUser", user)
	}
}
