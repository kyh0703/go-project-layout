package presenter

type Server interface {
	Listen() error
	Shutdown() error
}
