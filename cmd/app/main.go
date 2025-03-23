package main

import (
	"GorillaMuxProject/internal/database"
	"GorillaMuxProject/internal/handlers"
	"GorillaMuxProject/internal/taskService"
	"GorillaMuxProject/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}

	/*router := mux.NewRouter()

	router.HandleFunc("/api/tasks", handler.GetTasksHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/tasks", handler.PostTaskHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/tasks/{id}", handler.PatchTaskHandler).Methods(http.MethodPatch)
	router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskHandler).Methods(http.MethodDelete)

	http.ListenAndServe(":8080", router)*/
}
