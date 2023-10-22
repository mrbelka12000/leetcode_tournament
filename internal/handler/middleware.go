package handler

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
)

func (h *Handler) getCookie(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(consts.CookieName)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), consts.CKey, cookie.Value)
		next(w, r.WithContext(ctx))
		return
	}
}

func (h *Handler) limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		limiter, err := h.limiter.GetVisitor(ip)
		if err != nil || limiter.Allow() == false {
			h.limiter.Block(ip)
			fmt.Println(err)
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
