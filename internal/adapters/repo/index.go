package repo

import (
	"database/sql"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/interfaces"
)

type Repo struct {
	Usr       interfaces.Usr
	Footprint interfaces.Footprint
}

func New(db *sql.DB) *Repo {
	return &Repo{
		Usr:       newUsr(db),
		Footprint: newFootprint(db),
	}
}
