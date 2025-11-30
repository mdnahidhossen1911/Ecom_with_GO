package util

import (
	"net/http"
)

type PaginationResponse struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page      int64 `json:"page"`
	Limit     int64 `json:"limit"`
	TotalPage int64 `json:"total_page"`
	TotalItem int64 `json:"total_item"`
}

func SendPage(w http.ResponseWriter, data any, page, limit, totalItem int64) {

	pagination := PaginationResponse{
		Data: data,
		Pagination: Pagination{
			Page:      page,
			Limit:     limit,
			TotalItem: totalItem,
			TotalPage: (totalItem + limit - 1) / limit,
		},
	}

	SendData(w, pagination, http.StatusOK)

}
