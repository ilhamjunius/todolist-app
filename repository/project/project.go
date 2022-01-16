package project

import (
	"todolist-app/entities"

	"gorm.io/gorm"
)

type projectRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *projectRepository {
	return &projectRepository{db}
}

func (br *projectRepository) GetAllProject() ([]entities.Project, error) {
	projects := []entities.Project{}

	err := br.db.Find(&projects).Error

	if err != nil {
		return projects, err
	}

	return projects, nil
}

func (br *projectRepository) GetProjectByID(id int) (entities.Project, error) {
	project := entities.Project{}

	err := br.db.Where("ID = ?", id).Find(&project).Error

	if err != nil {
		return project, err
	}

	return project, nil
}

func (br *projectRepository) CreateProject(project entities.Project) (entities.Project, error) {
	err := br.db.Save(&project).Error

	if err != nil {
		return project, err
	}

	return project, nil
}

func (br *projectRepository) UpdateProject(project entities.Project) (entities.Project, error) {
	err := br.db.Save(&project).Error

	if err != nil {
		return project, err
	}

	return project, nil
}

func (br *projectRepository) DeleteProject(project entities.Project) (entities.Project, error) {
	err := br.db.Where("id = ?", project.ID).Delete(&project).Error

	if err != nil {
		return project, err
	}

	return project, nil
}
