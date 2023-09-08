package services

import (
	"github.com/awesomeProject/internal/api/dto"
	"github.com/awesomeProject/internal/storage/repositories"
	"github.com/sirupsen/logrus"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repo: repository}
}

func (u *userService) Get() ([]dto.GetUserDto, error) {
	users, err := u.repo.Get()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	var dtos []dto.GetUserDto = make([]dto.GetUserDto, len(users))
	for i, user := range users {
		dtos[i] = dto.GetUserDto{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
	}

	return dtos, nil
}

func (u *userService) GetByID(id int) (dto.GetUserDto, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) Create(dto dto.CreateUserDto) (dto.GetUserDto, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) Update(dto dto.UpdateUserDto) (dto.GetUserDto, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
