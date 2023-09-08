package services

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/awesomeProject/internal/api/dto"
	"github.com/awesomeProject/internal/storage/models"
	"github.com/awesomeProject/internal/storage/repositories"
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
	user, err := u.repo.GetByID(id)
	if err != nil {
		logrus.Error(err)
		return dto.GetUserDto{}, err
	}

	return dto.GetUserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *userService) Create(createUserDto dto.CreateUserDto) (dto.GetUserDto, error) {
	user, err := u.repo.Create(models.User{
		Name:     createUserDto.Name,
		Email:    createUserDto.Email,
		Password: createUserDto.Password,
	})

	if err != nil {
		logrus.Error(err)
		return dto.GetUserDto{}, err
	}

	return dto.GetUserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *userService) Update(updateUserDto dto.UpdateUserDto) (dto.GetUserDto, error) {
	user, err := u.repo.Update(models.User{
		Model: gorm.Model{
			ID: updateUserDto.ID,
		},
		Name:  updateUserDto.Name,
		Email: updateUserDto.Email,
	})

	if err != nil {
		logrus.Error(err)
		return dto.GetUserDto{}, err
	}

	return dto.GetUserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *userService) Delete(id int) error {
	err := u.repo.Delete(id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
