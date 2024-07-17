package main

import (
	"autThumbnails/workflow"
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "autoThumbnailsGenerationNew",
		TaskQueue: "autoThumbnails1",
	}

	videoPath := "/Users/pranavreddy/Desktop/autoThumbnailsTemporal/demo.mp4"
	log.Printf("Starting start.go with video path: %s", videoPath) // Ensure this log message appears

	we, err := c.ExecuteWorkflow(context.Background(), options, workflow.AutoThumbnailsCreationWorkflow, videoPath)
	if err != nil {
		log.Fatalln("Unable to start Workflow", err)
	}

	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())
}
