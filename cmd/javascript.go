package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func RunJs(args []string) {
	// Path to the JavaScript file
	// scriptPath := "cli.js"

	commandArgs := append([]string{"cli.mjs"}, args...)

	// Create the Node.js command
	cmd := exec.Command("node", commandArgs...)

	// Capture stdout and stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating stdout pipe:", err)
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Error creating stderr pipe:", err)
		return
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	// Create scanners for real-time logging
	stdoutScanner := bufio.NewScanner(stdout)
	stderrScanner := bufio.NewScanner(stderr)

	// Read stdout in a separate goroutine
	go func() {
		for stdoutScanner.Scan() {
			fmt.Println(stdoutScanner.Text()) // Print stdout logs
		}
	}()

	// Read stderr in a separate goroutine
	go func() {
		for stderrScanner.Scan() {
			fmt.Fprintln(os.Stderr, stderrScanner.Text()) // Print stderr logs
		}
	}()

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for command:", err)
	}
}
