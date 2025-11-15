package handlers

import (
	"ecom_project/database"
	"ecom_project/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("productID")

	productID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	deocder := json.NewDecoder(r.Body)
	var newProduct database.Product
	err = deocder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	newProduct.ID = productID
	UpdateProduct := database.Update(newProduct)
	if UpdateProduct == nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}

	util.SendData(w, UpdateProduct, 201)
}
