package handlers

import (
	"GorillaMuxProject/internal/userService"
	"GorillaMuxProject/internal/web/users"
	"context"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) GetApiUsers(_ context.Context, _ users.GetApiUsersRequestObject) (users.GetApiUsersResponseObject, error) {
	//TODO implement me
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetApiUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:    usr.Id,
			Email: usr.Email,
		}
		response = append(response, user)
	}
	return response, nil
}

func (h *UserHandler) PostApiUsers(_ context.Context, request users.PostApiUsersRequestObject) (users.PostApiUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := users.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	response := users.PostApiUsers201JSONResponse(createdUser)
	return response, nil
}

func (h *UserHandler) PatchApiUsersId(_ context.Context, request users.PatchApiUsersIdRequestObject) (users.PatchApiUsersIdResponseObject, error) {
	id := request.Id
	userRequest := request.Body

	userToUpdate := users.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	updatedUser, err := h.Service.UpdateUserByID(uint(id), userToUpdate)
	if err != nil {
		return nil, err
	}

	// Convert the updatedUser directly to the response type
	response := users.PatchApiUsersId200JSONResponse(updatedUser)
	return response, nil
}

func (h *UserHandler) DeleteApiUsersId(_ context.Context, request users.DeleteApiUsersIdRequestObject) (users.DeleteApiUsersIdResponseObject, error) {
	id := request.Id
	err := h.Service.DeleteUserByID(uint(id))
	if err != nil {
		return nil, err
	}
	return users.DeleteApiUsersId204Response{}, nil
}

func (h *UserHandler) GetApiUsersUserIdTasks(_ context.Context, request users.GetApiUsersUserIdTasksRequestObject) (users.GetApiUsersUserIdTasksResponseObject, error) {
	userId := request.UserId
	tasks, err := h.Service.GetUserTasks(userId)
	if err != nil {
		return nil, err
	}

	var response []users.Task
	for _, tsk := range tasks {
		user := users.Task{
			Id:     tsk.Id,
			Task:   tsk.Task,
			IsDone: tsk.IsDone,
		}
		response = append(response, user)
	}
	return users.GetApiUsersUserIdTasks200JSONResponse(response), nil
}
