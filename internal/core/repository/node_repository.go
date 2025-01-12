package repository

import "github.com/kyh0703/layout/internal/core/domain/repository"

type NodeRepository struct{}

func NewNodeRepository() repository.NodeRepository {
	return &NodeRepository{}
}

// CreateOne implements repository.NodeRepository.
func (n *NodeRepository) CreateOne() error {
	panic("unimplemented")
}

// DeleteOne implements repository.NodeRepository.
func (n *NodeRepository) DeleteOne() error {
	panic("unimplemented")
}

// UpdateOne implements repository.NodeRepository.
func (n *NodeRepository) UpdateOne() error {
	panic("unimplemented")
}
