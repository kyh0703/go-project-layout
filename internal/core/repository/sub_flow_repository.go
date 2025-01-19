package repository

import (
	"github.com/kyh0703/layout/internal/core/domain/repository"
)

type SubFlowRepository struct{}

func NewSubFlowRepository() repository.NodeRepository {
	return &NodeRepository{}
}
