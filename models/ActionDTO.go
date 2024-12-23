package models

type ActionDTO struct {
	ActionType  ActionType  `json:"action_type"`
	Geolocation Geolocation `json:"geolocation"`
	ActionID    string      `json:"action_id"`
}
