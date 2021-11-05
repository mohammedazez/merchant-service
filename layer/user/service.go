package user

import (
	"errors"
	"fmt"
	"merchant-service/entity"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(userUser entity.UserInput) (UserFormat, error)
	LoginUser(input entity.LoginUserInput) (UserFormat, error)
	ShowAllUser() ([]UserFormat, error)
	FindUserByID(userID string) (UserFormat, error)
	UpdateUserByID(userID string, input entity.UpdateUserInput) (UserFormat, error)
	DeleteUserByID(userID string) (interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(userUser entity.UserInput) (UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(userUser.Password), bcrypt.MinCost)

	if err != nil {
		return UserFormat{}, err
	}

	useruuid := uuid.NewV4()

	var newUser = entity.User{
		ID:        useruuid.String(),
		FullName:  userUser.FullName,
		Email:     userUser.Email,
		Password:  string(genPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := s.repository.RegisterUser(newUser)
	formatUser := Format(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *service) LoginUser(input entity.LoginUserInput) (UserFormat, error) {
	userUser, err := s.repository.FindUserByEmail(input.Email)

	if err != nil {
		return UserFormat{}, err
	}

	if len(userUser.ID) != 0 {
		if err := bcrypt.CompareHashAndPassword([]byte(userUser.Password), []byte(input.Password)); err != nil {
			return UserFormat{}, errors.New("password invalid")
		}

		formatter := Format(userUser)

		return formatter, nil
	}

	newError := fmt.Sprintf("user id %v not found", userUser.ID)
	return UserFormat{}, errors.New(newError)
}

func (s *service) ShowAllUser() ([]UserFormat, error) {
	userUser, err := s.repository.ShowAllUser()
	var formatuserUser []UserFormat

	for _, user := range userUser {
		formatUser := Format(user)
		formatuserUser = append(formatuserUser, formatUser)

	}
	if err != nil {
		return formatuserUser, err
	}

	return formatuserUser, nil
}

func (s *service) FindUserByID(userID string) (UserFormat, error) {
	userUser, err := s.repository.FindUserByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if len(userUser.ID) == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return UserFormat{}, errors.New(newError)
	}

	formatUser := Format(userUser)

	return formatUser, nil
}

func (s *service) UpdateUserByID(userID string, input entity.UpdateUserInput) (UserFormat, error) {
	var dataUpdate = map[string]interface{}{}

	userUser, err := s.repository.FindUserByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if len(userUser.ID) == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return UserFormat{}, errors.New(newError)
	}

	if input.FullName != "" || len(input.FullName) != 0 {
		dataUpdate["FullName"] = input.FullName
	}
	if input.Email != "" || len(input.Email) != 0 {
		dataUpdate["Email"] = input.Email
	}
	if input.Password != "" || len(input.Password) != 0 {
		dataUpdate["Password"] = input.Password
	}
	dataUpdate["updated_at"] = time.Now()

	userUpdated, err := s.repository.UpdateUserByID(userID, dataUpdate)

	if err != nil {
		return UserFormat{}, err
	}

	formatUser := Format(userUpdated)

	return formatUser, nil
}

func (s *service) DeleteUserByID(userID string) (interface{}, error) {
	userUser, err := s.repository.FindUserByID(userID)

	if err != nil {
		return nil, err
	}

	if len(userUser.ID) == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteUserByID(userID)

	if err != nil {
		return nil, err
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete user ID : %s", userID)

	formatDelete := FormatDelete(msg)

	return formatDelete, nil
}
