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
	Config         *config.Config
	ProductHandler *product.Handler
	UserHandler    *user.Handler
}

func NewServer(config *config.Config, productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		Config:         config,
		ProductHandler: productHandler,
		UserHandler:    userHandler,
	}
}

func (Server *Server) Start() {
	mux := http.NewServeMux()

    // Initialize middleware manager
	manager := middleware.NewManager()

	manager.Use(middleware.CORS)
	manager.Use(middleware.Preflight)
	manager.Use(middleware.Logger)
	manager.Use(middleware.RateLimit)

	Server.ProductHandler.RegisterRoutes(mux, manager)
	Server.UserHandler.RegisterRoutes(mux, manager)

	wrappedMux := manager.ApplyToMux(mux)
	addr := fmt.Sprintf(":%s", Server.Config.Port)

	fmt.Println("Server is listening on port", Server.Config.Port, "...")
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
