package models

type (
	UsrTournament struct {
		ID           int64 `json:"id" schema:"id"`
		UsrID        int64 `json:"usr_id" schema:"usr_id"`
		TournamentID int64 `json:"tournament_id" schema:"tournament_id"`
		Active       bool  `json:"active" schema:"active"`
		Winner       bool  `json:"winner" schema:"winner"`
	}
	UsrTournamentCU struct {
		UsrID        *int64 `json:"usr_id" schema:"usr_id"`
		TournamentID *int64 `json:"tournament_id" schema:"tournament_id"`
		Active       *bool  `json:"active" schema:"active"`
		Winner       *bool  `json:"winner" schema:"winner"`
	}
	UsrTournamentGetPars struct {
		ID    *int64 `json:"id" schema:"id"`
		UsrID *int64 `json:"usr_id" schema:"usr_id"`
	}
	UsrTournamentListPars struct {
		PaginationParams
		OnlyCount bool `json:"only_count" schema:"only_count"`

		IDs    *[]int64 `json:"ids" schema:"ids"`
		UsrIDs *[]int64 `json:"usr_ids" schema:"usr_ids"`
		Active *bool    `json:"active" schema:"active"`
		Winner *bool    `json:"winner" schema:"winner"`

		UsrTournamentGetPars
	}
)
