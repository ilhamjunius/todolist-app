package task

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todolist-app/delivery/response"
	"todolist-app/entities"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetTasksByUserID(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {

		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks")

		taskController := NewController(mockTaskRepository{})
		taskController.GetTasksByUserID(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data.([]interface{})[0].(map[string]interface{})

		name := data["name"]
		desc := data["desc"]
		status := data["status"]
		user_id := data["user_id"]
		project_id := data["project_id"]

		assert.Equal(t, name, "Task 1")
		assert.Equal(t, desc, "task 1 desc")
		assert.Equal(t, status, "reopen")
		assert.Equal(t, user_id, float64(1))
		assert.Equal(t, project_id, float64(1))
		assert.Equal(t, response.Status, http.StatusOK)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks")

		taskController := NewController(mockFalseTaskRepository{})
		taskController.GetTasksByUserID(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, response.Data, nil)
		assert.Equal(t, response.Status, http.StatusBadRequest)
	})
}

func TestGetTaskByID(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		context.Request().Header.Set(echo.HeaderAuthorization, "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIyLTAxLTE5VDA3OjMxOjM5LjM4MTkxMDgrMDc6MDAiLCJ1c2VyX2lkIjoxfQ.UNcVUvyKK58YoXUe6B9Yy6OwMg3RfE7EQL6mCk8043U")

		taskController := NewController(mockTaskRepository{})
		taskController.GetTaskByID(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data.(map[string]interface{})

		name := data["name"]
		desc := data["desc"]
		status := data["status"]
		user_id := data["user_id"]
		project_id := data["project_id"]

		assert.Equal(t, name, "Task 1")
		assert.Equal(t, desc, "task 1 desc")
		assert.Equal(t, status, "reopen")
		assert.Equal(t, user_id, float64(1))
		assert.Equal(t, project_id, float64(1))
		assert.Equal(t, response.Status, http.StatusOK)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		context.Request().Header.Set(echo.HeaderAuthorization, "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIyLTAxLTE5VDA3OjMxOjM5LjM4MTkxMDgrMDc6MDAiLCJ1c2VyX2lkIjoxfQ.UNcVUvyKK58YoXUe6B9Yy6OwMg3RfE7EQL6mCk8043U")

		taskController := NewController(mockFalseTaskRepository{})
		taskController.GetTaskByID(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, response.Data, nil)
		assert.Equal(t, response.Status, http.StatusNotFound)
	})
}

func TestCreateTask(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(entities.Task{
			Name:      "Task 2",
			Desc:      "task 2 desc",
			Status:    "reopen",
			ProjectID: 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/tasks")

		context.Request().Header.Set(echo.HeaderAuthorization, "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIyLTAxLTE5VDA3OjMxOjM5LjM4MTkxMDgrMDc6MDAiLCJ1c2VyX2lkIjoxfQ.UNcVUvyKK58YoXUe6B9Yy6OwMg3RfE7EQL6mCk8043U")

		taskController := NewController(mockTaskRepository{})
		taskController.CreateTask(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		fmt.Println(response)

		data := response.Data.(map[string]interface{})

		name := data["name"]
		desc := data["desc"]
		status := data["status"]
		user_id := data["user_id"]
		project_id := data["project_id"]

		assert.Equal(t, name, "Task 2")
		assert.Equal(t, desc, "task 2 desc")
		assert.Equal(t, status, "reopen")
		assert.Equal(t, user_id, float64(1))
		assert.Equal(t, project_id, float64(1))
		assert.Equal(t, response.Status, http.StatusOK)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(entities.Task{
			Name:      "Task 2",
			Desc:      "task 2 desc",
			Status:    "reopen",
			ProjectID: 1,
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/tasks")

		context.Request().Header.Set(echo.HeaderAuthorization, "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIyLTAxLTE5VDA3OjMxOjM5LjM4MTkxMDgrMDc6MDAiLCJ1c2VyX2lkIjoxfQ.UNcVUvyKK58YoXUe6B9Yy6OwMg3RfE7EQL6mCk8043U")

		taskController := NewController(mockFalseTaskRepository{})
		taskController.CreateTask(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, response.Data, nil)
		assert.Equal(t, response.Status, http.StatusBadRequest)
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(entities.Task{
			Name:      "Task 1 new",
			Desc:      "task 1 desc new",
			Status:    "reopen",
			ProjectID: 1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		context.Request().Header.Set(echo.HeaderAuthorization, "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIyLTAxLTE5VDA3OjMxOjM5LjM4MTkxMDgrMDc6MDAiLCJ1c2VyX2lkIjoxfQ.UNcVUvyKK58YoXUe6B9Yy6OwMg3RfE7EQL6mCk8043U")

		taskController := NewController(mockTaskRepository{})
		taskController.UpdateTask(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		fmt.Println(response.Message)

		data := response.Data.(map[string]interface{})

		name := data["name"]
		desc := data["desc"]
		status := data["status"]
		user_id := data["user_id"]
		project_id := data["project_id"]

		assert.Equal(t, name, "Task 1 new")
		assert.Equal(t, desc, "task 1 desc new")
		assert.Equal(t, status, "reopen")
		assert.Equal(t, user_id, float64(1))
		assert.Equal(t, project_id, float64(1))
		assert.Equal(t, response.Status, http.StatusOK)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(entities.Task{
			Name:      "Task 1 new",
			Desc:      "task 1 desc new",
			Status:    "reopen",
			ProjectID: 1,
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		context.Request().Header.Set(echo.HeaderAuthorization, "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIyLTAxLTE5VDA3OjMxOjM5LjM4MTkxMDgrMDc6MDAiLCJ1c2VyX2lkIjoxfQ.UNcVUvyKK58YoXUe6B9Yy6OwMg3RfE7EQL6mCk8043U")

		taskController := NewController(mockFalseTaskRepository{})
		taskController.UpdateTask(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, response.Data, nil)
		assert.Equal(t, response.Status, http.StatusNotFound)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		context.Request().Header.Set(echo.HeaderAuthorization, "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIyLTAxLTE5VDA3OjMxOjM5LjM4MTkxMDgrMDc6MDAiLCJ1c2VyX2lkIjoxfQ.UNcVUvyKK58YoXUe6B9Yy6OwMg3RfE7EQL6mCk8043U")

		taskController := NewController(mockTaskRepository{})
		taskController.DeleteTask(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data.(map[string]interface{})

		name := data["name"]
		desc := data["desc"]
		status := data["status"]
		user_id := data["user_id"]
		project_id := data["project_id"]

		assert.Equal(t, name, "Task 1")
		assert.Equal(t, desc, "task 1 desc")
		assert.Equal(t, status, "reopen")
		assert.Equal(t, user_id, float64(1))
		assert.Equal(t, project_id, float64(1))
		assert.Equal(t, response.Status, http.StatusOK)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		context.Request().Header.Set(echo.HeaderAuthorization, "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIyLTAxLTE5VDA3OjMxOjM5LjM4MTkxMDgrMDc6MDAiLCJ1c2VyX2lkIjoxfQ.UNcVUvyKK58YoXUe6B9Yy6OwMg3RfE7EQL6mCk8043U")

		taskController := NewController(mockFalseTaskRepository{})
		taskController.DeleteTask(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, response.Data, nil)
		assert.Equal(t, response.Status, http.StatusNotFound)
	})
}

type mockTaskRepository struct{}

func (m mockTaskRepository) GetTasksByUserID(UserID int) ([]entities.Task, error) {
	return []entities.Task{
		{ID: 1, Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1},
	}, nil
}

func (m mockTaskRepository) GetTaskByID(id int, user_id int) (entities.Task, error) {
	return entities.Task{
		ID: 1, Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}, nil
}

func (m mockTaskRepository) GetTaskByName(name string) (entities.Task, error) {
	if name != "" {
		return entities.Task{
			ID: 0, Name: "", Desc: "", Status: "", UserID: 0, ProjectID: 0}, nil
	}
	return entities.Task{
		ID: 1, Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}, nil
}

func (m mockTaskRepository) FindProjectID(id int) (entities.Project, error) {
	if id == 0 {
		return entities.Project{
			ID: 0, Name: ""}, errors.New("can't get project")
	}
	return entities.Project{
		ID: 1, Name: "Project Alpha"}, nil
}

func (m mockTaskRepository) CreateTask(entities.Task) (entities.Task, error) {
	return entities.Task{
		ID: 2, Name: "Task 2", Desc: "task 2 desc", Status: "reopen", UserID: 1, ProjectID: 1}, nil
}

func (m mockTaskRepository) UpdateTask(entities.Task) (entities.Task, error) {
	return entities.Task{
		ID: 1, Name: "Task 1 new", Desc: "task 1 desc new", Status: "reopen", UserID: 1, ProjectID: 1}, nil
}

func (m mockTaskRepository) DeleteTask(entities.Task) (entities.Task, error) {
	return entities.Task{
		ID: 1, Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}, nil
}

type mockFalseTaskRepository struct{}

func (m mockFalseTaskRepository) GetTasksByUserID(UserID int) ([]entities.Task, error) {
	return nil, errors.New("can't get tasks")
}

func (m mockFalseTaskRepository) GetTaskByID(id int, user_id int) (entities.Task, error) {
	return entities.Task{
		ID: 0, Name: "", Desc: "", Status: "", UserID: 0, ProjectID: 0}, errors.New("can't get task")
}

func (m mockFalseTaskRepository) GetTaskByName(name string) (entities.Task, error) {
	return entities.Task{
		ID: 0, Name: "", Desc: "", Status: "", UserID: 0, ProjectID: 0}, errors.New("can't get task")
}

func (m mockFalseTaskRepository) FindProjectID(id int) (entities.Project, error) {
	return entities.Project{
		ID: 0, Name: ""}, errors.New("can't get project")
}

func (m mockFalseTaskRepository) CreateTask(entities.Task) (entities.Task, error) {
	return entities.Task{
		ID: 0, Name: "", Desc: "", Status: "", UserID: 0, ProjectID: 0}, errors.New("no data")
}

func (m mockFalseTaskRepository) UpdateTask(entities.Task) (entities.Task, error) {
	return entities.Task{
		ID: 0, Name: " new", Desc: "", Status: "", UserID: 0, ProjectID: 0}, errors.New("no data")
}

func (m mockFalseTaskRepository) DeleteTask(entities.Task) (entities.Task, error) {
	return entities.Task{
		ID: 0, Name: "", Desc: "", Status: "", UserID: 0, ProjectID: 0}, errors.New("no data")
}
