package states

import (
	"errors"
	"journey/models"
)

type CreatedState struct{}

func (s *CreatedState) Name() string {
	return "Created"
}

func (s *CreatedState) AddAction(actionToAdd models.Action, actions *[]models.Action) (JourneyState, error) {
	if actionToAdd.ID == "" {
		return s, errors.New("action ID cannot be empty")
	}

	for _, action := range *actions {
		if action.ID == actionToAdd.ID {
			return s, errors.New("action already exists in the list")
		}
	}

	*actions = append(*actions, actionToAdd)
	// calculate route
	return s, nil
}

func (s *CreatedState) RemoveAction(actionToRemove models.Action, actions *[]models.Action) (JourneyState, error) {
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

func (s *CreatedState) AssignDriver(possibleDriver models.Driver, workflowDriver *models.Driver) (JourneyState, error) {
	if possibleDriver.ID == "" {
		return nil, errors.New("driver ID cannot be empty")
	}

	// check if driver is not in other journey
	*workflowDriver = possibleDriver
	return &AssignedState{}, nil
}

func (s *CreatedState) StartJourney() (JourneyState, error) {
	return nil, errors.New("cannot start journey in Created state")
}

func (s *CreatedState) EndJourney(reasonToEnd models.EndReason, workflowEndReason *models.EndReason) (JourneyState, error) {
	if reasonToEnd == models.ReasonCancelled {
		workflowEndReason = &reasonToEnd
		return &EndedState{}, nil
	}
	return nil, errors.New("cannot end journey in Created state unless the reason is 'cancel'")
}
