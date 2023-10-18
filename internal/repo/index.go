package repo

import (
	"database/sql"

	eventservice "github.com/mrbelka12000/leetcode_tournament/internal/service/event"
	scoreservice "github.com/mrbelka12000/leetcode_tournament/internal/service/score"
	tournamentservice "github.com/mrbelka12000/leetcode_tournament/internal/service/tournament"
	usrservice "github.com/mrbelka12000/leetcode_tournament/internal/service/usr"
	usreventservice "github.com/mrbelka12000/leetcode_tournament/internal/service/usrevent"
	usrtournamentservice "github.com/mrbelka12000/leetcode_tournament/internal/service/usrtournament"
)

type Repo struct {
	Usr           usrservice.Repo
	Score         scoreservice.Repo
	Event         eventservice.Repo
	Tournament    tournamentservice.Repo
	UsrEvent      usreventservice.Repo
	UsrTournament usrtournamentservice.Repo
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
