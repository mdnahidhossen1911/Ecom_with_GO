package user

import (
	"ecom_project/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	deocder := json.NewDecoder(r.Body)
	var user ReqLogin

	err := deocder.Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	newPassword := util.SecPass(user.Password)

	validuser, err := h.repo.Find(user.Email, newPassword)
	if err != nil {
		util.SendError(w, "Invalid Credential "+err.Error(), http.StatusNotFound)
		return
	}

	jwt, error := util.CreateJwt(h.cnf.JwtSecureKey, util.Payload{
		Sub:       validuser.ID,
		Name:      validuser.Name,
		Email:     validuser.Email,
		IsOwner:   validuser.IsOwner,
		CreatedAt: validuser.CreatedAt.String(),
		UpdatedAt: validuser.UpdatedAt.String(),
	})

	if error != nil {
		fmt.Println(error)
		return
	}

	util.SendData(w, jwt, 200)

}
