# data-model.md

## Entities

- Worktree
  - path: string (filesystem path)
  - branch: string or null (name of branch or null when detached)
  - start_point: commit-ish (branch/commit/tag used to create worktree)
  - created_at: timestamp

- Branch
  - name: string (e.g., feature/foo)
  - upstream: string or null (e.g., origin/feature/foo)
  - exists_local: boolean
  - exists_remote: boolean

- Remote
  - name: string (e.g., origin)
  - url: string

## Validation Rules

- Branch names must conform to Git ref name rules; reject invalid names.
- When creating a new named branch, ensure no local or remote branch with the same name exists unless user explicitly overrides.
- Upstream is only set when a branch name is explicitly created by the user (not on detached checkouts).

## State Transitions

- Creating worktree with new branch:
  - start: default branch (master/main)
  - action: create branch X and create worktree at path
  - end: worktree exists; branch X created; upstream configured per rules

- Creating worktree from commit (no branch):
  - start: commit hash
  - action: create worktree in detached HEAD
  - end: worktree exists; no branch created; no upstream configured
