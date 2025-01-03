package workflows

import (
	"go.temporal.io/sdk/temporal"
	"journey/models"
	"journey/workflows/states"

	"go.temporal.io/sdk/workflow"
)

func JourneyWorkflow(ctx workflow.Context) error {
	var currentState states.JourneyState = &states.CreatedState{}
	var endReason models.EndReason
	var assignedDriver models.Driver
	actions := []models.Action{}
	doneChan := workflow.NewChannel(ctx)

	// Search Attributes
	// command: tctl admin cluster add-search-attributes --name JourneyState --type Keyword
	journeyStateKey := temporal.NewSearchAttributeKeyKeyword("JourneyState")
	err := workflow.UpsertTypedSearchAttributes(ctx, journeyStateKey.ValueSet(currentState.Name()))
	if err != nil {
		return err
	}

	// Query Handlers
	workflow.SetQueryHandler(ctx, "currentState", func() (string, error) {
		return currentState.Name(), nil
	})
	workflow.SetQueryHandler(ctx, "actions", func() ([]models.Action, error) {
		return actions, nil
	})
	workflow.SetQueryHandler(ctx, "driver", func() (models.Driver, error) {
		return assignedDriver, nil
	})
	workflow.SetQueryHandler(ctx, "endReason", func() (models.EndReason, error) {
		return endReason, nil
	})

	// Update Handlers
	workflow.SetUpdateHandler(ctx, "addAction", func(action models.Action) (string, error) {
		nextState, err := currentState.AddAction(action, &actions)
		handleStateTransition(ctx, &currentState, nextState, err)
		return currentState.Name(), err
	})

	workflow.SetUpdateHandler(ctx, "removeAction", func(action models.Action) (string, error) {
		nextState, err := currentState.RemoveAction(action, &actions)
		handleStateTransition(ctx, &currentState, nextState, err)
		return currentState.Name(), err
	})

	workflow.SetUpdateHandler(ctx, "assignDriver", func(driver models.Driver) (string, error) {
		nextState, err := currentState.AssignDriver(driver, &assignedDriver)
		handleStateTransition(ctx, &currentState, nextState, err)
		return currentState.Name(), err
	})

	workflow.SetUpdateHandler(ctx, "startJourney", func() (string, error) {
		nextState, err := currentState.StartJourney()
		handleStateTransition(ctx, &currentState, nextState, err)
		return currentState.Name(), err
	})

	workflow.SetUpdateHandler(ctx, "endJourney", func(reason models.EndReason) (string, error) {
		nextState, err := currentState.EndJourney(reason, &endReason)
		handleStateTransition(ctx, &currentState, nextState, err)
		if nextState.Name() == "Finalized" {
			doneChan.Send(ctx, true)
		}
		return currentState.Name(), err
	})

	// Wait for the journey to end
	doneChan.Receive(ctx, nil)
	workflow.GetLogger(ctx).Info("Journey ended", "endReason", endReason)
	return nil
}

func handleStateTransition(ctx workflow.Context, currentState *states.JourneyState, nextState states.JourneyState, err error) {
	if err != nil {
		workflow.GetLogger(ctx).Error("error processing state transition", "error", err)
		return
	}
	if nextState != nil {
		*currentState = nextState

		// Update search attributes
		journeyStateKey := temporal.NewSearchAttributeKeyKeyword("JourneyState")
		updateErr := workflow.UpsertTypedSearchAttributes(ctx, journeyStateKey.ValueSet((*currentState).Name()))
		if updateErr != nil {
			workflow.GetLogger(ctx).Error("error updating search attributes", "error", updateErr)
			return
		}

		workflow.GetLogger(ctx).Info("State updated", "newState", nextState.Name())
	}
}
