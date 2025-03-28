// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// User defines model for User.
type User struct {
	Email    *openapi_types.Email `json:"email,omitempty"`
	Id       *uint                `json:"id,omitempty"`
	Password *string              `json:"password,omitempty"`
}

// PostApiUsersJSONRequestBody defines body for PostApiUsers for application/json ContentType.
type PostApiUsersJSONRequestBody = User

// PatchApiUsersIdJSONRequestBody defines body for PatchApiUsersId for application/json ContentType.
type PatchApiUsersIdJSONRequestBody = User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all users
	// (GET /api/users)
	GetApiUsers(ctx echo.Context) error
	// Create a new user
	// (POST /api/users)
	PostApiUsers(ctx echo.Context) error
	// Delete a user by ID
	// (DELETE /api/users/{id})
	DeleteApiUsersId(ctx echo.Context, id uint) error
	// Update a user by ID
	// (PATCH /api/users/{id})
	PatchApiUsersId(ctx echo.Context, id uint) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetApiUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetApiUsers(ctx)
	return err
}

// PostApiUsers converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostApiUsers(ctx)
	return err
}

// DeleteApiUsersId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteApiUsersId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteApiUsersId(ctx, id)
	return err
}

// PatchApiUsersId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchApiUsersId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchApiUsersId(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/users", wrapper.GetApiUsers)
	router.POST(baseURL+"/api/users", wrapper.PostApiUsers)
	router.DELETE(baseURL+"/api/users/:id", wrapper.DeleteApiUsersId)
	router.PATCH(baseURL+"/api/users/:id", wrapper.PatchApiUsersId)

}

type GetApiUsersRequestObject struct {
}

type GetApiUsersResponseObject interface {
	VisitGetApiUsersResponse(w http.ResponseWriter) error
}

type GetApiUsers200JSONResponse []User

func (response GetApiUsers200JSONResponse) VisitGetApiUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostApiUsersRequestObject struct {
	Body *PostApiUsersJSONRequestBody
}

type PostApiUsersResponseObject interface {
	VisitPostApiUsersResponse(w http.ResponseWriter) error
}

type PostApiUsers201JSONResponse User

func (response PostApiUsers201JSONResponse) VisitPostApiUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostApiUsers400Response struct {
}

func (response PostApiUsers400Response) VisitPostApiUsersResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type DeleteApiUsersIdRequestObject struct {
	Id uint `json:"id"`
}

type DeleteApiUsersIdResponseObject interface {
	VisitDeleteApiUsersIdResponse(w http.ResponseWriter) error
}

type DeleteApiUsersId204Response struct {
}

func (response DeleteApiUsersId204Response) VisitDeleteApiUsersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteApiUsersId404Response struct {
}

func (response DeleteApiUsersId404Response) VisitDeleteApiUsersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PatchApiUsersIdRequestObject struct {
	Id   uint `json:"id"`
	Body *PatchApiUsersIdJSONRequestBody
}

type PatchApiUsersIdResponseObject interface {
	VisitPatchApiUsersIdResponse(w http.ResponseWriter) error
}

type PatchApiUsersId200JSONResponse User

func (response PatchApiUsersId200JSONResponse) VisitPatchApiUsersIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchApiUsersId400Response struct {
}

func (response PatchApiUsersId400Response) VisitPatchApiUsersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type PatchApiUsersId404Response struct {
}

func (response PatchApiUsersId404Response) VisitPatchApiUsersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all users
	// (GET /api/users)
	GetApiUsers(ctx context.Context, request GetApiUsersRequestObject) (GetApiUsersResponseObject, error)
	// Create a new user
	// (POST /api/users)
	PostApiUsers(ctx context.Context, request PostApiUsersRequestObject) (PostApiUsersResponseObject, error)
	// Delete a user by ID
	// (DELETE /api/users/{id})
	DeleteApiUsersId(ctx context.Context, request DeleteApiUsersIdRequestObject) (DeleteApiUsersIdResponseObject, error)
	// Update a user by ID
	// (PATCH /api/users/{id})
	PatchApiUsersId(ctx context.Context, request PatchApiUsersIdRequestObject) (PatchApiUsersIdResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetApiUsers operation middleware
func (sh *strictHandler) GetApiUsers(ctx echo.Context) error {
	var request GetApiUsersRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetApiUsers(ctx.Request().Context(), request.(GetApiUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetApiUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetApiUsersResponseObject); ok {
		return validResponse.VisitGetApiUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostApiUsers operation middleware
func (sh *strictHandler) PostApiUsers(ctx echo.Context) error {
	var request PostApiUsersRequestObject

	var body PostApiUsersJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostApiUsers(ctx.Request().Context(), request.(PostApiUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostApiUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostApiUsersResponseObject); ok {
		return validResponse.VisitPostApiUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteApiUsersId operation middleware
func (sh *strictHandler) DeleteApiUsersId(ctx echo.Context, id uint) error {
	var request DeleteApiUsersIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteApiUsersId(ctx.Request().Context(), request.(DeleteApiUsersIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteApiUsersId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteApiUsersIdResponseObject); ok {
		return validResponse.VisitDeleteApiUsersIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchApiUsersId operation middleware
func (sh *strictHandler) PatchApiUsersId(ctx echo.Context, id uint) error {
	var request PatchApiUsersIdRequestObject

	request.Id = id

	var body PatchApiUsersIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchApiUsersId(ctx.Request().Context(), request.(PatchApiUsersIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchApiUsersId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchApiUsersIdResponseObject); ok {
		return validResponse.VisitPatchApiUsersIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
