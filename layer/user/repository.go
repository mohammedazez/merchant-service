package user

import (
	"merchant-service/entity"

	"gorm.io/gorm"
)

type Repository interface {
	RegisterUser(user entity.User) (entity.User, error)
	FindUserByEmail(email string) (entity.User, error)
	ShowAllUser() ([]entity.User, error)
	FindUserByID(ID string) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RegisterUser(userUser entity.User) (entity.User, error) {

	err := r.db.Create(&userUser).Error
	if err != nil {
		return userUser, err
	}

	return userUser, nil
}

func (r *repository) FindUserByEmail(email string) (entity.User, error) {
	var userUser entity.User

	err := r.db.Where("email = ?", email).Find(&userUser).Error
	if err != nil {
		return userUser, err
	}

	return userUser, nil
}

func (r *repository) ShowAllUser() ([]entity.User, error) {
	var userUser []entity.User

	err := r.db.Preload("Outlet").Find(&userUser).Error
	if err != nil {
		return userUser, err
	}

	return userUser, nil
}

func (r *repository) FindUserByID(ID string) (entity.User, error) {
	var userUser entity.User

	err := r.db.Where("id = ?", ID).Preload("Outlet").Find(&userUser).Error
	if err != nil {
		return userUser, err
	}

	return userUser, nil
}
