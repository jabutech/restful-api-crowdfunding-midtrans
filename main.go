package main

import (
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

	// Call NewRepository and set argument db
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Create new router
	router := gin.Default()
	// Create new group api version 1
	api := router.Group("/api/v1")

	// Endpoint
	// Endpoint register user
	api.POST("/users", userHandler.RegisterUser)

	// Run router
	router.Run()
}
