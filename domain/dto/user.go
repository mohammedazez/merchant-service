package dto

import (
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Outlet    []Outlet  `gorm:"ForeignKey:UserID"`
}

type UserInput struct {
	ID       string `json:"id"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
