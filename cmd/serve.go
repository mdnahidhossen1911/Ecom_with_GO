package cmd

import (
	"ecom_project/config"
	"ecom_project/rest"
	"ecom_project/rest/handlers/product"
	"ecom_project/rest/handlers/user"
	"ecom_project/rest/middleware"
)

func Serve() {

	config := config.GetConfig()

	middleware := middleware.NewConfigMiddleware(config)

	productsHandler := product.NewHandler(middleware)
	userHandler := user.NewHandler(middleware)

	server := rest.NewServer(config,productsHandler, userHandler)
	server.Start()

}
