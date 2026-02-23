# CLI Contract: Add Command

## Command Signature

```
wt add [-b] <branch|commit>
```

## Flags

| Flag | Shorthand | Type | Default | Description |
|------|-----------|------|---------|-------------|
| `-b` | `-b` | bool | false | Create a new branch with the given name before adding the worktree |

## Behavior

- `wt add <ref>` calls `git worktree add <path> <ref>` where path is `../worktrees/<repo>/<ref>`
- `wt add -b <new-branch> [<start-point>]` calls `git worktree add -b <new-branch> <path> [<start-point>]` and attempts to derive `start-point` when omitted.

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Invalid arguments |
| 3 | Not a git repository |
| 4 | Git command failed |
