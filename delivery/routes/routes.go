package routes

import (
	"todolist-app/configs"
	"todolist-app/delivery/controllers/project"
	"todolist-app/delivery/controllers/task"
	"todolist-app/delivery/controllers/user"
	m "todolist-app/delivery/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func RegisterPath(e *echo.Echo, uc *user.UserController, pc *project.ProjectController, tc *task.TaskController) {
	e.Validator = &CustomValidator{validator: validator.New()}
	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(configs.SecretKey)))

	eAuth.GET("/users/:id", uc.GetUserByID)
	e.POST("/users/register", uc.CreateUser)
	e.POST("/users/login", uc.Login)
	eAuth.PUT("/users/:id", uc.UpdateUser)
	eAuth.DELETE("/users/:id", uc.DeleteUser)

	eAuth.GET("/projects", pc.GetAllProject)
	eAuth.GET("/projects/:id", pc.GetProjectByID)
	eAuth.POST("/projects", pc.CreateProject)
	eAuth.PUT("/projects/:id", pc.UpdateProject)
	eAuth.DELETE("/projects/:id", pc.DeleteProject)

	eAuth.GET("/tasks", tc.GetTasksByUserID)
	eAuth.GET("/tasks/:id", tc.GetTaskByID)
	eAuth.POST("/tasks", tc.CreateTask)
	eAuth.PUT("/tasks/:id", tc.UpdateTask)
	eAuth.DELETE("/tasks/:id", tc.DeleteTask)

	m.LogMiddleware(e)
}
