package repo

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"leetcode_tournament/internal/domain/models"
)

type usr struct {
	db *sql.DB
}

func newUsr(db *sql.DB) *usr {
	return &usr{db: db}
}

func (u *usr) Create(ctx context.Context, obj *models.Usr) (string, error) {
	_, err := u.db.ExecContext(ctx, `
	INSERT INTO users (name, nickname,secret) 
		VALUES (?, ?, ?)`,
		obj.Name, obj.Nickname, obj.Secret)
	if err != nil {
		return "", fmt.Errorf("create usr: %w", err)
	}

	return obj.Secret, nil
}

func (u *usr) Update(ctx context.Context, obj *models.Usr) error {
	_, err := u.db.ExecContext(ctx, `
		UPDATE users 
		SET name=?, nickname=? 
		WHERE  secret =?`,
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
		FROM users 
		WHERE id=?`, id).Scan(
		&user.ID, &user.Name, &user.Nickname)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *usr) List(ctx context.Context) ([]*models.Usr, int64, error) {
	rows, err := u.db.QueryContext(ctx, `
		SELECT id, name, nickname 
		FROM users`)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []*models.Usr
	for rows.Next() {
		var user models.Usr
		if err := rows.Scan(&user.ID, &user.Name, &user.Nickname); err != nil {
			return nil, 0, err
		}
		users = append(users, &user)
	}

	var count int64
	err = u.db.QueryRowContext(ctx, `
		SELECT COUNT(*) 
		FROM users`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (u *usr) GetByField(ctx context.Context, field, value string) (*models.Usr, error) {
	query := fmt.Sprintf("SELECT id, name, nickname from users where %v = ?", field)

	var user models.Usr
	err := u.db.QueryRowContext(ctx, query, value).Scan(
		&user.ID, &user.Name, &user.Nickname)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
