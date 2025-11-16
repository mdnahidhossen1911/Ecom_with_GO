package product

import (
	"ecom_project/database"
	"ecom_project/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	deocder := json.NewDecoder(r.Body)
	var newProduct database.Product
	err := deocder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	CreateProduct := database.Stor(newProduct)

	util.SendData(w, CreateProduct, 201)
}
