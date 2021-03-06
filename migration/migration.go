package migration

import "time"

type User struct {
	ID        string `gorm:"PrimaryKey"`
	FullName  string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Outlet    []Outlet `gorm:"ForeignKey:UserID"`
}

type Outlet struct {
	ID         string `gorm:"PrimaryKey"`
	OutletName string
	Picture    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Product    Product `gorm:"ForeignKey:OutletID"`
	UserID     string  `gorm:"index"`
}

type Product struct {
	ID           string `gorm:"PrimaryKey"`
	ProductName  string
	Price        int64
	Sku          string
	Picture      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	OutletID     string       `gorm:"index"`
	DisplayImage ImageProduct `gorm:"ForeignKey:ProductID"`
}

type ImageProduct struct {
	ID           string
	DisplayImage string
	ProductID    string `gorm:"index"`
}
