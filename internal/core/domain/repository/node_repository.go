package repository

type NodeRepository interface {
	CreateOne() error
	UpdateOne() error
	DeleteOne() error
}
