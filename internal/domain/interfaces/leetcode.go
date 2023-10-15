package interfaces

import (
	"context"

	"github.com/mrbelka12000/leetcode_tournament/internal/adapters/leetcode"
)

type LeetCodeStats interface {
	Stats(ctx context.Context, nickname string) (resp leetcode.LCGetProblemsSolvedResp, err error)
}
