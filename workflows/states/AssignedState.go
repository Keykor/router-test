package states

import (
	"errors"
	"journey/models"
)

type AssignedState struct{}

func (s *AssignedState) Name() string {
	return "Assigned"
}

func (s *AssignedState) AddAction(actionToAdd models.Action, actions *[]models.Action) (JourneyState, error) {
	if actionToAdd.ID == "" {
		return s, errors.New("action ID cannot be empty")
	}

	for _, action := range *actions {
		if action.ID == actionToAdd.ID {
			return s, errors.New("action already exists in the list")
		}
	}
	// calculate route
	*actions = append(*actions, actionToAdd)
	return s, nil
}

func (s *AssignedState) RemoveAction(actionToRemove models.Action, actions *[]models.Action) (JourneyState, error) {
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

func (s *AssignedState) AssignDriver(possibleDriver models.Driver, workflowDriver *models.Driver) (JourneyState, error) {
	if possibleDriver.ID == "" {
		return nil, errors.New("driver ID cannot be empty")
	}

	// check if not in other journey

	*workflowDriver = possibleDriver
	return s, nil
}

func (s *AssignedState) StartJourney() (JourneyState, error) {
	// check if driver and actions?
	return &StartedState{}, nil
}

func (s *AssignedState) EndJourney(reasonToEnd models.EndReason, workflowEndReason *models.EndReason) (JourneyState, error) {
	if reasonToEnd == models.ReasonCancelled {
		workflowEndReason = &reasonToEnd
		return &EndedState{}, nil
	}
	return nil, errors.New("cannot end journey in Assigned State unless the reason is 'cancel'")
}
