package storage

import (
	"merchant-service/domain/dto"
	"merchant-service/storage/query"

	"gorm.io/gorm"
)

type UserDao interface {
	RegisterUser(user dto.User) (dto.User, error)
	FindUserByEmail(email string) (dto.User, error)
	ShowAllUser() ([]dto.User, error)
	FindUserByID(ID string) (dto.User, error)
	UpdateUserByID(ID string, input dto.UpdateUserInput) (dto.User, error)
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

func (r *dao) RegisterUser(user dto.User) (dto.User, error) {
	qry := query.RegisterUser

	err := r.db.Raw(qry,
		user.ID,
		user.FullName,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt).Scan(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *dao) FindUserByEmail(email string) (dto.User, error) {
	var user dto.User
	qry := query.LoginUser

	err := r.db.Raw(qry, email).Scan(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *dao) ShowAllUser() ([]dto.User, error) {
	var user []dto.User
	qry := query.GetAllUsers

	err := r.db.Raw(qry).Scan(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *dao) FindUserByID(ID string) (dto.User, error) {
	var user dto.User
	err := r.db.Where("id = ?", ID).Preload("Outlet").Find(&user).Error

	// qry := query.FindUserById

	// err := r.db.Raw(qry, ID).Scan(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *dao) UpdateUserByID(ID string, input dto.UpdateUserInput) (dto.User, error) {

	var user dto.User

	// if err := r.db.Model(&userUser).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
	// 	return userUser, err
	// }

	// if err := r.db.Where("id = ?", ID).Find(&userUser).Error; err != nil {
	// 	return userUser, err
	// }

	qry := query.UpdateUserByID
	err := r.db.Raw(qry, input.FullName, input.Email, input.Password, ID).Scan(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *dao) DeleteUserByID(ID string) (string, error) {
	// err := r.db.Where("id = ?", ID).Delete(&dto.User{}).Error;
	user := &dto.User{}
	qry := query.DeleteUserById

	err := r.db.Raw(qry, ID).Scan(&user).Error

	if err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *dao) CreateOutletUser(outlet dto.Outlet) (dto.Outlet, error) {

	// err := r.db.Create(&outlet).Error
	qry := query.CreateOutletbyUser

	err := r.db.Raw(qry,
		outlet.ID,
		outlet.OutletName,
		outlet.Picture,
		outlet.UserID,
		outlet.CreatedAt,
		outlet.UpdatedAt).Scan(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *dao) FindOutletUserByID(ID string) (dto.Outlet, error) {
	var outlet dto.Outlet

	// err := r.db.Where("id = ?", ID).Preload("Product").Find(&outlet).Error

	qry := query.FindOutletUserByID

	err := r.db.Raw(qry, ID).Scan(&outlet).Error

	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *dao) ShowAllOutletUser() ([]dto.Outlet, error) {
	var outlet []dto.Outlet

	// err := r.db.Preload("Product").Find(&outlet).Error
	qry := query.GetAllOutlets

	err := r.db.Raw(qry).Scan(&outlet).Error

	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
