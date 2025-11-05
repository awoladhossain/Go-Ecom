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

// func (mngr *Manager) Use(middleware ...Middleware) *Manager {
// 	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middleware...)
// 	return mngr
// }

func (mngr *Manager) Use(middleware ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middleware...)
}

// * this is called receiver function
func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {

	h := next
	// for i := len(middlewares) - 1; i >= 0; i-- {
	// 	middleware := middlewares[i]
	// 	h = middleware(h)
	// }
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	for _, globalMiddleware := range mngr.globalMiddlewares {
		h = globalMiddleware(h)
	}

	return h

}
