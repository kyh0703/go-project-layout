package repository

import (
	"context"

	"github.com/kyh0703/layout/internal/core/domain/model"
)

// counterfeiter:generate . SubFlowRepository

type SubFlowRepository interface {
	CreateOne(ctx context.Context, args model.CreateSubFlowParams) error
	FindOne(ctx context.Context, id int32) (*model.SubFlow, error)
	GetList(ctx context.Context, flowId int32) ([]*model.SubFlow, error)
	UpdateOne(ctx context.Context, args model.UpdateSubFlowParams) (*model.SubFlow, error)
	DeleteOne(ctx context.Context, id int32) error
}
