package user

import "ecom_project/rest/middleware"

type Handler struct{
	middleware *middleware.ConfigMiddleware
}

func NewHandler(middleware *middleware.ConfigMiddleware) *Handler {
	return &Handler{
		middleware: middleware,
	}
}
