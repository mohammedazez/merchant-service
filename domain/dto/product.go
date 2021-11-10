package dto

import (
	"time"
)

type Product struct {
	ID           string         `gorm:"PrimaryKey" json:"id"`
	ProductName  string         `json:"product_name"`
	Price        int            `json:"price"`
	Sku          string         `json:"sku"`
	Picture      string         `json:"picture"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	OutletID     string         `json:"outlet_id"`
	ImageProduct []ImageProduct `json:"imageproduct"`
}

type ImageProduct struct {
	ID           string `json:"id"`
	DisplayImage string `json:"display_image"`
	ProductID    string `json:"product_id"`
}

type ProductInput struct {
	ProductName string `json:"product_name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Sku         string `json:"sku" binding:"required"`
	Picture     string `json:"picture" binding:"required"`
	OutletID    string `json:"outlet_id" binding:"required"`
}

type UpdateProductInput struct {
	ProductName string    `json:"product_name"`
	Price       int64     `json:"price" `
	Sku         string    `json:"sku"`
	Picture     string    `json:"picture"`
	OutletID    string    `json:"outlet_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}
