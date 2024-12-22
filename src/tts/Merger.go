package tts

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
)

func (opts *Options) Merger(FilesPaths []string) (string, error) {
	outputFile := filepath.Join(opts.TempDir, fmt.Sprintf("%d.wav", rand.Int()))
	listFile := filepath.Join(opts.TempDir, fmt.Sprintf("%d.txt", rand.Int()))

	file, err := os.Create(listFile)
	if err != nil {
		return "", fmt.Errorf("error creating file list: %v", err)
	}
	defer file.Close()

	// Write each audio file to the list file
	for _, audioFile := range FilesPaths {
		// Ensure that each audio file exists
		if _, err := os.Stat(audioFile); err != nil {
			return "", fmt.Errorf("audio file does not exist: %v", audioFile)
		}
		_, err := fmt.Fprintf(file, "file '%s'\n", audioFile)
		if err != nil {
			return "", fmt.Errorf("error writing to list file: %v", err)
		}
	}

	// Construct the FFmpeg command
	cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", listFile, "-c", "copy", outputFile)

	// Suppress FFmpeg logs
	devNull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to open /dev/null: %w", err)
	}

	defer devNull.Close()
	cmd.Stdout = devNull
	cmd.Stderr = devNull

	// Run the command and capture errors if any
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error merging audio files with ffmpeg: %v", err)
	}

	// Verify that the output file exists
	if _, err := os.Stat(outputFile); err != nil {
		return "", fmt.Errorf("merged audio file not found: %v", err)
	}

	return outputFile, nil
}
