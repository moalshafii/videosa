package tts

import (
	"fmt"
	"math/rand"
	"path/filepath"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
)

func (opts *Options) Speecher() ([]string, error) {

	var (
		FileName   string
		FilePath   string
		FilesPaths []string
	)

	speech := htgotts.Speech{
		Folder:   opts.TempDir,
		Language: opts.Language,
		Handler:  &handlers.Native{},
	}

	for _, line := range opts.Lines {
		FileName = fmt.Sprintf("%d", rand.Int())
		FilePath = filepath.Join(opts.TempDir, FileName+".mp3")
		_, err := speech.CreateSpeechFile(line, FileName)
		if err != nil {
			return FilesPaths, fmt.Errorf("error generating speech: %v", err)
		}

		FilesPaths = append(FilesPaths, FilePath)
	}

	return FilesPaths, nil
}
