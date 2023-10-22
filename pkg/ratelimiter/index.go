package ratelimiter

import (
	"sync"
	"time"

	"golang.org/x/time/rate"

	"github.com/mrbelka12000/leetcode_tournament/internal/errs"
)

const (
	ipBLockTime = 30 * time.Second
	flushTime   = 15 * time.Second
)

type (
	Limiter struct {
		visitors map[string]*visitor
		mu       sync.RWMutex
		r, b     int
	}
	visitor struct {
		limiter  *rate.Limiter
		lastSeen time.Time
		block    bool
	}
)

func New(r, b int) *Limiter {
	l := &Limiter{
		visitors: make(map[string]*visitor),
		r:        r,
		b:        b,
	}

	go l.cleanupVisitors()

	return l
}

func (l *Limiter) GetVisitor(ip string) (*rate.Limiter, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	v, exists := l.visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(rate.Limit(l.r), l.b)
		l.visitors[ip] = &visitor{limiter: limiter, lastSeen: time.Now()}
		return limiter, nil
	}

	if v.block {
		return nil, errs.ErrRateLimiterWorked
	}

	v.lastSeen = time.Now()
	return v.limiter, nil
}

func (l *Limiter) cleanupVisitors() {
	for {
		time.Sleep(flushTime)

		l.mu.Lock()
		for ip, v := range l.visitors {
			if time.Since(v.lastSeen) > ipBLockTime {
				delete(l.visitors, ip)
			}
		}
		l.mu.Unlock()
	}
}

func (l *Limiter) Block(ip string) {
	l.visitors[ip].block = true
	l.visitors[ip].lastSeen = time.Now().Add(ipBLockTime)
}
