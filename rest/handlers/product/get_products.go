package product

import (
	"ecom_project/database"
	"ecom_project/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	util.SendData(w, database.List(), 200)
}
