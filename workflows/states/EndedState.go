package states

import (
	"errors"
	"journey/models"
)

type EndedState struct{}

func (s *EndedState) Name() string {
	return "Ended"
}

func (s *EndedState) AddAction(actionToAdd models.Action, actions *[]models.Action) (JourneyState, error) {
	if actionToAdd.ID == "" {
		return s, errors.New("action ID cannot be empty")
	}

	for _, action := range *actions {
		if action.ID == actionToAdd.ID {
			return s, errors.New("action already exists in the list")
		}
	}

	*actions = append(*actions, actionToAdd)
	return s, nil
}

func (s *EndedState) RemoveAction(actionToRemove models.Action, actions *[]models.Action) (JourneyState, error) {
	if actionToRemove.ID == "" {
		return s, errors.New("action ID cannot be empty")
	}

	for i, action := range *actions {
		if action.ID == actionToRemove.ID {
			*actions = append((*actions)[:i], (*actions)[i+1:]...)
			return s, nil
		}
	}

	return s, errors.New("action not found in the list")
}

func (s *EndedState) AssignDriver(possibleDriver models.Driver, workflowDriver *models.Driver) (JourneyState, error) {
	if possibleDriver.ID == "" {
		return nil, errors.New("driver ID cannot be empty")
	}

	// check if not in other journey
	*workflowDriver = possibleDriver
	return s, nil
}

func (s *EndedState) StartJourney() (JourneyState, error) {
	return &StartedState{}, nil
}

func (s *EndedState) EndJourney(reasonToEnd models.EndReason, workflowEndReason *models.EndReason) (JourneyState, error) {
	return nil, errors.New("cannot end an ended journey")
}
