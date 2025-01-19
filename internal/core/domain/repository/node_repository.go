package repository

import (
	"context"

	"github.com/kyh0703/layout/internal/core/domain/model"
)

// counterfeiter:generate . NodeRepository

type NodeRepository interface {
	CreateOne(ctx context.Context, args model.CreateNodeParams) error
	FindOne(ctx context.Context, id string) (*model.Node, error)
	GetList(ctx context.Context, subFlowId int32) ([]*model.Node, error)
	UpdateOne(ctx context.Context, args model.UpdateNodeParams) (*model.Node, error)
	DeleteOne(ctx context.Context, id string) error
}
