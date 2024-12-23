package models

type Action struct {
	id          int
	description string
	actionType  ActionType
	shipment    Shipment
}
