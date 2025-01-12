package repository

type SubflowRepository interface {
	CreateOne() error
	UpdateOne() error
	DeleteOne() error
}
