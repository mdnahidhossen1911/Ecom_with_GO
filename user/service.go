package user

import "ecom_project/domain"

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

// Create implements Service.
func (s *service) Create(user domain.User) (*domain.User, error) {
	usr, err := s.userRepo.Create(user)

	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil

}

// Delete implements Service.
func (s *service) Delete(userID string) error {
	panic("unimplemented")
}

// Find implements Service.
func (s *service) Find(email string, password string) (*domain.User, error) {

	usr, err := s.userRepo.Find(email, password)

	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}

	return usr, nil

}

// Get implements Service.
func (s *service) Get(userID string) (*domain.User, error) {
	panic("unimplemented")
}

// List implements Service.
func (s *service) List() ([]*domain.User, error) {
	panic("unimplemented")
}

// Update implements Service.
func (s *service) Update(user domain.User) (*domain.User, error) {
	panic("unimplemented")
}
