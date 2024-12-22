package initial

import (
	"fmt"
	"os"
	"path/filepath"
)

type dirs struct {
	Temp    string
	Outputs string
}

// Dirs initializes and returns a dirs struct with paths.
func Dirs() (dirs, error) {

	dirPaths := dirs{}
	dirPaths.Temp = filepath.Join(os.TempDir(), "videosa", "v1.00")
	dirPaths.Outputs = filepath.Join("outputs")

	// Create the directories
	if err := dirPaths.Create(); err != nil {
		return dirPaths, fmt.Errorf("error creating directories: %w", err)
	}

	return dirPaths, nil
}

// Create creates the directories defined in the dirs struct.
func (d dirs) Create() error {

	if err := os.MkdirAll(d.Temp, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create Temp directory %s: %w", d.Temp, err)
	}

	if err := os.MkdirAll(d.Outputs, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create Outputs directory %s: %w", d.Outputs, err)
	}

	return nil
}
