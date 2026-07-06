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

const (
	jwtTestKey = "testkey"
)

func main() {
	http.HandleFunc("/health-check", healthCheckHandler)
	http.HandleFunc("/users/register", userRegisterHandler)
	http.HandleFunc("/users/login", userLoginHandler)
	http.HandleFunc("/users/profile", userProfileHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}

// TODO - should get userid in a jwt
// TODO - GET shouldn't send body
func userProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprint(w, `{"error": "only Get method is allowed"}`)
		return
	}

	req := userService.ProfileRequest{UserID: 0}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, `{"error_onread": "%s"}`, err.Error())
		return
	}

	err = json.Unmarshal(data, &req)
	if err != nil {
		fmt.Fprintf(w, `{"error_unmarshal": "%s"}`, err.Error())
		return
	}

	mysqlRepo := mysql.New()
	userSvc := userService.New(mysqlRepo, jwtTestKey)

	response, err := userSvc.Profile(req)
	if err != nil {
		fmt.Fprintf(w, `{"error_profile": "%s"}`, err.Error())
		return
	}

	data, err = json.Marshal(response)
	if err != nil {
		fmt.Fprintf(w, `{"error_profile": "%s"}`, err.Error())
	}

	fmt.Fprintf(w, string(data))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"alive": true}`)
}

func userLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprint(w, `{"error": "only POST method is allowed"}`)
		return
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, `{"error_onread": "%s"}`, err.Error())
		return
	}

	var req userService.LoginRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		fmt.Fprintf(w, `{"error_unmarshal": "%s"}`, err.Error())
		return
	}

	mysqlRepo := mysql.New()
	userSvc := userService.New(mysqlRepo, jwtTestKey)

	response, err := userSvc.Login(req)
	if err != nil {
		fmt.Fprintf(w, `{"error_login": "%s"}`, err.Error())
		return
	}

	data, err = json.Marshal(response)
	if err != nil {
		fmt.Fprintf(w, `{"error_profile": "%s"}`, err.Error())
		return
	}

	w.Write(data)
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
	userSvc := userService.New(mysqlRepo, jwtTestKey)

	_, err = userSvc.Register(req)
	if err != nil {
		fmt.Fprintf(w, `{"error_register": "%s"}`, err.Error())
		return
	}

	fmt.Fprint(w, `{"message": "user registered"`)
}

func testUserCreationRepo() {
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
