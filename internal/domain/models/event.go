package models

import (
	"time"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/cns"
)

type (
	Event struct {
		ID        int64              `json:"id" schema:"id"`
		UsrID     int64              `json:"usr_id" schema:"usr_id"`
		StartTime time.Time          `json:"start_time"`
		EndTime   time.Time          `json:"end_time"`
		Goal      int                `json:"goal" schema:"goal"`
		Condition cns.EventCondition `json:"condition" schema:"condition"`
		StatusID  cns.EventStatus    `json:"status_id" schema:"status_id"`
		Winner    bool               `json:"winner" schema:"winner"`
		Active    bool               `json:"active" schema:"active"`
	}
	EventCU struct {
		UsrID     *int64              `json:"usr_id" schema:"usr_id"`
		StartTime *time.Time          `json:"start_time" schema:"start_time"`
		EndTime   *time.Time          `json:"end_time" schema:"end_time"`
		Goal      *int                `json:"goal" schema:"goal"`
		Condition *cns.EventCondition `json:"condition" schema:"condition"`
		StatusID  *cns.EventStatus    `json:"status_id" schema:"status_id"`
	}
	EventGetPars struct {
		ID    *int64 `json:"id" schema:"id"`
		UsrID *int64 `json:"usr_id" schema:"usr_id"`
	}
	EventListPars struct {
		PaginationParams
		OnlyCount bool `json:"only_count" schema:"only_count"`

		IDs       *[]int64            `json:"ids" schema:"ids"`
		UsrIDs    *[]int64            `json:"usr_ids" schema:"usr_ids"`
		StatusIDs *[]cns.EventStatus  `json:"status_ids" schema:"status_ids"`
		Condition *cns.EventCondition `json:"condition" schema:"condition"`

		EventGetPars
	}
)
