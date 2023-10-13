package repo

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/models"
)

type usr struct {
	db *sql.DB
}

func newUsr(db *sql.DB) *usr {
	return &usr{db: db}
}

func (u *usr) Create(ctx context.Context, obj *models.Usr) (string, int64, error) {
	var id int64
	err := u.db.QueryRowContext(ctx, `
	INSERT INTO usr (name, nickname,secret) 
		VALUES ($1, $2, $3) RETURNING id`,
		obj.Name, obj.Nickname, obj.Secret).Scan(&id)
	if err != nil {
		return "", 0, fmt.Errorf("create usr: %w", err)
	}

	return obj.Secret, id, nil
}

func (u *usr) Update(ctx context.Context, obj *models.Usr) error {
	_, err := u.db.ExecContext(ctx, `
		UPDATE usr 
		SET name=$1, nickname=$2 
		WHERE  secret =$3`,
		obj.Name, obj.Nickname, obj.Secret)
	if err != nil {
		return err
	}

	return nil
}

func (u *usr) Get(ctx context.Context, id int64) (*models.Usr, error) {
	var user models.Usr
	err := u.db.QueryRowContext(ctx, `
		SELECT id, name, nickname 
		FROM usr 
		WHERE id=$1`, id).Scan(
		&user.ID, &user.Name, &user.Nickname)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *usr) List(ctx context.Context) ([]*models.Usr, int64, error) {
	rows, err := u.db.QueryContext(ctx, `
		SELECT id, name, nickname 
		FROM usr`)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var usrs []*models.Usr
	for rows.Next() {
		var user models.Usr
		if err := rows.Scan(&user.ID, &user.Name, &user.Nickname); err != nil {
			return nil, 0, err
		}
		usrs = append(usrs, &user)
	}

	var count int64
	err = u.db.QueryRowContext(ctx, `
		SELECT COUNT(*) 
		FROM usr`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return usrs, count, nil
}

func (u *usr) GetByField(ctx context.Context, field, value string) (*models.Usr, error) {
	query := fmt.Sprintf("SELECT id, name, nickname from usr where %v = $1", field)

	var user models.Usr
	err := u.db.QueryRowContext(ctx, query, value).Scan(
		&user.ID, &user.Name, &user.Nickname)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
