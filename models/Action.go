package models

type Action struct {
	ID          string
	description string
	actionType  ActionType
	shipment    Shipment
}
