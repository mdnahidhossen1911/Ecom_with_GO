package user

import (
	"ecom_project/config"
	"ecom_project/repo"
)

type Handler struct {
	repo       repo.UserRepo
	cnf  config.Config
}

func NewHandler( repo repo.UserRepo, cnf config.Config) *Handler {
	return &Handler{
		repo:       repo,
		cnf:        cnf,
	}
}
