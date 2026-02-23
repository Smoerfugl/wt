# Tasks: List Command

**Input**: Design documents from `/specs/001-list-command/`
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, contracts/

**Tests**: Tests are included as specified in the feature requirements

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- **Single project**: `src/`, `tests/` at repository root

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [x] T001 Create project structure per implementation plan
- [x] T002 Initialize Go project with standard library dependencies
- [x] T003 [P] Configure Go formatting and vet tools
- [x] T004 [P] Setup test framework and coverage tools

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [x] T005 Create Worktree model in src/models/worktree.go
- [x] T006 Create Repository model in src/models/repository.go
- [x] T007 [P] Implement basic Git service interface in src/services/git_service.go
- [x] T008 [P] Implement error handling and logging infrastructure
- [x] T009 Create main CLI structure in cmd/wt/main.go

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - List All Worktrees (Priority: P1) 🎯 MVP

**Goal**: Implement basic worktree listing functionality

**Independent Test**: Run `wt list` and verify it displays all worktrees with names and paths

### Tests for User Story 1 (OPTIONAL - only if tests requested) ⚠️

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [ ] T010 [P] [US1] Unit test for Worktree model in tests/unit/models/worktree_test.go
- [ ] T011 [P] [US1] Unit test for Git service worktree listing in tests/unit/services/git_test.go
- [ ] T012 [P] [US1] Integration test for basic list command in tests/integration/list_test.go

### Implementation for User Story 1

- [x] T013 [P] [US1] Implement Git worktree listing in src/services/git_service.go
- [x] T014 [P] [US1] Implement basic output formatting in src/utils/formatting.go
- [x] T015 [US1] Implement list command in src/commands/list.go
- [x] T016 [US1] Add current worktree indication logic
- [x] T017 [US1] Add error handling for no worktrees case
- [x] T018 [US1] Add basic command help text

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase 4: User Story 2 - List Worktrees with Details (Priority: P2)

**Goal**: Add verbose mode with detailed worktree information

**Independent Test**: Run `wt list --verbose` and verify detailed information is displayed in table format

### Tests for User Story 2 (OPTIONAL - only if tests requested) ⚠️

- [ ] T019 [P] [US2] Unit test for detailed worktree parsing in tests/unit/services/git_test.go
- [ ] T020 [P] [US2] Unit test for table formatting in tests/unit/utils/formatting_test.go
- [ ] T021 [P] [US2] Integration test for verbose list command in tests/integration/list_test.go

### Implementation for User Story 2

- [x] T022 [P] [US2] Extend Git service to parse detailed worktree info in src/services/git_service.go
- [x] T023 [P] [US2] Implement table formatting for verbose output in src/utils/formatting.go
- [x] T024 [US2] Add --verbose flag handling in src/commands/list.go
- [x] T025 [US2] Add branch, commit, and status fields to Worktree model
- [x] T026 [US2] Implement clean/dirty status detection

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 5: User Story 3 - Filter and Search Worktrees (Priority: P3)

**Goal**: Add filtering capabilities by name and branch

**Independent Test**: Run `wt list --filter "feature"` and `wt list --branch "main"` to verify filtering works

### Tests for User Story 3 (OPTIONAL - only if tests requested) ⚠️

- [ ] T027 [P] [US3] Unit test for filter service in tests/unit/services/filter_test.go
- [ ] T028 [P] [US3] Integration test for name filtering in tests/integration/list_test.go
- [ ] T029 [P] [US3] Integration test for branch filtering in tests/integration/list_test.go

### Implementation for User Story 3

- [x] T030 [P] [US3] Implement filter service in src/commands/list.go (integrated directly)
- [x] T031 [US3] Add --filter flag handling in src/commands/list.go
- [x] T032 [US3] Add --branch flag handling in src/commands/list.go
- [x] T033 [US3] Implement name pattern matching logic in src/commands/list.go
- [x] T034 [US3] Implement branch name filtering logic in src/commands/list.go

**Checkpoint**: All user stories should now be independently functional

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Final improvements and quality assurance

- [x] T035 [P] Add JSON output format support in src/utils/formatting.go
- [x] T036 [P] Add --json flag handling in src/commands/list.go
- [ ] T037 [P] Implement comprehensive error handling for edge cases
- [ ] T038 [P] Add input validation for all flags
- [ ] T039 [P] Optimize performance for large numbers of worktrees
- [ ] T040 [P] Add comprehensive Go documentation to all public functions
- [ ] T041 [P] Update quickstart.md with final examples
- [ ] T042 [P] Run full test suite and achieve 80%+ coverage
- [ ] T043 [P] Validate all CLI contract requirements are met

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
- **User Story 2 (P2)**: Can start after Foundational (Phase 2) - Builds on US1 but independently testable
- **User Story 3 (P3)**: Can start after Foundational (Phase 2) - Builds on US1 but independently testable

### Within Each User Story

- Tests (if included) MUST be written and FAIL before implementation
- Models before services
- Services before commands
- Core implementation before integration
- Story complete before moving to next priority

### Parallel Opportunities

- All Setup tasks marked [P] can run in parallel
- All Foundational tasks marked [P] can run in parallel (within Phase 2)
- Once Foundational phase completes, all user stories can start in parallel (if team capacity allows)
- All tests for a user story marked [P] can run in parallel
- Models within a story marked [P] can run in parallel
- Different user stories can be worked on in parallel by different team members

---

## Parallel Example: User Story 1

```bash
# Launch all tests for User Story 1 together (if tests requested):
Task: "Unit test for Worktree model in tests/unit/models/worktree_test.go"
Task: "Unit test for Git service worktree listing in tests/unit/services/git_test.go"
Task: "Integration test for basic list command in tests/integration/list_test.go"

# Launch all models for User Story 1 together:
Task: "Create Worktree model in src/models/worktree.go"
Task: "Create Repository model in src/models/repository.go"
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
   - Developer A: User Story 1
   - Developer B: User Story 2
   - Developer C: User Story 3
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