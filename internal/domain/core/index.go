package core

import (
	"context"
	"fmt"

	"github.com/mrbelka12000/leetcode_tournament/internal/adapters/repo"
	"github.com/mrbelka12000/leetcode_tournament/internal/domain/interfaces"
	"github.com/mrbelka12000/leetcode_tournament/internal/domain/models"
)

type Core struct {
	Usr           *usr
	FootPrint     *footprint
	leetCodeStats interfaces.LeetCodeStats
}

func New(repo *repo.Repo, leetCodeStats interfaces.LeetCodeStats) *Core {
	return &Core{
		Usr:           newUsr(repo.Usr, leetCodeStats),
		FootPrint:     newFootprint(repo.Footprint),
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
