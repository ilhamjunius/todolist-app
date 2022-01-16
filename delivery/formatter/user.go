package formatter

import entities "todolist-app/entities"

type UserFormatter struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FormatUser(user entities.User) UserFormatter {
	userFormatter := UserFormatter{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}

	return userFormatter
}

func FormatUsers(user []entities.User) []UserFormatter {
	usersFormatter := []UserFormatter{}

	for _, user := range user {
		userFormatter := FormatUser(user)
		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}
