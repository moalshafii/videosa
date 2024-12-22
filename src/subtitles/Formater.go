package subtitles

import (
	"fmt"
	"time"
)

// Formater formats a time.Duration into SRT timestamp format (HH:MM:SS,ms)
func Formater(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	milliseconds := d.Milliseconds() % 1000

	return fmt.Sprintf("%02d:%02d:%02d,%03d", hours, minutes, seconds, milliseconds)
}
