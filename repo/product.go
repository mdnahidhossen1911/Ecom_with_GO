package repo

import (
	"errors"

	"github.com/google/uuid"
)

type Product struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
}

type ProductRepo interface {
	Create(product Product) (*Product, error)
	Get(productID string) (*Product, error)
	List() ([]*Product, error)
	Delete(productID string) error 
	Update(pr Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {

	repo := &productRepo{}
	generateProduct(repo)
	return repo
}

func (p *productRepo) Create(product Product) (*Product, error) {

	if product.ID != "" {
		return &product, nil
	}

	product.ID = uuid.New().String()
	p.productList = append(p.productList, &product)
	return &product, nil
}

func (p *productRepo) Get(productID string) (*Product, error) {
	for _, prod := range p.productList {
		if prod.ID == productID {
			return prod, nil
		}
	}
	return nil, errors.New("product not found")
}

func (p *productRepo) List() ([]*Product, error) {
	return p.productList, nil
}

func (p *productRepo) Delete(productID string) error {
	var tempList []*Product
	for _, prod := range p.productList {
		if prod.ID != productID {
			tempList = append(tempList, prod)
		}
	}
	p.productList = tempList
	return nil
}

func (p *productRepo) Update(pr Product) (*Product, error) {
	for index, prod := range p.productList {
		if prod.ID == pr.ID {
			p.productList[index] = &pr
			return &pr, nil
		}
	}
	return &pr, errors.New("product not found")
}

func generateProduct(r *productRepo) {

	prt := &Product{
		ID:          uuid.New().String(),
		Title:       "Sample Product",
		Description: "This is a sample product description.",
		Price:       19.99,
		ImageURL:    "http://example.com/image.jpg",
	}

	prt2 := &Product{
		ID:          uuid.New().String(),
		Title:       "Another Product",
		Description: "This is another product description.",
		Price:       29.99,
		ImageURL:    "http://example.com/image2.jpg",
	}

	prt3 := &Product{
		ID:          uuid.New().String(),
		Title:       "Third Product",
		Description: "This is the third product description.",
		Price:       39.99,
		ImageURL:    "http://example.com/image3.jpg",
	}

	prt4 := &Product{
		ID:          uuid.New().String(),
		Title:       "Fourth Product",
		Description: "This is the fourth product description.",
		Price:       49.99,
		ImageURL:    "http://example.com/image4.jpg",
	}

	r.productList = append(r.productList, prt, prt2, prt3, prt4)

}
