package test

import (
	"github.com/go-playground/validator/v10"
	"go-fiber-hex-arch/internal/core/domain"
	"go-fiber-hex-arch/internal/core/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Insert(product domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) FindByID(id uint) (domain.Product, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockProductRepository) Update(product domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockProductRepository) FindAll() ([]domain.Product, error) {
	args := m.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func TestInsertProduct_ValidInput(t *testing.T) {
	product := domain.Product{ProductName: "New Product", Price: 15000, Stock: 12}
	mockRepo := new(MockProductRepository)
	mockRepo.On("Insert", product).Return(nil)

	svc := service.ProductService{
		Repository: mockRepo,
		Validator:  validator.New(),
	}

	err := svc.InsertProduct(product)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestInsertProduct_InvalidInput(t *testing.T) {
	product := domain.Product{ProductName: "", Price: 15000, Stock: 12}
	mockRepo := new(MockProductRepository)

	svc := service.ProductService{
		Repository: mockRepo,
		Validator:  validator.New(),
	}

	err := svc.InsertProduct(product)

	assert.Error(t, err)
	assert.Equal(t, "Key: 'Product.ProductName' Error:Field validation for 'ProductName' failed on the 'required' tag", err.Error()) // Sesuaikan dengan pesan error yang diharapkan
	mockRepo.AssertNotCalled(t, "Insert", mock.Anything)
}
