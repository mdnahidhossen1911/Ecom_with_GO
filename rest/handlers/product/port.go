package product

import "ecom_project/domain"


type Service interface {
		Create(product domain.Product) (*domain.Product, error)
	Get(productID string) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Delete(productID string) error
	Update(pr domain.Product) (*domain.Product, error)

}