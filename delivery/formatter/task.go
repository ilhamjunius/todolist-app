package formatter

import (
	"time"
	entities "todolist-app/entities"
)

type TaskFormatter struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Status    string    `json:"status"`
	UserID    uint      `json:"user_id"`
	ProjectID uint      `json:"project_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatTask(task entities.Task) TaskFormatter {
	taskFormatter := TaskFormatter{
		ID:        task.ID,
		Name:      task.Name,
		Desc:      task.Desc,
		Status:    task.Status,
		UserID:    task.UserID,
		ProjectID: task.ProjectID,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return taskFormatter
}

func FormatTasks(task []entities.Task) []TaskFormatter {
	tasksFormatter := []TaskFormatter{}

	for _, task := range task {
		taskFormatter := FormatTask(task)
		tasksFormatter = append(tasksFormatter, taskFormatter)
	}

	return tasksFormatter
}
