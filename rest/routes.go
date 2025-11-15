package rest

import (
	"ecom_project/rest/handlers"
	"ecom_project/rest/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /products",
		manager.Apply(
			http.HandlerFunc(handlers.GetProducts),
			middleware.Demo,
		),
	)

	mux.Handle("POST /products",
		manager.Apply(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)

	mux.Handle("GET /products/{productID}",
		manager.Apply(
			http.HandlerFunc(handlers.GetProductByID),
		),
	)

	mux.Handle("PUT /products/{productID}",
		manager.Apply(
			http.HandlerFunc(handlers.UpdateProduct),
		),
	)

	mux.Handle("DELETE /products/{productID}",
		manager.Apply(
			http.HandlerFunc(handlers.DeleteProduct),
		),
	)

}
