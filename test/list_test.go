package test

import (
	"github.com/stretchr/testify/assert"
	"go-fiber-hex-arch/internal/core/domain"
	"go-fiber-hex-arch/internal/core/service"
	"testing"
)

func TestGetProducts_Success(t *testing.T) {
	mockRepo := new(MockProductRepository)
	products := []domain.Product{
		{
			ProductID:   1,
			ProductName: "Product 1",
		},
		{
			ProductID:   2,
			ProductName: "Product 2",
		},
	}
	mockRepo.On("FindAll").Return(products, nil)

	svc := service.ProductService{
		Repository: mockRepo,
	}

	result, err := svc.GetProducts()

	assert.NoError(t, err)
	assert.Equal(t, products, result)
	mockRepo.AssertExpectations(t)
}

func TestGetProducts_Empty(t *testing.T) {
	mockRepo := new(MockProductRepository)
	products := []domain.Product{}
	mockRepo.On("FindAll").Return(products, nil)

	svc := service.ProductService{
		Repository: mockRepo,
	}

	result, err := svc.GetProducts()

	assert.NoError(t, err)
	assert.Equal(t, products, result)
	mockRepo.AssertExpectations(t)
}
