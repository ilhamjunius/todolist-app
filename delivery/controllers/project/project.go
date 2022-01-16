package project

import (
	"fmt"
	"net/http"
	"strconv"
	"todolist-app/delivery/formatter"
	request "todolist-app/delivery/request"
	"todolist-app/delivery/response"
	"todolist-app/entities"
	"todolist-app/repository/project"

	"github.com/labstack/echo/v4"
)

type ProjectController struct {
	repo project.Project
}

func NewController(repo project.Project) *ProjectController {
	return &ProjectController{repo}
}

func (c *ProjectController) GetAllProject(e echo.Context) error {
	res := response.Response{}

	projects, err := c.repo.GetAllProject()

	if err != nil || len(projects) == 0 {
		res.Status = http.StatusNotFound
		res.Message = "no data found projects"
		res.Data = projects

		return e.JSON(http.StatusNotFound, res)
	}

	formatData := formatter.FormatProjects(projects)

	res.Status = http.StatusOK
	res.Message = "success get all project"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *ProjectController) GetProjectByID(e echo.Context) error {
	res := response.Response{}

	id := e.Param("id")

	project_id, _ := strconv.Atoi(id)

	project, err := c.repo.GetProjectByID(project_id)

	if err != nil || project.ID == 0 {
		res.Status = http.StatusNotFound
		res.Message = "project not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	formatData := formatter.FormatProject(project)

	res.Status = http.StatusOK
	res.Message = "success get project"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *ProjectController) CreateProject(e echo.Context) error {
	res := response.Response{}

	request := request.CreateProjectInput{}

	err := e.Bind(&request)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = fmt.Sprint(err.Error())
		res.Data = request

		return e.JSON(http.StatusBadRequest, res)
	}

	err = e.Validate(&request)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = fmt.Sprint(err.Error())
		res.Data = request

		return e.JSON(http.StatusBadRequest, res)
	}

	project := entities.Project{}

	project.Name = request.Name

	newProject, err := c.repo.CreateProject(project)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Data = nil
		res.Message = "can't add project"

		return e.JSON(http.StatusBadRequest, res)
	}

	formatData := formatter.FormatProject(newProject)

	res.Status = http.StatusOK
	res.Message = "success create project"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *ProjectController) UpdateProject(e echo.Context) error {
	res := response.Response{}

	id := e.Param("id")

	project := entities.Project{}

	project_id, _ := strconv.Atoi(id)

	project, err := c.repo.GetProjectByID(project_id)

	if err != nil || project.ID == 0 {
		res.Status = http.StatusNotFound
		res.Message = "project not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	request := request.UpdateProjectInput{}

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = fmt.Sprint(err.Error())
		res.Data = request

		return e.JSON(http.StatusBadRequest, res)
	}

	err = e.Validate(&request)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = fmt.Sprint(err.Error())
		res.Data = request

		return e.JSON(http.StatusBadRequest, res)
	}

	project.Name = request.Name

	updatedProject, err := c.repo.UpdateProject(project)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "can't update project"
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	formatData := formatter.FormatProject(updatedProject)

	res.Status = http.StatusOK
	res.Message = "success update project"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *ProjectController) DeleteProject(e echo.Context) error {
	res := response.Response{}

	id := e.Param("id")

	project := entities.Project{}

	project_id, _ := strconv.Atoi(id)

	project, err := c.repo.GetProjectByID(project_id)

	if err != nil || project.ID == 0 {
		res.Status = http.StatusNotFound
		res.Message = "project not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	oldProject := project

	_, err = c.repo.DeleteProject(project)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = fmt.Sprint(err.Error())
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	formatData := formatter.FormatProject(oldProject)

	res.Status = http.StatusOK
	res.Message = "success delete project"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}
