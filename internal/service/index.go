package service

import (
	"github.com/mrbelka12000/leetcode_tournament/internal/client/leetcode"
	"github.com/mrbelka12000/leetcode_tournament/internal/repo"
	eventservice "github.com/mrbelka12000/leetcode_tournament/internal/service/event"
	scoreservice "github.com/mrbelka12000/leetcode_tournament/internal/service/score"
	usrservice "github.com/mrbelka12000/leetcode_tournament/internal/service/usr"
)

type Core struct {
	Usr   *usrservice.Usr
	Event *eventservice.Event
	Score *scoreservice.Score
}

func New(repo *repo.Repo, lc *leetcode.LeetCode) *Core {
	return &Core{
		Usr:   usrservice.New(repo.Usr, lc),
		Score: scoreservice.New(repo.Score),
		Event: eventservice.New(repo.Event),
	}
}
