package storage

import (
	"merchant-service/domain/dto"

	"gorm.io/gorm"
)

type OutletDao interface {
	CreateOutlet(Outlet dto.Outlet) (dto.Outlet, error)
	ShowAllOutlet() ([]dto.Outlet, error)
	FindOutletByID(ID string) (dto.Outlet, error)
}

func NewDaoOutlet(db *gorm.DB) *dao {
	return &dao{db}
}

func (r *dao) CreateOutlet(outlet dto.Outlet) (dto.Outlet, error) {

	err := r.db.Create(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *dao) ShowAllOutlet() ([]dto.Outlet, error) {
	var outlet []dto.Outlet

	err := r.db.Preload("Product").Find(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *dao) FindOutletByID(ID string) (dto.Outlet, error) {
	var outlet dto.Outlet

	err := r.db.Where("id = ?", ID).Preload("Product").Find(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
