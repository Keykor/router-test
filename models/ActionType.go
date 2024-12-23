package models

type ActionType string

const (
	ActionCollect ActionType = "collect"
	ActionDeliver ActionType = "deliver"
	ActionWait    ActionType = "wait"
)
