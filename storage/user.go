package storage

import (
	"merchant-service/domain/dto"

	"gorm.io/gorm"
)

type UserDao interface {
	RegisterUser(user dto.User) (dto.User, error)
	FindUserByEmail(email string) (dto.User, error)
	ShowAllUser() ([]dto.User, error)
	FindUserByID(ID string) (dto.User, error)
	UpdateUserByID(ID string, dataUpdate map[string]interface{}) (dto.User, error)
	DeleteUserByID(ID string) (string, error)
	CreateOutletUser(Outlet dto.Outlet) (dto.Outlet, error)
	FindOutletUserByID(ID string) (dto.Outlet, error)
	ShowAllOutletUser() ([]dto.Outlet, error)
}

type dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *dao {
	return &dao{db}
}

func (r *dao) RegisterUser(userUser dto.User) (dto.User, error) {

	err := r.db.Create(&userUser).Error
	if err != nil {
		return userUser, err
	}

	return userUser, nil
}

func (r *dao) FindUserByEmail(email string) (dto.User, error) {
	var userUser dto.User

	err := r.db.Where("email = ?", email).Find(&userUser).Error
	if err != nil {
		return userUser, err
	}

	return userUser, nil
}

func (r *dao) ShowAllUser() ([]dto.User, error) {
	var userUser []dto.User

	err := r.db.Preload("Outlet").Find(&userUser).Error
	if err != nil {
		return userUser, err
	}

	return userUser, nil
}

func (r *dao) FindUserByID(ID string) (dto.User, error) {
	var userUser dto.User

	err := r.db.Where("id = ?", ID).Preload("Outlet").Find(&userUser).Error
	if err != nil {
		return userUser, err
	}

	return userUser, nil
}

func (r *dao) UpdateUserByID(ID string, dataUpdate map[string]interface{}) (dto.User, error) {

	var userUser dto.User

	if err := r.db.Model(&userUser).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return userUser, err
	}

	if err := r.db.Where("id = ?", ID).Find(&userUser).Error; err != nil {
		return userUser, err
	}

	return userUser, nil
}

func (r *dao) DeleteUserByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&dto.User{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *dao) CreateOutletUser(outlet dto.Outlet) (dto.Outlet, error) {

	err := r.db.Create(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *dao) FindOutletUserByID(ID string) (dto.Outlet, error) {
	var outlet dto.Outlet

	err := r.db.Where("id = ?", ID).Preload("Product").Find(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *dao) ShowAllOutletUser() ([]dto.Outlet, error) {
	var outlet []dto.Outlet

	err := r.db.Preload("Product").Find(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
