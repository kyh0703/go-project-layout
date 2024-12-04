package auth

import (
	"context"

	"github.com/kyh0703/go-project-layout/internal/app/service/auth/dto"
)

type authService struct{}

func (s *authService) SignIn(ctx context.Context, req *dto.SignIn) (*dto.Token, error) {
	return nil, nil
}

func (s *authService) SignUp(ctx context.Context, req *dto.SignUp) (*dto.Token, error) {
	return nil, nil
}

func (s *authService) RefreshToken(ctx context.Context, req *dto.Refresh) (*dto.Token, error) {
	return nil, nil
}

func (s *authService) SignOut(ctx context.Context) error {
	return nil
}
