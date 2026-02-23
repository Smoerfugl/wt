# Quickstart: Exec Command

Run a command inside a selected worktree:

```bash
wt exec <command> [args...]
# example
wt exec ls -la
```

You will be shown a numbered list of non-main worktrees to choose from. The command runs in the selected worktree directory with stdin/stdout/stderr connected.
