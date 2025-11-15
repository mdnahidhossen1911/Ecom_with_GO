package handlers

import (
	"ecom_project/database"
	"ecom_project/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	util.SendData(w, database.List(), 200)
}
