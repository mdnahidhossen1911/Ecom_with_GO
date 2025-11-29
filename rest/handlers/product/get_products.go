package product

import (
	"ecom_project/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	reqQuary := r.URL.Query()

	pageAsString := reqQuary.Get("page")
	limitAsString := reqQuary.Get("limit")

	page, _ := strconv.ParseInt(pageAsString, 10, 32)
	limit, _ := strconv.ParseInt(limitAsString, 10, 32)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	products, err := h.svc.List(page, limit)
	if err != nil {
		util.SendError(w, "Failed to retrieve products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if products == nil {
		util.SendError(w, "Not found products", http.StatusNotFound)
		return
	}

	cnt, err := h.svc.Count()
	if err != nil {
		util.SendError(w, "Failed to retrieve products count: "+err.Error(), http.StatusInternalServerError)
		return
	}


	util.SendPage(w, products, page, limit, cnt)
}
