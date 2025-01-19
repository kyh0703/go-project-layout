package repository

import (
	"context"

	"github.com/kyh0703/layout/internal/core/domain/model"
	"github.com/kyh0703/layout/internal/core/domain/repository"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (u *userRepository) CreateOne(ctx context.Context, args model.CreateUserParams) error {
	panic("unimplemented")
}

func (u *userRepository) FindOne(ctx context.Context, id int32) (*model.User, error) {
	panic("unimplemented")
}

func (u *userRepository) UpdateOne(ctx context.Context, args model.UpdateUserParams) (*model.User, error) {
	panic("unimplemented")
}

func (u *userRepository) DeleteOne(ctx context.Context, id int32) error {
	panic("unimplemented")
}
