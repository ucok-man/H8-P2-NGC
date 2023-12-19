package middleware

import (
	"net/http"
	"time"
)

func (md *Middleware) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		datetime := time.Now().Format("2006/01/02 15:04:05")
		md.Log.Printf("%s HTTP request sent to %v %v\n", datetime, r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

	})
}
