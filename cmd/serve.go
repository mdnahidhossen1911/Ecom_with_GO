package cmd

import (
	"ecom_project/config"
	"ecom_project/rest"
	"ecom_project/rest/handlers/product"
	"ecom_project/rest/handlers/user"
)

func Serve() {

	config := config.GetConfig()

	productsHandler := product.NewHandler()
	userHandler := user.NewHandler()

	server := rest.NewServer(productsHandler, userHandler)

	server.Start(config)

}
