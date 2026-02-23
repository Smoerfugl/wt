// GitService provides integration with Git commands for worktree management
package services

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/smoerfugl/wt/internal/models"
)

// GitService handles Git operations
type GitService struct {
	gitPath string
}

// NewGitService creates a new GitService instance
func NewGitService(gitPath string) *GitService {
	if gitPath == "" {
		gitPath = "git"
	}
	return &GitService{gitPath: gitPath}
}

// GetWorktrees retrieves all worktrees from the repository
func (gs *GitService) GetWorktrees(repoPath string) ([]models.Worktree, error) {
	cmd := exec.Command(gs.gitPath, "worktree", "list", "--porcelain")
	cmd.Dir = repoPath

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to get worktrees: %w", err)
	}

	return gs.parseWorktreeOutput(output)
}

// parseWorktreeOutput parses the output from 'git worktree list --porcelain'
func (gs *GitService) parseWorktreeOutput(output []byte) ([]models.Worktree, error) {
	var worktrees []models.Worktree
	var currentWorktree models.Worktree

	lines := bytes.Split(output, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		// Git porcelain format uses space-separated key-value pairs
		parts := bytes.Fields(line)
		if len(parts) < 2 {
			continue
		}

		key := string(parts[0])
		value := strings.TrimSpace(string(line[len(parts[0]):]))

		switch key {
		case "worktree":
			if currentWorktree.Path != "" {
				worktrees = append(worktrees, currentWorktree)
			}
			// Extract name from path (last component)
			name := filepath.Base(value)
			currentWorktree = models.Worktree{Path: value, Name: name}
		case "HEAD":
			// Extract branch name from refs/heads/branch
			if strings.HasPrefix(value, "refs/heads/") {
				currentWorktree.Branch = strings.TrimPrefix(value, "refs/heads/")
				currentWorktree.CommitHash = "" // Clear commit hash if we have a branch
			} else if len(value) >= 7 {
				currentWorktree.CommitHash = value[:7] // Short hash
			}
		case "branch":
			// Handle branch refs/heads/branch format
			if strings.HasPrefix(value, "refs/heads/") {
				currentWorktree.Branch = strings.TrimPrefix(value, "refs/heads/")
			} else {
				currentWorktree.Branch = value
			}
		case "bare":
			// Ignore bare indicator
		case "detached":
			// Handle detached HEAD
		case "locked":
			currentWorktree.IsLocked = true
		}
	}

	// Add the last worktree
	if currentWorktree.Path != "" {
		worktrees = append(worktrees, currentWorktree)
	}

	// Determine current worktree by checking which worktree path matches the repo path
	// For now, we'll mark the first worktree as current as a basic implementation
	if len(worktrees) > 0 {
		worktrees[0].IsCurrent = true
	}

	return worktrees, nil
}

// GetGitVersion retrieves the Git version
func (gs *GitService) GetGitVersion() (string, error) {
	cmd := exec.Command(gs.gitPath, "--version")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get git version: %w", err)
	}

	version := strings.TrimSpace(string(output))
	// Extract version number from "git version X.Y.Z"
	parts := strings.Fields(version)
	if len(parts) >= 3 {
		return parts[2], nil
	}
	return version, nil
}
