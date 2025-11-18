package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// -------------------- CLI COMMAND STRUCT --------------------
type cliCommand struct {
	name        string
	description string
	callback    func(args []string) error
}

// Declare commands map globally (but do NOT initialize here)
var commands map[string]cliCommand

func main() {
	// Initialize commands inside main to avoid initialization cycle
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the terminal",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays this help message",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Get current directory
		cwd, err := os.Getwd()
		if err != nil {
			cwd = "unknown"
		}

		// Print prompt
		fmt.Printf("%s> ", cwd)

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

		// Check for built-in commands
		if cmd, exists := commands[cmdName]; exists {
			err := cmd.callback(args)
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		}

		// Handle cd manually
		if cmdName == "cd" {
			handleCd(args)
			continue
		}

		// Run OS command
		runCommand(input)
	}
}

// -------------------- CD COMMAND --------------------
func handleCd(args []string) {
	if len(args) == 0 {
		fmt.Println("cd requires a directory")
		return
	}

	dir := args[0]

	// Expand ~ to home directory
	if dir == "~" {
		home, err := os.UserHomeDir()
		if err == nil {
			dir = home
		}
	}

	// Make path absolute
	if !filepath.IsAbs(dir) {
		cwd, _ := os.Getwd()
		dir = filepath.Join(cwd, dir)
	}

	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// -------------------- OS COMMANDS --------------------
func runCommand(input string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", input)
	} else {
		cmd = exec.Command("bash", "-c", input)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// -------------------- EXIT COMMAND --------------------
func commandExit(args []string) error {
	fmt.Println("Closing the terminal... Goodbye!")
	os.Exit(0)
	return nil
}

// -------------------- HELP COMMAND --------------------
func commandHelp(args []string) error {
	fmt.Println("Welcome to the terminal!")
	fmt.Println("Available commands:")

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println("Other OS commands are also supported.")
	fmt.Println("Use 'cd <dir>' to change directory.")
	return nil
}
