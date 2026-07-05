package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"quizGameGo/entities"
	"quizGameGo/repository/mysql"
	"quizGameGo/service/userService"
)

func main() {
	//mux := http.NewServeMux()
	http.HandleFunc("/health-check", healthCheckHandler)
	http.HandleFunc("/users/register", userRegisterHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"alive": true}`)
}

func userRegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprint(w, `{"error": "only POST method is allowed"}`)
		return
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, `{"error_onread": "%s"}`, err.Error())
		return
	}

	var req userService.RegisterRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		fmt.Fprintf(w, `{"error_unmarshal": "%s"}`, err.Error())
		return
	}

	mysqlRepo := mysql.New()
	userSvc := userService.New(mysqlRepo)

	_, err = userSvc.Register(req)
	if err != nil {
		fmt.Fprintf(w, `{"error_register": "%s"}`, err.Error())
		return
	}

	fmt.Fprint(w, `{"message": "user registered"`)
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
