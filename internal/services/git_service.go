// GitService provides integration with Git commands for worktree management
package services

import (
	"bytes"
	"fmt"
	"os"
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

// AddWorktree creates a new worktree at the specified path for the given ref
func (gs *GitService) AddWorktree(worktreePath, ref string) error {
	cmd := exec.Command(gs.gitPath, "worktree", "add", worktreePath, ref)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to add worktree: %w (output: %s)", err, string(output))
	}
	return nil
}

// AddWorktreeWithBranch creates a new worktree with a new branch
func (gs *GitService) AddWorktreeWithBranch(worktreePath, branchName, startPoint string) error {
	args := []string{"worktree", "add", "-b", branchName, worktreePath}
	if startPoint != "" {
		args = append(args, startPoint)
	}

	cmd := exec.Command(gs.gitPath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to add worktree with branch: %w (output: %s)", err, string(output))
	}
	return nil
}

// EnsureWorktreesDir ensures the worktrees directory exists
func (gs *GitService) EnsureWorktreesDir(repoPath string) error {
	parentDir := filepath.Dir(repoPath)
	repoName := filepath.Base(repoPath)
	worktreesDir := filepath.Join(parentDir, "worktrees", repoName)

	// Create worktrees directory if it doesn't exist
	if err := os.MkdirAll(worktreesDir, 0755); err != nil {
		return fmt.Errorf("failed to create worktrees directory: %w", err)
	}
	return nil
}

// GetDefaultRef gets the repository's default branch
func (gs *GitService) GetDefaultRef(repoPath string) (string, error) {
	// Try: git symbolic-ref refs/remotes/origin/HEAD
	cmd := exec.Command(gs.gitPath, "symbolic-ref", "refs/remotes/origin/HEAD")
	cmd.Dir = repoPath
	output, err := cmd.Output()
	if err == nil {
		ref := strings.TrimSpace(string(output))
		// refs/remotes/origin/HEAD -> refs/remotes/origin/main => use origin/main
		if strings.HasPrefix(ref, "refs/remotes/") {
			return strings.TrimPrefix(ref, "refs/remotes/"), nil
		}
		return ref, nil
	}

	// As a fallback, try `git remote show origin` and parse "HEAD branch: <name>"
	cmd2 := exec.Command(gs.gitPath, "remote", "show", "origin")
	cmd2.Dir = repoPath
	output2, err2 := cmd2.Output()
	if err2 == nil {
		lines := strings.Split(string(output2), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "HEAD branch:") {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) == 2 {
					branch := strings.TrimSpace(parts[1])
					// prefer origin/<branch>
					return "origin/" + branch, nil
				}
			}
		}
	}

	// Try to get current branch
	cmd3 := exec.Command(gs.gitPath, "symbolic-ref", "--short", "HEAD")
	cmd3.Dir = repoPath
	output3, err3 := cmd3.Output()
	if err3 == nil {
		branch := strings.TrimSpace(string(output3))
		return branch, nil
	}

	// Last resort: use 'main' or 'master' if they exist
	for _, cand := range []string{"main", "master"} {
		cmd := exec.Command(gs.gitPath, "show-ref", "--verify", "refs/heads/"+cand)
		cmd.Dir = repoPath
		if err := cmd.Run(); err == nil {
			return cand, nil
		}
	}

	return "", fmt.Errorf("could not determine default branch")
}
