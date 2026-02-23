// Worktree represents a Git worktree with its properties and state
package models

// Worktree represents a Git worktree
type Worktree struct {
	Name       string // The name of the worktree
	Path       string // Absolute path to the worktree directory
	Branch     string // The Git branch the worktree is on
	CommitHash string // The current commit hash (short form)
	IsCurrent  bool   // Whether this is the current worktree
	IsClean    bool   // Whether the worktree has no uncommitted changes
	IsLocked   bool   // Whether the worktree is locked
}

// NewWorktree creates a new Worktree instance
func NewWorktree(name, path, branch, commitHash string, isCurrent, isClean, isLocked bool) *Worktree {
	return &Worktree{
		Name:       name,
		Path:       path,
		Branch:     branch,
		CommitHash: commitHash,
		IsCurrent:  isCurrent,
		IsClean:    isClean,
		IsLocked:   isLocked,
	}
}

// GetStatus returns the status of the worktree
func (w *Worktree) GetStatus() string {
	if w.IsLocked {
		return "locked"
	}
	if w.IsClean {
		return "clean"
	}
	return "dirty"
}
