package session

import (
	"context"
	"fmt"
	"time"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type Session struct {
	sessionRepo Repo
}

func New(sessionRepo Repo) *Session {
	return &Session{
		sessionRepo: sessionRepo,
	}
}

func (s *Session) Build(ctx context.Context, obj models.Session) error {
	obj.ExpireAt = time.Now().AddDate(0, 1, 0)

	err := s.sessionRepo.Create(ctx, obj)
	if err != nil {
		return fmt.Errorf("session create: %w", err)
	}

	return nil
}

func (s *Session) Delete(ctx context.Context, token string) error {
	return s.sessionRepo.Delete(ctx, token)
}

func (s *Session) Get(ctx context.Context, token string) (models.Session, error) {
	return s.sessionRepo.Get(ctx, token)
}
