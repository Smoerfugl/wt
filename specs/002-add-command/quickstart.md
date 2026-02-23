# Quickstart: Add Command

## Basic Usage

Add an existing branch or commit as a worktree:

```bash
wt add <branch|commit>
```

Create a new branch and worktree:

```bash
wt add -b <new-branch> [<start-point>]
# examples:
wt add -b feature-x
wt add -b feature-x origin/main
```

Output on success:

```
Worktree created at: /path/to/worktrees/repo/feature-x
```

Errors and troubleshooting:

- If the target path exists the command will fail; remove or choose a different name.
- Ensure you are in a Git repository.
- Make sure Git is installed and in PATH.
