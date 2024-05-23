package ip2geo

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"
	"time"
)

type limiter struct {
	init    time.Time
	allowed int
	counter int
	lock    sync.Mutex
}

func (l *limiter) limit() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	if time.Now().UTC().Round(time.Second).After(l.init.UTC().Add(1 * time.Second)) {
		l.init = time.Now().Round(time.Second)
		l.counter = 1
		return false
	} else if l.counter < l.allowed {
		l.counter++
		return false
	} else {
		return true
	}
}

func createLimiter(allowed int) *limiter {
	return &limiter{
		time.Now(),
		allowed,
		0,
		sync.Mutex{},
	}
}

type RateLimiter struct {
	allowed        int
	limitsByOrigin map[string]*limiter
	handler        http.Handler
}

func (rl *RateLimiter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	println("request originated from %s", host)
	// TODO: clean cache every once in a while
	limiter, ok := rl.limitsByOrigin[host]
	if !ok {
		println("missing limiter, creating a new one...")
		limiter = createLimiter(rl.allowed)
		rl.limitsByOrigin[host] = limiter
	}

	if limiter.limit() {
		println("too many requests from %s. blocking", host)
		w.WriteHeader(http.StatusTooManyRequests)
		e, _ := json.Marshal(ErrorResponseBody{"try again later"})
		w.Write(e)

	} else {
		rl.handler.ServeHTTP(w, req)
	}
}

func WithRateLimiter(allowed int, handelr http.Handler) *RateLimiter {
	return &RateLimiter{
		allowed,
		make(map[string]*limiter),
		handelr,
	}
}
