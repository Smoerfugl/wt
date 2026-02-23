# Tasks: Add Command

## Phase 1: Tests
- [ ] Unit tests for defaultRef() behavior
- [ ] Integration test: `wt add -b new-branch` creates branch and worktree

## Phase 2: Implementation
- [x] T001 Ensure path creation logic exists in `cmd/wt/main.go` (already implemented)
- [ ] T002 Harden error handling when `git worktree add` fails
- [ ] T003 Add logging and user-friendly error messages

## Phase 3: Docs
- [ ] Update quickstart and CLI contract
