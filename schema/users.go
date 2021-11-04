package schema

type Users struct {
	BaseLegacy
	FullName string `gorm:"null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}

func (Users) TableName() string {
	return "users"
}
