package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	//* retrun http.HandlerFunc()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		diff := time.Since(start)
		log.Println(r.Method, r.URL.Path, diff)
	})
	return handler
}
