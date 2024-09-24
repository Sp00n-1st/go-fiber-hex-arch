package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-fiber-hex-arch/internal/core/domain"
	"go-fiber-hex-arch/internal/core/service"
	"testing"
)

func TestGetProductsById_Success(t *testing.T) {
	mockRepo := new(MockProductRepository)
	product := domain.Product{
		ProductID:   1,
		ProductName: "Kraft",
		Price:       15000,
		Stock:       4,
	}
	mockRepo.On("FindByID", uint(1)).Return(product, nil)

	svc := service.ProductService{
		Repository: mockRepo,
	}

	result, err := svc.GetProduct(uint(1))

	assert.NoError(t, err)
	assert.Equal(t, product, result)
	mockRepo.AssertExpectations(t)
}

func TestGetProductByID_NotFound(t *testing.T) {
	mockRepo := new(MockProductRepository)
	mockRepo.On("FindByID", uint(99)).Return(domain.Product{}, errors.New("record not found"))

	svc := service.ProductService{
		Repository: mockRepo,
	}

	_, err := svc.GetProduct(uint(99))

	assert.Error(t, err)
	assert.Equal(t, "record not found", err.Error())
	mockRepo.AssertExpectations(t)
}
