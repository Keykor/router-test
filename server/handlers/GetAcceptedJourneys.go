package handlers

import (
	"go.temporal.io/api/workflowservice/v1"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
)

func GetAcceptedJourneysHandler(c *gin.Context) {
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Temporal"})
		return
	}
	defer temporalClient.Close()

	request := &workflowservice.ListWorkflowExecutionsRequest{Query: `JourneyState="Accepted"`}

	resp, err := temporalClient.ListWorkflow(c, request)
	if err != nil {
		return
	}

	var workflows []map[string]string
	for _, exec := range resp.Executions {
		workflows = append(workflows, map[string]string{
			"workflowID": exec.Execution.WorkflowId,
			"runID":      exec.Execution.RunId,
		})
	}

	c.JSON(http.StatusOK, workflows)
}
