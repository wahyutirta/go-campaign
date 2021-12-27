package main

import (
	"fmt"
	"gocampaign/user"
	"log"

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

	userInput := user.RegisterUserInput{}
	userInput.Name = "Test Simpan"
	userInput.Email = "admin@email.com"
	userInput.Occupation = "Anak Singkong"
	userInput.Password = "password"

	userService.RegisterUser(userInput)

	// input dari user
	// handler -> mapping input -> struct input
	// service -> mapping struct input -> struct user
	// repository
	// db

}
