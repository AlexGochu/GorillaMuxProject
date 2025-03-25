package handlers

import (
	"GorillaMuxProject/internal/taskService"
	"GorillaMuxProject/internal/web/tasks"
	"context"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetApiTasks(_ context.Context, _ tasks.GetApiTasksRequestObject) (tasks.GetApiTasksResponseObject, error) {
	//TODO implement me
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetApiTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     tsk.Id,
			Task:   tsk.Task,
			IsDone: tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *Handler) PostApiTasks(_ context.Context, request tasks.PostApiTasksRequestObject) (tasks.PostApiTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := tasks.Task{
		Task:   taskRequest.Task,
		IsDone: taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}
	response := tasks.PostApiTasks201JSONResponse(createdTask)
	return response, nil
}

func (h *Handler) PatchApiTasksId(_ context.Context, request tasks.PatchApiTasksIdRequestObject) (tasks.PatchApiTasksIdResponseObject, error) {
	id := request.Id
	taskRequest := request.Body

	taskToUpdate := tasks.Task{
		Task:   taskRequest.Task,
		IsDone: taskRequest.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskByID(uint(id), taskToUpdate)
	if err != nil {
		return nil, err
	}

	// Convert the updatedTask directly to the response type
	response := tasks.PatchApiTasksId200JSONResponse(updatedTask)
	return response, nil
}

func (h *Handler) DeleteApiTasksId(_ context.Context, request tasks.DeleteApiTasksIdRequestObject) (tasks.DeleteApiTasksIdResponseObject, error) {
	id := request.Id
	err := h.Service.DeleteTaskByID(uint(id))
	if err != nil {
		return nil, err
	}
	return tasks.DeleteApiTasksId204Response{}, nil
}
