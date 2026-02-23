# Feature Specification: Add --exec Option to Worktree Add Command

**Feature Branch**: `002-exec`  
**Created**: 2026-02-23  
**Status**: Draft  
**Input**: User description: "add an option --exec opencode to add, it should execute the command on the worktree once it's created."

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

### User Story 1 - Execute Command After Worktree Creation (Priority: P1)

As a developer using worktree-manager, I want to execute a command immediately after creating a new worktree so that I can automate setup tasks like running tests, installing dependencies, or starting development servers without manual intervention.

**Why this priority**: This is the core functionality requested and provides immediate value by automating common post-creation tasks, saving developers time and reducing manual steps.

**Independent Test**: Can be fully tested by creating a worktree with the --exec flag and verifying the specified command executes successfully in the new worktree context.

**Acceptance Scenarios**:

1. **Given** a valid Git repository, **When** I run `wt add --exec "npm install" <branch>`, **Then** a new worktree should be created and `npm install` should execute in that worktree
2. **Given** a valid Git repository, **When** I run `wt add -b new-branch --exec "make test"`, **Then** a new worktree should be created on the new branch and `make test` should execute
3. **Given** an invalid command, **When** I run `wt add --exec "invalid-command" <branch>`, **Then** the worktree should still be created but the command execution should fail with an appropriate error message
4. **Given** a non-existent branch/commit, **When** I run `wt add --exec "echo hello" <invalid-ref>`, **Then** the command should not execute and an error should be shown

---

### User Story 2 - Command Execution in Correct Worktree Context (Priority: P2)

As a developer, I want commands executed via --exec to run in the correct worktree directory context so that file operations and environment variables work as expected.

**Why this priority**: Ensures the executed commands have the correct working directory and environment, which is essential for build tools, test runners, and other development commands.

**Independent Test**: Can be tested by creating a worktree with a command that outputs the current directory and verifying it matches the worktree path.

**Acceptance Scenarios**:

1. **Given** a Git repository, **When** I run `wt add --exec "pwd" <branch>`, **Then** the output should show the path to the newly created worktree directory
2. **Given** a Git repository, **When** I run `wt add --exec "echo $PWD" <branch>`, **Then** the PWD environment variable should reflect the worktree directory path

---

### User Story 3 - Error Handling and User Feedback (Priority: P3)

As a developer, I want clear error messages when --exec command execution fails so that I can quickly diagnose and fix issues.

**Why this priority**: Good error handling improves developer experience and reduces debugging time when things go wrong.

**Independent Test**: Can be tested by attempting to execute invalid commands or commands that fail and verifying appropriate error messages are displayed.

**Acceptance Scenarios**:

1. **Given** a Git repository, **When** I run `wt add --exec "nonexistent-command" <branch>`, **Then** an error message should indicate the command failed to execute
2. **Given** a Git repository, **When** I run `wt add --exec "exit 1" <branch>`, **Then** an error message should indicate the command returned non-zero exit code
3. **Given** a Git repository, **When** I run `wt add --exec "" <branch>`, **Then** an error message should indicate no command was provided

---

[Add more user stories as needed, each with an assigned priority]

### Edge Cases

- What happens when the command contains special characters or quotes that need escaping?
- How does system handle when the command execution takes a long time or hangs?
- Multiple --exec options are executed sequentially in the order provided
- When worktree creation succeeds but command execution fails, the worktree remains intact and an error is displayed
- Interactive commands are allowed but users are warned about potential issues with automated workflows

### Clarifications

#### Session 2026-02-23

- Q: How should the system handle multiple --exec options? → A: Run commands sequentially in the order provided
- Q: How should the system handle commands that require user input (interactive commands)? → A: Allow interactive commands but warn users about potential issues
- Q: What should be the timeout behavior for long-running commands? → A: Use a reasonable default timeout (e.g., 5 minutes) with option to override

## Requirements *(mandatory)*

<!--
  ACTION REQUIRED: The content in this section represents placeholders.
  Fill them out with the right functional requirements.
-->

### Functional Requirements

- **FR-001**: System MUST add a `--exec` (or `-e`) option to the `wt add` command
- **FR-002**: System MUST execute the specified command in the newly created worktree directory after successful worktree creation
- **FR-003**: System MUST preserve the original behavior of `wt add` when `--exec` is not provided
- **FR-004**: System MUST execute the command in the worktree's directory context (correct working directory)
- **FR-005**: System MUST handle command execution failures gracefully without affecting the created worktree
- **FR-006**: System MUST provide clear error messages when command execution fails
- **FR-007**: System MUST support commands with arguments and flags
- **FR-008**: System MUST handle quoted commands properly (supporting spaces and special characters)
- **FR-009**: System MUST validate that a command is provided when --exec flag is used
- **FR-010**: System MUST support multiple --exec options and execute them sequentially in the order provided
- **FR-011**: System MUST allow interactive commands but display a warning about potential issues
- **FR-012**: System MUST implement a default timeout of 5 minutes for command execution with option to override

### Key Entities *(include if feature involves data)*

- **Worktree**: Represents a Git worktree with path, branch, and associated metadata
- **Command**: The executable command and arguments to run in the worktree context
- **Execution Result**: Status and output of the command execution (success/failure, exit code, stdout/stderr)

## Success Criteria *(mandatory)*

<!--
  ACTION REQUIRED: Define measurable success criteria.
  These must be technology-agnostic and measurable.
-->

### Measurable Outcomes

- **SC-001**: Users can create a worktree and execute a command in a single operation
- **SC-002**: Command execution completes successfully in the correct worktree directory 95% of the time
- **SC-003**: Worktree creation succeeds even when command execution fails
- **SC-004**: Error messages clearly indicate command execution failures and their causes
- **SC-005**: The feature reduces manual steps by allowing automation of common post-creation tasks
- **SC-006**: Command execution time adds less than 2 seconds overhead to worktree creation for typical commands
- **SC-007**: Multiple commands execute sequentially with proper error handling for each
- **SC-008**: Long-running commands are handled with appropriate timeout behavior
