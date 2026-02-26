# Contract: `wt add` (create worktree + branch)

Purpose: Define the command schema, flags, behavior, and exit codes for creating a worktree and (optionally) a new branch.

Path: `/cmd/wt`

Command: `wt add [options] <start-point>`

Flags:
- `-b, --branch <name>`: Create and check out a new branch with the given name in the worktree. If omitted and `<start-point>` is a branch name, checks out that branch instead.
- `--force`: Explicitly override existing branch name conflicts (must be used intentionally).
- `--remote <name>`: Override the remote used to set upstream (defaults to base branch's tracked remote or `origin`).
- `--path <dir>`: Specify worktree path (optional; default auto-generated)
- `--no-upstream`: Explicitly request no upstream be configured even when a remote exists.

Behavior:
- If `--branch` is provided:
  - Validate branch name; if invalid, exit 2 with error message.
  - If branch already exists (local or remote) and `--force` not provided, exit 3 with a clear error explaining options.
  - Create the branch from `<start-point>` and create worktree at path.
  - Configure upstream locally to `remote/<branch>` based on rules in spec unless `--no-upstream` provided.
- If `--branch` not provided and `<start-point>` is a branch name:
  - Create worktree checking out the existing branch; preserve upstream configuration.
- If `<start-point>` is a commit/tag and `--branch` not provided:
  - Create detached worktree; do not create or configure upstream.

Exit codes:
- `0` - Success
- `1` - General error (unexpected)
- `2` - Invalid branch name
- `3` - Branch name exists and conflict not resolved (no `--force`)

Examples:
- `wt add -b feature/xyz master` — create branch `feature/xyz` from master in a new worktree and set upstream to chosen remote.
- `wt add release/1.2` — create a worktree checking out existing branch `release/1.2` and preserve upstream.
- `wt add 2a4f3b` — create a detached worktree at the given commit; no upstream configured.
