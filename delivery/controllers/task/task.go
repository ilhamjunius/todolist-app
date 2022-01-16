package task

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todolist-app/delivery/formatter"
	mw "todolist-app/delivery/middlewares"
	request "todolist-app/delivery/request"
	"todolist-app/delivery/response"
	"todolist-app/entities"
	"todolist-app/repository/task"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	repoTask task.Task
}

func NewController(repoTask task.Task) *TaskController {
	return &TaskController{repoTask}
}

func (c *TaskController) GetTasksByUserID(e echo.Context) error {
	res := response.Response{}

	currentUserID := mw.NewAuth().ExtractTokenUserID(e)

	tasks, err := c.repoTask.GetTasksByUserID(currentUserID)

	if err != nil || len(tasks) == 0 {
		res.Status = http.StatusBadRequest
		res.Message = "not found tasks"
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	formatData := formatter.FormatTasks(tasks)

	res.Status = http.StatusOK
	res.Message = "success get all task"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *TaskController) GetTaskByID(e echo.Context) error {
	res := response.Response{}

	id := e.Param("id")

	task_id, _ := strconv.Atoi(id)

	currentUserID := mw.NewAuth().ExtractTokenUserID(e)

	task, err := c.repoTask.GetTaskByID(task_id, currentUserID)

	if err != nil || task.ID == 0 {
		res.Status = http.StatusNotFound
		res.Message = "task not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	formatData := formatter.FormatTask(task)

	res.Status = http.StatusOK
	res.Message = "success get task"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *TaskController) CreateTask(e echo.Context) error {
	res := response.Response{}

	request := request.CreateTaskInput{}

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

	currentUserID := mw.NewAuth().ExtractTokenUserID(e)

	taskData, err := c.repoTask.GetTaskByName(request.Name)

	if err != nil || taskData.ID != uint(0) {
		res.Status = http.StatusBadRequest
		res.Message = "the task same on previous task"
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	task := entities.Task{}

	if request.ProjectID != 0 {
		projectData, err := c.repoTask.FindProjectID(request.ProjectID)

		if err != nil || projectData.ID == 0 {
			res.Status = http.StatusBadRequest
			res.Message = "project id not found"
			res.Data = nil

			return e.JSON(http.StatusBadRequest, res)
		}

		task.ProjectID = uint(request.ProjectID)

	}

	task.Name = request.Name
	task.Desc = request.Desc
	task.Status = request.Status
	task.UserID = uint(currentUserID)

	newTask, err := c.repoTask.CreateTask(task)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = fmt.Sprint(err)
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	formatData := formatter.FormatTask(newTask)

	res.Status = http.StatusOK
	res.Message = "success create task"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *TaskController) UpdateTask(e echo.Context) error {
	res := response.Response{}

	id := e.Param("id")

	task_id, _ := strconv.Atoi(id)

	input := request.UpdateTaskInput{}

	err := e.Bind(&input)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = fmt.Sprint(err.Error())
		res.Data = input

		return e.JSON(http.StatusBadRequest, res)
	}

	err = e.Validate(&input)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = fmt.Sprint(err.Error())
		res.Data = input

		return e.JSON(http.StatusBadRequest, res)
	}

	currentUserID := mw.NewAuth().ExtractTokenUserID(e)

	task, err := c.repoTask.GetTaskByID(task_id, currentUserID)

	if err != nil || task.ID == 0 || task.Status != "reopen" {
		res.Status = http.StatusNotFound
		res.Message = "task not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	taskData, _ := c.repoTask.GetTaskByName(input.Name)

	if taskData.ID != 0 {
		res.Status = http.StatusBadRequest
		res.Message = "the task same on previous task"
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	if input.ProjectID != 0 {
		projectData, err := c.repoTask.FindProjectID(input.ProjectID)

		if err != nil || projectData.ID == 0 {
			res.Status = http.StatusBadRequest
			res.Message = "project id not found"
			res.Data = nil

			return e.JSON(http.StatusBadRequest, res)
		}

		task.ProjectID = uint(input.ProjectID)

	}

	task.Name = input.Name
	task.Desc = input.Desc
	task.Status = strings.ToLower(input.Status)
	task.UserID = uint(currentUserID)

	updatedTask, err := c.repoTask.UpdateTask(task)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = fmt.Sprint(err.Error())
		res.Data = updatedTask

		return e.JSON(http.StatusBadRequest, res)
	}

	formatData := formatter.FormatTask(updatedTask)

	res.Status = http.StatusOK
	res.Message = "success update task"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *TaskController) DeleteTask(e echo.Context) error {
	res := response.Response{}

	id := e.Param("id")

	task_id, _ := strconv.Atoi(id)

	currentUserID := mw.NewAuth().ExtractTokenUserID(e)

	task, err := c.repoTask.GetTaskByID(task_id, currentUserID)

	if err != nil || task.ID == 0 {
		res.Status = http.StatusNotFound
		res.Message = "task not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	if task.Status != "reopen" {
		res.Status = http.StatusBadRequest
		res.Message = "can't delete completed task"
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	oldTask := task

	_, err = c.repoTask.DeleteTask(task)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = fmt.Sprint(err.Error())
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	formatData := formatter.FormatTask(oldTask)

	res.Status = http.StatusOK
	res.Message = "success delete task"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}
