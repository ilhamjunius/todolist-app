package project

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"todolist-app/delivery/response"
	"todolist-app/entities"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllProject(t *testing.T) {
	t.Run("1-success-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projects")

		projectController := NewController(mockProjectRepository{})
		projectController.GetAllProject(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data.([]interface{})[0].(map[string]interface{})

		name := data["name"]

		assert.Equal(t, name, "Project Alpha")
		assert.Equal(t, response.Status, http.StatusOK)
	})

	t.Run("2-error-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/projects")

		projectController := NewController(mockFalseProjectRepository{})
		projectController.GetAllProject(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, response.Data, nil)
		assert.Equal(t, response.Status, http.StatusNotFound)
	})
}

func TestGetProjectByID(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		projectController := NewController(mockProjectRepository{})
		projectController.GetProjectByID(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data.(map[string]interface{})

		name := data["name"]

		assert.Equal(t, name, "Project Alpha")
		assert.Equal(t, response.Status, http.StatusOK)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/?", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)

		context.SetPath("/projects")
		context.SetParamNames("id")
		context.SetParamValues("2")

		projectController := NewController(mockFalseProjectRepository{})
		projectController.GetProjectByID(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, response.Status, http.StatusNotFound)
	})
}

func TestPostProject(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"Name":   "Project Alpha",
			"Author": "penulis 1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/projects")

		projectController := NewController(mockProjectRepository{})
		projectController.CreateProject(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data.(map[string]interface{})

		name := data["name"]

		assert.Equal(t, name, "Project Alpha")
		assert.Equal(t, response.Status, http.StatusOK)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"Name": "Project Alpha",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/projects")

		projectController := NewController(mockFalseProjectRepository{})
		projectController.CreateProject(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, response.Data, nil)
		assert.Equal(t, response.Status, http.StatusBadRequest)
	})
}

func TestUpdateProject(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name": "Project Alpha new",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/projects")
		context.SetParamNames("id")
		context.SetParamValues("1")

		projectController := NewController(mockProjectRepository{})
		projectController.UpdateProject(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data.(map[string]interface{})

		name := data["name"]

		assert.Equal(t, name, "Project Alpha new")
		assert.Equal(t, response.Status, http.StatusOK)
	})
	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name": "Project Alpha new",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/projects")
		context.SetParamNames("id")
		context.SetParamValues("1")

		projectController := NewController(mockFalseProjectRepository{})
		projectController.UpdateProject(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, response.Data, nil)
		assert.Equal(t, response.Status, http.StatusNotFound)
	})
}

func TestDeleteProject(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/projects")
		context.SetParamNames("id")
		context.SetParamValues("1")

		projectController := NewController(mockProjectRepository{})
		projectController.DeleteProject(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, response.Message, "success delete project")
		assert.Equal(t, response.Status, http.StatusOK)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/projects")
		context.SetParamNames("id")
		context.SetParamValues("10")

		projectController := NewController(mockFalseProjectRepository{})
		projectController.DeleteProject(context)

		var response response.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, response.Data, nil)
		assert.Equal(t, response.Status, http.StatusNotFound)
	})
}

type mockProjectRepository struct{}

func (m mockProjectRepository) GetAllProject() ([]entities.Project, error) {
	return []entities.Project{
		{ID: 1, Name: "Project Alpha"},
	}, nil
}

func (m mockProjectRepository) GetProjectByID(id int) (entities.Project, error) {
	return entities.Project{
		ID: 1, Name: "Project Alpha"}, nil
}

func (m mockProjectRepository) CreateProject(entities.Project) (entities.Project, error) {
	return entities.Project{
		ID: 1, Name: "Project Alpha"}, nil
}

func (m mockProjectRepository) UpdateProject(entities.Project) (entities.Project, error) {
	return entities.Project{
		ID: 1, Name: "Project Alpha new"}, nil
}

func (m mockProjectRepository) DeleteProject(entities.Project) (entities.Project, error) {
	return entities.Project{
		ID: 1, Name: "Project Alpha"}, nil
}

type mockFalseProjectRepository struct{}

func (m mockFalseProjectRepository) GetAllProject() ([]entities.Project, error) {
	return nil, errors.New("no data")
}

func (m mockFalseProjectRepository) GetProjectByID(id int) (entities.Project, error) {
	return entities.Project{
		ID: 0, Name: ""}, errors.New("can't get project")
}

func (m mockFalseProjectRepository) CreateProject(entities.Project) (entities.Project, error) {
	return entities.Project{
		ID: 0, Name: ""}, errors.New("error create project")
}

func (m mockFalseProjectRepository) UpdateProject(entities.Project) (entities.Project, error) {
	return entities.Project{
		ID: 0, Name: ""}, errors.New("error update project")
}

func (m mockFalseProjectRepository) DeleteProject(entities.Project) (entities.Project, error) {
	return entities.Project{
		ID: 0, Name: ""}, errors.New("error delete project")
}
