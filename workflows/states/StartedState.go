package states

import (
	"errors"
	"journey/models"
)

type StartedState struct{}

func (s *StartedState) Name() string {
	return "Started"
}

func (s *StartedState) AddAction(actionToAdd models.Action, actions *[]models.Action) (JourneyState, error) {
	return s, errors.New("cannot add action to a journey in Started state")
}

func (s *StartedState) RemoveAction(actionToRemove models.Action, actions *[]models.Action) (JourneyState, error) {
	if actionToRemove.ID == "" {
		return s, errors.New("action ID cannot be empty")
	}

	for i, action := range *actions {
		if action.ID == actionToRemove.ID {
			*actions = append((*actions)[:i], (*actions)[i+1:]...)
			return s, nil
		}
	}

	// calculate route
	return s, errors.New("action not found in the list")
}

func (s *StartedState) AssignDriver(possibleDriver models.Driver, workflowDriver *models.Driver) (JourneyState, error) {
	return s, errors.New("cannot assign driver to a journey in Started state")
}

func (s *StartedState) StartJourney() (JourneyState, error) {
	return s, errors.New("cannot start journey in Started state")
}

func (s *StartedState) EndJourney(reasonToEnd models.EndReason, workflowEndReason *models.EndReason) (JourneyState, error) {
	// check if can end journey with actions remaining
	workflowEndReason = &reasonToEnd
	return &EndedState{}, nil
}
