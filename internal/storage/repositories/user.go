package repositories

import (
	"gorm.io/gorm"

	"github.com/awesomeProject/internal/storage/models"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Get() ([]models.User, error) {
	users := []models.User{}
	return users, r.db.Find(&users).Error
}

func (r *userRepository) GetByID(id int) (models.User, error) {
	user := models.User{}
	return user, r.db.First(&user, id).Error
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	return user, r.db.Create(&user).Error
}

func (r *userRepository) Update(user models.User) (models.User, error) {
	return user, r.db.Model(&user).Updates(user).Error
}

func (r *userRepository) Delete(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}
