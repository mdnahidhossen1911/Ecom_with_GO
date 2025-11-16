package product

import (
	"ecom_project/database"
	"ecom_project/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("productID")

	productID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	isFound := database.Delete(productID)

	if isFound {
		util.SendData(w, "Product "+id+" deleted successfully.", http.StatusOK)
		return
	}

	util.SendError(w, "Product not found", http.StatusNotFound)
}
