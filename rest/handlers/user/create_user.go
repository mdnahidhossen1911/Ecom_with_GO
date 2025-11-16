package user

import (
	"ecom_project/database"
	"ecom_project/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	deocder := json.NewDecoder(r.Body)
	var newUser database.User
	err := deocder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	user, error := database.CreateUser(newUser)
	if error != nil {
		util.SendError(w, error.Error(),http.StatusBadRequest)
		return
	}

	util.SendData(w, user, 200)
}
