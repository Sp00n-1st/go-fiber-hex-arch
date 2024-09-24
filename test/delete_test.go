package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-fiber-hex-arch/internal/core/service"
	"testing"
)

func TestDeleteProduct_Success(t *testing.T) {
	mockRepo := new(MockProductRepository)
	mockRepo.On("Delete", mock.AnythingOfType("uint")).Return(nil)

	svc := service.ProductService{
		Repository: mockRepo,
	}

	err := svc.DeleteProduct(uint(1))

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct_Fail(t *testing.T) {
	mockRepo := new(MockProductRepository)
	mockRepo.On("Delete", mock.AnythingOfType("uint")).Return(errors.New("record not found"))

	svc := service.ProductService{
		Repository: mockRepo,
	}

	err := svc.DeleteProduct(uint(1))

	assert.Error(t, err)
	assert.Equal(t, "record not found", err.Error())
	mockRepo.AssertExpectations(t)
}
