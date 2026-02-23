# Feature Specification: Prune Command

**Feature Branch**: `004-prune-command`  
**Created**: 2026-02-23  
**Status**: Draft  
**Input**: User description: "We need a prune command"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Prune Stale Worktrees (Priority: P1)

As a developer, I want to run `wt prune` to prune stale worktrees using `git worktree prune` so the repository stays clean.

**Independent Test**: Run `wt prune` and verify `git worktree prune` is invoked and exit code reflects success/failure.

**Acceptance Scenarios**:
1. **Given** repository with stale worktrees, **When** I run `wt prune`, **Then** `git worktree prune` is executed and stale worktrees are cleaned up.

---

### Edge Cases

- Git command not available or failing
- No stale worktrees found (should return success with helpful message)

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST provide `prune` command invoking `git worktree prune` in repo root.
- **FR-002**: System MUST return non-zero exit code and clear error message when Git fails.

## Success Criteria *(mandatory)*

- **SC-001**: `wt prune` runs `git worktree prune` and returns 0 on success
- **SC-002**: When there are no stale worktrees, command prints a helpful message and returns 0
