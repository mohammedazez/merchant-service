package domain

type GetProductsResponse struct {
	Status  int64            `json:"status"`
	Meta    MetaDetail       `json:"meta"`
	Data    []GetAllProducts `json:"data"`
	Message string           `json:"message"`
}

type GetAllProducts struct {
	Id           string `gorm:"primaryKey" json:"id"`
	ProductName  string `json:"product_name"`
	Price        int64  `json:"price"`
	Sku          int64  `json:"sku"`
	DisplayImage string `json:"display_image"`
	OutletId     string `json:"outlet_id"`
}

type Product struct {
	Id           string `gorm:"primaryKey" json:"id"`
	ProductName  string `json:"product_name"`
	Price        int64  `json:"price"`
	Sku          int64  `json:"sku"`
	DisplayImage string `json:"display_image"`
	OutletId     string `json:"outlet_id"`
}

type InputProduct struct {
	ProductName  string `json:"product_name" binding:"required"`
	Price        int64  `json:"price" binding:"required"`
	Sku          int64  `json:"sku" binding:"required"`
	DisplayImage string `json:"display_image" binding:"required"`
	OutletId     string `json:"outlet_id" binding:"required"`
}

type UpdateProduct struct {
	ProductName  string `json:"product_name"`
	Price        int64  `json:"price"`
	Sku          int64  `json:"sku"`
	DisplayImage string `json:"display_image"`
	OutletId     string `json:"outlet_id"`
}
