package ptr

import "github.com/mrbelka12000/leetcode_tournament/internal/consts"

func UsrStatusPointer(v consts.UsrStatus) *consts.UsrStatus {
	return &v
}

func UsrTypePointer(v consts.UsrType) *consts.UsrType {
	return &v
}

func EventStatusPointer(v consts.EventStatus) *consts.EventStatus {
	return &v
}

func TournamentStatusPointer(v consts.TournamentStatus) *consts.TournamentStatus {
	return &v
}
