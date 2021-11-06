package service

import (
	"errors"
	"fmt"
	"merchant-service/domain/dto"
	"merchant-service/storage"
	"merchant-service/utils/formatter"
	"time"

	"github.com/gofrs/uuid"
)

type ProductService interface {
	CreateProduct(product dto.ProductInput) (formatter.ProductFormat, error)
	ShowAllProduct() ([]formatter.ProductFormat, error)
	FindProductByID(productID string) (formatter.ProductFormat, error)
	UpdateProductByID(productID string, input dto.UpdateProductInput) (formatter.ProductFormat, error)
	DeleteProductByID(productID string) (interface{}, error)
	FindOutletUserByID(outletID string) (dto.Outlet, error)
}

type productservice struct {
	dao storage.ProductDao
}

func NewProductService(dao storage.ProductDao) *productservice {
	return &productservice{dao}
}

func (s *productservice) CreateProduct(product dto.ProductInput) (formatter.ProductFormat, error) {

	productuuid, err := uuid.NewV4()

	if err != nil {
		return formatter.ProductFormat{}, err
	}

	var newProduct = dto.Product{
		ID:          productuuid.String(),
		ProductName: product.ProductName,
		Price:       product.Price,
		Sku:         product.Sku,
		Picture:     product.Picture,
		OutletID:    product.OutletID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createProduct, err := s.dao.CreateProduct(newProduct)

	formatProduct := formatter.FormatProduct(createProduct)

	if err != nil {
		return formatProduct, err
	}

	return formatProduct, nil
}

func (s *productservice) ShowAllProduct() ([]formatter.ProductFormat, error) {
	product, err := s.dao.ShowAllProduct()

	var formatuserProduct []formatter.ProductFormat

	for _, products := range product {
		formatProduct := formatter.FormatProduct(products)
		formatuserProduct = append(formatuserProduct, formatProduct)

	}
	if err != nil {
		return formatuserProduct, err
	}

	return formatuserProduct, nil
}

func (s *productservice) FindProductByID(productID string) (formatter.ProductFormat, error) {
	product, err := s.dao.FindProductByID(productID)

	if err != nil {
		return formatter.ProductFormat{}, err
	}

	if len(product.ID) == 0 {
		newError := fmt.Sprintf("product id %s not found", productID)
		return formatter.ProductFormat{}, errors.New(newError)
	}

	formatProduct := formatter.FormatProduct(product)

	return formatProduct, nil
}

func (s *productservice) UpdateProductByID(productID string, input dto.UpdateProductInput) (formatter.ProductFormat, error) {
	var dataUpdate = map[string]interface{}{}

	product, err := s.dao.FindProductByID(productID)

	if err != nil {
		return formatter.ProductFormat{}, err
	}

	if len(product.ID) == 0 {
		newError := fmt.Sprintf("product id %s not found", productID)
		return formatter.ProductFormat{}, errors.New(newError)
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
		return formatter.ProductFormat{}, err
	}

	formatProduct := formatter.FormatProduct(productUpdated)

	return formatProduct, nil
}

func (s *productservice) DeleteProductByID(productID string) (interface{}, error) {

	product, err := s.dao.FindProductByID(productID)

	if err != nil {
		return nil, err
	}

	if len(product.ID) == 0 {
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

	formatDelete := formatter.FormatDeleteProduct(msg)

	return formatDelete, nil
}

func (s *productservice) FindOutletUserByID(outletID string) (dto.Outlet, error) {
	outlet, err := s.dao.FindOutletProductByID(outletID)

	if err != nil {
		return outlet, err
	}

	if len(outlet.ID) == 0 {
		newError := fmt.Sprintf("Outlet id %s not found", outletID)
		return outlet, errors.New(newError)
	}

	return outlet, nil
}
