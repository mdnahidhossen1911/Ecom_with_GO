package handlers

import (
	"ecom_project/database"
	"ecom_project/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("productID")

	productID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range database.ProductsList {
		if product.ID == productID {
			util.SendJSONResponse(w, product, http.StatusOK)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}