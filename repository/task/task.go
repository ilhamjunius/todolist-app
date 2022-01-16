package task

import (
	"todolist-app/entities"

	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (br *taskRepository) GetTasksByUserID(UserID int) ([]entities.Task, error) {
	tasks := []entities.Task{}

	err := br.db.Where("user_id = ?", UserID).Find(&tasks).Error

	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (br *taskRepository) GetTaskByID(id int, user_id int) (entities.Task, error) {
	task := entities.Task{}

	err := br.db.Where("ID = ? AND user_id = ? ", id, user_id).Find(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

func (br *taskRepository) GetTaskByName(name string) (entities.Task, error) {
	task := entities.Task{}

	err := br.db.Where("name = ?", name).Find(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

func (br *taskRepository) FindProjectID(id int) (entities.Project, error) {
	project := entities.Project{}

	err := br.db.Where("id = ?", id).Find(&project).Error

	if err != nil {
		return project, err
	}

	return project, nil
}

func (br *taskRepository) CreateTask(task entities.Task) (entities.Task, error) {
	err := br.db.Save(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

func (br *taskRepository) UpdateTask(task entities.Task) (entities.Task, error) {
	err := br.db.Save(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

func (br *taskRepository) DeleteTask(task entities.Task) (entities.Task, error) {
	err := br.db.Delete(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}
