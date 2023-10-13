package models

type Footprint struct {
	ID     int64  `json:"id,omitempty"`
	UsrID  int64  `json:"usr_id,omitempty"`
	Easy   uint64 `json:"easy,omitempty"`
	Medium uint64 `json:"medium,omitempty"`
	Hard   uint64 `json:"hard,omitempty"`
	Total  uint64 `json:"total,omitempty"`
	Active bool   `json:"active"`
}

type FootprintGetPars struct {
	ID    *int64
	UsrID *int64
}

type FootprintListPars struct {
	PaginationParams
	OnlyCount bool

	FootprintGetPars
}
