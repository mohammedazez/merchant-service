package product

import (
	"merchant-service/entity"

	"gorm.io/gorm"
)

type Repository interface {
	CreateProduct(product entity.Product) (entity.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateProduct(product entity.Product) (entity.Product, error) {

	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
