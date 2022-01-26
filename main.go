package main

import (
	"fmt"
	"gocampaign/auth"
	"gocampaign/campaign"
	"gocampaign/handler"
	"gocampaign/helper"
	"gocampaign/user"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
	campaignRepository := campaign.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	campaignService := campaign.NewService(campaignRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	router.Static("/images", "./images") //static image

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/session", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvaibility)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)

	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
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

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		fmt.Println(authHeader)

		// add authorization bearer token in postman in testing
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error no bearer detected", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // terminate sebelum melanjutkan ke proses selanjutnya
			return
		}

		var tokenString string = ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error validate", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // terminate sebelum melanjutkan ke proses selanjutnya
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // terminate sebelum melanjutkan ke proses selanjutnya
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {
			response := helper.APIResponse("Internal Server Error", http.StatusInternalServerError, "error", nil)
			c.AbortWithStatusJSON(http.StatusInternalServerError, response) // terminate sebelum melanjutkan ke proses selanjutnya
			return
		}
		c.Set("currentUser", user)
	}
}
