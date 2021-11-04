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
