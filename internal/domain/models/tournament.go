package models

import (
	"time"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/cns"
)

type (
	Tournament struct {
		ID        int64
		UsrID     int64
		StartTime time.Time
		EndTime   time.Time
		Goal      int
		PrizePool int
		Cost      int
		StatusID  cns.EventStatus
	}
	TournamentCU struct {
		UsrID     *int64
		StartTime *time.Time
		EndTime   *time.Time
		Goal      *int
		Cost      *int
		PrizePool *int
		StatusID  *cns.TournamentStatus
	}
	TournamentGetPars struct {
		ID    *int64
		UsrID *int64
	}
	TournamentListPars struct {
		PaginationParams
		OnlyCount bool

		IDs       *[]int64
		UsrIDs    *[]int64
		StatusIDs *[]cns.TournamentStatus

		TournamentGetPars
	}
)
