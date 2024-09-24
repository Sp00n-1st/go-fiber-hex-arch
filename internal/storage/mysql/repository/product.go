package repository

import (
	"go-fiber-hex-arch/internal/core/domain"
	"gorm.io/gorm"
)

type ProductRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) *ProductRepositoryDB {
	return &ProductRepositoryDB{db: db}
}

func (r *ProductRepositoryDB) Insert(product domain.Product) error {
	tx := r.db.Begin()
	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *ProductRepositoryDB) FindByID(id uint) (domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r *ProductRepositoryDB) Update(product domain.Product) error {
	tx := r.db.Begin()
	if err := tx.Model(&domain.Product{}).Where("product_id = ?", product.ProductID).Updates(product).Error; err != nil {
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryDB) Delete(id uint) error {
	if err := r.db.Delete(&domain.Product{}, id).Error; err != nil {
		return err
	}
	if err := r.db.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryDB) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Find(&products).Error
	return products, err
}
