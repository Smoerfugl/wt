package utils

import (
	"errors"
	"fmt"
)

// Custom error types for command execution
var (
	ErrCommandNotFound        = errors.New("command not found")
	ErrPermissionDenied       = errors.New("permission denied")
	ErrTimeoutExceeded        = errors.New("command timeout exceeded")
	ErrNonZeroExit            = errors.New("command returned non-zero exit code")
	ErrInvalidCommand         = errors.New("invalid command")
	ErrEmptyCommand           = errors.New("empty command provided")
	ErrWorktreeCreationFailed = errors.New("worktree creation failed")
	ErrCommandExecutionFailed = errors.New("command execution failed")
)

// CommandError represents an error that occurred during command execution
type CommandError struct {
	Command   string
	Err       error
	ExitCode  int
	IsTimeout bool
}

func (e *CommandError) Error() string {
	if e.IsTimeout {
		return fmt.Sprintf("command '%s' timed out", e.Command)
	}
	if e.ExitCode != 0 {
		return fmt.Sprintf("command '%s' failed with exit code %d: %v", e.Command, e.ExitCode, e.Err)
	}
	return fmt.Sprintf("command '%s' failed: %v", e.Command, e.Err)
}

func (e *CommandError) Unwrap() error {
	return e.Err
}

// NewCommandError creates a new CommandError
func NewCommandError(command string, err error, exitCode int, isTimeout bool) *CommandError {
	return &CommandError{
		Command:   command,
		Err:       err,
		ExitCode:  exitCode,
		IsTimeout: isTimeout,
	}
}

// WorktreeError represents an error related to worktree operations
type WorktreeError struct {
	Operation string
	Path      string
	Err       error
}

func (e *WorktreeError) Error() string {
	return fmt.Sprintf("worktree %s failed for '%s': %v", e.Operation, e.Path, e.Err)
}

func (e *WorktreeError) Unwrap() error {
	return e.Err
}

// NewWorktreeError creates a new WorktreeError
func NewWorktreeError(operation, path string, err error) *WorktreeError {
	return &WorktreeError{
		Operation: operation,
		Path:      path,
		Err:       err,
	}
}

// IsCommandNotFound checks if an error is due to command not found
func IsCommandNotFound(err error) bool {
	return errors.Is(err, ErrCommandNotFound)
}

// IsPermissionDenied checks if an error is due to permission denied
func IsPermissionDenied(err error) bool {
	return errors.Is(err, ErrPermissionDenied)
}

// IsTimeout checks if an error is due to timeout
func IsTimeout(err error) bool {
	return errors.Is(err, ErrTimeoutExceeded)
}

// IsNonZeroExit checks if an error is due to non-zero exit code
func IsNonZeroExit(err error) bool {
	return errors.Is(err, ErrNonZeroExit)
}

// GetExitCode extracts exit code from error if available
func GetExitCode(err error) int {
	var cmdErr *CommandError
	if errors.As(err, &cmdErr) {
		return cmdErr.ExitCode
	}
	return -1
}
