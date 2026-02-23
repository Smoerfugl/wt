# CLI Contract: Remove Command

## Command Signature

```
wt remove [path]
```

If `path` is omitted the command runs interactively to select a removable worktree.

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Invalid arguments |
| 3 | Not a git repository |
| 4 | Git command failed |

## Guarantees

- The repository main worktree is never offered for removal.
- The command is read-only with the exception of removal operations invoked by git.
