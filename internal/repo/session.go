package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type session struct {
	db *sql.DB
}

func newSession(db *sql.DB) *session {
	return &session{db: db}
}

func (s *session) Save(ctx context.Context, obj models.Session) error {
	_, err := Exec(ctx, s.db, `
		INSERT INTO session
			(usr_id, token,type_id, expire_at)
		VALUES 
			($1,$2,$3,$4)		
`, obj.UsrID, obj.Token, obj.TypeID, obj.ExpireAt)
	if err != nil {
		return fmt.Errorf("create session in db : %w", err)
	}

	return nil
}

func (s *session) Get(ctx context.Context, pars models.SessionGetPars) (models.Session, error) {
	var filterValues []interface{}

	querySelect := `
	SELECT s.id, s.usr_id, s.token, s.type_id, s.expire_at
`
	queryFrom := ` FROM session s`
	queryWhere := ` WHERE 1 = 1`

	if pars.ID != nil {
		filterValues = append(filterValues, *pars.ID)
		queryWhere += ` AND s.id = $` + strconv.Itoa(len(filterValues))
	}

	if pars.Token != nil {
		filterValues = append(filterValues, *pars.Token)
		queryWhere += ` AND s.token = $` + strconv.Itoa(len(filterValues))
	}
	if pars.UsrID != nil {
		filterValues = append(filterValues, *pars.UsrID)
		queryWhere += ` AND s.usr_id = $` + strconv.Itoa(len(filterValues))
	}

	var obj models.Session

	row := QueryRow(ctx, s.db, querySelect+queryFrom+queryWhere, filterValues...)
	err := row.Scan(
		&obj.ID,
		&obj.UsrID,
		&obj.Token,
		&obj.TypeID,
		&obj.ExpireAt,
	)
	if err != nil {
		return models.Session{}, fmt.Errorf("get session from db: %w", err)
	}

	return obj, nil
}

func (s *session) Delete(ctx context.Context, token string) error {
	_, err := Exec(ctx, s.db, `
	DELETE 
		FROM session
	WHERE token =$1
`, token)
	if err != nil {
		return fmt.Errorf("delete session from db: %w", err)
	}

	return nil
}

func (s *session) Update(ctx context.Context, obj models.Session) error {
	_, err := Exec(ctx, s.db, `
		UPDATE session
		SET token = $1
		WHERE usr_id = $2
`, obj.Token, obj.UsrID)
	if err != nil {
		return fmt.Errorf("update session in db: %w", err)
	}

	return nil
}
