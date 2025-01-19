package repository

import (
	"context"

	"github.com/kyh0703/layout/internal/core/domain/model"
)

// counterfeiter:generate . UserRepository

type UserRepository interface {
	CreateOne(ctx context.Context, args model.CreateUserParams) error
	FindOne(ctx context.Context, id int32) (*model.User, error)
	UpdateOne(ctx context.Context, args model.UpdateUserParams) (*model.User, error)
	DeleteOne(ctx context.Context, id int32) error
}
