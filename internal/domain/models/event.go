package models

import (
	"time"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/cns"
)

type (
	Event struct {
		ID        int64
		UsrID     int64
		StartTime time.Time
		EndTime   time.Time
		Goal      int
		Condition cns.EventCondition
		StatusID  cns.EventStatus
	}

	EventCU struct {
		UsrID     *int64
		StartTime *time.Time
		EndTime   *time.Time
		Goal      *int
		Condition *cns.EventCondition
		StatusID  *cns.EventStatus
	}

	EventGetPars struct {
		ID    *int64
		UsrID *int64
	}
	EventListPars struct {
		PaginationParams
		OnlyCount bool

		IDs       *[]int64
		UsrIDs    *[]int64
		StatusIDs *[]cns.EventStatus
		Condition *cns.EventCondition

		EventGetPars
	}
)
