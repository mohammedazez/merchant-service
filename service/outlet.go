package service

import (
	"errors"
	"fmt"
	"merchant-service/domain/dto"
	"merchant-service/storage"
	"time"
)

type OutletService interface {
	CreateOutlet(outlet dto.OutletInput, userID string) (dto.Outlet, error)
	ShowAllOutlet() ([]dto.Outlet, error)
	FindOutletByID(outletID string) (dto.Outlet, error)
}

type outletservice struct {
	dao storage.OutletDao
}

func NewOutletService(dao storage.OutletDao) *outletservice {
	return &outletservice{dao}
}

func (s *outletservice) CreateOutlet(outlet dto.OutletInput, userID string) (dto.Outlet, error) {

	checkStatus, err := s.dao.FindOutletByID(userID)

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

	createOutlet, err := s.dao.CreateOutlet(newOutlet)

	if err != nil {
		return createOutlet, err
	}

	return createOutlet, nil
}

func (s *outletservice) ShowAllOutlet() ([]dto.Outlet, error) {
	outlet, err := s.dao.ShowAllOutlet()

	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (s *outletservice) FindOutletByID(outletID string) (dto.Outlet, error) {
	outlet, err := s.dao.FindOutletByID(outletID)

	if err != nil {
		return outlet, err
	}

	if outlet.ID == 0 {
		newError := fmt.Sprintf("Outlet id %s not found", outletID)
		return outlet, errors.New(newError)
	}

	return outlet, nil
}
