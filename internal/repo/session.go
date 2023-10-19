package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type session struct {
	db *sql.DB
}

func newSession(db *sql.DB) *session {
	return &session{db: db}
}

func (s *session) Create(ctx context.Context, obj models.Session) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO session
			(usr_id, token, expire_at)
		VALUES 
			($1,$2,$3)
`, obj.UsrID, obj.Token, obj.ExpireAt)
	if err != nil {
		return fmt.Errorf("create session in db : %w", err)
	}

	return nil
}

func (s *session) Get(ctx context.Context, token string) (models.Session, error) {
	var obj models.Session

	err := s.db.QueryRowContext(ctx, `
	SELECT
	    id, usr_id, token, expire_at
	FROM session
	WHERE token = $1
`, token).Scan(&obj.ID, &obj.UsrID, &obj.Token, &obj.ExpireAt)
	if err != nil {
		return models.Session{}, fmt.Errorf("get session from db: %w", err)
	}

	return obj, nil
}

func (s *session) Delete(ctx context.Context, token string) error {
	_, err := s.db.ExecContext(ctx, `
	DELETE 
		FROM session
	WHERE token =$1
`, token)
	if err != nil {
		return fmt.Errorf("delete session from db: %w", err)
	}

	return nil
}
