package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/layout/internal/core/domain/model"
	"github.com/kyh0703/layout/internal/core/domain/repository"
	"github.com/kyh0703/layout/internal/core/dto/auth"
	"github.com/kyh0703/layout/internal/pkg/password"
)

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(
	userRepository repository.UserRepository,
) Service {
	return &authService{
		userRepository: userRepository,
	}
}

func (a *authService) SignUp(ctx context.Context, req *auth.SignUp) (*model.User, error) {
	existUser, err := a.userRepository.FindOneByEmail(ctx, req.Email)
	if err != nil {
		return nil, fiber.NewError(500, err.Error())
	}

	if existUser.ID != 0 {
		return nil, fiber.NewError(409, "email already exists")
	}

	if req.Password != req.PasswordConfirm {
		return nil, fiber.NewError(400, "password and password confirm do not match")
	}

	hash, err := password.Hashed(req.Password)
	if err != nil {
		return nil, fiber.NewError(500, err.Error())
	}

	createdUser, err := a.userRepository.CreateOne(ctx, model.CreateUserParams{
		Email:    req.Email,
		Password: hash,
		Name:     req.Name,
		Bio:      req.Bio,
	})

	return &createdUser, nil
}

func (a *authService) RefreshToken(ctx context.Context, req *auth.Refresh) (*auth.Token, error) {
	panic("unimplemented")
}

func (a *authService) SignIn(ctx context.Context, req *auth.SignIn) (*auth.Token, error) {
	panic("unimplemented")
}

func (a *authService) SignOut(ctx context.Context) error {
	panic("unimplemented")
}
