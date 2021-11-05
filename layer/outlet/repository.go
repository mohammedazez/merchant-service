package outlet

import (
	"merchant-service/entity"

	"gorm.io/gorm"
)

type Repository interface {
	CreateOutlet(Outlet entity.Outlet) (entity.Outlet, error)
	ShowAllOutlet() ([]entity.Outlet, error)
	FindOutletByID(ID string) (entity.Outlet, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateOutlet(outlet entity.Outlet) (entity.Outlet, error) {

	err := r.db.Create(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *repository) ShowAllOutlet() ([]entity.Outlet, error) {
	var outlet []entity.Outlet

	err := r.db.Preload("Product").Find(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *repository) FindOutletByID(ID string) (entity.Outlet, error) {
	var outlet entity.Outlet

	err := r.db.Where("id = ?", ID).Preload("Product").Find(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
