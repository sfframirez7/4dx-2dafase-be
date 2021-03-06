package server

import (
	"net/http"

	"../routes"
	"../services"
)

func Serve() {

	myRouter := routes.Routes()

	// http.Handle("/", services.IsLogginMiddleWare(myRouter))

	http.Handle("/", CORS(services.IsLogginMiddleWare(myRouter)))

	http.ListenAndServe(":8087", nil)
}

func CORS(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST, OPTIONS")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization, X-Requested-With, Content-Length, Accept-Encoding,Screen")
			w.WriteHeader(http.StatusOK)
			return
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
