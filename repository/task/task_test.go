package task

import (
	"os"
	"testing"
	"todolist-app/configs"
	"todolist-app/entities"
	"todolist-app/repository/project"
	"todolist-app/repository/user"
	"todolist-app/utils"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGetTasksByUserID(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	repo := NewRepository(db)
	repoProject := project.NewRepository(db)
	repoUser := user.NewRepository(db)

	t.Run("success-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Task{})
		db.AutoMigrate(&entities.Task{})

		mockTask := entities.Task{Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}
		mockProject := entities.Project{Name: "Project Alpha"}
		mockUser := entities.User{Name: "ela", Email: "ela@gmail.com", Password: "ela123"}
		_, _ = repoProject.CreateProject(mockProject)
		_, _ = repoUser.CreateUser(mockUser)

		createData, _ := repo.CreateTask(mockTask)
		taskData, err := repo.GetTasksByUserID(int(createData.UserID))

		assert.Nil(t, err)
		assert.Equal(t, mockTask.Name, taskData[0].Name)
		assert.Equal(t, mockTask.Desc, taskData[0].Desc)
		assert.Equal(t, mockTask.Status, taskData[0].Status)
		assert.Equal(t, 1, int(taskData[0].ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Task{})

		mockTask := entities.Task{Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}
		mockProject := entities.Project{Name: "Project Alpha"}
		mockUser := entities.User{Name: "ela", Email: "ela@gmail.com", Password: "ela123"}
		_, _ = repoProject.CreateProject(mockProject)
		_, _ = repoUser.CreateUser(mockUser)

		createData, _ := repo.CreateTask(mockTask)
		taskData, err := repo.GetTasksByUserID(int(createData.UserID))

		assert.Nil(t, err)
		assert.Equal(t, taskData, entities.Task{})

	})
}

func TestGetTaskByID(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	repo := NewRepository(db)
	repoProject := project.NewRepository(db)
	repoUser := user.NewRepository(db)

	t.Run("success-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Task{})
		db.AutoMigrate(&entities.Task{})

		mockTask := entities.Task{Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}
		mockProject := entities.Project{Name: "Project Alpha"}
		mockUser := entities.User{Name: "ela", Email: "ela@gmail.com", Password: "ela123"}
		_, _ = repoProject.CreateProject(mockProject)
		userData, _ := repoUser.CreateUser(mockUser)

		createData, _ := repo.CreateTask(mockTask)
		taskData, err := repo.GetTaskByID(int(createData.ID), int(userData.ID))

		assert.Nil(t, err)
		assert.Equal(t, mockTask.Name, taskData.Name)
		assert.Equal(t, mockTask.Desc, taskData.Desc)
		assert.Equal(t, mockTask.Status, taskData.Status)
		assert.Equal(t, 1, int(taskData.ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Task{})

		mockTask := entities.Task{Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}
		mockProject := entities.Project{Name: "Project Alpha"}
		mockUser := entities.User{Name: "ela", Email: "ela@gmail.com", Password: "ela123"}
		_, _ = repoProject.CreateProject(mockProject)
		userData, _ := repoUser.CreateUser(mockUser)

		createData, _ := repo.CreateTask(mockTask)
		taskData, err := repo.GetTaskByID(int(createData.ID), int(userData.ID))

		assert.Nil(t, err)
		assert.Equal(t, taskData, entities.Task{})

	})
}

func TestGetTaskByName(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	repo := NewRepository(db)
	repoProject := project.NewRepository(db)
	repoUser := user.NewRepository(db)

	t.Run("success-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Task{})
		db.AutoMigrate(&entities.Task{})

		mockTask := entities.Task{Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}
		mockProject := entities.Project{Name: "Project Alpha"}
		mockUser := entities.User{Name: "ela", Email: "ela@gmail.com", Password: "ela123"}
		_, _ = repoProject.CreateProject(mockProject)
		_, _ = repoUser.CreateUser(mockUser)

		createData, _ := repo.CreateTask(mockTask)
		taskData, err := repo.GetTaskByName(createData.Name)

		assert.Nil(t, err)
		assert.Equal(t, mockTask.Name, taskData.Name)
		assert.Equal(t, mockTask.Desc, taskData.Desc)
		assert.Equal(t, mockTask.Status, taskData.Status)
		assert.Equal(t, 1, int(taskData.ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Task{})

		mockTask := entities.Task{Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}
		mockProject := entities.Project{Name: "Project Alpha"}
		mockUser := entities.User{Name: "ela", Email: "ela@gmail.com", Password: "ela123"}
		_, _ = repoProject.CreateProject(mockProject)
		_, _ = repoUser.CreateUser(mockUser)

		createData, _ := repo.CreateTask(mockTask)
		taskData, err := repo.GetTaskByName(createData.Name)

		assert.Nil(t, err)
		assert.Equal(t, taskData, entities.Task{})

	})
}

func TestCreateProject(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	repo := NewRepository(db)
	repoProject := project.NewRepository(db)
	repoUser := user.NewRepository(db)

	t.Run("success-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Task{})
		db.AutoMigrate(&entities.Task{})

		mockTask := entities.Task{Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}
		mockProject := entities.Project{Name: "Project Alpha"}
		mockUser := entities.User{Name: "ela", Email: "ela@gmail.com", Password: "ela123"}
		_, _ = repoProject.CreateProject(mockProject)
		_, _ = repoUser.CreateUser(mockUser)

		createData, err := repo.CreateTask(mockTask)

		assert.Nil(t, err)
		assert.Equal(t, mockTask.Name, createData.Name)
		assert.Equal(t, mockTask.Desc, createData.Desc)
		assert.Equal(t, mockTask.Status, createData.Status)
		assert.Equal(t, 1, int(createData.ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Task{})

		mockTask := entities.Task{Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}
		mockProject := entities.Project{Name: "Project Alpha"}
		mockUser := entities.User{Name: "ela", Email: "ela@gmail.com", Password: "ela123"}
		_, _ = repoProject.CreateProject(mockProject)
		_, _ = repoUser.CreateUser(mockUser)

		createData, err := repo.CreateTask(mockTask)

		assert.Nil(t, err)
		assert.Equal(t, createData, entities.Task{})

	})
}

func TestFindProjectByID(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	repo := NewRepository(db)
	repoProject := project.NewRepository(db)

	t.Run("success-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Task{})
		db.AutoMigrate(&entities.Task{})

		mockProject := entities.Project{Name: "Project Alpha"}
		createProject, _ := repoProject.CreateProject(mockProject)

		projectData, err := repo.FindProjectID(int(createProject.ID))

		assert.Nil(t, err)
		assert.Equal(t, mockProject.Name, projectData.Name)
		assert.Equal(t, 1, int(projectData.ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Task{})

		mockProject := entities.Project{Name: "Project Alpha"}
		createProject, _ := repoProject.CreateProject(mockProject)

		projectData, err := repo.FindProjectID(int(createProject.ID))

		assert.Nil(t, err)
		assert.Equal(t, projectData, entities.Project{})

	})
}

func TestUpdateTask(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	repo := NewRepository(db)
	repoProject := project.NewRepository(db)
	repoUser := user.NewRepository(db)

	t.Run("success-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Task{})
		db.AutoMigrate(&entities.Task{})

		mockTaskCreate := entities.Task{Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}
		mockTaskUpdate := entities.Task{Name: "Task 1 new", Desc: "task 1 desc new", Status: "reopen", UserID: 1, ProjectID: 1}
		mockProject := entities.Project{Name: "Project Alpha"}
		mockUser := entities.User{Name: "ela", Email: "ela@gmail.com", Password: "ela123"}
		_, _ = repoProject.CreateProject(mockProject)
		_, _ = repoUser.CreateUser(mockUser)

		_, _ = repo.CreateTask(mockTaskCreate)
		updateTask, err := repo.UpdateTask(mockTaskUpdate)

		assert.Nil(t, err)
		assert.Equal(t, mockTaskUpdate.Name, updateTask.Name)
		assert.Equal(t, mockTaskUpdate.Desc, updateTask.Desc)
		assert.Equal(t, mockTaskUpdate.Status, updateTask.Status)
		assert.Equal(t, 1, int(updateTask.ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Project{})

		mockTaskCreate := entities.Task{Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}
		mockTaskUpdate := entities.Task{Name: "Task 1 new", Desc: "task 1 desc new", Status: "reopen", UserID: 1, ProjectID: 1}
		mockProject := entities.Project{Name: "Project Alpha"}
		mockUser := entities.User{Name: "ela", Email: "ela@gmail.com", Password: "ela123"}
		_, _ = repoProject.CreateProject(mockProject)
		_, _ = repoUser.CreateUser(mockUser)

		_, _ = repo.CreateTask(mockTaskCreate)
		updateTask, err := repo.UpdateTask(mockTaskUpdate)

		assert.Nil(t, err)
		assert.Equal(t, updateTask, entities.Task{})
	})
}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	repo := NewRepository(db)
	repoProject := project.NewRepository(db)
	repoUser := user.NewRepository(db)

	t.Run("success-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Task{})
		db.AutoMigrate(&entities.Task{})

		mockTask := entities.Task{Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}
		mockProject := entities.Project{Name: "Project Alpha"}
		mockUser := entities.User{Name: "ela", Email: "ela@gmail.com", Password: "ela123"}
		_, _ = repoProject.CreateProject(mockProject)
		_, _ = repoUser.CreateUser(mockUser)

		deleteData, err := repo.DeleteTask(mockTask)

		assert.Nil(t, err)
		assert.Equal(t, mockTask.Name, deleteData.Name)
		assert.Equal(t, mockTask.Desc, deleteData.Desc)
		assert.Equal(t, mockTask.Status, deleteData.Status)
		assert.Equal(t, 1, int(deleteData.ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Task{})

		mockTask := entities.Task{Name: "Task 1", Desc: "task 1 desc", Status: "reopen", UserID: 1, ProjectID: 1}
		mockProject := entities.Project{Name: "Project Alpha"}
		mockUser := entities.User{Name: "ela", Email: "ela@gmail.com", Password: "ela123"}
		_, _ = repoProject.CreateProject(mockProject)
		_, _ = repoUser.CreateUser(mockUser)

		deleteData, err := repo.DeleteTask(mockTask)

		assert.Nil(t, err)
		assert.Equal(t, deleteData, entities.Task{})

	})
}
