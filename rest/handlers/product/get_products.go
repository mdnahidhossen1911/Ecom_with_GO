package product

import (
	"ecom_project/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	products, err := h.svc.List()
	if err != nil {
		util.SendError(w, "Failed to retrieve products: " + err.Error(), http.StatusInternalServerError)
		return
	}

	if products == nil {
		util.SendError(w, "Not found products", http.StatusNotFound)
		return
	}

	util.SendData(w, products, http.StatusOK)
}
