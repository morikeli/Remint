package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	inputDir  = "videos"
	outputDir = "videos/h264"
)

var videoExtensions = map[string]bool{
	".mp4":  true,
	".mkv":  true,
	".avi":  true,
	".mov":  true,
	".webm": true,
	".flv":  true,
}

func convertToH264(inputFile, outputFile string) error {
	// ffmpeg must be installed and available on the system PATH.
	args := []string{
		"-y", // Overwrite the output file if it already exists.
		"-i", inputFile,
		"-c:v", "libx264",
		"-preset", "medium",
		"-crf", "23",
		"-pix_fmt", "yuv420p",
		"-c:a", "aac",
		"-movflags", "+faststart",
		outputFile,
	}

	cmd := exec.Command("ffmpeg", args...)
	// Keep ffmpeg's progress and error messages visible in the terminal.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func run() error {
	// Check for ffmpeg before starting the batch.
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return fmt.Errorf("ffmpeg was not found; install it and make sure it is on your PATH: %w", err)
	}

	entries, err := os.ReadDir(inputDir)
	if err != nil {
		return fmt.Errorf("could not read input directory %q: %w", inputDir, err)
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("could not create output directory %q: %w", outputDir, err)
	}

	var videos []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if videoExtensions[ext] {
			videos = append(videos, filepath.Join(inputDir, entry.Name()))
		}
	}

	fmt.Printf("Found %d video(s).\n", len(videos))

	failed := 0
	for index, inputFile := range videos {
		baseName := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))
		outputFile := filepath.Join(outputDir, baseName+".mp4")

		fmt.Printf("[%d/%d] Converting: %s\n", index+1, len(videos), filepath.Base(inputFile))

		if err := convertToH264(inputFile, outputFile); err != nil {
			failed++
			fmt.Fprintf(os.Stderr, "  ✘ Failed: %s: %v\n", filepath.Base(inputFile), err)
			continue
		}

		fmt.Printf(" ✔ Saved: %s\n", outputFile)
	}

	fmt.Printf("Done. %d succeeded, %d failed.\n", len(videos)-failed, failed)
	if failed > 0 {
		return fmt.Errorf("%d video(s) failed to convert", failed)
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
