package middleware

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	maxReqPerSec   = 6
	blockDuration  = 5 * time.Minute
	windowDuration = 1 * time.Second
)

type clientInfo struct {
	requests int
	window   time.Time
	blocked  time.Time
}

var (
	mu      sync.Mutex
	reqData = make(map[string]*clientInfo)
)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := getIP(r)
		now := time.Now()

		mu.Lock()
		ci, exists := reqData[ip]

		if !exists {
			reqData[ip] = &clientInfo{
				requests: 1,
				window:   now,
			}
			mu.Unlock()
			next.ServeHTTP(w, r)
			return
		}

		// Block check
		if ci.blocked.After(now) {
			mu.Unlock()
			http.Error(w, "Too many requests. IP blocked.", http.StatusTooManyRequests)
			fmt.Printf("Blocked request from IP: %s\n", ip)
			return
		}

		// Reset window
		if now.Sub(ci.window) > windowDuration {
			ci.window = now
			ci.requests = 0
		}

		ci.requests++

		// Exceeded limit
		if ci.requests > maxReqPerSec {
			ci.blocked = now.Add(blockDuration)
			mu.Unlock()
			http.Error(w, "Too many requests! IP blocked for 5 minutes.", http.StatusTooManyRequests)
			return
		}

		mu.Unlock()
		next.ServeHTTP(w, r)
	})
}

func getIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
