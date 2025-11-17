package product

import (
	"ecom_project/util"
	"net/http"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("productID")

	if id == "" {
		util.SendError(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	err := h.productRepo.Delete(id)

	if err != nil {
		util.SendError(w, err.Error(), http.StatusNotFound)
		return
	}

	util.SendData(w, "Product deleted successfully", http.StatusOK)
}
