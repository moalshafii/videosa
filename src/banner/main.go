package banner

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/common-nighthawk/go-figure"
)

func Print() {
	var cmd *exec.Cmd

	// Check the operating system and use the appropriate clear command
	switch {
	case os.Getenv("OS") == "Windows_NT":
		cmd = exec.Command("cls") // Windows command
	default:
		cmd = exec.Command("clear") // Unix/Linux command
	}

	cmd.Stdout = os.Stdout
	cmd.Run()

	figure.NewFigure("VIDEOSA", "doom", true).Print()
	fmt.Printf(">>>-- https://github.com/MoAlshafii/Videosa -->>>\n\n")
}
