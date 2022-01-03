package main

import (
	"fmt"
	"gocampaign/auth"
	"gocampaign/handler"
	"gocampaign/helper"
	"gocampaign/user"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/gocampaign?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Success connecting to the database")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/session", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvaibility)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()

	// input dari user
	// handler -> mapping input -> struct input
	// service -> mapping struct input -> struct user
	// repository
	// db

	//middle ware
	// ambil value header authorization
	// dari header authorization, kita ambil nikai token
	// validasi token
	// ambil user id, ambil user dari db berdasarkan user id
	// set context berisi user

}

func authMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if !strings.Contains(authHeader, "Bearer") {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response) // terminate sebelum melanjutkan ke proses selanjutnya
		return
	}

	var tokenString string = ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}
	token, err := 
}
