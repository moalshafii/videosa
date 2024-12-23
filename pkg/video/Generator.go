package video

import (
	"fmt"
	"math/rand"
	"os/exec"
)

func (opts *Options) Generator() (string, error) {

	if err := opts.Validation(); err != nil {
		return "", err
	}

	// Build the output file path dynamically
	outputFile := fmt.Sprintf("%s/%d.mp4", opts.OutputDir, rand.Int())

	// Build the FFmpeg command
	subtitleFilter := opts.SubtitleFilter()
	overlayFilter := opts.OverlayFilter()

	cmd := fmt.Sprintf(
		"ffmpeg -stream_loop -1 -i %s -i %s -i %s -filter_complex \"%s,%s[video_out];[0:a][1:a]amix=inputs=2[audio_out]\" -map \"[video_out]\" -map \"[audio_out]\" -t %f -c:v libx264 -c:a aac %s",
		opts.VideoFile,
		opts.AudioFile,
		opts.Overlay.File,
		subtitleFilter,
		overlayFilter,
		opts.Duration.Seconds(),
		outputFile,
	)

	// Run the command using exec
	command := exec.Command("bash", "-c", cmd)
	output, err := command.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("failed to generate video: %v\nOutput: %s", err, string(output))
	}

	return outputFile, nil
}
