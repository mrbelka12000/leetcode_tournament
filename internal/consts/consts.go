package consts

const (
	TimeFormat = "2006-01-02"
)

type UsrStatus uint

const (
	UsrStatusCreated UsrStatus = iota + 1
	UsrStatusConfirmed
	UsrStatusDeleted
)

func IsValidUsrStatus(v UsrStatus) bool {
	return v == UsrStatusCreated ||
		v == UsrStatusConfirmed ||
		v == UsrStatusDeleted
}

type UsrType uint

const (
	UsrTypeDeveloper UsrType = iota + 1 // игрок
	UsrTypeClient                       // заказчик турниров
	UsrTypeAdmin
)

func IsValidUsrType(v UsrType) bool {
	return v == UsrTypeDeveloper ||
		v == UsrTypeClient ||
		v == UsrTypeAdmin
}

type EventCondition string

const (
	EventConditionOnFirst      EventCondition = "on_first"       // кто первый наберет нужное количество задач, тот выйграл
	EventConditionOnMax        EventCondition = "on_max"         // кто решит больше всего задач за определнный период
	EventConditionOnTimeExceed EventCondition = "on_time_exceed" // может быть много победителей, победители определяется путем если превысил нужное количество задач
)

func IsValidEventCondition(v EventCondition) bool {
	return v == EventConditionOnFirst ||
		v == EventConditionOnMax ||
		v == EventConditionOnTimeExceed
}

type EventStatus uint

const (
	EventStatusCreated EventStatus = iota + 1
	EventStatusStarted
	EventStatusCanceled

	EventStatusFinished EventStatus = 100
)

func IsValidEventStatus(v EventStatus) bool {
	return v == EventStatusCreated ||
		v == EventStatusStarted ||
		v == EventStatusCanceled ||
		v == EventStatusFinished
}

type TournamentStatus uint

const (
	TournamentStatusCreated TournamentStatus = iota + 1
	TournamentStatusStarted
	TournamentStatusCanceled

	TournamentStatusFinished TournamentStatus = 100
)

func IsValidTournamentStatus(v TournamentStatus) bool {
	return v == TournamentStatusCreated ||
		v == TournamentStatusStarted ||
		v == TournamentStatusCanceled ||
		v == TournamentStatusFinished
}

// COOKIE

type CookieKey string

var CKey CookieKey = "token"

const (
	CookieName = "session"
)
