package user

import (
	"todolist-app/entities"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetAllUser() ([]entities.User, error) {
	users := []entities.User{}

	ur.db.Find(&users)

	return users, nil
}

func (ur *userRepository) GetUserByID(id int) (entities.User, error) {
	user := entities.User{}

	ur.db.Where("ID = ?", id).Find(&user)

	return user, nil
}

func (ur *userRepository) GetUserByEmail(email string) (entities.User, error) {
	user := entities.User{}

	ur.db.Where("email = ?", email).Find(&user)

	return user, nil
}

func (ur *userRepository) CreateUser(user entities.User) (entities.User, error) {
	ur.db.Save(&user)

	return user, nil
}

func (ur *userRepository) UpdateUser(user entities.User) (entities.User, error) {
	ur.db.Save(&user)

	return user, nil
}

func (ur *userRepository) DeleteUser(user entities.User) (entities.User, error) {
	ur.db.Delete(&user)

	return user, nil
}
