package states

import (
	"router/models"
)

type JourneyState interface {
	addAction(actionDTO models.ActionDTO) (bool, error)
	removeAction(actionDTO models.ActionDTO) (bool, error)
	cancelJourney(id string) (bool, error)
	assignDriver(driver models.Driver) (bool, error)
	startJourney(id string) (bool, error)
	endJourney(id string, reason models.EndReason) (bool, error)
	fusionJourney(id string) (bool, error)
}
