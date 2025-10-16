package global_router

import "net/http"

func GlobalRouters(mux *http.ServeMux) http.Handler {
	handleGlobalRoutesFoRequest := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		mux.ServeHTTP(w, r)
	}
	allHandlers := http.HandlerFunc(handleGlobalRoutesFoRequest)
	return allHandlers
}
