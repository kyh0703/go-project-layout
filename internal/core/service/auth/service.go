package auth

import (
	"context"

	"github.com/kyh0703/layout/internal/core/service/auth/dto"
)

type Service interface {
	SignIn(ctx context.Context, req *dto.SignIn) (*dto.Token, error)
	SignUp(ctx context.Context, req *dto.SignUp) (*dto.Token, error)
	RefreshToken(ctx context.Context, req *dto.Refresh) (*dto.Token, error)
	SignOut(ctx context.Context) error
}
