package service

import (
	"errors"
	"fmt"
	"merchant-service/domain/dto"
	"merchant-service/storage"
	"merchant-service/utils/formatter"
	"time"
)

type ProductService interface {
	CreateProduct(product dto.ProductInput) (dto.Product, error)
	ShowAllProduct() ([]dto.Product, error)
	FindProductByID(productID string) (dto.Product, error)
	UpdateProductByID(productID string, input dto.UpdateProductInput) (dto.Product, error)
	DeleteProductByID(productID string) (interface{}, error)
}

type productservice struct {
	dao storage.ProductDao
}

func NewProductService(dao storage.ProductDao) *productservice {
	return &productservice{dao}
}

func (s *productservice) CreateProduct(product dto.ProductInput) (dto.Product, error) {

	var newProduct = dto.Product{
		ProductName: product.ProductName,
		Price:       product.Price,
		Sku:         product.Sku,
		Picture:     product.Picture,
		OutletID:    product.OutletID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createProduct, err := s.dao.CreateProduct(newProduct)

	if err != nil {
		return createProduct, err
	}

	return createProduct, nil
}

func (s *productservice) ShowAllProduct() ([]dto.Product, error) {
	product, err := s.dao.ShowAllProduct()

	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *productservice) FindProductByID(productID string) (dto.Product, error) {
	product, err := s.dao.FindProductByID(productID)

	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		newError := fmt.Sprintf("product id %s not found", productID)
		return product, errors.New(newError)
	}

	return product, nil
}

func (s *productservice) UpdateProductByID(productID string, input dto.UpdateProductInput) (dto.Product, error) {
	var dataUpdate = map[string]interface{}{}

	product, err := s.dao.FindProductByID(productID)

	if err != nil {
		return dto.Product{}, err
	}

	if product.ID == 0 {
		newError := fmt.Sprintf("product id %s not found", productID)
		return dto.Product{}, errors.New(newError)
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

	productUpdated, err := s.dao.UpdateProductByID(productID, dataUpdate)

	if err != nil {
		return productUpdated, err
	}

	return productUpdated, nil
}

func (s *productservice) DeleteProductByID(productID string) (interface{}, error) {

	product, err := s.dao.FindProductByID(productID)

	if err != nil {
		return nil, err
	}

	if product.ID == 0 {
		newError := fmt.Sprintf("Product id %s not found", productID)
		return nil, errors.New(newError)
	}

	status, err := s.dao.DeleteProductByID(productID)

	if err != nil {
		return nil, err
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete Product ID : %s", productID)

	formatDelete := formatter.FormatDelete(msg)

	return formatDelete, nil
}
