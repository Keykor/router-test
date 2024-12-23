package models

import "time"

type Journey struct {
	id               int
	creationDate     time.Time
	scheduledDate    time.Time
	calculatedRoutes []Route
	activeRoute      Route
}
