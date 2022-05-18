package main

import (
	"bwacroudfunding/auth"
	"bwacroudfunding/handler"
	"bwacroudfunding/user"
	"log"

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
	api.POST("/avatars", userHandler.UploadAvatar)

	// Run router
	router.Run()
}
