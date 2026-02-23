# CLI Contract: List Command

## Command Signature

```
wt list [flags]
```

## Flags

| Flag | Shorthand | Type | Default | Description |
|------|-----------|------|---------|-------------|
| `--verbose` | `-v` | bool | false | Show detailed information including branch, commit, and status |
| `--filter` | `-f` | string | "" | Filter worktrees by name pattern (substring match) |
| `--branch` | `-b` | string | "" | Filter worktrees by branch name |
| `--json` | `-j` | bool | false | Output in JSON format for programmatic use |

## Output Formats

### Basic Format (default)
```
worktree1 (current)
worktree2
worktree3
```

### Verbose Format (`--verbose`)
```
NAME        PATH                    BRANCH    COMMIT    STATUS    CURRENT
worktree1   /path/to/worktree1       main      abc1234   clean     ✓
worktree2   /path/to/worktree2       feature   def5678   dirty     
worktree3   /path/to/worktree3       develop   ghi9012   clean     
```

### JSON Format (`--json`)
```json
{
  "repository": "/path/to/main/repo",
  "worktrees": [
    {
      "name": "worktree1",
      "path": "/path/to/worktree1",
      "branch": "main",
      "commit": "abc1234",
      "status": "clean",
      "current": true,
      "locked": false
    },
    {
      "name": "worktree2",
      "path": "/path/to/worktree2",
      "branch": "feature",
      "commit": "def5678",
      "status": "dirty",
      "current": false,
      "locked": false
    }
  ]
}
```

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Invalid arguments |
| 3 | Git repository not found |
| 4 | Git command failed |

## Error Messages

- "No worktrees found in repository"
- "Failed to execute git command: {error}"
- "Invalid filter pattern: {pattern}"
- "Branch not found: {branch}"
- "Not a git repository: {path}"

## Guarantees

1. **Idempotency**: Running `wt list` multiple times with same parameters produces same output
2. **Safety**: Command is read-only and never modifies repository state
3. **Performance**: Response time under 1 second for repositories with ≤50 worktrees
4. **Compatibility**: Works with Git versions 2.5+
5. **Cross-platform**: Consistent behavior on Linux, macOS, and Windows