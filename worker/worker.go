package worker

import (
	"journey/workflows"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func StartWorker() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("Error connecting to Temporal: %v", err)
	}
	defer c.Close()

	w := worker.New(c, "journey-task-queue", worker.Options{})
	w.RegisterWorkflow(workflows.JourneyWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalf("Error running worker: %v", err)
	}
}
