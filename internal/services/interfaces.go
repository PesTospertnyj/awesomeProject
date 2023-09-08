package services

import "awesomeProject/internal/storage/models"

type UserService interface {
	Get() ([]models.User, error)
	GetByID(id int) (models.User, error)
	Create() (models.User, error)
	Update(models models.User) (models.User, error)
	Delete(id int) error
}
