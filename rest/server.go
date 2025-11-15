package rest

import (
	"ecom_project/config"
	"ecom_project/rest/middleware"
	"fmt"
	"net/http"
	"os"
)

func Start(config *config.Config) {

	manager := middleware.NewManager()
	mux := http.NewServeMux()

	manager.Use(middleware.CORS)
	manager.Use(middleware.Preflight)
	manager.Use(middleware.Logger)

	initRoutes(mux, manager)

	wrappedMux := manager.ApplyToMux(mux)

	addr := fmt.Sprintf(":%s", config.Port)

	fmt.Println("Server is listening on port", config.Port, "...")
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}

}
