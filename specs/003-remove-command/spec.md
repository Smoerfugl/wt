# Feature Specification: Remove Command

**Feature Branch**: `003-remove-command`  
**Created**: 2026-02-23  
**Status**: Draft  
**Input**: User description: "We need a remove command"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Non-interactive Remove (Priority: P1)

As a developer, I want to remove a worktree by path so I can clean up unused worktrees quickly with `wt remove <path>`.

**Independent Test**: Run `wt remove /path/to/wt` and verify `git worktree remove` is invoked and the worktree is removed.

**Acceptance Scenarios**:
1. **Given** a valid worktree path, **When** I run `wt remove <path>`, **Then** the worktree is removed and command exits 0.
2. **Given** the path is the main repository, **When** I run `wt remove <path>`, **Then** the command refuses to remove the main working tree and returns an error.

---

### User Story 2 - Interactive Remove (Priority: P2)

As a developer, I want to run `wt remove` without arguments to get an interactive list and choose which worktree to remove to avoid mistakes.

**Independent Test**: Run `wt remove` and simulate input selection to remove an entry.

**Acceptance Scenarios**:
1. **Given** multiple removable worktrees, **When** I run `wt remove`, **Then** the tool lists removable worktrees (excluding main) and prompts me to select one.
2. **Given** I input `q`, **When** prompted, **Then** the command cancels without removing anything.

---

### User Story 3 - Safety Checks (Priority: P3)

As a developer, I want safety checks that prevent accidental removal of the main worktree and confirm destructive actions.

**Independent Test**: Try to remove main worktree and verify command refuses; remove a non-main worktree and verify operation completes.

**Acceptance Scenarios**:
1. The tool must never remove the repository top-level worktree.
2. The tool should show a confirmation message when removing a worktree interactively.

---

### Edge Cases

- User selects invalid number in interactive mode
- Worktree path is missing or has special characters
- Git fails to remove worktree (locked, permission denied)

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST provide `remove` command supporting `wt remove [path]` and interactive flow if no path is provided.
- **FR-002**: System MUST exclude the repository's main worktree from removal options.
- **FR-003**: System MUST call `git worktree remove <path>` and surface errors.
- **FR-004**: System MUST prompt for confirmation in interactive mode.

### Key Entities

- **Worktree Entry**: path, branch, HEAD

## Success Criteria *(mandatory)*

- **SC-001**: `wt remove <path>` removes the worktree for valid non-main paths
- **SC-002**: `wt remove` interactive mode allows selection and removal with cancellation pathway
- **SC-003**: Attempts to remove main worktree are safely refused
