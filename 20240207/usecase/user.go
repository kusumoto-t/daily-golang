package usecase

import "github.com/kusumoto-t/cleanarchitecture/domain"

type UserUsecase struct {
	UserRepository domain.UserRepository
}

func (uc *UserUsecase) GetUserByID(id string) (*domain.User, error) {
	return uc.UserRepository.FindByID(id)
}

func (uc *UserUsecase) SaveUser(user *domain.User) error {
	return uc.UserRepository.Save(user)
}