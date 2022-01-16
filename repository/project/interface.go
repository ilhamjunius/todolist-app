package project

import "todolist-app/entities"

type Project interface {
	GetAllProject() ([]entities.Project, error)
	GetProjectByID(id int) (entities.Project, error)
	CreateProject(project entities.Project) (entities.Project, error)
	UpdateProject(project entities.Project) (entities.Project, error)
	DeleteProject(project entities.Project) (entities.Project, error)
}
