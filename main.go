package main

import (
	"fmt"
	"todolist-app/configs"
	_projectController "todolist-app/delivery/controllers/project"
	_taskController "todolist-app/delivery/controllers/task"
	_userController "todolist-app/delivery/controllers/user"
	"todolist-app/delivery/routes"
	_projectRepository "todolist-app/repository/project"
	_taskRepository "todolist-app/repository/task"
	_userRepository "todolist-app/repository/user"
	"todolist-app/utils"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	userRepository := _userRepository.NewRepository(db)
	userController := _userController.NewController(userRepository)

	projectRepository := _projectRepository.NewRepository(db)
	projectController := _projectController.NewController(projectRepository)

	taskRepository := _taskRepository.NewRepository(db)
	taskController := _taskController.NewController(taskRepository)

	e := echo.New()

	routes.RegisterPath(e, userController, projectController, taskController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
