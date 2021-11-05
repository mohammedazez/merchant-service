package user

import (
	"merchant-service/entity"
	"time"
)

type UserFormat struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Outlet   []entity.Outlet
}

type UserDeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func Format(user entity.User) UserFormat {
	var formatUser = UserFormat{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Outlet:   user.Outlet,
	}

	return formatUser
}

func FormatDelete(msg string) UserDeleteFormat {
	var deleteFormat = UserDeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
