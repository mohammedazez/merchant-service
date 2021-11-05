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
	UpdateProductByID(productID string, input entity.UpdateProductInput) (entity.Product, error)
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

func (s *service) UpdateProductByID(productID string, input entity.UpdateProductInput) (entity.Product, error) {
	var dataUpdate = map[string]interface{}{}

	product, err := s.repository.FindProductByID(productID)

	if err != nil {
		return entity.Product{}, err
	}

	if product.ID == 0 {
		newError := fmt.Sprintf("product id %s not found", productID)
		return entity.Product{}, errors.New(newError)
	}

	if input.ProductName != "" || len(input.ProductName) != 0 {
		dataUpdate["product_name"] = input.ProductName
	}
	if input.Price != 0 {
		dataUpdate["price"] = input.Price
	}
	if input.Sku != "" || len(input.Sku) != 0 {
		dataUpdate["sku"] = input.Sku
	}

	if input.Picture != "" || len(input.Picture) != 0 {
		dataUpdate["picture"] = input.Picture
	}
	if input.OutletID != "" || len(input.OutletID) != 0 {
		dataUpdate["OutletID"] = input.OutletID
	}
	dataUpdate["updated_at"] = time.Now()

	productUpdated, err := s.repository.UpdateProductByID(productID, dataUpdate)

	if err != nil {
		return productUpdated, err
	}

	return productUpdated, nil
}
