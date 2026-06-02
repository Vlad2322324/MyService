package services

import (
	"context"
	"errors"
	"myservice/internal/domain"
	"myservice/internal/repositories"
)

type UserService struct {
	ur *repositories.UserRepo
}

var ErrUserNotFound = errors.New("Target User Not Found")

func NewUserService(ur *repositories.UserRepo) (*UserService, error) {
	return &UserService{ur: ur}, nil
}

func (us *UserService) CreateUser(ctx context.Context, name string) (*domain.Users, error) {
	// FIXME
	return us.ur.InsertUser(ctx, name)
}

func (ur *UserService) DeleteUser(ctx context.Context, id int) error {
	rowsAffected, err := ur.ur.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {

		return ErrUserNotFound
	}

	return nil
}

func (ur *UserService) UpdateUser(ctx context.Context, id int, name string) error {
	return ur.ur.UpdateUser(ctx, id, name)
}

func (ur *UserService) GetUser(ctx context.Context, id int) (domain.Users, error) {
	return ur.ur.GetUserById(ctx, id)
}
