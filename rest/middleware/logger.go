package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = r.RemoteAddr
		}

		next.ServeHTTP(w, r)

		fmt.Printf(
			"User IP: %s | Method: %s | URL: %s | Duration: %v\n",
			ip,
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
