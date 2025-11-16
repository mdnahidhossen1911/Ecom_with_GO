package user

import (
	"ecom_project/config"
	"ecom_project/database"
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

	validuser := database.FindUser(user.Email, user.Password)
	if validuser == nil {
		util.SendError(w, "Invaild Cradiential", http.StatusNotFound)
		return
	}

	jwt, error := util.CreateJwt(config.GetConfig().JwtSecureKey, util.Payload{
		Sub:     validuser.ID,
		Name:    validuser.Name,
		Email:   validuser.Email,
		IsOwner: validuser.IsOwner,
	})

	if error != nil {
		fmt.Println(err)
		return
	}

	util.SendData(w, jwt, 200)

}
