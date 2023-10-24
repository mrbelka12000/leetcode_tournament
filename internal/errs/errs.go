package errs

type Error string

func (e Error) Error() string {
	return string(e)
}

const (

	// Service errors

	ErrPermissionDenied  = Error("permission_denied")
	ErrRateLimiterWorked = Error("rate_limiter_worked")
	ErrNotFound          = Error("not_found")
	// Usr errors

	ErrUsrIDNotFound     = Error("usr_id_not_found")
	ErrEmailNotFound     = Error("email_not_found")
	ErrInvalidEmail      = Error("invalid_email")
	ErrEmailExists       = Error("email_exists")
	ErrUsernameNotFound  = Error("username_not_found")
	ErrInvalidUsername   = Error("invalid_username")
	ErrUsernameExists    = Error("username_exists")
	ErrNameNotFound      = Error("name_not_found")
	ErrInvalidName       = Error("invalid_name")
	ErrPasswordNotFound  = Error("password_not_found")
	ErrInvalidPassword   = Error("invalid_password")
	ErrPasswordDontMatch = Error("password_dont_match")
	ErrUsrStatusNotFound = Error("usr_not_found")
	ErrInvalidUsrStatus  = Error("invalid_status")
	ErrUsrTypeNotFound   = Error("type_not_found")
	ErrInvalidUsrType    = Error("invalid_type")
	ErrUsrNotFound       = Error("usr_not_found")

	// Event errors

	ErrEventIDNotFound        = Error("event_id_not_found")
	ErrStartTimeNotFound      = Error("start_time_not_found")
	ErrInvalidStartTime       = Error("invalid_start_time")
	ErrEndTimeNotFound        = Error("end_time_not_found")
	ErrInvalidEndTime         = Error("invalid_end_time")
	ErrGoalNotFound           = Error("goal_not_found")
	ErrInvalidGoal            = Error("invalid_goal")
	ErrEventConditionNotFound = Error("condition_not_found")
	ErrInvalidEventCondition  = Error("invalid_condition")
	ErrEventStatusNotFound    = Error("status_not_found")
	ErrInvalidEventStatus     = Error("invalid_status")
	ErrEventNotFound          = Error("event_not_found")

	// UsrEvent errors

	ErrUsrEventNotFound = Error("usr_event_not_found")
	ErrActiveNotFound   = Error("active_not_found")

	// Tournament errors

	ErrTournamentIDNotFound     = Error("event_id_not_found")
	ErrTournamentStatusNotFound = Error("status_not_found")
	ErrInvalidTournamentStatus  = Error("invalid_status")
	ErrTournamentNotFound       = Error("tournament_not_found")
	ErrTournamentCostNotFound   = Error("cost_not_found")
	ErrInvalidTournamentCost    = Error("invalid_cost")
	ErrTournamentPrizeNotFound  = Error("prize_pool_not_found")
	ErrInvalidTournamentPrize   = Error("invalid_prize_pool")

	ErrUsrTournamentNotFound = Error("usr_event_not_found")
)
