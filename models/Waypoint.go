package models

type Waypoint struct {
	id          int
	order       int
	Geolocation Geolocation
	actions     []Action
}
