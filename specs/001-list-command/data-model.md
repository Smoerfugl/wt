# Data Model: List Command

## Entities

### Worktree
Represents a Git worktree with its properties and state.

**Attributes**:
- `Name` (string): The name of the worktree
- `Path` (string): Absolute path to the worktree directory
- `Branch` (string): The Git branch the worktree is on
- `CommitHash` (string): The current commit hash (short form)
- `IsCurrent` (bool): Whether this is the current worktree
- `IsClean` (bool): Whether the worktree has no uncommitted changes
- `IsLocked` (bool): Whether the worktree is locked

**Relationships**:
- Belongs to one Repository

### Repository
Represents the Git repository containing worktrees.

**Attributes**:
- `MainPath` (string): Path to the main repository directory
- `GitVersion` (string): Version of Git being used
- `Worktrees` ([]Worktree): List of all worktrees in this repository

## Data Flow

1. **Input**: User runs `wt list` command with optional flags
2. **Processing**:
   - GitService retrieves worktree information using `git worktree list`
   - Parser converts Git output to Worktree objects
   - FilterService applies any filter criteria
   - Formatter prepares output based on verbose flag
3. **Output**: Formatted worktree list displayed to user

## Validation Rules

- Worktree paths must be absolute and exist
- Worktree names must be non-empty
- Repository main path must be valid Git repository
- Branch names must be valid Git references

## State Transitions

N/A (This feature is read-only, no state changes)