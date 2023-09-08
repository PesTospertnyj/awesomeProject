package services

import (
	"github.com/awesomeProject/internal/api/dto"
)

type UserService interface {
	Get() ([]dto.GetUserDto, error)
	GetByID(id int) (dto.GetUserDto, error)
	Create(dto dto.CreateUserDto) (dto.GetUserDto, error)
	Update(dto dto.UpdateUserDto) (dto.GetUserDto, error)
	Delete(id int) error
}
