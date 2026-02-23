// List command implementation
package commands

import (
	"fmt"
	"strings"

	"worktree-manager/models"
	"worktree-manager/services"
	"worktree-manager/utils"
)

// ListCommand handles the 'wt list' command
type ListCommand struct {
	gitService *services.GitService
	verbose    bool
	filter     string
	branch     string
	jsonOutput bool
}

// NewListCommand creates a new ListCommand instance
func NewListCommand(gitService *services.GitService) *ListCommand {
	return &ListCommand{
		gitService: gitService,
	}
}

// Execute runs the list command
func (lc *ListCommand) Execute(repoPath string) error {
	worktrees, err := lc.gitService.GetWorktrees(repoPath)
	if err != nil {
		return fmt.Errorf("failed to get worktrees: %w", err)
	}

	// Apply filters
	filteredWorktrees := lc.applyFilters(worktrees)

	// Format output
	var output string
	if lc.jsonOutput {
		output = utils.FormatJSON(filteredWorktrees, repoPath)
	} else if lc.verbose {
		output = utils.FormatVerbose(filteredWorktrees)
	} else {
		output = utils.FormatBasic(filteredWorktrees)
	}

	fmt.Print(output)
	return nil
}

// applyFilters applies name and branch filters to worktrees
func (lc *ListCommand) applyFilters(worktrees []models.Worktree) []models.Worktree {
	result := worktrees

	// Apply branch filter
	if lc.branch != "" {
		var filtered []models.Worktree
		for _, wt := range result {
			if wt.Branch == lc.branch {
				filtered = append(filtered, wt)
			}
		}
		result = filtered
	}

	// Apply name filter
	if lc.filter != "" {
		var filtered []models.Worktree
		for _, wt := range result {
			if containsSubstring(wt.Name, lc.filter) {
				filtered = append(filtered, wt)
			}
		}
		result = filtered
	}

	return result
}

// containsSubstring checks if a string contains a substring (case-insensitive)
func containsSubstring(str, substr string) bool {
	strLower := strings.ToLower(str)
	substrLower := strings.ToLower(substr)
	return strings.Contains(strLower, substrLower)
}

// SetVerbose sets verbose output mode
func (lc *ListCommand) SetVerbose(verbose bool) {
	lc.verbose = verbose
}

// SetFilter sets the name filter
func (lc *ListCommand) SetFilter(filter string) {
	lc.filter = filter
}

// SetBranch sets the branch filter
func (lc *ListCommand) SetBranch(branch string) {
	lc.branch = branch
}

// SetJSONOutput sets JSON output mode
func (lc *ListCommand) SetJSONOutput(jsonOutput bool) {
	lc.jsonOutput = jsonOutput
}

// RunListCommand is the entry point for the list command
func RunListCommand(repoPath, gitPath string, verbose, jsonOutput bool, filter, branch string) error {
	gitService := services.NewGitService(gitPath)
	listCmd := NewListCommand(gitService)
	listCmd.SetVerbose(verbose)
	listCmd.SetJSONOutput(jsonOutput)
	listCmd.SetFilter(filter)
	listCmd.SetBranch(branch)

	return listCmd.Execute(repoPath)
}
