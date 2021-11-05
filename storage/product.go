package storage

import (
	"merchant-service/domain/dto"

	"gorm.io/gorm"
)

type ProductDao interface {
	CreateProduct(product dto.Product) (dto.Product, error)
	ShowAllProduct() ([]dto.Product, error)
	FindProductByID(ID string) (dto.Product, error)
	UpdateProductByID(ID string, dataUpdate map[string]interface{}) (dto.Product, error)
	DeleteProductByID(ID string) (string, error)
}

func NewProductDao(db *gorm.DB) *dao {
	return &dao{db}
}

func (r *dao) CreateProduct(product dto.Product) (dto.Product, error) {

	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *dao) ShowAllProduct() ([]dto.Product, error) {
	var product []dto.Product

	err := r.db.Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *dao) FindProductByID(ID string) (dto.Product, error) {
	var product dto.Product

	err := r.db.Where("id = ?", ID).Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *dao) UpdateProductByID(ID string, dataUpdate map[string]interface{}) (dto.Product, error) {

	var product dto.Product

	if err := r.db.Model(&product).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return product, err
	}

	if err := r.db.Where("id = ?", ID).Find(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *dao) DeleteProductByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&dto.Product{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
