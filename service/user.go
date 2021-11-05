package service

import (
	"errors"
	"fmt"
	"merchant-service/domain/dto"
	"merchant-service/storage"
	"merchant-service/utils/formatter"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(userUser dto.UserInput) (formatter.UserFormat, error)
	LoginUser(input dto.LoginUserInput) (formatter.UserFormat, error)
	ShowAllUser() ([]formatter.UserFormat, error)
	FindUserByID(userID string) (formatter.UserFormat, error)
	UpdateUserByID(userID string, input dto.UpdateUserInput) (formatter.UserFormat, error)
	DeleteUserByID(userID string) (interface{}, error)
	CreateOutletUser(outlet dto.OutletInput, userID string) (dto.Outlet, error)
	ShowAllOutletUser() ([]dto.Outlet, error)
}

type userservice struct {
	dao storage.UserDao
}

func NewUserService(dao storage.UserDao) *userservice {
	return &userservice{dao}
}

func (s *userservice) RegisterUser(userUser dto.UserInput) (formatter.UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(userUser.Password), bcrypt.MinCost)

	if err != nil {
		return formatter.UserFormat{}, err
	}

	useruuid, err := uuid.NewV4()

	if err != nil {
		return formatter.UserFormat{}, err
	}

	var newUser = dto.User{
		ID:        useruuid.String(),
		FullName:  userUser.FullName,
		Email:     userUser.Email,
		Password:  string(genPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := s.dao.RegisterUser(newUser)
	formatUser := formatter.Format(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *userservice) LoginUser(input dto.LoginUserInput) (formatter.UserFormat, error) {
	userUser, err := s.dao.FindUserByEmail(input.Email)

	if err != nil {
		return formatter.UserFormat{}, err
	}

	if len(userUser.ID) != 0 {
		if err := bcrypt.CompareHashAndPassword([]byte(userUser.Password), []byte(input.Password)); err != nil {
			return formatter.UserFormat{}, errors.New("password invalid")
		}

		formatter := formatter.Format(userUser)

		return formatter, nil
	}

	newError := fmt.Sprintf("user id %v not found", userUser.ID)
	return formatter.UserFormat{}, errors.New(newError)
}

func (s *userservice) ShowAllUser() ([]formatter.UserFormat, error) {
	userUser, err := s.dao.ShowAllUser()
	var formatuserUser []formatter.UserFormat

	for _, users := range userUser {
		formatUser := formatter.Format(users)
		formatuserUser = append(formatuserUser, formatUser)

	}
	if err != nil {
		return formatuserUser, err
	}

	return formatuserUser, nil
}

func (s *userservice) FindUserByID(userID string) (formatter.UserFormat, error) {
	userUser, err := s.dao.FindUserByID(userID)

	if err != nil {
		return formatter.UserFormat{}, err
	}

	if len(userUser.ID) == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return formatter.UserFormat{}, errors.New(newError)
	}

	formatUser := formatter.Format(userUser)

	return formatUser, nil
}

func (s *userservice) UpdateUserByID(userID string, input dto.UpdateUserInput) (formatter.UserFormat, error) {
	var dataUpdate = map[string]interface{}{}

	userUser, err := s.dao.FindUserByID(userID)

	if err != nil {
		return formatter.UserFormat{}, err
	}

	if len(userUser.ID) == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return formatter.UserFormat{}, errors.New(newError)
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

	userUpdated, err := s.dao.UpdateUserByID(userID, dataUpdate)

	if err != nil {
		return formatter.UserFormat{}, err
	}

	formatUser := formatter.Format(userUpdated)

	return formatUser, nil
}

func (s *userservice) DeleteUserByID(userID string) (interface{}, error) {
	userUser, err := s.dao.FindUserByID(userID)

	if err != nil {
		return nil, err
	}

	if len(userUser.ID) == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return nil, errors.New(newError)
	}

	status, err := s.dao.DeleteUserByID(userID)

	if err != nil {
		return nil, err
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete user ID : %s", userID)

	formatDelete := formatter.FormatDelete(msg)

	return formatDelete, nil
}

func (s *userservice) CreateOutletUser(outlet dto.OutletInput, userID string) (dto.Outlet, error) {

	checkStatus, err := s.dao.FindOutletUserByID(userID)

	if err != nil {
		return checkStatus, err
	}

	if checkStatus.UserID == userID {
		errorStatus := fmt.Sprintf("Outlet for user id : %s has been created", userID)
		return checkStatus, errors.New(errorStatus)
	}

	var newOutlet = dto.Outlet{
		OutletName: outlet.OutletName,
		Picture:    outlet.Picture,
		UserID:     userID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	createOutlet, err := s.dao.CreateOutletUser(newOutlet)

	if err != nil {
		return createOutlet, err
	}

	return createOutlet, nil
}

func (s *userservice) ShowAllOutletUser() ([]dto.Outlet, error) {
	outlet, err := s.dao.ShowAllOutletUser()

	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
