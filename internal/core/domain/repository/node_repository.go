package repository

// counterfeiter:generate . NodeRepository

type NodeRepository interface {
	CreateOne() error
	UpdateOne() error
	DeleteOne() error
}
