# Feature Specification: Set upstream branch on worktree checkout

**Feature Branch**: `001-set-upstream-branch`
**Created**: 2026-02-26
**Status**: Draft
**Input**: User description: "the worktree supports checkout from master, however we need to also set the upstream branch to the branch created, and not the default ref"

## Clarifications


### Session 2026-02-26

- Q: When creating a new branch via the worktree, should the tool configure upstream immediately or defer? → A: Option A - Configure upstream immediately: set the new branch to track the chosen remote/branch locally (no automatic push).

- Q: If a requested branch name already exists locally or remotely, how should the tool behave? → A: Option A - Fail with a clear error and suggest a remedy (do not create the branch; instruct the user to pick a different name or use an explicit override).

- Q: If the base branch has no tracked remote, which remote should be selected for upstream? → A: Option A - Prefer remote named `origin` when present; if `origin` is absent, leave upstream unset and document that no remote was chosen.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Create new feature worktree (Priority: P1)

As a developer I want to create a new worktree for a branch derived from the repository default branch (e.g., master/main) so that subsequent push/pull operations target the newly created branch rather than the repository's default ref.

**Why this priority**: This is the primary workflow for feature development; incorrect upstream configuration causes accidental pushes to the default ref and breaks expected collaboration flows.

**Independent Test**: Create a new worktree with a new branch name, commit a change, then push without specifying remote/branch. Verify the remote receives the new branch and that the created branch's upstream is the created branch (not the default ref).

**Acceptance Scenarios**:

1. **Given** a repository with a configured remote and a local default branch (master/main), **When** a developer creates a new worktree with a new branch name, **Then** the new worktree exists, the branch's upstream is configured locally to track the chosen remote branch (no automatic push), and a subsequent push without explicit remote/branch targets the configured upstream.
2. **Given** a repository without any remote configured, **When** a developer creates a new worktree with a new branch name, **Then** the new worktree is created and no upstream to the repository default ref is created (upstream remains unset until the user configures/pushes).

---

### User Story 2 - Create worktree for an existing remote branch (Priority: P2)

As a developer I want to create a worktree that checks out an existing branch (remote or local) without changing its existing upstream configuration.

**Why this priority**: Users often create local worktrees for in-progress remote branches; changing upstream in this case can break collaboration or surprise other contributors.

**Independent Test**: Create a worktree that checks out an existing branch known to have an upstream; verify that the branch's upstream value is unchanged after worktree creation.

**Acceptance Scenarios**:

1. **Given** branch `feature/x` exists remotely and locally and already tracks `origin/feature/x`, **When** a developer creates a worktree for `feature/x`, **Then** the upstream remains `origin/feature/x`.

---

### User Story 3 - Detached-start-point or commit-based worktree (Priority: P3)

As a developer I want predictable upstream behavior when creating a worktree from a commit or tag: upstream should only be configured when a new branch name is explicitly created by the user.

**Why this priority**: Checkout-from-commit workflows are common for debugging; silently creating/updating upstreams during these flows is surprising.

**Independent Test**: Create a worktree from a commit (without -b/new-branch), verify no upstream is configured; then create a new branch from that commit with an explicit branch name and verify upstream is configured according to the rules in FR-001.

**Acceptance Scenarios**:

1. **Given** a commit hash and no -b flag, **When** a developer creates a worktree from that commit, **Then** no upstream for a branch is created.
2. **Given** a commit hash and an explicit new branch name, **When** a developer creates a worktree, **Then** the new branch is created and its upstream is configured following FR-001.

---

### Edge Cases

- Creating a branch name that already exists locally or remotely: ensure the command fails with a clear error or prompts to reuse the existing branch rather than silently changing upstreams.
- Repositories with multiple remotes: determine which remote is chosen for upstream (see Assumptions).
- Repositories where the default ref is not named "master" (e.g., "main"): behavior must be equivalent.
- Repositories without any remotes: ensure no upstream is set and behavior is documented.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: When a new branch is created via the worktree command (explicit -b/new-branch), and a remote is configured, the system MUST configure the branch's upstream locally to track the chosen remote/branch (for example, origin/<branch>) without performing an automatic push to the remote; pushes remain an explicit user action. This ensures push/pull without explicit remotes target the intended branch.
- **FR-002**: If the branch already has an upstream configured prior to creating the worktree, the system MUST NOT change that existing upstream mapping.
- **FR-003**: When creating a worktree from a commit or tag without an explicit new branch name, the system MUST NOT create or modify upstream configuration.
- **FR-004**: If multiple remotes exist, the system MUST use the remote selected by the base branch's tracking configuration (if any). If the base branch has no tracked remote, the system MUST prefer the remote named `origin` when present; if `origin` does not exist, the system MUST leave upstream unset and document that no remote was chosen.
- **FR-005**: The CLI MUST return clear, actionable messages when it cannot configure upstream (e.g., remote unreachable, name conflict) so users can safely proceed.
- **FR-006**: If the requested branch name already exists locally or remotely, the CLI MUST refuse to create a new branch and present a clear error describing resolution options (choose another name, explicitly reuse an existing branch, or retry with an explicit override). The tool MUST NOT silently reuse or overwrite existing branches.

### Key Entities *(include if feature involves data)*

- **Worktree**: Representation of a checkout at a path; associated with a branch or commit.
- **Branch**: The named branch created or checked out in the worktree; relevant attributes: name, upstream mapping.
- **Remote**: Optional remote repository (e.g., origin) used for tracking; used to determine upstream target.
- **User (developer)**: Person invoking the CLI commands; expectations for predictable push/pull behavior.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: When running the primary flow (User Story 1) in a repository with a configured remote, 100% of acceptance tests must observe the new branch's upstream is the newly created branch (i.e., pushing without explicit remote/branch targets the new branch).
- **SC-002**: When creating a worktree for an existing branch (User Story 2), 100% of acceptance tests must observe that the existing upstream mapping is preserved.
- **SC-003**: When creating a worktree from a commit without an explicit branch name (User Story 3), 100% of acceptance tests must observe no upstream is created or modified.
- **SC-004**: CLI error/messages for upstream configuration issues are present and rated as "clear" by at least one developer tester in manual verification (qualitative acceptance).

## Assumptions

- The repository may use "master" or "main" as its default ref; the feature must behave identically regardless of default ref name.
- If multiple remotes exist, the system will prefer the remote already associated with the base branch (if any). If there is no tracked remote on the base branch, the system will prefer a remote named "origin" when present; if no remotes exist, upstream is left unset.
- Repository tooling (Git) is available for verification by developer tests.

## Dependencies

- Behavior depends on the presence and reachability of configured remotes; tests must include both remote-present and remote-absent scenarios.

## Notes

- This specification focuses on user-visible behavior (what upstream is configured and when). It intentionally avoids prescribing the internal implementation steps; those belong in the planning phase.
