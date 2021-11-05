package product

import (
	"errors"
	"fmt"
	"merchant-service/entity"
	"time"
)

type Service interface {
	CreateProduct(product entity.ProductInput) (entity.Product, error)
	ShowAllProduct() ([]entity.Product, error)
	FindProductByID(productID string) (entity.Product, error)
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

func (s *service) ShowAllProduct() ([]entity.Product, error) {
	product, err := s.repository.ShowAllProduct()

	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) FindProductByID(productID string) (entity.Product, error) {
	product, err := s.repository.FindProductByID(productID)

	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		newError := fmt.Sprintf("product id %s not found", productID)
		return product, errors.New(newError)
	}

	return product, nil
}
