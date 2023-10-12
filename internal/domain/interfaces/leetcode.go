package interfaces

import (
	"context"

	"leetcode_tournament/internal/adapters/leetcode"
)

type LeetCodeStats interface {
	Stats(ctx context.Context, nickname string) (resp leetcode.LCGetProblemsSolvedResp, err error)
}
