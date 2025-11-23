package user

import (
	"ecom_project/domain"
	"ecom_project/util"
	"encoding/json"
	"net/http"
)

type ReqCreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsOwner  bool   `json:"is_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var newUser ReqCreateUser
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	newPassword := util.SecPass(newUser.Password)

	user, err := h.svc.Create(domain.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newPassword,
		IsOwner:  newUser.IsOwner,
	})

	if err != nil {
		util.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.SendData(w, user, 200)
}
