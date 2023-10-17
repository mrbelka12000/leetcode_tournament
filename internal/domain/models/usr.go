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
		ID       int64         `json:"ID,omitempty" schema:"ID,omitempty"`
		Name     string        `json:"name,omitempty" schema:"name,omitempty"`
		Username string        `json:"username,omitempty" schema:"username,omitempty"`
		Email    string        `json:"email,omitempty" schema:"email,omitempty"`
		Password string        `json:"password,omitempty" schema:"password,omitempty"`
		Group    *string       `json:"group,omitempty" schema:"group,omitempty"`
		StatusID cns.UsrStatus `json:"statusID,omitempty" schema:"statusID,omitempty"`
		TypeID   cns.UsrType   `json:"typeID,omitempty" schema:"typeID,omitempty"`
	}

	// UsrCU
	UsrCU struct {
		Name     *string        `json:"name,omitempty" schema:"name,omitempty"`
		Username *string        `json:"username,omitempty" schema:"username,omitempty"`
		Email    *string        `json:"email,omitempty" schema:"email,omitempty"`
		Password *string        `json:"password,omitempty" schema:"password,omitempty"`
		Group    *string        `json:"group,omitempty" schema:"group,omitempty"`
		StatusID *cns.UsrStatus `json:"statusID,omitempty" schema:"statusID,omitempty"`
		TypeID   *cns.UsrType   `json:"typeID,omitempty" schema:"typeID,omitempty"`
	}

	// UsrListPars
	UsrListPars struct {
		PaginationParams

		UsrGetPars

		IDs       *[]int64
		StatusIDs *[]cns.UsrStatus
		TypeIDs   *[]cns.UsrType

		OnlyCount bool
	}

	// UsrGetPars
	UsrGetPars struct {
		ID            *int64
		Username      *string
		Email         *string
		StatusID      *cns.UsrStatus
		TypeID        *cns.UsrType
		Group         *string
		UsernameEmail *string
	}
)
