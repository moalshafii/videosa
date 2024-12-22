package subtitles

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"videosa/src/audios"
)

func (opts *Options) Generator() (string, error) {
	// Ensure the number of lines matches the number of audio files
	if len(opts.Lines) != len(opts.Audios) {
		return "", fmt.Errorf("number of lines must match the number of audio files")
	}

	// Create a single subtitle file
	outputFile := fmt.Sprintf("%s/%d.srt", opts.TempDir, rand.Int())
	file, err := os.Create(outputFile)
	if err != nil {
		return "", fmt.Errorf("error creating subtitle file: %v", err)
	}
	defer file.Close()

	var start time.Duration

	// Write subtitles for each line and its corresponding audio
	for i, line := range opts.Lines {
		audio := opts.Audios[i]

		// Get the duration of the current audio file
		duration, err := audios.Duration(audio)
		if err != nil {
			return "", fmt.Errorf("error getting audio duration for %s: %v", audio, err)
		}

		// Calculate the end time for this subtitle line
		end := start + duration

		// Write subtitle entry
		_, err = fmt.Fprintf(file, "%d\n%s --> %s\n%s\n\n",
			i+1, Formater(start), Formater(end), line)
		if err != nil {
			return "", fmt.Errorf("error writing to subtitle file: %v", err)
		}

		// Update the start time for the next subtitle line
		start = end
	}

	// Return the generated subtitle file path
	return outputFile, nil
}
