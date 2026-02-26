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

// GitServiceIface defines the subset of Git operations used by commands
type GitServiceIface interface {
	EnsureWorktreesDir(repoPath string) error
	AddWorktree(worktreePath, ref string) error
	AddWorktreeWithBranch(worktreePath, branchName, startPoint string) error
	GetDefaultRef(repoPath string) (string, error)
	BranchExistsLocal(repoPath, branch string) (bool, error)
	BranchExistsRemote(repoPath, remote, branch string) (bool, error)
	HasRemote(repoPath, remote string) (bool, error)
	SetBranchUpstream(repoPath, branch, remoteBranch string) error
	GetBranchUpstream(repoPath, branch string) (string, error)
	GetGitVersion() (string, error)
}

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

// BranchExistsLocal checks whether a local branch exists in the repository
func (gs *GitService) BranchExistsLocal(repoPath, branch string) (bool, error) {
	if branch == "" {
		return false, fmt.Errorf("branch name required")
	}
	cmd := exec.Command(gs.gitPath, "show-ref", "--verify", "--", "refs/heads/"+branch)
	cmd.Dir = repoPath
	if err := cmd.Run(); err != nil {
		// Non-zero exit indicates ref missing
		if exitErr, ok := err.(*exec.ExitError); ok {
			_ = exitErr
			return false, nil
		}
		return false, fmt.Errorf("failed to check local branch: %w", err)
	}
	return true, nil
}

// BranchExistsRemote checks whether a branch exists on the given remote
func (gs *GitService) BranchExistsRemote(repoPath, remote, branch string) (bool, error) {
	if branch == "" || remote == "" {
		return false, fmt.Errorf("remote and branch required")
	}
	ref := fmt.Sprintf("refs/remotes/%s/%s", remote, branch)
	cmd := exec.Command(gs.gitPath, "show-ref", "--verify", "--", ref)
	cmd.Dir = repoPath
	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return false, nil
		}
		return false, fmt.Errorf("failed to check remote branch: %w", err)
	}
	return true, nil
}

// HasRemote reports whether the repository has a remote with the given name
func (gs *GitService) HasRemote(repoPath, remote string) (bool, error) {
	if remote == "" {
		return false, fmt.Errorf("remote name required")
	}
	cmd := exec.Command(gs.gitPath, "remote")
	cmd.Dir = repoPath
	out, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("failed to list remotes: %w", err)
	}
	remotes := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, r := range remotes {
		if strings.TrimSpace(r) == remote {
			return true, nil
		}
	}
	return false, nil
}

// SetBranchUpstream configures the upstream for a local branch to the given remote/branch
// remoteBranch should be in form "<remote>/<branch>" or a full ref
func (gs *GitService) SetBranchUpstream(repoPath, branch, remoteBranch string) error {
	if branch == "" || remoteBranch == "" {
		return fmt.Errorf("branch and remoteBranch required")
	}
	// Try the user-friendly command first. This can fail if the remote/<branch>
	// ref does not exist locally (we don't want to force a push here), so fall
	// back to setting the branch.* config values directly.
	cmd := exec.Command(gs.gitPath, "branch", "--set-upstream-to", remoteBranch, branch)
	cmd.Dir = repoPath
	output, err := cmd.CombinedOutput()
	if err == nil {
		return nil
	}

	// Fallback: set branch.<branch>.remote and branch.<branch>.merge
	parts := strings.SplitN(remoteBranch, "/", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid remoteBranch format: %s (output: %s)", remoteBranch, string(output))
	}
	remoteName := parts[0]
	mergeBranch := parts[1]

	// git config branch.<branch>.remote <remote>
	cfg1 := exec.Command(gs.gitPath, "config", "branch."+branch+".remote", remoteName)
	cfg1.Dir = repoPath
	if out1, err1 := cfg1.CombinedOutput(); err1 != nil {
		return fmt.Errorf("failed to set branch remote config: %w (output: %s)", err1, string(out1))
	}

	// git config branch.<branch>.merge refs/heads/<mergeBranch>
	cfg2 := exec.Command(gs.gitPath, "config", "branch."+branch+".merge", "refs/heads/"+mergeBranch)
	cfg2.Dir = repoPath
	if out2, err2 := cfg2.CombinedOutput(); err2 != nil {
		return fmt.Errorf("failed to set branch merge config: %w (output: %s)", err2, string(out2))
	}

	return nil
}

// GetBranchUpstream returns the upstream (remote/branch) for a local branch, or empty string if none
func (gs *GitService) GetBranchUpstream(repoPath, branch string) (string, error) {
	if branch == "" {
		return "", fmt.Errorf("branch name required")
	}
	// Use for-each-ref to get upstream short name
	ref := "refs/heads/" + branch
	cmd := exec.Command(gs.gitPath, "for-each-ref", "--format=%(upstream:short)", ref)
	cmd.Dir = repoPath
	out, err := cmd.Output()
	if err != nil {
		// If the ref doesn't exist, return an error
		if _, ok := err.(*exec.ExitError); ok {
			return "", nil
		}
		return "", fmt.Errorf("failed to get branch upstream: %w", err)
	}
	up := strings.TrimSpace(string(out))
	return up, nil
}
