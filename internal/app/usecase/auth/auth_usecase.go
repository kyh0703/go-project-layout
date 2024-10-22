package auth

import (
	"context"

	"github.com/kyh0703/go-project-layout/internal/app/usecase/auth/dto"
)

type authUsecase struct{}

func (u *authUsecase) SignIn(ctx context.Context, req *dto.SignIn) (*dto.Token, error) {
	return nil, nil
}

func (u *authUsecase) SignUp(ctx context.Context, req *dto.SignUp) (*dto.Token, error) {
	return nil, nil
}

func (u *authUsecase) RefreshToken(ctx context.Context, req *dto.Refresh) (*dto.Token, error) {
	return nil, nil
}

func (u *authUsecase) SignOut(ctx context.Context) error {
	return nil
}
