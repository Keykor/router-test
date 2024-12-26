package workflows

import (
	"errors"
	"fmt"
	"router/models"
	"router/workflows/states"

	"go.temporal.io/sdk/workflow"
)

type Signal struct {
	Method  string
	Payload interface{}
}

func JourneyWorkflow(ctx workflow.Context, journeyID string) error {
	var currentState states.JourneyState = &states.AssignedState{}
	var endReason models.EndReason
	var driver models.Driver
	actions := []models.Action{}

	for currentState.Name() != "Finalized" {
		var signal Signal
		workflow.GetSignalChannel(ctx, "signal").Receive(ctx, &signal)

		var nextState states.JourneyState
		var err error

		switch signal.Method {
		case "addAction":
			action, ok := signal.Payload.(models.Action)
			if !ok {
				err = errors.New("invalid payload for addAction")
			} else {
				nextState, err = currentState.AddAction(action, &actions)
			}
		case "removeAction":
			action, ok := signal.Payload.(models.Action)
			if !ok {
				err = errors.New("invalid payload for removeAction")
			} else {
				nextState, err = currentState.RemoveAction(action, &actions)
			}
		case "assignDriver":
			possibleDriver, ok := signal.Payload.(models.Driver)
			if !ok {
				err = errors.New("invalid payload for assignDriver")
			} else {
				nextState, err = currentState.AssignDriver(possibleDriver, &driver)
			}
		case "startJourney":
			nextState, err = currentState.StartJourney()
		case "endJourney":
			possibleReason, ok := signal.Payload.(models.EndReason)
			if !ok {
				err = errors.New("invalid payload for endJourney")
			} else {
				nextState, err = currentState.EndJourney(possibleReason, &endReason)
			}
		default:
			err = fmt.Errorf("unknown method: %s", signal.Method)
		}

		if err != nil {
			workflow.GetLogger(ctx).Error("error managing the signal", "signal", signal.Method, "error", err)
			continue
		}

		if nextState != nil {
			currentState = nextState
		}
		workflow.GetLogger(ctx).Info("State updated", "newState", currentState.Name())
	}

	workflow.GetLogger(ctx).Info("Journey ended", "journeyID", journeyID)
	return nil
}
