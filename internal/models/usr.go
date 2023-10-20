package models

import (
	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
)

type (
	// Usr
	Usr struct {
		ID       int64            `json:"id" schema:"id"`
		Name     string           `json:"name" schema:"name"`
		Username string           `json:"username" schema:"username"`
		Email    string           `json:"email" schema:"email"`
		Password string           `json:"password" schema:"password"`
		Group    *string          `json:"group" schema:"group"`
		StatusID consts.UsrStatus `json:"status_id" schema:"status_id"`
		TypeID   consts.UsrType   `json:"type_id" schema:"type_id"`
		Score    Score            `json:"score" schema:"score"`
	}

	// UsrCU
	UsrCU struct {
		Name     *string           `json:"name" schema:"name"`
		Username *string           `json:"username" schema:"username"`
		Email    *string           `json:"email" schema:"email"`
		Password *string           `json:"password" schema:"password"`
		Group    *string           `json:"group" schema:"group"`
		StatusID *consts.UsrStatus `json:"status_id" schema:"status_id"`
		TypeID   *consts.UsrType   `json:"type_id" schema:"type_id"`
	}

	// UsrListPars
	UsrListPars struct {
		PaginationParams

		UsrGetPars

		IDs       *[]int64            `json:"ids" schema:"ids"`
		StatusIDs *[]consts.UsrStatus `json:"status_ids" schema:"status_ids"`
		TypeIDs   *[]consts.UsrType   `json:"type_ids" schema:"type_ids"`

		OnlyCount bool `json:"onlyCount" schema:"onlyCount"`
	}

	// UsrGetPars
	UsrGetPars struct {
		ID            *int64            `json:"id" schema:"id"`
		Username      *string           `json:"username" schema:"username"`
		Email         *string           `json:"email" schema:"email"`
		StatusID      *consts.UsrStatus `json:"status_id" schema:"status_id"`
		TypeID        *consts.UsrType   `json:"type_id" schema:"type_id"`
		Group         *string           `json:"group" schema:"group"`
		UsernameEmail *string           `json:"username_email" schema:"username_email"`
	}

	UsrLogin struct {
		UsernameEmail string `json:"username_email" schema:"username_email"`
		Password      string `json:"password" schema:"password"`
	}
)
