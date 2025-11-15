package database

import (
	"fmt"

	"github.com/google/uuid"
)

var users []User

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsOwner  bool   `json:"is_owner"`
}

func CreateUser(user User) (*User, error) {
	for _, u := range users {
		if u.Email == user.Email {
			return nil, fmt.Errorf("email already exists")
		}
	}

	user.ID = uuid.New().String()
	users = append(users, user)
	return &users[len(users)-1], nil
}

func FindUser(email string, password string) *User {

	for _, usr := range users {
		if usr.Email == email && usr.Password == password {
			return &usr
		}
	}

	return nil
}
