package models

import "time"

type (
	Session struct {
		ID       int64
		UsrID    int64
		Token    string
		ExpireAt time.Time
	}
)
