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
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *Handler) PostApiTasks(_ context.Context, request tasks.PostApiTasksRequestObject) (tasks.PostApiTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}
	response := tasks.PostApiTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

func (h *Handler) PatchApiTasksId(_ context.Context, request tasks.PatchApiTasksIdRequestObject) (tasks.PatchApiTasksIdResponseObject, error) {
	id := request.Id

	taskRequest := request.Body

	taskToUpdate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	updatedTask, err := h.Service.UpdateTaskByID(uint(id), taskToUpdate)
	if err != nil {
		// Return an error response if the update fails
		return nil, err
	}
	response := tasks.PatchApiTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
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
