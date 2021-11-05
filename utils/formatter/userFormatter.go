package formatter

import (
	"merchant-service/domain/dto"
	"time"
)

type UserFormat struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Outlet   []dto.Outlet
}

type UserDeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func Format(user dto.User) UserFormat {
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
