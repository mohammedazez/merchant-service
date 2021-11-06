package storage

import (
	"merchant-service/domain/dto"
	"merchant-service/storage/query"
	"time"

	"gorm.io/gorm"
)

type ProductDao interface {
	CreateProduct(product dto.Product) (dto.Product, error)
	ShowAllProduct() ([]dto.Product, error)
	FindProductByID(ID string) (dto.Product, error)
	UpdateProductByID(ID string, input dto.UpdateProductInput) (dto.Product, error)
	DeleteProductByID(ID string) (string, error)
	FindOutletProductByID(ID string) (dto.Outlet, error)
}

func NewProductDao(db *gorm.DB) *dao {
	return &dao{db}
}

func (r *dao) CreateProduct(product dto.Product) (dto.Product, error) {

	qry := query.QueryCreateProduct

	err := r.db.Raw(qry,
		product.ID,
		product.ProductName,
		product.Price,
		product.Sku,
		product.Picture,
		product.CreatedAt,
		product.UpdatedAt,
		product.OutletID).Scan(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *dao) ShowAllProduct() ([]dto.Product, error) {
	var product []dto.Product

	qry := query.QueryFindAllProduct

	err := r.db.Raw(qry).Scan(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *dao) FindProductByID(ID string) (dto.Product, error) {
	var product dto.Product

	qry := query.QueryFindProductById

	err := r.db.Raw(qry, ID).Scan(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *dao) UpdateProductByID(ID string, input dto.UpdateProductInput) (dto.Product, error) {

	var product dto.Product

	input.UpdatedAt = time.Now()

	qry := query.QueryUpdateProductByID
	err := r.db.Raw(qry,
		input.ProductName,
		input.Price,
		input.Sku,
		input.Picture,
		input.OutletID,
		input.UpdatedAt,
		ID).Scan(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *dao) DeleteProductByID(ID string) (string, error) {
	product := &dto.Product{}
	qry := query.QueryDeleteProductById

	err := r.db.Raw(qry, ID).Scan(&product).Error
	if err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *dao) FindOutletProductByID(ID string) (dto.Outlet, error) {
	var outlet dto.Outlet

	err := r.db.Where("id = ?", ID).Preload("Product").Find(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
