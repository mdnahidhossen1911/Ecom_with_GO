package product

import (
	"ecom_project/repo"
	"ecom_project/rest/middleware"
)

type Handler struct {
	middleware *middleware.ConfigMiddleware
	productRepo repo.ProductRepo
}

func NewHandler(
	middleware *middleware.ConfigMiddleware,
	 productRepo repo.ProductRepo,
	 ) *Handler {
	return &Handler{
		middleware: middleware,
		productRepo: productRepo,
	}
}
