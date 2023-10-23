package handler

import "golang.org/x/time/rate"

type rateLimit interface {
	GetVisitor(ip string) (*rate.Limiter, error)
	Block(ip string)
}
