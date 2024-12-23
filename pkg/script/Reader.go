package script

import (
	"bufio"
	"os"
	"strings"
)

// Cleaner trims whitespace from a given line.
func Cleaner(line string) string {
	return strings.TrimSpace(line)
}

// Reader reads a file line by line, trims whitespace, and skips empty lines.
func Reader(scriptFile string) ([]string, error) {
	// Open the file
	file, err := os.Open(scriptFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// Read and clean lines
	for scanner.Scan() {
		line := Cleaner(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
