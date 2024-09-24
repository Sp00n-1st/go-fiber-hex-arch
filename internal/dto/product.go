package dto

import "go-fiber-hex-arch/internal/core/domain"

type ProductResponse struct {
	ProductID   uint   `json:"product_id,omitempty"`
	ProductName string `gorm:"size:255"`
	Price       int
	Stock       int
}

func ToProductResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		ProductID:   product.ProductID,
		ProductName: product.ProductName,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}

func ToProductResponses(products []domain.Product) []ProductResponse {
	responses := make([]ProductResponse, 0, len(products))

	for _, product := range products {
		responses = append(responses, ToProductResponse(product))
	}

	return responses
}
