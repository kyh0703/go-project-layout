package repository

import (
	"context"

	"github.com/kyh0703/layout/internal/core/domain/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . PostRepository
type PostRepository interface {
	CreateOne(ctx context.Context, arg model.CreatePostParams) (model.Post, error)
	FindOne(ctx context.Context, id int64) (model.Post, error)
	GetList(ctx context.Context, flowID int64) ([]model.Post, error)
	UpdateOne(ctx context.Context, arg model.UpdatePostParams) error
	DeleteOne(ctx context.Context, id int64) error
}
