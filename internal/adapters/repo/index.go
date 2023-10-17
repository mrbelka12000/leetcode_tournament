package repo

import (
	"database/sql"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/interfaces"
)

type Repo struct {
	Usr        interfaces.UsrRepo
	Score      interfaces.Score
	Event      interfaces.Event
	Tournament interfaces.Tournament
}

func New(db *sql.DB) *Repo {
	return &Repo{
		Usr:        newUsr(db),
		Score:      newScore(db),
		Event:      newEvent(db),
		Tournament: newTournament(db),
	}
}
