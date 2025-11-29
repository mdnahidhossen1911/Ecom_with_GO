package product

import "ecom_project/domain"

type service struct {
	prdRepo ProductRepo
}



func NewService(prdRepo ProductRepo) Service {
	return &service{
		prdRepo: prdRepo,
	}
}

func (s *service) Create(product domain.Product) (*domain.Product, error) {
	return s.prdRepo.Create(product)
}

func (s *service) Delete(productID string) error {
	return s.prdRepo.Delete(productID)
}

func (s *service) Get(productID string) (*domain.Product, error) {
	return s.prdRepo.Get(productID)
}

func (s *service) List(page, limit int64) ([]*domain.Product, error) {
	return s.prdRepo.List(page, limit)
}


func (s *service) Count() (int64, error) {
	return s.prdRepo.Count()
}

func (s *service) Update(pr domain.Product) (*domain.Product, error) {
	return s.prdRepo.Update(pr)
}
