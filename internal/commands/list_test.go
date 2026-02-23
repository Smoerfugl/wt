package commands

import (
	"github.com/smoerfugl/wt/internal/services"
	"os/exec"
	"testing"
)

func TestNewListCommandCreation(t *testing.T) {
	// Ensure we can create a ListCommand and set flags
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not installed in test environment")
	}
	gs := services.NewGitService("")
	lc := NewListCommand(gs)
	lc.SetVerbose(true)
	lc.SetFilter("foo")
	lc.SetBranch("main")
	lc.SetJSONOutput(false)
	if !lc.verbose {
		t.Fatalf("verbose flag not set")
	}
}
