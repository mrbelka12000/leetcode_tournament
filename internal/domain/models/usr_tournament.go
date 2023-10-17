package models

type (
	UsrTournament struct {
		ID           int64
		UsrID        int64
		TournamentID int64
		Active       bool
		Winner       bool
	}
	UsrTournamentCU struct {
		UsrID        *int64
		TournamentID *int64
		Active       *bool
		Winner       *bool
	}
	UsrTournamentGetPars struct {
		ID    *int64
		UsrID *int64
	}
	UsrTournamentListPars struct {
		PaginationParams
		OnlyCount bool

		IDs    *[]int64
		UsrIDs *[]int64
		Active *bool
		Winner *bool

		UsrTournamentGetPars
	}
)
