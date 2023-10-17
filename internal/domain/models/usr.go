package models

import (
	"github.com/mrbelka12000/leetcode_tournament/internal/adapters/leetcode"
	"github.com/mrbelka12000/leetcode_tournament/internal/domain/cns"
)

type UsrOld struct {
	ID             int64                            `json:"id"`
	Name           string                           `json:"name"`
	Nickname       string                           `json:"nickname"`
	Secret         string                           `json:"-"`
	SolvedProblems leetcode.LCGetProblemsSolvedResp `json:"solved_problems"`
	Footprint      *Score                           `json:"footprint"`
}

type (
	// Usr
	Usr struct {
		ID       int64         `json:"ID" schema:"ID"`
		Name     string        `json:"name" schema:"name"`
		Username string        `json:"username" schema:"username"`
		Email    string        `json:"email" schema:"email"`
		Password string        `json:"password" schema:"password"`
		Group    *string       `json:"group" schema:"group"`
		StatusID cns.UsrStatus `json:"status_id" schema:"status_id"`
		TypeID   cns.UsrType   `json:"type_id" schema:"type_id"`
	}

	// UsrCU
	UsrCU struct {
		Name     *string        `json:"name" schema:"name"`
		Username *string        `json:"username" schema:"username"`
		Email    *string        `json:"email" schema:"email"`
		Password *string        `json:"password" schema:"password"`
		Group    *string        `json:"group" schema:"group"`
		StatusID *cns.UsrStatus `json:"status_id" schema:"status_id"`
		TypeID   *cns.UsrType   `json:"type_id" schema:"type_id"`
	}

	// UsrListPars
	UsrListPars struct {
		PaginationParams

		UsrGetPars

		IDs       *[]int64         `json:"ids" schema:"ids"`
		StatusIDs *[]cns.UsrStatus `json:"status_ids" schema:"status_ids"`
		TypeIDs   *[]cns.UsrType   `json:"type_ids" schema:"type_ids"`

		OnlyCount bool `json:"onlyCount" schema:"onlyCount"`
	}

	// UsrGetPars
	UsrGetPars struct {
		ID            *int64         `json:"id" schema:"id"`
		Username      *string        `json:"username" schema:"username"`
		Email         *string        `json:"email" schema:"email"`
		StatusID      *cns.UsrStatus `json:"status_id" schema:"status_id"`
		TypeID        *cns.UsrType   `json:"type_id" schema:"type_id"`
		Group         *string        `json:"group" schema:"group"`
		UsernameEmail *string        `json:"username_email" schema:"username_email"`
	}
)
