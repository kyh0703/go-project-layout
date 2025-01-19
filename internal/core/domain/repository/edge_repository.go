package repository

import (
	"context"

	"github.com/kyh0703/layout/internal/core/domain/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// counterfeiter:generate . EdgeRepository

type EdgeRepository interface {
	CreateOne(ctx context.Context, args model.CreateEdgeParams) error
	FindOne(ctx context.Context, id string) (*model.Edge, error)
	GetList(ctx context.Context, subFlowId int32) ([]*model.Edge, error)
	UpdateOne(ctx context.Context, args model.UpdateEdgeParams) (*model.Edge, error)
	DeleteOne(ctx context.Context, id string) error
}
