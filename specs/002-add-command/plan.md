# Implementation Plan: Add Command

**Branch**: `002-add-command` | **Date**: 2026-02-23 | **Spec**: specs/002-add-command/spec.md

## Summary

The Add Command feature implements `wt add <branch|commit>` and `wt add -b <new-branch> [<start-point>]`. The command creates worktrees under `../worktrees/<repo>/<ref>` by default and integrates with native Git (`git worktree add`). When `-b` is used, the tool creates a new branch and the worktree for it, attempting to use the repository default ref when no start point is provided.

## Technical Context

- Language: Go 1.21+
- Integrations: Git CLI (`git worktree add`, `git symbolic-ref`, `git remote show origin`)
- Key files:
  - `cmd/wt/main.go` (add command implementation and path creation)
  - `services/git_service.go` (git integration utilities)

## Project Structure

```
cmd/wt/main.go          # CLI wiring, 'add' handling
services/git_service.go # Git helpers used by other commands
```

## Plan

1. Write failing unit tests around the behavior of path creation and default ref selection (if tests are included).
2. Implement path creation and `git worktree add` invocation in `cmd/wt/main.go` (already present) and harden error handling.
3. Add tests for `-b` flow: creating a branch and using default ref fallback logic (`defaultRef()` exists in `cmd/wt/main.go`).
4. Add quickstart, contracts, and documentation.

## Risks & Mitigations

- Race conditions creating directories: use `os.MkdirAll` and check for existing paths (do not overwrite).
- Default ref detection may fail in unusual remotes: fallback to HEAD then `main`/`master`.
