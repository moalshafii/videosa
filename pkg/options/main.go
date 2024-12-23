package options

import (
	"errors"
	"flag"
)

// Types defines the structure for options.
type Types struct {
	Script  string
	BGVideo string
}

// Parse parses and validates command-line arguments and populates the Types struct.
func (opts *Types) Parse() error {
	// Define flags
	flag.StringVar(&opts.Script, "script", "", "Path to the script file (required)")
	flag.StringVar(&opts.BGVideo, "bgvideo", "", "Path to the background video file (required)")

	// Parse command-line arguments
	flag.Parse()

	// Validate inputs
	return opts.validate()
}

// validate checks if the required fields are provided and valid.
func (opts *Types) validate() error {
	if opts.Script == "" {
		return errors.New("missing required field: script")
	}
	if opts.BGVideo == "" {
		return errors.New("missing required field: bgvideo")
	}

	return nil
}
