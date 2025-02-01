package repository

import (
	"context"

	"github.com/kyh0703/layout/internal/core/domain/model"
	"github.com/kyh0703/layout/internal/core/domain/repository"
)

type postRepository struct {
	queries *model.Queries
}

func NewPostRepository(
	queries *model.Queries,
) repository.PostRepository {
	return &postRepository{
		queries: queries,
	}
}

func (s *postRepository) CreateOne(ctx context.Context, param model.CreatePostParams) (model.Post, error) {
	return s.queries.CreatePost(ctx, param)
}

func (s *postRepository) FindOne(ctx context.Context, id int64) (model.Post, error) {
	return s.queries.GetPost(ctx, id)
}

func (s *postRepository) GetList(ctx context.Context, userID int64) ([]model.Post, error) {
	return s.queries.ListPostByUserID(ctx, userID)
}

func (s *postRepository) UpdateOne(ctx context.Context, param model.UpdatePostParams) error {
	return s.queries.UpdatePost(ctx, param)
}

func (s *postRepository) DeleteOne(ctx context.Context, id int64) error {
	return s.queries.DeletePosts(ctx, id)
}
