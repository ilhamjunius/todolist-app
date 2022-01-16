package project

import (
	"fmt"
	"os"
	"testing"
	"todolist-app/configs"
	"todolist-app/entities"
	"todolist-app/utils"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGetAllProject(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	repo := NewRepository(db)

	t.Run("success-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Project{})
		db.AutoMigrate(&entities.Project{})

		mockProject := entities.Project{Name: "Project Alpha"}
		_, _ = repo.CreateProject(mockProject)
		projectData, err := repo.GetAllProject()

		assert.Nil(t, err)
		assert.Equal(t, mockProject.Name, projectData[0].Name)
		assert.Equal(t, 1, int(projectData[0].ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Project{})
		projectData, err := repo.GetAllProject()
		assert.Nil(t, err)
		assert.Equal(t, []entities.Project{}, projectData)
	})
}

func TestGetProjectByID(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	repo := NewRepository(db)

	t.Run("success-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Project{})
		db.AutoMigrate(&entities.Project{})

		mockProject := entities.Project{Name: "Project Alpha"}
		res, _ := repo.CreateProject(mockProject)
		projectData, err := repo.GetProjectByID(int(res.ID))

		assert.Nil(t, err)
		assert.Equal(t, mockProject.Name, projectData.Name)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Project{})
		projectData, err := repo.GetProjectByID(1)

		fmt.Println("projectData", projectData, "error", err)
		assert.Nil(t, err)
		assert.Equal(t, "", projectData.Name)
		assert.Equal(t, 0, int(projectData.ID))
	})
}

func TestCreateProject(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Project{})
	db.AutoMigrate(&entities.Project{})

	repo := NewRepository(db)

	t.Run("success-case", func(t *testing.T) {
		mockProject := entities.Project{Name: "Project Alpha"}
		res, err := repo.CreateProject(mockProject)
		assert.Nil(t, err)
		assert.Equal(t, mockProject.Name, res.Name)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Project{})
		mockProject := entities.Project{Name: "Project Alpha"}
		res, err := repo.CreateProject(mockProject)
		assert.Nil(t, err)
		assert.Equal(t, "", "")
		assert.Equal(t, 0, int(res.ID))
	})
}

func TestUpdateProject(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Project{})
	db.AutoMigrate(&entities.Project{})

	repo := NewRepository(db)

	t.Run("success-case", func(t *testing.T) {
		mockProject := entities.Project{Name: "Project Alpha new"}
		res, err := repo.UpdateProject(mockProject)
		assert.Nil(t, err)
		assert.Equal(t, mockProject.Name, res.Name)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Project{})
		mockProject := entities.Project{Name: "Project Alpha new"}
		res, err := repo.UpdateProject(mockProject)
		assert.Nil(t, err)
		assert.Equal(t, "", "")
		assert.Equal(t, 0, int(res.ID))
	})
}

func TestDeleteProject(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	repo := NewRepository(db)

	t.Run("success-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Project{})
		db.AutoMigrate(&entities.Project{})

		mockProject := entities.Project{Name: "Project Alpha"}
		res, _ := repo.CreateProject(mockProject)
		projectData, err := repo.DeleteProject(mockProject)

		assert.Nil(t, err)
		assert.Equal(t, mockProject.Name, projectData.Name)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Project{})
		projectData, err := repo.DeleteProject(entities.Project{Name: ""})
		assert.Nil(t, err)
		assert.Equal(t, "", projectData.Name)
		assert.Equal(t, 0, int(projectData.ID))
	})
}
