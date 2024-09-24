package domain

type Product struct {
	ProductID   uint   `gorm:"primary_key"`
	ProductName string `json:"product_name" gorm:"uniqueIndex;size:255" validate:"required,min=3,max=255"`
	Price       int    `json:"price" validate:"required"`
	Stock       int    `json:"stock"`
}
