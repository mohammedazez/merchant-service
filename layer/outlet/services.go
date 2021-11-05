package outlet

import (
	"errors"
	"fmt"
	"merchant-service/entity"
	"time"
)

type Service interface {
	CreateOutlet(outlet entity.OutletInput, userID string) (entity.Outlet, error)
	ShowAllOutlet() ([]entity.Outlet, error)
	FindOutletByID(outletID string) (entity.Outlet, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateOutlet(outlet entity.OutletInput, userID string) (entity.Outlet, error) {

	checkStatus, err := s.repository.FindOutletByID(userID)

	if err != nil {
		return checkStatus, err
	}

	if checkStatus.UserID == userID {
		errorStatus := fmt.Sprintf("Outlet for user id : %s has been created", userID)
		return checkStatus, errors.New(errorStatus)
	}

	var newOutlet = entity.Outlet{
		OutletName: outlet.OutletName,
		Picture:    outlet.Picture,
		UserID:     userID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	createOutlet, err := s.repository.CreateOutlet(newOutlet)

	if err != nil {
		return createOutlet, err
	}

	return createOutlet, nil
}

func (s *service) ShowAllOutlet() ([]entity.Outlet, error) {
	outlet, err := s.repository.ShowAllOutlet()

	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (s *service) FindOutletByID(outletID string) (entity.Outlet, error) {
	outlet, err := s.repository.FindOutletByID(outletID)

	if err != nil {
		return outlet, err
	}

	if outlet.ID == 0 {
		newError := fmt.Sprintf("Outlet id %s not found", outletID)
		return outlet, errors.New(newError)
	}

	return outlet, nil
}
