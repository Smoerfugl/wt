# Feature Specification: Add Command

**Feature Branch**: `002-add-command`  
**Created**: 2026-02-23  
**Status**: Draft  
**Input**: User description: "We need an add command"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Add Existing Branch or Commit (Priority: P1)

As a developer, I want to add an existing branch or a commit as a new worktree so I can work on it in parallel without affecting the main working tree.

**Independent Test**: Run `wt add <branch|commit>` and verify a worktree directory is created and `git worktree add` was called.

**Acceptance Scenarios**:
1. **Given** a valid branch name, **When** I run `wt add feature-x`, **Then** a new worktree is created at `../worktrees/<repo>/<feature-x>` and the command exits 0.
2. **Given** a commit hash, **When** I run `wt add a1b2c3d`, **Then** a detached worktree is created at the expected path and the commit is checked out there.
3. **Given** the target worktree path already exists, **When** I run `wt add <ref>`, **Then** the command returns an error and does not overwrite existing data.

---

### User Story 2 - Create New Branch + Worktree (Priority: P2)

As a developer, I want `wt add -b <new-branch> [<start-point>]` to create a new branch and an associated worktree in a sensible location.

**Independent Test**: Run `wt add -b new-feature` with and without a start point and verify the branch is created and the worktree is added.

**Acceptance Scenarios**:
1. **Given** no start point, **When** I run `wt add -b new-feature`, **Then** the tool tries to use repository default ref (e.g. `origin/main`) or falls back to `HEAD`, and creates the new branch and worktree.
2. **Given** a start point `origin/main`, **When** I run `wt add -b new-feature origin/main`, **Then** the new branch is created at `origin/main` and the worktree is created.

---

### User Story 3 - Robust Error Handling (Priority: P3)

As a developer, I want sensible errors when Git fails, permissions are missing, or arguments are invalid.

**Independent Test**: Simulate Git failures or pass invalid args and verify the command returns non-zero and prints a clear error.

**Acceptance Scenarios**:
1. Invalid args return exit code 2 and usage text.
2. Not a git repository returns a clear error.
3. Git command failures surface the underlying error message.

---

### Edge Cases

- Branch name conflicts (branch already exists)
- Start point reference not found
- Insufficient filesystem permissions creating worktree directories
- Very long branch names or refs with unusual characters

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST provide `add` command: `wt add <branch|commit>` and `wt add -b <new-branch> [<start-point>]`.
- **FR-002**: System MUST create worktree directories under `../worktrees/<repoName>/<ref>` relative to repository top-level.
- **FR-003**: System MUST create a new branch when `-b` is used and optionally use a start point.
- **FR-004**: System MUST attempt to use repository default ref when no start point is provided.
- **FR-005**: System MUST not overwrite existing directories and must fail with a helpful error.
- **FR-006**: System MUST surface Git errors and return non-zero exit codes on failure.

### Key Entities

- **Worktree**: name, path, branch/ref, commit hash, metadata
- **Repository**: top-level path, default ref detection

## Success Criteria *(mandatory)*

- **SC-001**: `wt add <ref>` creates a new worktree directory and returns 0 for valid refs
- **SC-002**: `wt add -b <branch>` creates a branch and worktree using the default ref when start point omitted
- **SC-003**: Errors for invalid inputs are returned with a clear message and non-zero exit code
- **SC-004**: Creation completes in under 2 seconds for typical repos (small files)
