package models

type Score struct {
	ID        int64  `json:"id,omitempty"`
	UsrID     int64  `json:"usr_id,omitempty"`
	Current   Points `json:"current"`
	Footprint Points `json:"footprint"`
	Active    bool   `json:"active"`
}

type ScoreCU struct {
	UsrID     *int64  `json:"usr_id" schema:"usr_id"`
	Current   *Points `json:"current" schema:"current"`
	Footprint *Points `json:"footprint" schema:"footprint"`
	Active    *bool   `json:"active" schema:"active"`
}

type ScoreGetPars struct {
	ID    *int64 `schema:"id"`
	UsrID *int64 `schema:"usr_id"`
}

type ScoreListPars struct {
	PaginationParams
	OnlyCount bool `json:"only_count" schema:"only_count"`

	IDs    *[]int64 `json:"ids" schema:"ids"`
	UsrIDs *[]int64 `json:"usr_ids" schema:"usr_ids"`

	ScoreGetPars `json:"score_get_pars"`
}
