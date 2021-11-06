package storage

import (
	"merchant-service/domain/dto"
	"merchant-service/storage/query"
	"time"

	"golang.org/x/crypto/bcrypt"
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

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *dao) UpdateUserByID(ID string, input dto.UpdateUserInput) (dto.User, error) {

	var user dto.User
	genPassword, err2 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err2 != nil {
		return user, err2
	}

	input.UpdatedAt = time.Now()

	qry := query.UpdateUserByID
	err := r.db.Raw(qry, input.FullName, input.Email, genPassword, input.UpdatedAt, ID).Scan(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *dao) DeleteUserByID(ID string) (string, error) {
	user := &dto.User{}
	qry := query.DeleteUserById

	err := r.db.Raw(qry, ID).Scan(&user).Error

	if err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *dao) CreateOutletUser(outlet dto.Outlet) (dto.Outlet, error) {
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

	qry := query.FindOutletUserByID

	err := r.db.Raw(qry, ID).Scan(&outlet).Error

	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *dao) ShowAllOutletUser() ([]dto.Outlet, error) {
	var outlet []dto.Outlet

	qry := query.GetAllOutlets

	err := r.db.Raw(qry).Scan(&outlet).Error

	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
