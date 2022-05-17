package main

import (
	"bwacroudfunding/handler"
	"bwacroudfunding/user"
	"fmt"
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

	input := user.LoginInput{
		Email:    "darmawanrizky43@gmail.com",
		Password: "password",
	}
	user, err := userService.Login(input)
	if err != nil {
		fmt.Println("Terjadi kesalahan")
		fmt.Println(err.Error())
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)

	userHandler := handler.NewUserHandler(userService)

	// Create new router
	router := gin.Default()
	// Create new group api version 1
	api := router.Group("/api/v1")

	// Endpoint
	// Endpoint register user
	api.POST("/users", userHandler.RegisterUser)
	// Endpoint login user
	api.POST("/sessions", userHandler.Login)

	// Run router
	router.Run()
}
