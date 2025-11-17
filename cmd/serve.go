package cmd

import (
	"ecom_project/config"
	"ecom_project/infra/db"
	"ecom_project/repo"
	"ecom_project/rest"
	"ecom_project/rest/handlers/product"
	"ecom_project/rest/handlers/user"
	"ecom_project/rest/middleware"
	"fmt"
)

func Serve() {

	config := config.GetConfig()

	_, err := db.NewDBConnection()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	userRepo := repo.NewUserRepo()
	productRepo := repo.NewProductRepo()

	middleware := middleware.NewConfigMiddleware(config)

	productsHandler := product.NewHandler(middleware, productRepo)
	userHandler := user.NewHandler(userRepo, *config)

	server := rest.NewServer(config, productsHandler, userHandler)
	server.Start()

}
