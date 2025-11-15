package handlers

import (
	"ecom_project/util"
	"ecom_project/database"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendJSONResponse(w, database.ProductsList, 200)
}