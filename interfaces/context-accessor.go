package interfaces

type ContextAccessor interface {
	Secrets() map[string]string
}
