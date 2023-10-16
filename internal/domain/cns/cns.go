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
