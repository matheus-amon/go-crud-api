package usecase

import (
	"go-crud-api/model"
	"go-crud-api/repository"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repo,
	}
}

func (uu *UserUseCase) GetUsers() ([]model.User, error) {
	return uu.repository.GetUsers()
}

func (uu *UserUseCase) CreateUser(user model.User) (model.User, error) {
	userId, err := uu.repository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}

	user.ID = userId
	return user, nil
}

func (uu *UserUseCase) GetUserByID(userId int) (*model.User, error) {
	user, err := uu.repository.GetUserByID(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
