package main

import (
	"GorillaMuxProject/internal/database"
	"GorillaMuxProject/internal/handlers"
	"GorillaMuxProject/internal/taskService"
	"GorillaMuxProject/internal/userService"
	"GorillaMuxProject/internal/web/tasks"
	"GorillaMuxProject/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	tasksRepo := taskService.NewTaskRepository(database.DB)
	usersRepo := userService.NewUserRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	usersService := userService.NewService(usersRepo)

	tasksHandler := handlers.NewTaskHandler(tasksService)
	usersHandler := handlers.NewUserHandler(usersService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, taskStrictHandler)
	userStrictHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, userStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
