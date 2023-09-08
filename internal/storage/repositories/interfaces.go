package repositories

import "github.com/awesomeProject/internal/storage/models"

type UserRepository interface {
	Get() ([]models.User, error)
	GetByID(id int) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(models.User) (models.User, error)
	Delete(id int) error
}
