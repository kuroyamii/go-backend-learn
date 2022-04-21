package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ContentTypeMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func CorsMiddlerware(whitelistedUrls map[string]bool) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,POST,PUT,DELETE,PATCH")
			rw.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-CSRF-Token, Authorization")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			requestOriginUrl := r.Header.Get("Origin")
			log.Printf("INFO CorsMiddleware: received request from %s %v", requestOriginUrl, whitelistedUrls[requestOriginUrl])
			if whitelistedUrls[requestOriginUrl] {
				rw.Header().Set("Access-Control-Allow-Origin", requestOriginUrl)
			}

			if r.Method != http.MethodOptions {
				handler.ServeHTTP(rw, r)
				return
			}
			rw.Write([]byte("Hello Mother Father"))
		})
	}
}
