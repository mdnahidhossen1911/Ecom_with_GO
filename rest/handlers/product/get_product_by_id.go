package product

import (
	"ecom_project/util"
	"net/http"
)

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("productID")

	if id == "" {
		util.SendError(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	pdt, err := h.productRepo.Get(id)
	if err != nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}

	util.SendData(w, pdt, http.StatusOK)

}
