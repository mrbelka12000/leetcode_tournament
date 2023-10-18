package models

import (
	"time"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
)

type (
	Tournament struct {
		ID        int64              `json:"id" schema:"id"`
		UsrID     int64              `json:"usr_id" schema:"usr_id"`
		StartTime time.Time          `json:"start_time" schema:"start_time"`
		EndTime   time.Time          `json:"end_time" schema:"end_time"`
		Goal      int                `json:"goal" schema:"goal"`
		PrizePool int                `json:"prize_pool" schema:"prize_pool"`
		Cost      int                `json:"cost" schema:"cost"`
		StatusID  consts.EventStatus `json:"status_id" schema:"status_id"`
	}
	TournamentCU struct {
		UsrID     *int64                   `json:"usr_id" schema:"usr_id"`
		StartTime *time.Time               `json:"start_time" schema:"start_time"`
		EndTime   *time.Time               `json:"end_time" schema:"end_time"`
		Goal      *int                     `json:"goal" schema:"goal"`
		Cost      *int                     `json:"cost" schema:"cost"`
		PrizePool *int                     `json:"prize_pool" schema:"prize_pool"`
		StatusID  *consts.TournamentStatus `json:"status_id" schema:"status_id"`
	}
	TournamentGetPars struct {
		ID    *int64 `json:"id" schema:"id"`
		UsrID *int64 `json:"usr_id" schema:"usr_id"`
	}
	TournamentListPars struct {
		PaginationParams
		OnlyCount bool `json:"onlyCount" schema:"onlyCount"`

		IDs       *[]int64                   `json:"ids" schema:"ids"`
		UsrIDs    *[]int64                   `json:"usr_ids" schema:"usr_ids"`
		StatusIDs *[]consts.TournamentStatus `json:"status_ids" schema:"status_ids"`

		TournamentGetPars
	}
)
