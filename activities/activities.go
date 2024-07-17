package activities

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func AutoThumbnailsCreationActivity(ctx context.Context, videoPath string, outputDir string) error {
	// Check if ffmpeg is available
	_, err := exec.LookPath("ffmpeg")
	if err != nil {
		return fmt.Errorf("ffmpeg not found: %v", err)
	}

	fmt.Printf("Starting activity with video path: %s", videoPath)

	videoPath = "/Users/pranavreddy/Desktop/autoThumbnailsTemporal/demo.mp4"

	// Ensure the video file exists
	if _, err := os.Stat(videoPath); os.IsNotExist(err) {
		return fmt.Errorf("video file does not exist: %s", videoPath)
	}

	// Ensure the output directory exists
	err = exec.Command("mkdir", "-p", outputDir).Run()
	if err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	fmt.Printf("Starting activity with video path: %s and output directory: %s\n", videoPath, outputDir)

	intervalSeconds := 10
	videoFileName := filepath.Base(videoPath)
	outputPattern := filepath.Join(outputDir, strings.TrimSuffix(videoFileName, filepath.Ext(videoFileName))+"_thumbnail_%03d.png")

	// ffmpeg command: extract frames at regular intervals
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vf", fmt.Sprintf("fps=1/%d", intervalSeconds), outputPattern)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to generate thumbnails: %v: %s", err, stderr.String())
	}

	fmt.Printf("Command output: %s\n", out.String())
	return nil
}
