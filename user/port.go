package user

import (
	"ecom_project/domain"
	userrHandiler "ecom_project/rest/handlers/user"
)

type Service interface {
	userrHandiler.Service
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Get(userID string) (*domain.User, error)
	Find(email string, password string) (*domain.User, error)

	List() ([]*domain.User, error)
	Delete(userID string) error
	Update(user domain.User) (*domain.User, error)
}
