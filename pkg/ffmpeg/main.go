package ffmpeg

import (
	"os/exec"
)

// Installed checks if ffmpeg is available in the PATH.
func Installed() bool {
	cmd := exec.Command("ffmpeg", "-version")
	err := cmd.Run()
	return err == nil
}
