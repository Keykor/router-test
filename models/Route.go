package models

import "time"

type Route struct {
	id                int
	estimatedTime     time.Time
	estimatedDistance float64
	waypoints         []Waypoint
}
