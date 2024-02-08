package usecase

import "testClean/domain"

type UserUsecase struct {
	UserRepository domain.UserRepository
}

func (uc *UserUsecase) GetUserByID(id string) (*domain.User, error) {
	user, err := uc.UserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}