package services

import (
	"os/exec"
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
