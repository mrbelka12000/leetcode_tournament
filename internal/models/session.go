package models

import (
	"time"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
)

type (
	Session struct {
		ID       int64
		UsrID    int64
		Token    string
		TypID    consts.UsrType
		ExpireAt time.Time
	}
	SessionGetPars struct {
		ID    *int64
		UsrID *int64
		Token *string
	}
)
