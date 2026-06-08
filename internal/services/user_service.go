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

var ErrUserNotFound = errors.New("target User Not Found")

func NewUserService(ur *repositories.UserRepo) (*UserService, error) {
	return &UserService{ur: ur}, nil
}

func (us *UserService) CreateUser(ctx context.Context, name string) (*domain.Users, error) {
	// FIXME
	return us.ur.InsertUser(ctx, name)
}

func (us *UserService) DeleteUser(ctx context.Context, id int) error {
	rowsAffected, err := us.ur.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {

		return ErrUserNotFound
	}

	return nil
}

func (us *UserService) UpdateUser(ctx context.Context, id int, name string) error {
	return us.ur.UpdateUser(ctx, id, name)
}

func (us *UserService) GetUser(ctx context.Context, id int) (domain.Users, error) {
	return us.ur.GetUserById(ctx, id)
}
