package user

import "ecom_project/domain"

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
