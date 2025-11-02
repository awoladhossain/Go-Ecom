package middleware

import (
	"net/http"
)

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

// * this is called receiver function
func (mngr *Manager) With(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		h := next
		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]
			h = middleware(h)
		}
		return h

	}
}
