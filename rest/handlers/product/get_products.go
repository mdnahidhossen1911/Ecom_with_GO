package product

import (
	"ecom_project/util"
	"net/http"
	"strconv"
	"sync"
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

	var (
		products any
		count    int64
		errList  error
		errCount error
	)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		products, errList = h.svc.List(page, limit)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		count, errCount = h.svc.Count()
	}()

	wg.Wait()

	if errList != nil {
		util.SendError(w, "Failed to retrieve products: "+errList.Error(), http.StatusInternalServerError)
		return
	}

	if products == nil {
		util.SendError(w, "Not found products", http.StatusNotFound)
		return
	}

	if errCount != nil {
		util.SendError(w, "Failed to retrieve products count: "+errCount.Error(), http.StatusInternalServerError)
		return
	}

	util.SendPage(w, products, page, limit, count)
}
