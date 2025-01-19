package repository

import (
	"context"

	"github.com/kyh0703/layout/internal/core/domain/model"
	"github.com/kyh0703/layout/internal/core/domain/repository"
)

type EdgeRepository struct{}

func NewEdgeRepository() repository.EdgeRepository {
	return &EdgeRepository{}
}

func (e *EdgeRepository) CreateOne(ctx context.Context, args model.CreateEdgeParams) error {
	panic("unimplemented")
}

func (e *EdgeRepository) FindOne(ctx context.Context, id string) (*model.Edge, error) {
	panic("unimplemented")
}

func (e *EdgeRepository) GetList(ctx context.Context, subFlowId int32) ([]*model.Edge, error) {
	panic("unimplemented")
}

func (e *EdgeRepository) UpdateOne(ctx context.Context, args model.UpdateEdgeParams) (*model.Edge, error) {
	panic("unimplemented")
}

func (e *EdgeRepository) DeleteOne(ctx context.Context, id string) error {
	panic("unimplemented")
}
