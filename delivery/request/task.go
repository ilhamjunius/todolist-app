package request

type CreateTaskInput struct {
	UserID    int    `json:"user_id" form:"user_id"`
	ProjectID int    `json:"project_id" form:"project_id"`
	Name      string `json:"name" form:"name" validate:"required"`
	Desc      string `json:"desc" form:"desc" validate:"required"`
	Status    string `json:"status" form:"status" validate:"required"`
}

type UpdateTaskInput struct {
	UserID    int    `json:"user_id" form:"user_id"`
	ProjectID int    `json:"project_id" form:"project_id"`
	Name      string `json:"name" form:"name" validate:"required"`
	Desc      string `json:"desc" form:"desc" validate:"required"`
	Status    string `json:"status" form:"status" validate:"required"`
}

type GetTaskInput struct {
	ID     uint `json:"id" form:"name" validate:"required"`
	UserID int
}
