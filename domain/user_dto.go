package domain

type MetaDetail struct {
	TotalPage   int64 `json:"total_page"`
	CurrentPage int64 `json:"current_page"`
	PageSize    int64 `json:"page_size"`
	Total       int64 `json:"total"`
}

type GetUsersResponse struct {
	Status  int64         `json:"status"`
	Meta    MetaDetail    `json:"meta"`
	Data    []GetAllUsers `json:"data"`
	Message string        `json:"message"`
}

type GetAllUsers struct {
	Id       string `gorm:"primaryKey" json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type User struct {
	Id       string   `gorm:"primaryKey" json:"id"`
	FullName string   `json:"full_name"`
	Email    string   `json:"email"`
	Outlets  []Outlet `gorm:"ForeignKey:UserId"`
}

type InputUser struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUser struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
