package rest

import (
	"ecom_project/config"
	"ecom_project/rest/handlers/product"
	"ecom_project/rest/handlers/user"
	"ecom_project/rest/middleware"
	"fmt"
	"net/http"
	"os"
)

type Server struct {
	ProductHandler *product.Handler
	UserHandler    *user.Handler
}

func NewServer(productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		ProductHandler: productHandler,
		UserHandler:    userHandler,
	}
}

func (Server *Server) Start(config *config.Config) {

	manager := middleware.NewManager()
	mux := http.NewServeMux()

	manager.Use(middleware.CORS)
	manager.Use(middleware.Preflight)
	manager.Use(middleware.Logger)

	Server.ProductHandler.RegisterRoutes(mux, manager)
	Server.UserHandler.RegisterRoutes(mux, manager)

	wrappedMux := manager.ApplyToMux(mux)

	addr := fmt.Sprintf(":%s", config.Port)

	fmt.Println("Server is listening on port", config.Port, "...")
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}

}
