// Repository represents a Git repository containing worktrees
package models

// Repository represents a Git repository
type Repository struct {
	MainPath   string     // Path to the main repository directory
	GitVersion string     // Version of Git being used
	Worktrees  []Worktree // List of all worktrees in this repository
}

// NewRepository creates a new Repository instance
func NewRepository(mainPath, gitVersion string) *Repository {
	return &Repository{
		MainPath:   mainPath,
		GitVersion: gitVersion,
		Worktrees:  []Worktree{},
	}
}

// AddWorktree adds a worktree to the repository
func (r *Repository) AddWorktree(worktree Worktree) {
	r.Worktrees = append(r.Worktrees, worktree)
}

// GetCurrentWorktree returns the current worktree
func (r *Repository) GetCurrentWorktree() *Worktree {
	for i := range r.Worktrees {
		if r.Worktrees[i].IsCurrent {
			return &r.Worktrees[i]
		}
	}
	return nil
}

// GetWorktreesByBranch returns worktrees filtered by branch
func (r *Repository) GetWorktreesByBranch(branch string) []Worktree {
	var result []Worktree
	for _, wt := range r.Worktrees {
		if wt.Branch == branch {
			result = append(result, wt)
		}
	}
	return result
}

// GetWorktreesByNamePattern returns worktrees filtered by name pattern
func (r *Repository) GetWorktreesByNamePattern(pattern string) []Worktree {
	var result []Worktree
	for _, wt := range r.Worktrees {
		if containsSubstring(wt.Name, pattern) {
			result = append(result, wt)
		}
	}
	return result
}

// containsSubstring checks if a string contains a substring (case-insensitive)
func containsSubstring(str, substr string) bool {
	return containsCI(str, substr)
}

// containsCI performs case-insensitive substring check
func containsCI(str, substr string) bool {
	for i := 0; i <= len(str)-len(substr); i++ {
		if matchSubstringCI(str[i:], substr) {
			return true
		}
	}
	return false
}

// matchSubstringCI checks if string starts with substring (case-insensitive)
func matchSubstringCI(str, substr string) bool {
	if len(str) < len(substr) {
		return false
	}
	for i := 0; i < len(substr); i++ {
		if toLower(str[i]) != toLower(substr[i]) {
			return false
		}
	}
	return true
}

// toLower converts byte to lowercase
func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}
