package repository

type EdgeRepository interface {
	CreateOne() error
	UpdateOne() error
	DeleteOne() error
}
