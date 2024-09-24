package port

import (
	"go-fiber-hex-arch/internal/core/domain"
)

type ProductRepository interface {
	Insert(product domain.Product) error
	FindByID(id uint) (domain.Product, error)
	Update(product domain.Product) error
	Delete(id uint) error
	FindAll() ([]domain.Product, error)
}
