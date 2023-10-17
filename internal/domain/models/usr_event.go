package models

type (
	UsrEvent struct {
		ID      int64
		UsrID   int64
		EventID int64
		Active  bool
		Winner  bool
	}
	UsrEventCU struct {
		UsrID   *int64
		EventID *int64
		Active  *bool
		Winner  *bool
	}
	UsrEventGetPars struct {
		ID    *int64
		UsrID *int64
	}
	UsrEventListPars struct {
		PaginationParams
		OnlyCount bool

		IDs    *[]int64
		UsrIDs *[]int64
		Active *bool
		Winner *bool

		UsrEventGetPars
	}
)
