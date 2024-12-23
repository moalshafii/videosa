package tts

import (
	"fmt"
)

func (opts *Options) Generator() (string, []string, error) {
	TTSFilesPaths, err := opts.Speecher()
	if err != nil {
		return "", []string{}, fmt.Errorf("TTS Generator Error: %w", err)
	}

	MergedFile, err := opts.Merger(TTSFilesPaths)
	if err != nil {
		return "", []string{}, fmt.Errorf("TTS Generator Error: %w", err)
	}

	return MergedFile, TTSFilesPaths, nil
}
