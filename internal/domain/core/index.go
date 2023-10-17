package core

import (
	"github.com/mrbelka12000/leetcode_tournament/internal/adapters/repo"
	"github.com/mrbelka12000/leetcode_tournament/internal/domain/interfaces"
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
