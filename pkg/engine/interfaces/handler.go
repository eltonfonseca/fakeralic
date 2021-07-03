package interfaces

type Handler interface {
	Execute() (Handler, error)
}
