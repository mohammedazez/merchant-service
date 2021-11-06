package formatter

import (
	"merchant-service/domain/dto"
	"time"
)

type UserFormat struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type UserDeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatUser(user dto.User) UserFormat {
	var formatUser = UserFormat{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
	}

	return formatUser
}

func FormatDeleteUser(msg string) UserDeleteFormat {
	var deleteFormat = UserDeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
