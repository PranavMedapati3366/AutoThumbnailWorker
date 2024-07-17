package workflow

import (
	"autThumbnails/activities"
	"log"
	"time"

	"go.temporal.io/sdk/workflow"
)

func AutoThumbnailsCreationWorkflow(ctx workflow.Context, videoPath string) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 5, // Set the timeout duration according to your needs
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	log.Printf("Workflow is started with video path: %s", videoPath) // Ensure this log message is present

	outputDir := "/Users/pranavreddy/Desktop/autoThumbnailsTemporal/output"

	err := workflow.ExecuteActivity(ctx, activities.AutoThumbnailsCreationActivity, videoPath, outputDir).Get(ctx, nil)

	if err != nil {
		return err
	}

	log.Printf("Workflow execution completed successfully")
	return nil
}
