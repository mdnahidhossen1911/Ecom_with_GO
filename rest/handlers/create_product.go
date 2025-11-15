package handlers

import (
	"ecom_project/util"
	"ecom_project/database"
	"encoding/json"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	deocder := json.NewDecoder(r.Body)
	var newProduct database.Product
	err := deocder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}
	newProduct.ID = len(database.ProductsList) + 1
	database.ProductsList = append(database.ProductsList, newProduct)
	util.SendJSONResponse(w, newProduct, 201)
}