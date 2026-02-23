# Feature Specification: List Command

**Feature Branch**: `001-list-command`  
**Created**: 2026-02-23  
**Status**: Draft  
**Input**: User description: "We need a list comand"

## User Scenarios & Testing *(mandatory)*

<!--
  IMPORTANT: User stories should be PRIORITIZED as user journeys ordered by importance.
  Each user story/journey must be INDEPENDENTLY TESTABLE - meaning if you implement just ONE of them,
  you should still have a viable MVP (Minimum Viable Product) that delivers value.
  
  Assign priorities (P1, P2, P3, etc.) to each story, where P1 is the most critical.
  Think of each story as a standalone slice of functionality that can be:
  - Developed independently
  - Tested independently
  - Deployed independently
  - Demonstrated to users independently
-->

### User Story 1 - List All Worktrees (Priority: P1)

As a developer using the worktree manager CLI, I want to see a list of all existing worktrees so I can quickly identify and navigate to the worktree I need to work on.

**Why this priority**: This is the most fundamental operation for worktree management - users need to see what worktrees exist before they can work with them. Without this, users cannot effectively use the tool.

**Independent Test**: Can be fully tested by running the list command and verifying it displays all existing worktrees with their basic information.

**Acceptance Scenarios**:

1. **Given** a repository with multiple worktrees, **When** I run `wt list`, **Then** I should see all worktrees listed with their names and paths
2. **Given** a repository with no worktrees, **When** I run `wt list`, **Then** I should see a message indicating no worktrees exist
3. **Given** a repository with worktrees, **When** I run `wt list`, **Then** the current worktree should be clearly indicated

---

### User Story 2 - List Worktrees with Details (Priority: P2)

As a developer, I want to see detailed information about each worktree including branch name, commit hash, and status so I can understand the state of each worktree before switching to it.

**Why this priority**: While less critical than basic listing, detailed information helps users make informed decisions about which worktree to use.

**Independent Test**: Can be fully tested by running the list command with a verbose flag and verifying detailed information is displayed.

**Acceptance Scenarios**:

1. **Given** a repository with worktrees, **When** I run `wt list --verbose`, **Then** I should see additional details like branch name, commit hash, and clean/dirty status for each worktree
2. **Given** a repository with worktrees, **When** I run `wt list --verbose`, **Then** the output should be formatted in a readable table format

---

### User Story 3 - Filter and Search Worktrees (Priority: P3)

As a developer with many worktrees, I want to filter or search the list of worktrees by name or branch so I can quickly find the specific worktree I need.

**Why this priority**: This enhances usability for power users with many worktrees but is not essential for basic functionality.

**Independent Test**: Can be fully tested by running the list command with filter options and verifying only matching worktrees are displayed.

**Acceptance Scenarios**:

1. **Given** a repository with multiple worktrees, **When** I run `wt list --filter "feature"`, **Then** I should only see worktrees whose names contain "feature"
2. **Given** a repository with multiple worktrees, **When** I run `wt list --branch "main"`, **Then** I should only see worktrees based on the main branch

---

[Add more user stories as needed, each with an assigned priority]

### Edge Cases

- What happens when the repository has no worktrees? (Should display helpful message)
- How does system handle worktrees with very long names or paths? (Should truncate or wrap appropriately)
- What happens when worktree paths contain special characters or spaces? (Should handle properly)
- How does system handle permission issues when accessing worktree directories? (Should show error for inaccessible worktrees)

## Requirements *(mandatory)*

<!--
  ACTION REQUIRED: The content in this section represents placeholders.
  Fill them out with the right functional requirements.
-->

### Functional Requirements

- **FR-001**: System MUST provide a `list` command that displays all existing worktrees in the repository
- **FR-002**: System MUST show basic information for each worktree including name and path
- **FR-003**: System MUST clearly indicate which worktree is currently active
- **FR-004**: System MUST handle repositories with no worktrees gracefully with appropriate messaging
- **FR-005**: System MUST support a verbose mode that shows additional details (branch, commit, status)
- **FR-006**: System MUST support filtering worktrees by name pattern
- **FR-007**: System MUST support filtering worktrees by branch name
- **FR-008**: System MUST handle edge cases like long paths, special characters, and permission issues
- **FR-009**: System MUST format output in a readable, consistent manner
- **FR-010**: System MUST maintain compatibility with native Git worktree functionality

### Key Entities *(include if feature involves data)*

- **Worktree**: Represents a Git worktree with attributes: name, path, branch, commit hash, status (clean/dirty), current flag
- **Repository**: Represents the Git repository containing worktrees

## Success Criteria *(mandatory)*

<!--
  ACTION REQUIRED: Define measurable success criteria.
  These must be technology-agnostic and measurable.
-->

### Measurable Outcomes

- **SC-001**: Users can view all worktrees in under 1 second for repositories with up to 50 worktrees
- **SC-002**: Users can identify the current worktree at a glance with clear visual indication
- **SC-003**: Users can find specific worktrees using filter options in under 2 seconds
- **SC-004**: System handles edge cases gracefully with appropriate error messages
- **SC-005**: 95% of users can successfully identify and navigate to their desired worktree using the list command
