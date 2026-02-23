---

description: "Task list for version command implementation"
---

# Tasks: Version Command

**Input**: Design documents from `/specs/001-version-command/`
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, contracts/

**Tests**: Tests are included as this is a core CLI feature requiring validation

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [x] T001 Create version model structure in internal/models/version.go
- [x] T002 Create version command structure in internal/commands/version.go
- [x] T003 [P] Update main.go to include version command case in switch statement

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [x] T004 Create VersionInfo model with semantic version validation in internal/models/version.go
- [x] T005 [P] Create version formatter utility in internal/utils/version_formatter.go
- [x] T006 Create basic version command implementation in internal/commands/version.go
- [x] T007 Add version command to main command switch in cmd/wt/main.go
- [x] T008 Add version command to help output in cmd/wt/main.go

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Display Application Version (Priority: P1) 🎯 MVP

**Goal**: Implement basic version command that displays application version

**Independent Test**: Run `wt version` and verify it outputs version information

### Tests for User Story 1 (OPTIONAL - only if tests requested) ⚠️

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [x] T009 [P] [US1] Unit test for version parsing in internal/models/version_test.go
- [x] T010 [P] [US1] Unit test for text formatting in internal/utils/version_formatter_test.go
- [x] T011 [P] [US1] Integration test for basic version command in integration/version_test.go

### Implementation for User Story 1

- [x] T012 [P] [US1] Implement VersionInfo struct with validation in internal/models/version.go
- [x] T013 [P] [US1] Implement text formatter for version output in internal/utils/version_formatter.go
- [x] T014 [US1] Implement basic version command execution in internal/commands/version.go
- [x] T015 [US1] Add default version fallback (0.0.1-dev) when no version info available
- [x] T016 [US1] Implement error handling for version command
- [x] T017 [US1] Add version command to main switch with basic flag parsing

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase 4: User Story 2 - Version Information Format (Priority: P2)

**Goal**: Add JSON output format and ensure semantic versioning compliance

**Independent Test**: Run `wt version --json` and verify valid JSON output

### Tests for User Story 2 (OPTIONAL - only if tests requested) ⚠️

- [x] T018 [P] [US2] Unit test for JSON formatting in internal/utils/version_formatter_test.go
- [x] T019 [P] [US2] Integration test for JSON output in integration/version_test.go
- [x] T020 [P] [US2] Contract test for semantic version validation in tests/contract/version_contract_test.go

### Implementation for User Story 2

- [x] T021 [P] [US2] Implement JSON formatter for version output in internal/utils/version_formatter.go
- [x] T022 [US2] Add --json flag parsing to version command in internal/commands/version.go
- [x] T023 [US2] Implement semantic version validation regex in internal/models/version.go
- [x] T024 [US2] Add JSON output format to version command execution
- [x] T025 [US2] Update version command help text to include --json flag

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 5: User Story 3 - Help Integration (Priority: P3)

**Goal**: Ensure version command is discoverable through help system

**Independent Test**: Run `wt help` and `wt version --help` to verify help integration

### Tests for User Story 3 (OPTIONAL - only if tests requested) ⚠️

- [x] T026 [P] [US3] Integration test for help integration in integration/version_help_test.go

### Implementation for User Story 3

- [x] T027 [US3] Add version command description to main help output in cmd/wt/main.go
- [x] T028 [US3] Implement detailed help for version command in internal/commands/version.go
- [x] T029 [US3] Add --help flag handling to version command
- [x] T030 [US3] Update usage() function to include version command

**Checkpoint**: All user stories should now be independently functional

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [x] T031 [P] Add build information extraction using debug.ReadBuildInfo
- [x] T032 [P] Add optional fields (buildDate, gitCommit, goVersion, platform) to VersionInfo
- [x] T033 Update version formatter to include optional build information
- [x] T034 Add comprehensive error handling and exit codes
- [x] T035 [P] Add unit tests for edge cases (missing version info, invalid formats)
- [x] T036 [P] Add integration tests for all output formats
- [ ] T037 Update README.md with version command usage examples
- [ ] T038 Run quickstart.md validation tests
- [x] T039 Performance optimization - ensure <100ms execution time

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
- **User Story 2 (P2)**: Can start after Foundational (Phase 2) - Enhances US1 but should be independently testable
- **User Story 3 (P3)**: Can start after Foundational (Phase 2) - UI enhancement, independent of core functionality

### Within Each User Story

- Tests (if included) MUST be written and FAIL before implementation
- Models before services
- Services before command implementation
- Core implementation before integration
- Story complete before moving to next priority

### Parallel Opportunities

- All Setup tasks marked [P] can run in parallel
- All Foundational tasks marked [P] can run in parallel (within Phase 2)
- Once Foundational phase completes, all user stories can start in parallel (if team capacity allows)
- All tests for a user story marked [P] can run in parallel
- Model and utility implementations within a story marked [P] can run in parallel
- Different user stories can be worked on in parallel by different team members

---

## Parallel Example: User Story 1

```bash
# Launch all tests for User Story 1 together:
Task: "Unit test for version parsing in internal/models/version_test.go"
Task: "Unit test for text formatting in internal/utils/version_formatter_test.go"
Task: "Integration test for basic version command in integration/version_test.go"

# Launch all models and utilities for User Story 1 together:
Task: "Implement VersionInfo struct with validation in internal/models/version.go"
Task: "Implement text formatter for version output in internal/utils/version_formatter.go"
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
5. Each story adds value without breaking previous stories

### Parallel Team Strategy

With multiple developers:

1. Team completes Setup + Foundational together
2. Once Foundational is done:
   - Developer A: User Story 1 (Core version display)
   - Developer B: User Story 2 (JSON format enhancement)
   - Developer C: User Story 3 (Help integration)
3. Stories complete and integrate independently

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story for traceability
- Each user story should be independently completable and testable
- Verify tests fail before implementing
- Commit after each task or logical group
- Stop at any checkpoint to validate story independently
- Avoid: vague tasks, same file conflicts, cross-story dependencies that break independence