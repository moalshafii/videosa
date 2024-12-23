package audios

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
)

// GetAudioDuration returns the duration of an audio file in seconds.
func Duration(filePath string) (time.Duration, error) {
	// Open the MP3 file
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Decode the MP3 file
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		return 0, fmt.Errorf("error decoding MP3 file: %v", err)
	}
	defer streamer.Close()

	// Calculate the duration of the audio
	duration := float64(streamer.Len()) / float64(format.SampleRate)

	// Return the duration as a time.Duration
	return time.Duration(duration * float64(time.Second)), nil
}
