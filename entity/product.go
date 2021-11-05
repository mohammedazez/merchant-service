package entity

import "time"

type Product struct {
	ID          int       `gorm:"PrimaryKey" json:"id"`
	ProductName string    `json:"product_name"`
	Price       int       `json:"price"`
	Sku         string    `json:"sku"`
	Picture     string    `json:"picture"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	OutletID    int       `json:"outlet_id"`
}

type ProductInput struct {
	ProductName string `json:"product_name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Sku         string `json:"sku"`
	Picture     string `json:"picture"`
	OutletID    int    `json:"outlet_id" binding:"required"`
}

type UpdateProductInput struct {
	ProductName string `json:"product_name"`
	Price       int64  `json:"price" `
	Sku         string `json:"sku"`
	Picture     string `json:"picture"`
	OutletID    int    `json:"outlet_id"`
}
