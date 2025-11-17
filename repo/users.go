package repo

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsOwner  bool   `json:"is_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Get(userID string) (*User, error)
	Find(email string, password string) (*User, error)

	List() ([]*User, error)
	Delete(userID string) error
	Update(user User) (*User, error)
}

type userRepo struct {
	users []User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (u *userRepo) Create(user User) (*User, error) {

	if user.ID != "" {
		return &user, nil
	}

	user.ID = uuid.New().String()
	u.users = append(u.users, user)

	return &user, nil
}

func (u *userRepo) Find(email string, password string) (*User, error) {
	for _, usr := range u.users {
		if usr.Email == email && usr.Password == password {
			return &usr, nil
		}
	}
	return nil, errors.New("user not found")
}

func (u *userRepo) Get(userID string) (*User, error) {
	for _, usr := range u.users {
		if usr.ID == userID {
			return &usr, nil
		}
	}
	return nil, errors.New("user not found")
}

func (u *userRepo) List() ([]*User, error) {
	var userList []*User
	for i := range u.users {
		userList = append(userList, &u.users[i])
	}
	return userList, nil
}

func (u *userRepo) Update(user User) (*User, error) {
	for index, usr := range u.users {
		if usr.ID == user.ID {
			u.users[index] = user
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (u *userRepo) Delete(userID string) error {
	var tempList []User
	for _, usr := range u.users {
		if usr.ID != userID {
			tempList = append(tempList, usr)
		}
	}
	u.users = tempList
	return nil
}

// func CreateUser(user User) (*User, error) {
// 	for _, u := range users {
// 		if u.Email == user.Email {
// 			return nil, fmt.Errorf("email already exists")
// 		}
// 	}

// 	user.ID = uuid.New().String()
// 	users = append(users, user)
// 	return &users[len(users)-1], nil
// }

// func FindUser(email string, password string) *User {

// 	for _, usr := range users {
// 		if usr.Email == email && usr.Password == password {
// 			return &usr
// 		}
// 	}

// 	return nil
// }
