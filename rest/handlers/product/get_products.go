package product

import (
	"ecom_project/domain"
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

	var productList = make(chan []*domain.Product)
	var cont = make(chan int64)

	go func() {
		products, errList := h.svc.List(page, limit)
		if errList != nil {
			util.SendError(w, "Failed to retrieve products: "+errList.Error(), http.StatusInternalServerError)
			return
		}
		if products == nil {
			util.SendError(w, "Not found products", http.StatusNotFound)
			return
		}

		productList <- products

	}()

	go func() {
		count, errCount := h.svc.Count()
		if errCount != nil {
			util.SendError(w, "Failed to retrieve products count: "+errCount.Error(), http.StatusInternalServerError)
			return

		}

		cont <- count

	}()

	pdtList := <-productList
	count := <-cont

	util.SendPage(w, pdtList, page, limit, count)
}
