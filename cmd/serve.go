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
	"os"
)

func Serve() {

	config := config.GetConfig()

	dbCon, err := db.NewDBConnection(config.DBConfig)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		os.Exit(1)
		return
	}

	error := db.MigrateDB(dbCon, "./migrations")
	if error != nil {
		fmt.Println("Failed to migrate the database:", error)
		os.Exit(1)
		return
	}

	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)

	middleware := middleware.NewConfigMiddleware(config)

	productsHandler := product.NewHandler(middleware, productRepo)
	userHandler := user.NewHandler(userRepo, *config)

	server := rest.NewServer(config, productsHandler, userHandler)
	server.Start()

}
