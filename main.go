package main

import (
	"fmt"
	"quizGameGo/entities"
	"quizGameGo/repository/mysql"
)

func main() {
	testUserCreationRepo()
}

func testUserCreationRepo() {
	// TODO - failed user registration increments the id - fix it
	mysqlRepo := mysql.New()

	createdUser, err := mysqlRepo.Register(entities.User{
		ID:          0,
		PhoneNumber: "09323",
		Name:        "hiu",
	})

	if err != nil {
		fmt.Println("register user error:", err)
	} else {
		fmt.Println("created user:", createdUser)
	}

	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber)
	if err != nil {
		fmt.Println("unique user phone number error:", err)
	}
	fmt.Println("isUnique:", isUnique)
}
