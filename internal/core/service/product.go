package service

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"go-fiber-hex-arch/internal/core/domain"
	"go-fiber-hex-arch/internal/core/port"
)

type ProductService struct {
	Repository port.ProductRepository
	Validator  *validator.Validate
}

func NewProductService(repository port.ProductRepository) *ProductService {
	return &ProductService{
		Repository: repository,
		Validator:  validator.New(),
	}
}

func (s *ProductService) InsertProduct(product domain.Product) error {
	if err := s.Validator.Struct(product); err != nil {
		return err
	}
	err := s.Repository.Insert(product)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return errors.New(fmt.Sprintf("%s already exists", product.ProductName))
	} else {
		return err
	}
}

func (s *ProductService) UpdateProduct(product domain.Product) error {
	if err := s.Validator.Struct(product); err != nil {
		return err
	}
	err := s.Repository.Update(product)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return errors.New(fmt.Sprintf("%s already exists", product.ProductName))
	} else {
		return err
	}
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.Repository.Delete(id)
}

func (s *ProductService) GetProducts() ([]domain.Product, error) {
	return s.Repository.FindAll()
}

func (s *ProductService) GetProduct(id uint) (domain.Product, error) {
	return s.Repository.FindByID(id)
}
