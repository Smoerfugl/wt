# Quickstart: List Command

## Installation

Ensure you have the worktree manager CLI installed and in your PATH.

## Basic Usage

### List all worktrees
```bash
wt list
```

### List worktrees with detailed information
```bash
wt list --verbose
# or
wt list -v
```

### Filter worktrees by name
```bash
wt list --filter "feature"
# or
wt list -f "feature"
```

### Filter worktrees by branch
```bash
wt list --branch "main"
# or
wt list -b "main"
```

### Get JSON output for programmatic use
```bash
wt list --json
# or
wt list -j
```

## Examples

### Example 1: Basic listing
```bash
$ wt list
main-worktree (current)
feature-login
bugfix-auth
develop
```

### Example 2: Verbose listing
```bash
$ wt list --verbose
NAME            PATH                        BRANCH    COMMIT    STATUS    CURRENT
main-worktree   /home/user/project           main      a1b2c3   clean     ✓
feature-login  /home/user/project-login     feature   d4e5f6   dirty     
bugfix-auth    /home/user/project-auth      main      g7h8i9   clean     
develop        /home/user/project-develop   develop   j0k1l2   clean     
```

### Example 3: Filtered listing
```bash
$ wt list --filter "feature"
feature-login
feature-payment
feature-notifications
```

### Example 4: JSON output
```bash
$ wt list --json
{
  "repository": "/home/user/project",
  "worktrees": [
    {
      "name": "main-worktree",
      "path": "/home/user/project",
      "branch": "main",
      "commit": "a1b2c3",
      "status": "clean",
      "current": true,
      "locked": false
    },
    {
      "name": "feature-login",
      "path": "/home/user/project-login",
      "branch": "feature",
      "commit": "d4e5f6",
      "status": "dirty",
      "current": false,
      "locked": false
    }
  ]
}
```

## Troubleshooting

### No worktrees found
If you see "No worktrees found", ensure:
- You're in a Git repository
- You have worktrees created (use `git worktree list` to check)
- You have permission to access the worktree directories

### Permission denied errors
Ensure you have read access to all worktree directories and the main repository.

### Git command failed
Make sure Git is installed and available in your PATH.