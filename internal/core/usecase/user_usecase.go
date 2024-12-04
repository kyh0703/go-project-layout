package user

import (
	"context"

	"github.com/kyh0703/go-project-layout/internal/app/service/auth/dto"
)

type UserUsecase interface {
	Create(ctx context.Context, email, username, password string)
	Get(ctx context.Context, email string)
	GetById(ctx context.Context, id int)
	Update(ctx context.Context, id int, dto *dto.SignUp)
	Remove(ctx context.Context, id int) error
}
