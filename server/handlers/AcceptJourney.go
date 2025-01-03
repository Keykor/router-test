package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
	"journey/models"
)

func AcceptJourneyHandler(c *gin.Context) {
	journeyID := c.Param("id")
	token := c.GetString("token")
	driverID := extractDriverIDFromToken(token)
	if driverID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid driver ID"})
		return
	}

	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Temporal"})
		return
	}
	defer temporalClient.Close()

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	updateHandle, err := temporalClient.UpdateWorkflow(ctxWithTimeout, client.UpdateWorkflowOptions{
		WorkflowID:   journeyID,
		UpdateName:   "assignDriver",
		Args:         []interface{}{models.Driver{ID: driverID}},
		WaitForStage: client.WorkflowUpdateStageAccepted,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send update", "details": err.Error()})
		return
	}

	var result string
	err = updateHandle.Get(ctxWithTimeout, &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get update result", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Driver assigned successfully",
		"state":     result,
		"driverID":  driverID,
		"journeyID": journeyID,
	})
}

// extractDriverIDFromToken is a helper function that extracts the driver ID from the token
// it's a fake implementation used for testing purposes
func extractDriverIDFromToken(token string) string {
	return "driver-123"
}
