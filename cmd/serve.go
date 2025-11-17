package cmd

import (
	"ecom_project/config"
	"ecom_project/repo"
	"ecom_project/rest"
	"ecom_project/rest/handlers/product"
	"ecom_project/rest/handlers/user"
	"ecom_project/rest/middleware"
)

func Serve() {

	config := config.GetConfig()

	userRepo := repo.NewUserRepo()
	productRepo := repo.NewProductRepo()

	middleware := middleware.NewConfigMiddleware(config)

	productsHandler := product.NewHandler(middleware, productRepo)
	userHandler := user.NewHandler(userRepo, *config)

	server := rest.NewServer(config, productsHandler, userHandler)
	server.Start()

}
