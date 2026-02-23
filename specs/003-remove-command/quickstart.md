# Quickstart: Remove Command

## Remove by path

```bash
wt remove /path/to/worktree
```

## Interactive remove

```bash
wt remove
# Select the number of the worktree to remove, or 'q' to cancel
```

Notes:
- The tool will not list the primary repository worktree for removal.
- It runs `git worktree remove <path>` under the hood.
