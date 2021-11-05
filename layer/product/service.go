package product

import (
	"merchant-service/entity"
	"time"
)

type Service interface {
	CreateProduct(product entity.ProductInput) (entity.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateProduct(product entity.ProductInput) (entity.Product, error) {

	var newProduct = entity.Product{
		ProductName: product.ProductName,
		Price:       product.Price,
		Sku:         product.Sku,
		Picture:     product.Picture,
		OutletID:    product.OutletID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createProduct, err := s.repository.CreateProduct(newProduct)

	if err != nil {
		return createProduct, err
	}

	return createProduct, nil
}
