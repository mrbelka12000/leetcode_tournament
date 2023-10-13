package models

import "github.com/mrbelka12000/leetcode_tournament/internal/adapters/leetcode"

type Usr struct {
	ID             int64                            `json:"id"`
	Name           string                           `json:"name"`
	Nickname       string                           `json:"nickname"`
	Secret         string                           `json:"-"`
	SolvedProblems leetcode.LCGetProblemsSolvedResp `json:"solved_problems"`
	Footprint      *Footprint                       `json:"footprint"`
}
