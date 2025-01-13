package repository

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// counterfeiter:generate . EdgeRepository

type EdgeRepository interface {
	CreateOne() error
	UpdateOne() error
	DeleteOne() error
}
