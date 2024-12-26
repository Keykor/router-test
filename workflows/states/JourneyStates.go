package states

import (
	"router/models"
)

type JourneyState interface {
	AddAction(actionToAdd models.Action, actions *[]models.Action) (JourneyState, error)
	RemoveAction(actionToRemove models.Action, actions *[]models.Action) (JourneyState, error)
	AssignDriver(possibleDriver models.Driver, workflowDriver *models.Driver) (JourneyState, error)
	StartJourney() (JourneyState, error)
	EndJourney(reasonToEnd models.EndReason, workflowEndReason *models.EndReason) (JourneyState, error)
	Name() string
}
