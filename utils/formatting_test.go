package utils

import (
	"strings"
	"testing"
	"worktree-manager/models"
)

func TestFormatBasicAndVerbose(t *testing.T) {
	wts := []models.Worktree{
		{Name: "main", Path: "/repo", Branch: "main", CommitHash: "a1b2c3", IsCurrent: true, IsClean: true},
		{Name: "feature", Path: "/repo-feature", Branch: "feature", CommitHash: "d4e5f6", IsCurrent: false, IsClean: false},
	}

	basic := FormatBasic(wts)
	if !strings.Contains(basic, "main") || !strings.Contains(basic, "feature") {
		t.Fatalf("basic output missing expected names: %s", basic)
	}

	verbose := FormatVerbose(wts)
	if !strings.Contains(verbose, "NAME") || !strings.Contains(verbose, "STATUS") {
		t.Fatalf("verbose output missing header or status: %s", verbose)
	}
}

func TestFormatJSON(t *testing.T) {
	wts := []models.Worktree{{Name: "main", Path: "/repo", Branch: "main", CommitHash: "a1b2c3", IsCurrent: true, IsClean: true}}
	js := FormatJSON(wts, "/repo")
	if !strings.Contains(js, "\"repository\": \"/repo\"") || !strings.Contains(js, "\"name\": \"main\"") {
		t.Fatalf("json output unexpected: %s", js)
	}
}
