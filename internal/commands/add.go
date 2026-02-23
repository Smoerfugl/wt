// Add command implementation
package commands

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/smoerfugl/wt/internal/services"
	"github.com/smoerfugl/wt/internal/utils"
)

// AddCommand handles the 'wt add' command
type AddCommand struct {
	gitService   *services.GitService
	createBranch bool
	branchName   string
	startPoint   string
	execCommands []*utils.Command // Commands to execute after worktree creation
	worktreePath string
	verbose      bool
}

// NewAddCommand creates a new AddCommand instance
func NewAddCommand(gitService *services.GitService) *AddCommand {
	return &AddCommand{
		gitService: gitService,
	}
}

// SetCreateBranch sets whether to create a new branch
func (ac *AddCommand) SetCreateBranch(createBranch bool) {
	ac.createBranch = createBranch
}

// SetBranchName sets the branch name
func (ac *AddCommand) SetBranchName(branchName string) {
	ac.branchName = branchName
}

// SetStartPoint sets the start point (commit/branch)
func (ac *AddCommand) SetStartPoint(startPoint string) {
	ac.startPoint = startPoint
}

// AddExecCommand adds a command to execute after worktree creation
func (ac *AddCommand) AddExecCommand(cmd *utils.Command) {
	ac.execCommands = append(ac.execCommands, cmd)
}

// SetWorktreePath sets the worktree path
func (ac *AddCommand) SetWorktreePath(path string) {
	ac.worktreePath = path
}

// SetVerbose sets verbose output mode
func (ac *AddCommand) SetVerbose(verbose bool) {
	ac.verbose = verbose
}

// Execute runs the add command
func (ac *AddCommand) Execute(repoPath string) error {
	// Validate required fields
	if ac.branchName == "" {
		return fmt.Errorf("branch name is required")
	}

	// Determine the worktree path if not set
	if ac.worktreePath == "" {
		parentDir := filepath.Dir(repoPath)
		repoName := filepath.Base(repoPath)
		worktreesDir := filepath.Join(parentDir, "worktrees", repoName)
		ac.worktreePath = filepath.Join(worktreesDir, ac.branchName)
	}

	// Create worktrees directory if it doesn't exist
	if err := ac.gitService.EnsureWorktreesDir(repoPath); err != nil {
		return fmt.Errorf("failed to create worktrees directory: %w", err)
	}

	// Create the worktree
	var err error
	if ac.createBranch {
		// Create worktree with new branch
		if ac.startPoint != "" {
			err = ac.gitService.AddWorktreeWithBranch(ac.worktreePath, ac.branchName, ac.startPoint)
		} else {
			// Try to use default branch if no start point provided
			defaultRef, refErr := ac.gitService.GetDefaultRef(repoPath)
			if refErr == nil {
				err = ac.gitService.AddWorktreeWithBranch(ac.worktreePath, ac.branchName, defaultRef)
			} else {
				// Fall back to HEAD
				err = ac.gitService.AddWorktreeWithBranch(ac.worktreePath, ac.branchName, "")
			}
		}
	} else {
		// Create worktree from existing ref
		err = ac.gitService.AddWorktree(ac.worktreePath, ac.branchName)
	}

	if err != nil {
		return fmt.Errorf("failed to create worktree: %w", err)
	}

	fmt.Printf("Worktree created at: %s\n", ac.worktreePath)

	// Execute commands if any
	if len(ac.execCommands) > 0 {
		if err := ac.executeCommands(); err != nil {
			return fmt.Errorf("command execution failed: %w", err)
		}
	}

	return nil
}

// executeCommands executes all commands in the worktree directory
func (ac *AddCommand) executeCommands() error {
	if len(ac.execCommands) == 0 {
		return nil
	}

	fmt.Printf("Executing %d command(s) in worktree...\n", len(ac.execCommands))

	// Set working directory for all commands
	for _, cmd := range ac.execCommands {
		cmd.Dir = ac.worktreePath

		// Add worktree-specific environment variables
		cmd.Env = append(cmd.Env,
			fmt.Sprintf("WT_WORKTREE=%s", ac.worktreePath),
			fmt.Sprintf("WT_BRANCH=%s", ac.branchName),
		)
	}

	// Execute commands sequentially
	results, err := utils.ExecuteCommands(ac.execCommands)
	if err != nil {
		return fmt.Errorf("failed to execute commands: %w", err)
	}

	// Report results
	allSuccess := true
	for i, result := range results {
		cmdDesc := ac.execCommands[i].Name + " " + strings.Join(ac.execCommands[i].Args, " ")
		if result.Success {
			fmt.Printf("✓ %s completed successfully\n", cmdDesc)
		} else {
			fmt.Printf("✗ %s failed: %s\n", cmdDesc, result.Error)
			allSuccess = false
		}
	}

	if !allSuccess {
		return fmt.Errorf("one or more commands failed")
	}

	return nil
}

// RunAddCommand is the entry point for the add command
func RunAddCommand(repoPath, gitPath string, createBranch, verbose bool, branchName, startPoint string, execCommands []*utils.Command) error {
	gitService := services.NewGitService(gitPath)
	addCmd := NewAddCommand(gitService)
	addCmd.SetCreateBranch(createBranch)
	addCmd.SetBranchName(branchName)
	addCmd.SetStartPoint(startPoint)
	addCmd.SetVerbose(verbose)

	// Add exec commands
	for _, cmd := range execCommands {
		addCmd.AddExecCommand(cmd)
	}

	return addCmd.Execute(repoPath)
}
