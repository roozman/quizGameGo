package main

import (
	"fmt"
	"quizGameGo/entities"
	"quizGameGo/repository/mysql"
)

func main() {
	// TODO - failed user registration increments the id - fix it
	mysqlRepo := mysql.New()

	createdUser, err := mysqlRepo.Register(entities.User{
		ID:          0,
		PhoneNumber: "0933",
		Name:        "hiu",
	})

	if err != nil {
		fmt.Println("register user error:", err)
	} else {
		fmt.Println("created user:", createdUser)
	}
}
