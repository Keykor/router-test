package main

import (
	"journey/server"
	"journey/worker"
	"log"
)

func main() {
	go func() {
		log.Println("Starting Temporal Worker...")
		worker.StartWorker()
	}()

	log.Println("Starting Gin Server on port 8080...")
	r := server.NewServer()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting Gin server: %v", err)
	}
}
