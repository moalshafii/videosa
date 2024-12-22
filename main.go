package main

import (
	"fmt"
	"log"
	"os"
	"videosa/src/audios"
	"videosa/src/banner"
	"videosa/src/ffmpeg"
	"videosa/src/initial"
	"videosa/src/logo"
	"videosa/src/options"
	"videosa/src/script"
	"videosa/src/subtitles"
	"videosa/src/tts"
	"videosa/src/video"
)

// Error logs the error and exits the program.
func Error(err error, message string) {
	banner.Print()
	log.Fatalf("Error: %s - %v", message, err)
}

func main() {
	// Ensure FFmpeg is installed
	if !ffmpeg.Installed() {
		log.Fatalln("FFmpeg is not installed on the system.")
	}

	// Print banner
	banner.Print()

	// Parse options
	opts := &options.Types{}
	if err := opts.Parse(); err != nil {
		Error(err, "parsing options")
	}

	// Print parsed options
	fmt.Printf("-- Options:\n#~> Script: %s\n#~> Background Video: %s\n--\n\n",
		opts.Script, opts.BGVideo,
	)

	// Initialize directories
	dirs, err := initial.Dirs()
	if err != nil {
		Error(err, "initializing directories")
	}

	// Retrieve logo path
	logo, err := logo.Path(dirs.Temp)
	if err != nil {
		Error(err, "retrieving logo path")
	}

	// Read script lines
	lines, err := script.Reader(opts.Script)
	if err != nil {
		Error(err, "reading script")
	}
	fmt.Println("-- Script")
	for i, line := range lines {
		fmt.Printf("#~> line %d : %s\n", i, line)
	}
	fmt.Printf("--\n\n")
	fmt.Printf("-- Start Generating \n")

	// Generate text-to-speech
	ttso := tts.Options{
		Lines:    lines,
		Language: "en-US",
		TempDir:  dirs.Temp,
	}
	ttsFile, ttsPartsFiles, err := ttso.Generator()
	if err != nil {
		Error(err, "generating text-to-speech")
	}

	// Generate subtitles
	SubsOpts := subtitles.Options{
		Lines:   lines,
		Audios:  ttsPartsFiles,
		TempDir: dirs.Temp,
	}
	SubsFile, err := SubsOpts.Generator()
	if err != nil {
		Error(err, "generating subtitles")
	}

	// Get audio duration
	Duration, err := audios.Duration(ttsFile)
	if err != nil {
		Error(err, "retrieving audio duration")
	}

	// Generate video
	VideoOpts := video.Options{
		VideoFile: opts.BGVideo,
		Duration:  Duration,
		OutputDir: dirs.Outputs,

		AudioFile: ttsFile,
		Overlay:   video.Overlay{File: logo, X: 15, Y: 15},

		SubsFile:       SubsFile,
		SubFontName:    "More Sugar",
		SubFontColor:   "&HFFFFFF",
		SubBGColor:     "&H000",
		SubFontSize:    20,
		SubBorderWidth: 3,
		SubAlignment:   2,
	}
	video, err := VideoOpts.Generator()
	if err != nil {
		Error(err, "generating video")
	}

	// Output final video path
	fmt.Println("-- Video Generated :", video)

	// Clean up temporary directory
	if err := os.RemoveAll(dirs.Temp); err != nil {
		Error(err, "cleaning up")
	}

}
