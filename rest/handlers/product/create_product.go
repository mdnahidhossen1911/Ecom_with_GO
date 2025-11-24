package product

import (
	"ecom_project/domain"
	"ecom_project/util"
	"encoding/json"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newProduct ReqCreateProduct
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	CreateProduct, err := h.svc.Create(domain.Product{
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImageURL:    newProduct.ImageURL,
	})

	if err != nil {
		util.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.SendData(w, CreateProduct, 201)
}
