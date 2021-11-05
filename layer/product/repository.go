package product

import (
	"merchant-service/entity"

	"gorm.io/gorm"
)

type Repository interface {
	CreateProduct(product entity.Product) (entity.Product, error)
	ShowAllProduct() ([]entity.Product, error)
	FindProductByID(ID string) (entity.Product, error)
	UpdateProductByID(ID string, dataUpdate map[string]interface{}) (entity.Product, error)
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

func (r *repository) ShowAllProduct() ([]entity.Product, error) {
	var product []entity.Product

	err := r.db.Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindProductByID(ID string) (entity.Product, error) {
	var product entity.Product

	err := r.db.Where("id = ?", ID).Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) UpdateProductByID(ID string, dataUpdate map[string]interface{}) (entity.Product, error) {

	var product entity.Product

	if err := r.db.Model(&product).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return product, err
	}

	if err := r.db.Where("id = ?", ID).Find(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}
