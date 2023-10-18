package service

import (
	"github.com/mrbelka12000/leetcode_tournament/internal/client/leetcode"
	"github.com/mrbelka12000/leetcode_tournament/internal/repo"

	eventservice "github.com/mrbelka12000/leetcode_tournament/internal/service/event"
	scoreservice "github.com/mrbelka12000/leetcode_tournament/internal/service/score"
	tournamentservice "github.com/mrbelka12000/leetcode_tournament/internal/service/tournament"
	usrservice "github.com/mrbelka12000/leetcode_tournament/internal/service/usr"
	usreventservice "github.com/mrbelka12000/leetcode_tournament/internal/service/usrevent"
	usrtournamentservice "github.com/mrbelka12000/leetcode_tournament/internal/service/usrtournament"
)

type Core struct {
	Usr           *usrservice.Usr
	Event         *eventservice.Event
	Score         *scoreservice.Score
	Tournament    *tournamentservice.Tournament
	UsrEvent      *usreventservice.UsrEvent
	UsrTournament *usrtournamentservice.UsrTournament
}

func New(repo *repo.Repo, lc *leetcode.LeetCode) *Core {
	return &Core{
		Usr:           usrservice.New(repo.Usr, lc),
		Score:         scoreservice.New(repo.Score),
		Event:         eventservice.New(repo.Event),
		Tournament:    tournamentservice.New(repo.Tournament),
		UsrEvent:      usreventservice.New(repo.UsrEvent),
		UsrTournament: usrtournamentservice.New(repo.UsrTournament),
	}
}
