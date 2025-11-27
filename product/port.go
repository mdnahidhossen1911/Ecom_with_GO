package product

import (
	"ecom_project/domain"
	productHandler "ecom_project/rest/handlers/product"
)

type Service interface {
	productHandler.Service
}

type ProductRepo interface {
	Create(product domain.Product) (*domain.Product, error)
	Get(productID string) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Delete(productID string) error
	Update(pr domain.Product) (*domain.Product, error)
}
