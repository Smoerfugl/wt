// Formatting provides utilities for formatting worktree output
package utils

import (
	"fmt"
	"strings"

	"worktree-manager/models"
)

// FormatBasic outputs worktrees in basic format
func FormatBasic(worktrees []models.Worktree) string {
	var result strings.Builder
	for _, wt := range worktrees {
		line := wt.Name
		if wt.IsCurrent {
			line += " (current)"
		}
		result.WriteString(line + "\n")
	}
	return result.String()
}

// FormatVerbose outputs worktrees in verbose table format
func FormatVerbose(worktrees []models.Worktree) string {
	if len(worktrees) == 0 {
		return "No worktrees found\n"
	}

	// Calculate column widths
	nameWidth := len("NAME")
	pathWidth := len("PATH")
	branchWidth := len("BRANCH")
	commitWidth := len("COMMIT")
	statusWidth := len("STATUS")
	currentWidth := len("CURRENT")

	for _, wt := range worktrees {
		if len(wt.Name) > nameWidth {
			nameWidth = len(wt.Name)
		}
		if len(wt.Path) > pathWidth {
			pathWidth = len(wt.Path)
		}
		if len(wt.Branch) > branchWidth {
			branchWidth = len(wt.Branch)
		}
		if len(wt.CommitHash) > commitWidth {
			commitWidth = len(wt.CommitHash)
		}
		if len(wt.GetStatus()) > statusWidth {
			statusWidth = len(wt.GetStatus())
		}
	}

	// Create header
	header := fmt.Sprintf("%-*s  %-*s  %-*s  %-*s  %-*s  %-*s\n",
		nameWidth, "NAME",
		pathWidth, "PATH",
		branchWidth, "BRANCH",
		commitWidth, "COMMIT",
		statusWidth, "STATUS",
		currentWidth, "CURRENT")

	result := header

	// Add rows
	for _, wt := range worktrees {
		currentMark := ""
		if wt.IsCurrent {
			currentMark = "✓"
		}
		row := fmt.Sprintf("%-*s  %-*s  %-*s  %-*s  %-*s  %-*s\n",
			nameWidth, wt.Name,
			pathWidth, wt.Path,
			branchWidth, wt.Branch,
			commitWidth, wt.CommitHash,
			statusWidth, wt.GetStatus(),
			currentWidth, currentMark)
		result += row
	}

	return result
}

// FormatJSON outputs worktrees in JSON format
func FormatJSON(worktrees []models.Worktree, repoPath string) string {
	var result strings.Builder
	result.WriteString("{\n")
	result.WriteString(fmt.Sprintf("  \"repository\": \"%s\",\n", repoPath))
	result.WriteString("  \"worktrees\": [\n")

	for i, wt := range worktrees {
		result.WriteString("    {\n")
		result.WriteString(fmt.Sprintf("      \"name\": \"%s\",\n", wt.Name))
		result.WriteString(fmt.Sprintf("      \"path\": \"%s\",\n", wt.Path))
		result.WriteString(fmt.Sprintf("      \"branch\": \"%s\",\n", wt.Branch))
		result.WriteString(fmt.Sprintf("      \"commit\": \"%s\",\n", wt.CommitHash))
		result.WriteString(fmt.Sprintf("      \"status\": \"%s\",\n", wt.GetStatus()))
		result.WriteString(fmt.Sprintf("      \"current\": %t,\n", wt.IsCurrent))
		result.WriteString(fmt.Sprintf("      \"locked\": %t\n", wt.IsLocked))
		if i < len(worktrees)-1 {
			result.WriteString("    },\n")
		} else {
			result.WriteString("    }\n")
		}
	}

	result.WriteString("  ]\n")
	result.WriteString("}\n")
	return result.String()
}
