package product

import (
	"ecom_project/repo"
	"ecom_project/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("productID")

	if id == "" {
		util.SendError(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var newProduct ReqCreateProduct
	
	err := decoder.Decode(&newProduct)
	if err != nil {
		util.SendError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	UpdateProduct, err := h.productRepo.Update(repo.Product{
		ID:          id,
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImageURL:    newProduct.ImageURL,
	})
	if err != nil {
		util.SendError(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if UpdateProduct == nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}

	util.SendData(w, UpdateProduct, 201)
}
