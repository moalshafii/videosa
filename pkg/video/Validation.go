package video

import (
	"fmt"
	"os"
)

func (opts *Options) Validation() error {
	// Check if the video file exists
	if _, err := os.Stat(opts.VideoFile); err != nil {
		return fmt.Errorf("video file does not exist: %v", err)
	}

	// Check if the duration is valid
	if opts.Duration.Seconds() <= 0 {
		return fmt.Errorf("duration must be greater than zero")
	}

	return nil
}
