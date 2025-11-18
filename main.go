package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Get current directory
		cwd, err := os.Getwd()
		if err != nil {
			cwd = "unknown"
		}

		// Print prompt like Windows
		fmt.Printf("your are on: %s> ", cwd)

		// Read input
		if !scanner.Scan() {
			return
		}
		input := scanner.Text()
		if input == "" {
			continue
		}

		// Clean input
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]
		args := words[1:]

		// Handle cd command manually
		if cmdName == "cd" {
			if len(args) == 0 {
				fmt.Println("cd requires a directory")
				continue
			}
			err := os.Chdir(args[0])
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		}

		// Execute other commands
		cmd := exec.Command(cmdName, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// On Windows
		if isWindows() {
			cmd = exec.Command("cmd", "/C", input)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}

		err = cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

// Helper function to detect Windows
func isWindows() bool {
	return strings.Contains(strings.ToLower(os.Getenv("OS")), "windows")
}
