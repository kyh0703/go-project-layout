package repository

import "github.com/kyh0703/layout/internal/core/domain/repository"

type EdgeRepository struct{}

func NewEdgeRepository() repository.EdgeRepository {
	return &EdgeRepository{}
}

func (e *EdgeRepository) CreateOne() error {
	panic("unimplemented")
}

func (e *EdgeRepository) DeleteOne() error {
	panic("unimplemented")
}

func (e *EdgeRepository) UpdateOne() error {
	panic("unimplemented")
}
