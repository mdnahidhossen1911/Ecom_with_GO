package product

import (
	"ecom_project/rest/middleware"
	"net/http"

)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

		mux.Handle("GET /products",
		manager.Apply(
			http.HandlerFunc(h.GetProducts),
		),
	)

	mux.Handle("POST /products",
		manager.Apply(
			http.HandlerFunc(h.CreateProduct),
			middleware.AuthenticateJWT,
		),
	)

	mux.Handle("GET /products/{productID}",
		manager.Apply(
			http.HandlerFunc(h.GetProductByID),
		),
	)

	mux.Handle("PUT /products/{productID}",
		manager.Apply(
			http.HandlerFunc(h.UpdateProduct),
			middleware.AuthenticateJWT,
		),
	)

	mux.Handle("DELETE /products/{productID}",
		manager.Apply(
			http.HandlerFunc(h.DeleteProduct),
			middleware.AuthenticateJWT,
		),
	)

}