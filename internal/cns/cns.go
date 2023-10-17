package cns

type UsrStatus uint

const (
	UsrStatusCreated UsrStatus = iota + 1
	UsrStatusConfirmed
	UsrStatusDeleted
)

type UsrType uint

const (
	UsrTypeDeveloper UsrType = iota + 1 // игрок
	UsrTypeClient                       // заказчик турниров
	UsrTypeAdmin
)

type EventCondition string

const (
	EventConditionOnFirst      EventCondition = "on_first"       // кто первый наберет нужное количество задач, тот выйграл
	EventConditionOnMax        EventCondition = "on_max"         // кто решит больше всего задач за определнный период
	EventConditionOnTimeExceed EventCondition = "on_time_exceed" // может быть много победителей, победители определяется путем если превысил нужное количество задач
)

type EventStatus uint

const (
	EventStatusCreated EventStatus = iota + 1
	EventStatusStarted
	EventStatusCanceled

	EventStatusFinished EventStatus = 100
)

type TournamentStatus uint

const (
	TournamentStatusCreated TournamentStatus = iota + 1
	TournamentStatusStarted
	TournamentStatusCanceled

	TournamentStatusFinished TournamentStatus = 100
)
