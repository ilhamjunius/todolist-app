package task

import "todolist-app/entities"

type Task interface {
	GetTaskByID(id int, UserID int) (entities.Task, error)
	GetTaskByName(name string) (entities.Task, error)
	GetTasksByUserID(UserID int) ([]entities.Task, error)
	CreateTask(task entities.Task) (entities.Task, error)
	UpdateTask(task entities.Task) (entities.Task, error)
	DeleteTask(task entities.Task) (entities.Task, error)
	FindProjectID(id int) (entities.Project, error)
}
