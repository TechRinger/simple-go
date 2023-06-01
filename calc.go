package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("calc")
	case "linux":
		cmd = exec.Command("gnome-calculator")
	case "darwin":
		cmd = exec.Command("open", "-a", "Calculator")
	default:
		fmt.Println("Unsupported operating system")
		os.Exit(1)
	}

	err := cmd.Start()
	if err != nil {
		fmt.Printf("Failed to start calculator: %v\n", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Calculator exited with error: %v\n", err)
	}
}
