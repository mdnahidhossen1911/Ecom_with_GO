package user

import (
	"ecom_project/rest/middleware"
	"net/http"

)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("POST /users",
		manager.Apply(
			http.HandlerFunc(h.CreateUser),
		),
	)

	mux.Handle("POST /users/login",
		manager.Apply(
			http.HandlerFunc(h.Login),
		),
	)


}