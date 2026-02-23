package models

import "testing"

func TestRepositoryAddAndFilters(t *testing.T) {
	r := NewRepository("/tmp/main", "2.39")

	r.AddWorktree(Worktree{Name: "main-wt", Path: "/tmp/main", Branch: "main", CommitHash: "a1b2c3", IsCurrent: true, IsClean: true, IsLocked: false})
	r.AddWorktree(Worktree{Name: "feature-x", Path: "/tmp/feature-x", Branch: "feature", CommitHash: "d4e5f6", IsCurrent: false, IsClean: false, IsLocked: false})

	cur := r.GetCurrentWorktree()
	if cur == nil {
		t.Fatalf("expected current worktree, got nil")
	}
	if cur.Name != "main-wt" {
		t.Fatalf("expected current worktree name main-wt, got %s", cur.Name)
	}

	byBranch := r.GetWorktreesByBranch("feature")
	if len(byBranch) != 1 || byBranch[0].Name != "feature-x" {
		t.Fatalf("GetWorktreesByBranch returned unexpected result: %#v", byBranch)
	}

	byName := r.GetWorktreesByNamePattern("FEATURE")
	if len(byName) != 1 || byName[0].Name != "feature-x" {
		t.Fatalf("GetWorktreesByNamePattern returned unexpected result: %#v", byName)
	}
}
