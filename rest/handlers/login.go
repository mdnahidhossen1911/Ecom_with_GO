package handlers

import (
	"ecom_project/database"
	"ecom_project/util"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {

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
	util.SendData(w, validuser, 200)

}
