package repository

type authRepository struct{}

func NewAuthRepository() *authRepository {
	return &authRepository{}
}
