package video

import "time"

type Overlay struct {
	File string
	X    int
	Y    int
}

type Options struct {
	OutputDir      string
	Duration       time.Duration
	VideoFile      string
	AudioFile      string
	SubsFile       string
	Overlay        Overlay
	SubFontName    string
	SubFontColor   string
	SubBGColor     string
	SubFontSize    int
	SubBorderWidth int
	SubAlignment   int
}
