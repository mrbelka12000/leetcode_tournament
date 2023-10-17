package models

type Score struct {
	ID        int64  `json:"id,omitempty"`
	UsrID     int64  `json:"usr_id,omitempty"`
	Current   Points `json:"current"`
	Footprint Points `json:"footprint"`
	Active    bool   `json:"active"`
}

type ScoreCU struct {
	UsrID     *int64  `json:"usr_id,omitempty"`
	Current   *Points `json:"current"`
	Footprint *Points `json:"footprint"`
	Active    *bool   `json:"active"`
}

type ScoreGetPars struct {
	ID    *int64
	UsrID *int64
}

type ScoreListPars struct {
	PaginationParams
	OnlyCount bool

	IDs    *[]int64
	UsrIDs *[]int64

	ScoreGetPars
}
