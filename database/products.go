package database

import (
)

var productsList []Product

type Product struct {
	ID          int     `json:"id"` // Adding json tags for better JSON representation
	Image       string  `json:"image"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func Stor(p Product) Product {
	p.ID = len(productsList) + 1
	productsList = append(productsList, p)

	return p
}

func List() []Product {
	return productsList
}

func Get(productId int) *Product {
	for _, product := range productsList {
		if product.ID == productId {
			return &product
		}
	}

	return nil
}

func Update(p Product) *Product {
	for index, product := range productsList {
		if product.ID == p.ID {
			productsList[index] = p
			return &productsList[index]
		}
	}

	return nil
}

func Delete(productID int) bool {
	var temp []Product
	found := false

	for _, product := range productsList {
		if product.ID != productID {
			temp = append(temp, product)
		} else {
			found = true
		}
	}

	productsList = temp

	return found
}

func init() {
	p1 := Product{
		ID:          1,
		Image:       "image1.jpg",
		Name:        "Product 1",
		Description: "Description for product 1",
		Price:       19.99,
	}
	p2 := Product{
		ID:          2,
		Image:       "image2.jpg",
		Name:        "Product 2",
		Description: "Description for product 2",
		Price:       29.99,
	}
	p3 := Product{
		ID:          3,
		Image:       "image3.jpg",
		Name:        "Product 3",
		Description: "Description for product 3",
		Price:       39.99,
	}
	p4 := Product{
		ID:          4,
		Image:       "image4.jpg",
		Name:        "Product 4",
		Description: "Description for product 4",
		Price:       49.99,
	}
	productsList = []Product{p1, p2, p3, p4}
}
