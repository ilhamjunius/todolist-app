package request

type CreateProjectInput struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type UpdateProjectInput struct {
	ID   int
	Name string `json:"name" form:"name" validate:"required"`
}
