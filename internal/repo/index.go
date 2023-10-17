package repo

import (
	"database/sql"

	"github.com/mrbelka12000/leetcode_tournament/internal/interfaces"
	eventservice "github.com/mrbelka12000/leetcode_tournament/internal/service/event"
	scoreservice "github.com/mrbelka12000/leetcode_tournament/internal/service/score"
	usrservice "github.com/mrbelka12000/leetcode_tournament/internal/service/usr"
)

type Repo struct {
	Usr           usrservice.Repo
	Score         scoreservice.Repo
	Event         eventservice.Repo
	Tournament    interfaces.Tournament
	UsrEvent      interfaces.UsrEvent
	UsrTournament interfaces.UsrTournament
}

func New(db *sql.DB) *Repo {
	return &Repo{
		Usr:           newUsr(db),
		Score:         newScore(db),
		Event:         newEvent(db),
		Tournament:    newTournament(db),
		UsrEvent:      newUsrEvent(db),
		UsrTournament: newUsrTournament(db),
	}
}
