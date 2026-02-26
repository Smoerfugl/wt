package services

import (
	"os/exec"
	"path/filepath"
	"testing"
)

// These tests exercise the git helper functions. They require `git` in PATH
// and will be skipped when git is not available.

func run(t *testing.T, dir string, args ...string) {
	t.Helper()
	cmd := exec.Command("git", args...)
	if dir != "" {
		cmd.Dir = dir
	}
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git %v failed: %v\noutput: %s", args, err, string(out))
	}
}

func TestBranchExistsLocal(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not installed in test environment")
	}

	tempDir := t.TempDir()
	// init repo and commit
	run(t, tempDir, "init")
	run(t, tempDir, "config", "user.email", "test@example.com")
	run(t, tempDir, "config", "user.name", "tester")
	// create initial commit
	run(t, tempDir, "commit", "--allow-empty", "-m", "initial")

	gs := NewGitService("")

	// branch should not exist yet
	exists, err := gs.BranchExistsLocal(tempDir, "feature-x")
	if err != nil {
		t.Fatalf("BranchExistsLocal error: %v", err)
	}
	if exists {
		t.Fatalf("expected branch to not exist yet")
	}

	// create branch and test
	run(t, tempDir, "checkout", "-b", "feature-x")
	exists, err = gs.BranchExistsLocal(tempDir, "feature-x")
	if err != nil {
		t.Fatalf("BranchExistsLocal error: %v", err)
	}
	if !exists {
		t.Fatalf("expected branch to exist after creation")
	}
}

func TestHasRemoteAndBranchExistsRemote(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not installed in test environment")
	}

	// Create bare remote
	remoteBare := t.TempDir()
	run(t, remoteBare, "init", "--bare")

	// Producer repo: create a branch and push it to bare
	producer := t.TempDir()
	run(t, producer, "init")
	run(t, producer, "config", "user.email", "test@example.com")
	run(t, producer, "config", "user.name", "tester")
	run(t, producer, "commit", "--allow-empty", "-m", "initial")
	run(t, producer, "remote", "add", "origin", remoteBare)
	// push branch 'feature-remote'
	run(t, producer, "checkout", "-b", "feature-remote")
	run(t, producer, "push", "origin", "feature-remote")

	// Consumer: clone the bare remote
	consumer := t.TempDir()
	run(t, "", "clone", remoteBare, consumer)

	gs := NewGitService("")

	// HasRemote should see 'origin'
	has, err := gs.HasRemote(consumer, "origin")
	if err != nil {
		t.Fatalf("HasRemote error: %v", err)
	}
	if !has {
		t.Fatalf("expected origin remote to exist in clone")
	}

	// BranchExistsRemote should find the branch on origin
	exists, err := gs.BranchExistsRemote(consumer, "origin", "feature-remote")
	if err != nil {
		t.Fatalf("BranchExistsRemote error: %v", err)
	}
	if !exists {
		t.Fatalf("expected remote branch 'feature-remote' to exist on origin")
	}

	// Also ensure the refs actually exist
	// show-ref should find refs/remotes/origin/feature-remote (clone may not have fetched it as remote ref but branches should be present)
	// As a sanity check, list refs
	// (not asserting further here)
	_ = filepath.Join(consumer)
}
