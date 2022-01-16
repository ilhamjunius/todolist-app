package formatter

import (
	"time"
	entities "todolist-app/entities"
)

type ProjectFormatter struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatProject(task entities.Project) ProjectFormatter {
	taskFormatter := ProjectFormatter{
		ID:        task.ID,
		Name:      task.Name,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return taskFormatter
}

func FormatProjects(task []entities.Project) []ProjectFormatter {
	tasksFormatter := []ProjectFormatter{}

	for _, task := range task {
		taskFormatter := FormatProject(task)
		tasksFormatter = append(tasksFormatter, taskFormatter)
	}

	return tasksFormatter
}
