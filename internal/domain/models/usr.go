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
	Footprint      *Footprint                       `json:"footprint"`
}

type Usr struct {
	ID       int64
	Name     string
	Username string
	Email    string
	Password string
	Group    string
	StatusID cns.UsrStatus
	TypeID   cns.UsrType
}
