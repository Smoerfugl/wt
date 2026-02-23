package models

import "testing"

func TestNewWorktreeAndStatus(t *testing.T) {
	w := NewWorktree("wt1", "/tmp/wt1", "main", "abc1234", true, true, false)
	if w.Name != "wt1" {
		t.Fatalf("unexpected name: %s", w.Name)
	}
	if w.Path != "/tmp/wt1" {
		t.Fatalf("unexpected path: %s", w.Path)
	}

	// IsCurrent true should not affect status other than marking current
	if !w.IsCurrent {
		t.Fatalf("expected IsCurrent true")
	}

	// Clean worktree should report "clean"
	if got := w.GetStatus(); got != "clean" {
		t.Fatalf("expected clean status, got %s", got)
	}

	// Dirty
	w.IsClean = false
	if got := w.GetStatus(); got != "dirty" {
		t.Fatalf("expected dirty status, got %s", got)
	}

	// Locked overrides clean/dirty
	w.IsLocked = true
	if got := w.GetStatus(); got != "locked" {
		t.Fatalf("expected locked status, got %s", got)
	}
}
