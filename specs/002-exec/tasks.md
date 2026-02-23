# Tasks: Add --exec Option to Worktree Add Command

**Input**: Design documents from `/specs/002-exec/`
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, contracts/

**Tests**: Tests are included as this is a CLI tool where testing is critical for reliability

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

Single project structure following worktree-manager conventions:
- `src/commands/` - CLI command implementations
- `src/services/` - Core business logic
- `src/models/` - Data models
- `src/utils/` - Utility functions
- `tests/unit/` - Unit tests
- `tests/integration/` - Integration tests

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [ ] T001 Verify Go 1.21+ environment setup
- [ ] T002 [P] Run `go mod tidy` to ensure dependencies are clean
- [ ] T003 [P] Configure linting with `go fmt` and `go vet`
- [ ] T004 [P] Setup test framework structure in tests/

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [ ] T005 Create command execution utility in src/utils/exec.go
- [ ] T006 [P] Implement Command struct with proper fields in src/utils/exec.go
- [ ] T007 [P] Implement ExecutionResult struct in src/utils/exec.go
- [ ] T008 Implement basic command execution with timeout in src/utils/exec.go
- [ ] T009 [P] Add error handling utilities in src/utils/errors.go
- [ ] T010 [P] Create unit tests for exec utilities in tests/unit/utils/exec_test.go

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Execute Command After Worktree Creation (Priority: P1) 🎯 MVP

**Goal**: Enable users to execute a single command automatically after worktree creation

**Independent Test**: Can be tested by running `wt add --exec "echo hello" main` and verifying the command executes in the new worktree

### Tests for User Story 1 (OPTIONAL - only if tests requested) ⚠️

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [ ] T011 [P] [US1] Unit test for --exec flag parsing in tests/unit/commands/add_test.go
- [ ] T012 [P] [US1] Unit test for single command execution in tests/unit/services/worktree_test.go
- [ ] T013 [P] [US1] Integration test for basic --exec functionality in tests/integration/add_exec_test.go

### Implementation for User Story 1

- [ ] T014 [P] [US1] Extend AddCommand struct with ExecCommands field in src/commands/add.go
- [ ] T015 [US1] Add --exec flag parsing to AddCommand in src/commands/add.go
- [ ] T016 [US1] Modify worktree creation service to accept commands in src/services/worktree.go
- [ ] T017 [US1] Implement single command execution after worktree creation in src/services/worktree.go
- [ ] T018 [US1] Add command validation in src/services/worktree.go
- [ ] T019 [US1] Implement error handling for command failures in src/services/worktree.go
- [ ] T020 [US1] Add CLI output messages for command execution status
- [ ] T021 [US1] Update help text for add command to include --exec documentation

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase 4: User Story 2 - Command Execution in Correct Worktree Context (Priority: P2)

**Goal**: Ensure commands execute in the correct worktree directory with proper environment

**Independent Test**: Can be tested by running `wt add --exec "pwd" main` and verifying output shows worktree path

### Tests for User Story 2 (OPTIONAL - only if tests requested) ⚠️

- [ ] T022 [P] [US2] Unit test for working directory handling in tests/unit/utils/exec_test.go
- [ ] T023 [P] [US2] Integration test for worktree context in tests/integration/add_exec_test.go

### Implementation for User Story 2

- [ ] T024 [US2] Enhance command execution to set working directory explicitly in src/utils/exec.go
- [ ] T025 [US2] Add worktree path resolution in src/services/worktree.go
- [ ] T026 [US2] Implement environment variable setup for commands (WT_WORKTREE, WT_BRANCH)
- [ ] T027 [US2] Add validation to ensure worktree exists before command execution
- [ ] T028 [US2] Update error messages to include worktree path context

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 5: User Story 3 - Error Handling and User Feedback (Priority: P3)

**Goal**: Provide clear error messages and graceful handling for command execution failures

**Independent Test**: Can be tested by running `wt add --exec "nonexistent-command" main` and verifying appropriate error messages

### Tests for User Story 3 (OPTIONAL - only if tests requested) ⚠️

- [ ] T029 [P] [US3] Unit test for error scenarios in tests/unit/utils/exec_test.go
- [ ] T030 [P] [US3] Integration test for error handling in tests/integration/add_exec_test.go

### Implementation for User Story 3

- [ ] T031 [US3] Implement comprehensive error classification in src/utils/errors.go
- [ ] T032 [US3] Add specific error messages for different failure types in src/utils/exec.go
- [ ] T033 [US3] Implement warning system for interactive commands
- [ ] T034 [US3] Add timeout handling with clear error messages
- [ ] T035 [US3] Implement command validation for empty commands
- [ ] T036 [US3] Add exit code handling and reporting

**Checkpoint**: All user stories should now be independently functional

---

## Phase 6: Additional Features from Clarifications

**Goal**: Implement features identified during clarification process

### Multiple Command Support

- [ ] T037 [P] Extend AddCommand to support multiple --exec flags in src/commands/add.go
- [ ] T038 [P] Implement command queue processing in src/services/worktree.go
- [ ] T039 [P] Add sequential execution logic in src/services/worktree.go
- [ ] T040 [P] Update error handling for multiple commands
- [ ] T041 [P] Add integration test for multiple commands

### Timeout Configuration

- [ ] T042 [P] Add timeout configuration to Command struct
- [ ] T043 [P] Implement default 5-minute timeout
- [ ] T044 [P] Add timeout override capability (future enhancement)
- [ ] T045 [P] Add timeout tests

---

## Phase 7: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [ ] T046 [P] Documentation updates in docs/commands/add.md
- [ ] T047 [P] Add examples to quickstart.md for common use cases
- [ ] T048 Code cleanup and refactoring for consistency
- [ ] T049 [P] Performance optimization for command execution
- [ ] T050 [P] Additional unit tests for edge cases in tests/unit/
- [ ] T051 Security review of command execution logic
- [ ] T052 [P] Cross-platform compatibility testing
- [ ] T053 Run quickstart.md validation and update examples
- [ ] T054 [P] Final integration testing across all scenarios

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Stories (Phase 3+)**: All depend on Foundational phase completion
  - User stories can then proceed in parallel (if staffed)
  - Or sequentially in priority order (P1 → P2 → P3)
- **Polish (Final Phase)**: Depends on all desired user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (Phase 2) - No dependencies on other stories
- **User Story 2 (P2)**: Can start after Foundational (Phase 2) - Enhances US1 but independently testable
- **User Story 3 (P3)**: Can start after Foundational (Phase 2) - Enhances US1/US2 but independently testable
- **Additional Features**: Can start after core user stories complete

### Within Each User Story

- Tests (if included) MUST be written and FAIL before implementation
- Models before services
- Services before command integration
- Core implementation before integration
- Story complete before moving to next priority

### Parallel Opportunities

- All Setup tasks marked [P] can run in parallel
- All Foundational tasks marked [P] can run in parallel (within Phase 2)
- Once Foundational phase completes, all user stories can start in parallel (if team capacity allows)
- All tests for a user story marked [P] can run in parallel
- Utility functions marked [P] can run in parallel
- Different user stories can be worked on in parallel by different team members

---

## Parallel Example: User Story 1

```bash
# Launch all tests for User Story 1 together:
Task: "Unit test for --exec flag parsing in tests/unit/commands/add_test.go"
Task: "Unit test for single command execution in tests/unit/services/worktree_test.go"
Task: "Integration test for basic --exec functionality in tests/integration/add_exec_test.go"

# Launch utility functions in parallel:
Task: "Create command execution utility in src/utils/exec.go"
Task: "Implement Command struct with proper fields in src/utils/exec.go"
Task: "Implement ExecutionResult struct in src/utils/exec.go"
```

---

## Parallel Example: Foundational Phase

```bash
# Launch foundational tasks in parallel:
Task: "Implement Command struct with proper fields in src/utils/exec.go"
Task: "Implement ExecutionResult struct in src/utils/exec.go"
Task: "Add error handling utilities in src/utils/errors.go"
Task: "Create unit tests for exec utilities in tests/unit/utils/exec_test.go"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational (CRITICAL - blocks all stories)
3. Complete Phase 3: User Story 1
4. **STOP and VALIDATE**: Test User Story 1 independently
5. Deploy/demo if ready

### Incremental Delivery

1. Complete Setup + Foundational → Foundation ready
2. Add User Story 1 → Test independently → Deploy/Demo (MVP!)
3. Add User Story 2 → Test independently → Deploy/Demo
4. Add User Story 3 → Test independently → Deploy/Demo
5. Add Additional Features → Test independently → Deploy/Demo
6. Each story adds value without breaking previous stories

### Parallel Team Strategy

With multiple developers:

1. Team completes Setup + Foundational together
2. Once Foundational is done:
   - Developer A: User Story 1 (Core --exec functionality)
   - Developer B: User Story 2 (Worktree context handling)
   - Developer C: User Story 3 (Error handling enhancements)
3. Stories complete and integrate independently
4. Additional features can be developed in parallel

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story for traceability
- Each user story should be independently completable and testable
- Verify tests fail before implementing
- Commit after each task or logical group
- Stop at any checkpoint to validate story independently
- Avoid: vague tasks, same file conflicts, cross-story dependencies that break independence
- Follow Go formatting standards (`go fmt`) after each implementation task
- Run `go vet` to catch potential issues early

## Task Summary

**Total Tasks**: 54
**Parallel Tasks**: 32 (59% parallelizable)
**User Story Breakdown**:
- US1 (P1 - MVP): 14 tasks (11 parallelizable)
- US2 (P2): 8 tasks (6 parallelizable)  
- US3 (P3): 8 tasks (6 parallelizable)
- Additional Features: 10 tasks (8 parallelizable)
- Setup/Foundational: 10 tasks (6 parallelizable)
- Polish: 12 tasks (8 parallelizable)

**Suggested MVP Scope**: Phase 1 + Phase 2 + Phase 3 (User Story 1) = 24 tasks
**Estimated MVP Completion**: 70% of core functionality with 14/54 tasks