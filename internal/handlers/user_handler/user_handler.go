package user_handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	api "myservice/internal/api/generated"
	"myservice/internal/services"
)

type UserHandler struct {
	us *services.UserService
}

func NewUserHandler(us *services.UserService) *UserHandler {
	return &UserHandler{
		us: us,
	}
}

// compile-time проверка
//var _ openapi.ServerInterface = (*UserHandler)(nil)

// GET /user/{id}
func (h *UserHandler) GetUserById(
	ctx context.Context,
	request api.GetUserByIdRequestObject,
) (api.GetUserByIdResponseObject, error) {
	user, err := h.us.GetUser(ctx, request.Id)

	if errors.Is(err, pgx.ErrNoRows) {
		return api.GetUserById404JSONResponse{
			Code:    404,
			Message: "User not found",
		}, nil
	}
	if err != nil {
		return api.GetUserById500JSONResponse{
			Code:    500,
			Message: "Internal error",
		}, nil
	}

	return api.GetUserById200JSONResponse{
		Id:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}, nil

}

// POST /user
func (h *UserHandler) InsertUser(
	ctx context.Context,
	request api.InsertUserRequestObject,
) (api.InsertUserResponseObject, error) {

	if request.Body == nil {
		return api.InsertUser400JSONResponse{
			Code:    400,
			Message: "Empty request body",
		}, nil
	}

	if request.Body.Name == "" {
		return api.InsertUser400JSONResponse{
			Code:    400,
			Message: "Empty User Name",
		}, nil
	}
	user, err := h.us.CreateUser(ctx, request.Body.Name)
	fmt.Print(err)
	if err != nil {
		return api.InsertUser500JSONResponse{
			Code:    500,
			Message: "Internal error",
		}, nil
	}

	return api.InsertUser200JSONResponse{
		Id:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}, nil
}

// POST /user
func (h *UserHandler) UpdateUser(
	ctx context.Context,
	request api.UpdateUserRequestObject,
) (api.UpdateUserResponseObject, error) {
	//todo: валидация?
	if request.Body == nil {
		return api.UpdateUser400JSONResponse{
			Code:    400,
			Message: "Empty request body",
		}, nil
	}

	if request.Body.Name == "" {
		return api.UpdateUser400JSONResponse{
			Code:    400,
			Message: "Empty user name",
		}, nil
	}

	err := h.us.UpdateUser(ctx, request.Id, request.Body.Name)
	if err != nil {
		return api.UpdateUser500JSONResponse{
			Code:    500,
			Message: "Internal error",
		}, nil
	}
	user, err := h.us.GetUser(ctx, request.Id)
	if err != nil {
		return api.UpdateUser404JSONResponse{
			Code:    404,
			Message: "User not found",
		}, nil
	}

	return api.UpdateUser200JSONResponse{
		Id:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (h *UserHandler) DeleteUser(
	ctx context.Context,
	request api.DeleteUserRequestObject,
) (api.DeleteUserResponseObject, error) {

	err := h.us.DeleteUser(ctx, request.Id)

	if errors.Is(err, services.ErrUserNotFound) {
		return api.DeleteUser404JSONResponse{
			Code:    404,
			Message: "User not found",
		}, nil

	}

	if err != nil {
		return api.DeleteUser500JSONResponse{
			Code:    500,
			Message: "Internal error",
		}, err
	}

	return api.DeleteUser204Response{}, nil

}
