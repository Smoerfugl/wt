# Feature Specification: Exec Command

**Feature Branch**: `005-exec-command`  
**Created**: 2026-02-23  
**Status**: Draft  
**Input**: User description: "We need an exec command"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Execute Command in Selected Worktree (Priority: P1)

As a developer, I want to run `wt exec <command> [args...]` and choose a worktree to run the command in so I can operate within that worktree's context.

**Independent Test**: Run `wt exec echo hello` and select a worktree; verify the command runs in the selected worktree directory and outputs expected results.

**Acceptance Scenarios**:
1. **Given** multiple non-main worktrees, **When** I run `wt exec ls`, **Then** I am prompted to select a worktree and the command runs in that directory.
2. **Given** no non-main worktrees exist, **When** I run `wt exec <cmd>`, **Then** the tool prints a helpful message and exits 0.

---

### Edge Cases

- Command not found (should surface error)
- Command requires stdin (should pass through os.Stdin)
- Worktree selection invalid input

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST provide `exec` command: `wt exec <command> [args...]`.
- **FR-002**: System MUST list non-main worktrees for selection and then execute the given command in the selected worktree directory with stdin/stdout/stderr attached.

## Success Criteria *(mandatory)*

- **SC-001**: Commands execute in chosen worktree with outputs forwarded to the user
- **SC-002**: Interactive selection is cancelable and safe
