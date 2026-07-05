package userService

import (
	"fmt"
	"quizGameGo/entities"
	"quizGameGo/pkg/phoneNumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entities.User) (entities.User, error)
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterResponse struct {
	User entities.User
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// TODO - phone number should be verified via sms
	// validate phone number
	if !phoneNumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("invalid phone number")
	}

	// checking phone number uniqueness
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error %w", err)
		}
		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}

	// validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name is too short")
	}

	// create new user
	user := entities.User{
		ID:          0, //db handles id management
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
	}
	createdUser, err := s.repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	// return created user
	return RegisterResponse{
		User: createdUser,
	}, nil
}
