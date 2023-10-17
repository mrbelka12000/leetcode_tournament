package service

import (
	"github.com/mrbelka12000/leetcode_tournament/internal/interfaces"
	"github.com/mrbelka12000/leetcode_tournament/internal/repo"
)

type Core struct {
	Usr           *usr
	leetCodeStats interfaces.LeetCodeStats
}

func New(repo *repo.Repo, leetCodeStats interfaces.LeetCodeStats) *Core {
	return &Core{
		Usr:           newUsr(repo.Usr, leetCodeStats),
		leetCodeStats: leetCodeStats,
	}
}
