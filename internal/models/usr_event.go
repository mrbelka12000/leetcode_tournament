package models

type (
	UsrEvent struct {
		ID      int64 `json:"id" schema:"id"`
		UsrID   int64 `json:"usr_id" schema:"usr_id"`
		EventID int64 `json:"event_id" schema:"event_id"`
		Active  bool  `json:"active" schema:"active"`
		Winner  bool  `json:"winner" schema:"winner"`
	}
	UsrEventCU struct {
		UsrID   *int64 `json:"usr_id" schema:"usr_id"`
		EventID *int64 `json:"event_id" schema:"event_id"`
		Active  *bool  `json:"active" schema:"active"`
		Winner  *bool  `json:"winner" schema:"winner"`
	}
	UsrEventGetPars struct {
		ID      *int64 `json:"id" schema:"id"`
		UsrID   *int64 `json:"usr_id" schema:"usr_id"`
		EventID *int64 `json:"event_id" schema:"event_id"`
		Active  *int64 `json:"active" schema:"active"`
		Winner  *int64 `json:"winner" schema:"winner"`
	}
	UsrEventListPars struct {
		PaginationParams
		OnlyCount bool `json:"only_count" schema:"only_count"`

		IDs    *[]int64 `json:"ids" schema:"ids"`
		UsrIDs *[]int64 `json:"usr_ids" schema:"usr_ids"`
		Active *bool    `json:"active" schema:"active"`
		Winner *bool    `json:"winner" schema:"winner"`

		UsrEventGetPars
	}
	UsrGetEventsPars struct {
		EventListPars

		Active *bool `json:"active" schema:"active"`
		Winner *bool `json:"winner" schema:"winner"`
	}
)
