package cmd

import (
	"ecom_project/config"
	"ecom_project/infra/db"
	"ecom_project/product"
	"ecom_project/repo"
	"ecom_project/rest"
	productHandler "ecom_project/rest/handlers/product"
	userHandler "ecom_project/rest/handlers/user"
	"ecom_project/rest/middleware"
	"ecom_project/user"
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

	// repos
	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)

	//domains
	usrService := user.NewService(userRepo)
	prdService := product.NewService(productRepo)
	middleware := middleware.NewConfigMiddleware(config)

	productsHandler := productHandler.NewHandler(middleware, prdService)
	userHandler := userHandler.NewHandler(config, usrService)

	server := rest.NewServer(config, productsHandler, userHandler)
	server.Start()

}
