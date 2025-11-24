package product

import (
	"ecom_project/rest/middleware"
)

type Handler struct {
	middleware *middleware.ConfigMiddleware
	svc        Service
}

func NewHandler(
	middleware *middleware.ConfigMiddleware,
	svc Service,
) *Handler {
	return &Handler{
		middleware: middleware,
		svc:        svc,
	}
}
