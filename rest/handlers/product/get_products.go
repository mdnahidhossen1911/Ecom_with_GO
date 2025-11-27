package product

import (
	"ecom_project/domain"
	"ecom_project/util"
	"net/http"
	"strconv"
)

type Pagination struct {
	Data []*domain.Product
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
	TotalPage int64 `json:"total_page"`
	TotalItem int64 `json:"total_item"`
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {


	reqQuary := r.URL.Query() 

	pageAsString := reqQuary.Get("page")
	limitAsString := reqQuary.Get("limit")

	page , _  := strconv.ParseInt(pageAsString, 10, 32)
	limit , _  := strconv.ParseInt(limitAsString, 10, 32)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	products, err := h.svc.List(page, limit)
	if err != nil {
		util.SendError(w, "Failed to retrieve products: " + err.Error(), http.StatusInternalServerError)
		return
	}

	if products == nil {
		util.SendError(w, "Not found products", http.StatusNotFound)
		return
	}

	pagination := Pagination{
		Data:  products,
		Page:  page,
		Limit: limit,
	}

	util.SendData(w, pagination, http.StatusOK)
}