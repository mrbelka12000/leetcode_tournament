package core

import (
	"context"
	"fmt"

	"leetcode_tournament/internal/adapters/repo"
	"leetcode_tournament/internal/domain/interfaces"
	"leetcode_tournament/internal/domain/models"
)

type Core struct {
	Usr           *usr
	leetCodeStats interfaces.LeetCodeStats
}

func NewCore(repo *repo.Repo, leetCodeStats interfaces.LeetCodeStats) *Core {
	return &Core{
		Usr:           newUsr(repo.Usr, leetCodeStats),
		leetCodeStats: leetCodeStats,
	}
}

func (c *Core) MainPage(ctx context.Context) (*models.Index, error) {
	index := &models.Index{}

	users, _, err := c.Usr.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("usr list: %w", err)
	}

	index.Users = users

	for i := 0; i < len(index.Users); i++ {
		index.Users[i].SolvedProblems, err = c.leetCodeStats.Stats(ctx, index.Users[i].Nickname)
		if err != nil {
			return nil, fmt.Errorf("get leetcode solved problems: %w", err)
		}
	}

	return index, nil
}
