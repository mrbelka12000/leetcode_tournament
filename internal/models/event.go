package models

import (
	"time"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
)

type (
	Event struct {
		ID              int64                 `json:"id" schema:"id"`
		UsrID           int64                 `json:"usr_id" schema:"usr_id"`
		StartTime       time.Time             `json:"start_time" schema:"start_time"`
		EndTime         time.Time             `json:"end_time" schema:"end_time"`
		Goal            int                   `json:"goal" schema:"goal"`
		Condition       consts.EventCondition `json:"condition" schema:"condition"`
		StatusID        consts.EventStatus    `json:"status_id" schema:"status_id"`
		IsParticipating bool                  `json:"is_participating" schema:"is_participating"`
	}
	EventCU struct {
		UsrID     *int64                 `json:"-" schema:"-"`
		StartTime *time.Time             `json:"start_time" schema:"start_time"`
		EndTime   *time.Time             `json:"end_time" schema:"end_time"`
		Goal      *int                   `json:"goal" schema:"goal"`
		Condition *consts.EventCondition `json:"condition" schema:"condition"`
		StatusID  *consts.EventStatus    `json:"-" schema:"-"`
	}
	EventGetPars struct {
		ID    *int64 `json:"id" schema:"id"`
		UsrID *int64 `json:"usr_id" schema:"usr_id"`
	}
	EventListPars struct {
		PaginationParams
		OnlyCount bool `json:"only_count" schema:"only_count"`

		IDs       *[]int64               `json:"ids" schema:"ids"`
		UsrIDs    *[]int64               `json:"usr_ids" schema:"usr_ids"`
		StatusIDs *[]consts.EventStatus  `json:"status_ids" schema:"status_ids"`
		Condition *consts.EventCondition `json:"condition" schema:"condition"`

		EventGetPars
	}
)
