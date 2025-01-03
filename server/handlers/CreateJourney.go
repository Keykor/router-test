package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
	"net/http"
)

func CreateJourneyHandler(c *gin.Context) {
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Temporal"})
		return
	}
	defer temporalClient.Close()

	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: "journey-task-queue",
	}

	we, err := temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, "JourneyWorkflow")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start workflow", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Workflow started successfully",
		"workflowID": we.GetID(),
		"runID":      we.GetRunID(),
	})
}
