package repository

// counterfeiter:generate . SubFlowRepository

type SubFlowRepository interface {
	CreateOne() error
	UpdateOne() error
	DeleteOne() error
}
