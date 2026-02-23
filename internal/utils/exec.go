package utils

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Command represents a command to execute in a worktree context
type Command struct {
	Name        string        // Command name/executable
	Args        []string      // Command arguments
	Dir         string        // Working directory (worktree path)
	Env         []string      // Environment variables (optional)
	Timeout     time.Duration // Execution timeout (default: 5 minutes)
	Interactive bool          // Whether command requires user input
}

// ExecutionResult represents the result of command execution
type ExecutionResult struct {
	Command  string        // The command that was executed
	Success  bool          // Whether execution succeeded
	ExitCode int           // Process exit code
	Stdout   string        // Standard output
	Stderr   string        // Standard error
	Duration time.Duration // Execution time
	Error    error         // Any execution error
}

// NewCommand creates a new Command with default values
func NewCommand(name string, args []string) *Command {
	return &Command{
		Name:        name,
		Args:        args,
		Timeout:     5 * time.Minute, // Default 5-minute timeout
		Interactive: false,
	}
}

// Execute runs the command with the configured settings
func (c *Command) Execute() (*ExecutionResult, error) {
	if c.Name == "" {
		return nil, errors.New("command name cannot be empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, c.Name, c.Args...)

	// Set working directory if specified
	if c.Dir != "" {
		cmd.Dir = c.Dir
	}

	// Set environment variables if specified
	if len(c.Env) > 0 {
		cmd.Env = c.Env
	} else {
		// Inherit parent process environment
		cmd.Env = os.Environ()
	}

	// Set up pipes for capturing output
	var stdoutBuf, stderrBuf strings.Builder
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	// For interactive commands, connect to terminal
	if c.Interactive {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	startTime := time.Now()
	err := cmd.Run()
	executionTime := time.Since(startTime)

	result := &ExecutionResult{
		Command:  c.Name + " " + strings.Join(c.Args, " "),
		Success:  err == nil,
		Duration: executionTime,
		Error:    err,
	}

	if err != nil {
		// Handle timeout error specifically
		if errors.Is(err, context.DeadlineExceeded) {
			result.Error = errors.New("command timed out after " + c.Timeout.String())
		} else {
			// Try to get exit code from error
			if exitErr, ok := err.(*exec.ExitError); ok {
				result.ExitCode = exitErr.ExitCode()
			} else if execErr, ok := err.(*exec.Error); ok {
				// Command not found or other exec error
				result.ExitCode = -1
				result.Error = fmt.Errorf("command execution failed: %w", execErr)
			} else {
				result.ExitCode = -1
			}
		}
	}

	// Capture output
	if !c.Interactive {
		result.Stdout = stdoutBuf.String()
		result.Stderr = stderrBuf.String()
	}

	return result, result.Error
}

// Validate checks if the command is valid for execution
func (c *Command) Validate() error {
	if c.Name == "" {
		return errors.New("command name cannot be empty")
	}

	// Check for path traversal attempts
	if strings.Contains(c.Name, "../") || strings.Contains(c.Name, "..\\") {
		return errors.New("command name contains invalid path traversal characters")
	}

	for _, arg := range c.Args {
		if strings.Contains(arg, "../") || strings.Contains(arg, "..\\") {
			return errors.New("command arguments contain invalid path traversal characters")
		}
	}

	if c.Timeout <= 0 {
		return errors.New("timeout must be positive")
	}

	return nil
}

// IsInteractiveCommand checks if a command is likely to be interactive
func IsInteractiveCommand(name string) bool {
	interactiveCommands := []string{
		"bash", "sh", "zsh", "fish", "ksh", "csh", "tcsh",
		"python", "python3", "ruby", "perl", "node", "rails", "console",
		"vim", "emacs", "nano", "pico", "less", "more", "man",
		"ssh", "sftp", "telnet", "mysql", "psql", "sqlite3",
	}

	for _, cmd := range interactiveCommands {
		if strings.HasPrefix(name, cmd) || strings.Contains(name, "/"+cmd) {
			return true
		}
	}
	return false
}

// ExecuteCommands runs multiple commands sequentially
func ExecuteCommands(commands []*Command) ([]*ExecutionResult, error) {
	var results []*ExecutionResult

	for _, cmd := range commands {
		result, err := cmd.Execute()
		if err != nil {
			// Continue with next command even if this one fails
			// but still record the error in the result
		}
		results = append(results, result)
	}

	return results, nil
}
