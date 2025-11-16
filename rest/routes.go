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
		),
	)

	mux.Handle("POST /products",
		manager.Apply(
			http.HandlerFunc(handlers.CreateProduct),
			middleware.AuthenticateJWT,
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
			middleware.AuthenticateJWT,
		),
	)

	mux.Handle("DELETE /products/{productID}",
		manager.Apply(
			http.HandlerFunc(handlers.DeleteProduct),
			middleware.AuthenticateJWT,
		),
	)

	mux.Handle("POST /users",
		manager.Apply(
			http.HandlerFunc(handlers.CreateUser),
		),
	)

	mux.Handle("POST /users/login",
		manager.Apply(
			http.HandlerFunc(handlers.Login),
		),
	)

}
