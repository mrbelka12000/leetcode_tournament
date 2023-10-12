package repo

import (
	"database/sql"

	"leetcode_tournament/internal/domain/interfaces"
)

type Repo struct {
	Usr interfaces.User
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{Usr: newUsr(db)}
}
