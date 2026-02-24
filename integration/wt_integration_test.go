package integration

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// findRepoRoot walks up from the current directory until it finds go.mod
func findRepoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found in parents")
		}
		dir = parent
	}
}

func runCmd(dir, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func TestWtAddListRemoveFlow(t *testing.T) {
	// Skip if git not available
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not installed; skipping integration test")
	}

	repoRoot, err := findRepoRoot()
	if err != nil {
		t.Fatalf("failed to find repo root: %v", err)
	}

	// Build binary
	binDir := t.TempDir()
	binPath := filepath.Join(binDir, "wtbin")
	if out, err := runCmd(repoRoot, "go", "build", "-o", binPath, "."); err != nil {
		t.Fatalf("failed to build wt binary: %v\noutput:\n%s", err, out)
	}
	defer os.Remove(binPath)

	// Prepare temporary repository
	base := t.TempDir()
	repoDir := filepath.Join(base, "repo")
	if err := os.Mkdir(repoDir, 0755); err != nil {
		t.Fatalf("mkdir repo: %v", err)
	}

	if out, err := runCmd(repoDir, "git", "init"); err != nil {
		t.Fatalf("git init failed: %v\n%s", err, out)
	}

	// Configure user so commits succeed
	if out, err := runCmd(repoDir, "git", "config", "user.email", "test@example.com"); err != nil {
		t.Fatalf("git config email failed: %v\n%s", err, out)
	}
	if out, err := runCmd(repoDir, "git", "config", "user.name", "Test User"); err != nil {
		t.Fatalf("git config name failed: %v\n%s", err, out)
	}

	// initial commit
	readme := filepath.Join(repoDir, "README.md")
	if err := os.WriteFile(readme, []byte("hello"), 0644); err != nil {
		t.Fatalf("write readme: %v", err)
	}
	if out, err := runCmd(repoDir, "git", "add", "."); err != nil {
		t.Fatalf("git add failed: %v\n%s", err, out)
	}
	if out, err := runCmd(repoDir, "git", "commit", "-m", "init"); err != nil {
		t.Fatalf("git commit failed: %v\n%s", err, out)
	}

	// Run wt add -b feature-1
	if out, err := runCmd(repoDir, binPath, "add", "-b", "feature-1"); err != nil {
		t.Fatalf("wt add failed: %v\n%s", err, out)
	}

	// Validate worktree directory exists
	parent := filepath.Dir(repoDir)
	repoName := filepath.Base(repoDir)
	wtPath := filepath.Join(parent, "worktrees", repoName, "feature-1")
	if _, err := os.Stat(wtPath); err != nil {
		t.Fatalf("expected worktree path %s to exist: %v", wtPath, err)
	}

	// Run wt list -v and check output contains feature-1
	if out, err := runCmd(repoDir, binPath, "list", "-v"); err != nil {
		t.Fatalf("wt list failed: %v\n%s", err, out)
	} else if !strings.Contains(out, "feature-1") {
		t.Fatalf("wt list output did not contain feature-1:\n%s", out)
	}

	// Run wt list -j and check json contains repo and feature
	if out, err := runCmd(repoDir, binPath, "list", "-j"); err != nil {
		t.Fatalf("wt list -j failed: %v\n%s", err, out)
	} else if !strings.Contains(out, "feature-1") || !strings.Contains(out, "worktrees") {
		t.Fatalf("wt list -j output unexpected:\n%s", out)
	}

	// Remove the created worktree
	if out, err := runCmd(repoDir, binPath, "remove", wtPath); err != nil {
		t.Fatalf("wt remove failed: %v\n%s", err, out)
	}

	// Ensure git worktree list does not include the removed path
	if out, err := runCmd(repoDir, "git", "worktree", "list", "--porcelain"); err != nil {
		t.Fatalf("git worktree list failed: %v\n%s", err, out)
	} else if strings.Contains(out, wtPath) {
		t.Fatalf("worktree path still listed after removal:\n%s", out)
	}

	// Run prune to ensure command works
	if out, err := runCmd(repoDir, binPath, "prune"); err != nil {
		t.Fatalf("wt prune failed: %v\n%s", err, out)
	}
}
