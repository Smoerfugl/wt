package utils

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestNewCommand(t *testing.T) {
	cmd := NewCommand("echo", []string{"hello"})

	if cmd.Name != "echo" {
		t.Errorf("Expected name 'echo', got '%s'", cmd.Name)
	}
	if len(cmd.Args) != 1 || cmd.Args[0] != "hello" {
		t.Errorf("Expected args ['hello'], got %v", cmd.Args)
	}
	if cmd.Timeout != 5*time.Minute {
		t.Errorf("Expected timeout 5 minutes, got %v", cmd.Timeout)
	}
	if cmd.Interactive != false {
		t.Errorf("Expected Interactive false, got %v", cmd.Interactive)
	}
}

func TestCommandValidate(t *testing.T) {
	tests := []struct {
		name        string
		cmd         *Command
		expectedErr error
	}{
		{
			name:        "valid command",
			cmd:         NewCommand("echo", []string{"hello"}),
			expectedErr: nil,
		},
		{
			name:        "empty command name",
			cmd:         &Command{Name: "", Args: []string{"hello"}},
			expectedErr: errors.New("command name cannot be empty"),
		},
		{
			name:        "command with path traversal",
			cmd:         &Command{Name: "../../bin/echo", Args: []string{"hello"}},
			expectedErr: errors.New("command name contains invalid path traversal characters"),
		},
		{
			name:        "args with path traversal",
			cmd:         &Command{Name: "echo", Args: []string{"../../hello"}},
			expectedErr: errors.New("command arguments contain invalid path traversal characters"),
		},
		{
			name:        "zero timeout",
			cmd:         &Command{Name: "echo", Args: []string{"hello"}, Timeout: 0},
			expectedErr: errors.New("timeout must be positive"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cmd.Validate()
			if tt.expectedErr != nil {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if err.Error() != tt.expectedErr.Error() {
					t.Errorf("Expected error '%s', got '%s'", tt.expectedErr.Error(), err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}

func TestIsInteractiveCommand(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		expected bool
	}{
		{"bash", "bash", true},
		{"sh", "sh", true},
		{"python", "python", true},
		{"vim", "vim", true},
		{"echo", "echo", false},
		{"ls", "ls", false},
		{"/usr/bin/python3", "/usr/bin/python3", true},
		{"python3 script.py", "python3", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsInteractiveCommand(tt.command)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestCommandExecute(t *testing.T) {
	tests := []struct {
		name          string
		cmd           *Command
		expectSuccess bool
		expectError   bool
		errorContains string
	}{
		{
			name:          "simple echo command",
			cmd:           NewCommand("echo", []string{"hello"}),
			expectSuccess: true,
			expectError:   false,
		},
		{
			name:          "command with args",
			cmd:           NewCommand("echo", []string{"hello", "world"}),
			expectSuccess: true,
			expectError:   false,
		},
		{
			name:          "nonexistent command",
			cmd:           NewCommand("nonexistent-command-xyz", []string{}),
			expectSuccess: false,
			expectError:   true,
			errorContains: "executable file not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.cmd.Execute()

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if tt.errorContains != "" && !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("Expected error to contain '%s', got: %v", tt.errorContains, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}

			if result.Success != tt.expectSuccess {
				t.Errorf("Expected success %v, got %v", tt.expectSuccess, result.Success)
			}
			if result.Duration <= 0 {
				t.Errorf("Expected non-zero duration, got %v", result.Duration)
			}
		})
	}
}

func TestCommandExecuteWithTimeout(t *testing.T) {
	// Test short timeout
	cmd := NewCommand("sleep", []string{"10"})
	cmd.Timeout = 100 * time.Millisecond // Very short timeout

	result, err := cmd.Execute()

	if err == nil {
		t.Errorf("Expected error but got none")
	} else if !strings.Contains(err.Error(), "timed out") && !strings.Contains(err.Error(), "signal: killed") {
		t.Errorf("Expected error to contain 'timed out' or 'signal: killed', got: %v", err)
	}
	if result.Duration >= time.Second {
		t.Errorf("Expected duration < 1 second, got %v", result.Duration)
	}
}

func TestCommandExecuteWithWorkingDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	cmd := NewCommand("pwd", []string{})
	cmd.Dir = tmpDir

	result, err := cmd.Execute()

	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}
	if !result.Success {
		t.Errorf("Expected success true, got false")
	}
	if !strings.Contains(result.Stdout, tmpDir) {
		t.Errorf("Expected stdout to contain '%s', got: %s", tmpDir, result.Stdout)
	}
}

func TestExecuteCommands(t *testing.T) {
	commands := []*Command{
		NewCommand("echo", []string{"first"}),
		NewCommand("echo", []string{"second"}),
		NewCommand("nonexistent-command", []string{}), // This should fail
	}

	results, err := ExecuteCommands(commands)

	// Should not return error for ExecuteCommands itself, individual command errors are in results
	if err != nil {
		t.Errorf("Expected no error from ExecuteCommands but got: %v", err)
	}
	if len(results) != 3 {
		t.Errorf("Expected 3 results, got %d", len(results))
	}

	// First two should succeed
	if !results[0].Success {
		t.Errorf("Expected first command to succeed")
	}
	if !results[1].Success {
		t.Errorf("Expected second command to succeed")
	}

	// Third should fail
	if results[2].Success {
		t.Errorf("Expected third command to fail")
	}
	if results[2].Error == nil {
		t.Errorf("Expected third command to have error")
	}
}

func TestCommandError(t *testing.T) {
	err := NewCommandError("test-command", errors.New("test error"), 127, false)

	if err.Command != "test-command" {
		t.Errorf("Expected command 'test-command', got '%s'", err.Command)
	}
	if err.ExitCode != 127 {
		t.Errorf("Expected exit code 127, got %d", err.ExitCode)
	}
	if err.IsTimeout != false {
		t.Errorf("Expected IsTimeout false, got %v", err.IsTimeout)
	}
	expectedError := "command 'test-command' failed with exit code 127: test error"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}

	// Test timeout error
	timeoutErr := NewCommandError("sleep", errors.New("timeout"), 0, true)
	if !timeoutErr.IsTimeout {
		t.Errorf("Expected IsTimeout true, got false")
	}
	if !strings.Contains(timeoutErr.Error(), "timed out") {
		t.Errorf("Expected error to contain 'timed out', got: %s", timeoutErr.Error())
	}
}

func TestWorktreeError(t *testing.T) {
	err := NewWorktreeError("creation", "/path/to/worktree", errors.New("git error"))

	if err.Operation != "creation" {
		t.Errorf("Expected operation 'creation', got '%s'", err.Operation)
	}
	if err.Path != "/path/to/worktree" {
		t.Errorf("Expected path '/path/to/worktree', got '%s'", err.Path)
	}
	expectedError := "worktree creation failed for '/path/to/worktree': git error"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}

func TestErrorUtilities(t *testing.T) {
	// Test error checking functions
	cmdErr := NewCommandError("test", ErrCommandNotFound, 1, false)
	if !IsCommandNotFound(cmdErr) {
		t.Errorf("Expected IsCommandNotFound to return true")
	}

	permErr := NewCommandError("test", ErrPermissionDenied, 1, false)
	if !IsPermissionDenied(permErr) {
		t.Errorf("Expected IsPermissionDenied to return true")
	}

	timeoutErr := NewCommandError("test", ErrTimeoutExceeded, 0, true)
	if !IsTimeout(timeoutErr) {
		t.Errorf("Expected IsTimeout to return true")
	}

	exitErr := NewCommandError("test", ErrNonZeroExit, 1, false)
	if !IsNonZeroExit(exitErr) {
		t.Errorf("Expected IsNonZeroExit to return true")
	}

	// Test GetExitCode
	if GetExitCode(exitErr) != 1 {
		t.Errorf("Expected GetExitCode to return 1, got %d", GetExitCode(exitErr))
	}
	if GetExitCode(errors.New("regular error")) != -1 {
		t.Errorf("Expected GetExitCode to return -1 for regular error")
	}
}
