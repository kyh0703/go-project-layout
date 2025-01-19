package repository

import (
	"context"

	"github.com/kyh0703/layout/internal/core/domain/model"
	"github.com/kyh0703/layout/internal/core/domain/repository"
)

type NodeRepository struct{}

func NewNodeRepository() repository.NodeRepository {
	return &NodeRepository{}
}

func (n *NodeRepository) CreateOne(ctx context.Context, args model.CreateNodeParams) error {
	panic("unimplemented")
}

func (n *NodeRepository) FindOne(ctx context.Context, id string) (*model.Node, error) {
	panic("unimplemented")
}

func (n *NodeRepository) GetList(ctx context.Context, subFlowId int32) ([]*model.Node, error) {
	panic("unimplemented")
}

func (n *NodeRepository) UpdateOne(ctx context.Context, args model.UpdateNodeParams) (*model.Node, error) {
	panic("unimplemented")
}

func (n *NodeRepository) DeleteOne(ctx context.Context, id string) error {
	panic("unimplemented")
}
