package user

import "ecom_project/domain"

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
