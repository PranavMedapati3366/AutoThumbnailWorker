package main

import (
	"autThumbnails/activities"
	"autThumbnails/workflow"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "autoThumbnails1", worker.Options{})

	w.RegisterWorkflow(workflow.AutoThumbnailsCreationWorkflow)
	w.RegisterActivity(activities.AutoThumbnailsCreationActivity)

	log.Println("worker started running ")

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start Worker", err)
	}
}
