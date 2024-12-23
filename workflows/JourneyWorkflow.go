package workflows

import (
	"errors"
	"fmt"
	"router/models"
	"router/workflows/states"

	"go.temporal.io/sdk/workflow"
)

func JourneyWorkflow(ctx workflow.Context, journeyID string) error {
	var state states.JourneyState = &
}
