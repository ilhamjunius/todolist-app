package user

import "todolist-app/entities"

type User interface {
	GetAllUser() ([]entities.User, error)
	GetUserByID(id int) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	CreateUser(user entities.User) (entities.User, error)
	UpdateUser(user entities.User) (entities.User, error)
	DeleteUser(user entities.User) (entities.User, error)
	// Login(email string, password string) (entities.User, error)
}
