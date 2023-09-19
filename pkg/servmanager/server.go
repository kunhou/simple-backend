package servmanager

// Server is transport server.
type Server interface {
	Start() error
	Shutdown() error
}
