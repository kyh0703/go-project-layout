package repository

import "go.uber.org/fx"

var RepositoryModule = fx.Provide(
	"repository",
	NewAuthRepository,
	NewEdgeRepository,
	NewAuthRepository,
	NewUserRepository,
)
