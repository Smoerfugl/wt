package services

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// These tests are light-weight and avoid invoking real git in CI environments.
// They verify helper behavior where possible.

func TestNewGitServiceDefault(t *testing.T) {
	gs := NewGitService("")
	if gs.gitPath == "" {
		t.Fatalf("expected gitPath to default to 'git'")
	}
}

func TestGetGitVersionCommandExists(t *testing.T) {
	// If git is not available, skip this test rather than fail CI
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not installed in test environment")
	}
	gs := NewGitService("")
	v, err := gs.GetGitVersion()
	if err != nil {
		t.Fatalf("GetGitVersion failed: %v", err)
	}
	if v == "" {
		t.Fatalf("empty git version returned")
	}
}

func TestSetUpstreamBranch(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not installed in test environment")
	}

	// Create a temp bare "remote" repo and a local clone to test upstream tracking.
	tmpDir := t.TempDir()
	remoteDir := filepath.Join(tmpDir, "remote.git")
	localDir := filepath.Join(tmpDir, "local")

	// Init bare remote
	if out, err := exec.Command("git", "init", "--bare", remoteDir).CombinedOutput(); err != nil {
		t.Fatalf("git init bare: %v: %s", err, out)
	}

	// Clone remote
	if out, err := exec.Command("git", "clone", remoteDir, localDir).CombinedOutput(); err != nil {
		t.Fatalf("git clone: %v: %s", err, out)
	}

	// Create an initial commit so the repo is non-empty
	configArgs := [][]string{
		{"config", "user.email", "test@example.com"},
		{"config", "user.name", "Test"},
	}
	for _, args := range configArgs {
		cmd := exec.Command("git", args...)
		cmd.Dir = localDir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v: %s", args, err, out)
		}
	}
	readmeFile := filepath.Join(localDir, "README.md")
	if err := os.WriteFile(readmeFile, []byte("hello"), 0644); err != nil {
		t.Fatalf("write file: %v", err)
	}
	for _, args := range [][]string{
		{"add", "."},
		{"commit", "-m", "init"},
	} {
		cmd := exec.Command("git", args...)
		cmd.Dir = localDir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v: %s", args, err, out)
		}
	}
	// Determine current branch name (may be "main" or "master" depending on config)
	headCmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")
	headCmd.Dir = localDir
	headOut, err := headCmd.Output()
	if err != nil {
		t.Fatalf("git symbolic-ref: %v", err)
	}
	defaultBranch := strings.TrimSpace(string(headOut))
	pushArgs := []string{"push", "-u", "origin", defaultBranch}
	if out, err := exec.Command("git", append([]string{"-C", localDir}, pushArgs...)...).CombinedOutput(); err != nil {
		t.Fatalf("git push: %v: %s", err, out)
	}

	// Create a new local branch that mirrors what AddWorktreeWithBranch would do
	branchName := "feature-upstream-test"
	if out, err := exec.Command("git", "-C", localDir, "branch", branchName).CombinedOutput(); err != nil {
		t.Fatalf("git branch: %v: %s", err, out)
	}

	gs := NewGitService("")
	// SetUpstreamBranch should return nil (no error) when origin/<branch> does not exist yet.
	if err := gs.SetUpstreamBranch(localDir, branchName); err != nil {
		t.Fatalf("SetUpstreamBranch should not error when remote branch does not exist: %v", err)
	}

	// Push the branch to remote so the upstream can be set
	if out, err := exec.Command("git", "-C", localDir, "push", "origin", branchName).CombinedOutput(); err != nil {
		t.Fatalf("git push branch: %v: %s", err, out)
	}

	if err := gs.SetUpstreamBranch(localDir, branchName); err != nil {
		t.Fatalf("SetUpstreamBranch failed: %v", err)
	}

	// Verify upstream was configured
	cmd := exec.Command("git", "for-each-ref", "--format=%(upstream:short)", "refs/heads/"+branchName)
	cmd.Dir = localDir
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("git for-each-ref: %v", err)
	}
	got := strings.TrimSpace(string(out))
	want := "origin/" + branchName
	if got != want {
		t.Fatalf("upstream = %q, want %q", got, want)
	}
}

