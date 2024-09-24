package test

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/mock"
	"go-fiber-hex-arch/internal/core/domain"
	"go-fiber-hex-arch/internal/core/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateProduct_ValidInput(t *testing.T) {
	product := domain.Product{ProductID: 1, ProductName: "Updated Product", Price: 14000, Stock: 12}
	mockRepo := new(MockProductRepository)
	mockRepo.On("Update", mock.AnythingOfType("domain.Product")).Return(nil)

	svc := service.ProductService{
		Repository: mockRepo,
		Validator:  validator.New(),
	}

	err := svc.UpdateProduct(product)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProduct_NotFound(t *testing.T) {
	product := domain.Product{ProductID: 99, ProductName: "Non Exist Product", Price: 14000, Stock: 12}
	mockRepo := new(MockProductRepository)
	mockRepo.On("Update", mock.AnythingOfType("domain.Product")).Return(errors.New("record not found"))

	svc := service.ProductService{
		Repository: mockRepo,
		Validator:  validator.New(),
	}

	err := svc.UpdateProduct(product)

	assert.Error(t, err)
	assert.Equal(t, "record not found", err.Error())
	mockRepo.AssertExpectations(t)
}
