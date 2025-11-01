package middleware

import "net/http"

// this is a signature func(http.Handler) http.Handler of Middleware
type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	mngr := Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
	return &mngr
	
}
