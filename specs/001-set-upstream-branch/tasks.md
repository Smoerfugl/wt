---

description: "Task list for 'Set upstream branch on worktree checkout'"
---

# Tasks: Set upstream branch on worktree checkout

**Input**: Design documents from `/home/chrisjust/Projects/worktree-manager/specs/001-set-upstream-branch/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure required to implement feature

- [ ] T001 [P] Create integration test directory and helper at `tests/integration/helpers.go` (scaffold a TestRepo helper that can init a temp git repo for integration tests)
- [ ] T002 Update `cmd/wt/main.go` add and parse new flags for `add`: `--force` (bool), `--remote <name>` (string), `--path <dir>` (string), `--no-upstream` (bool); wire these variables into the `RunAddCommand` call so they can be passed to command implementation (modify call site in `cmd/wt/main.go`)
- [ ] T003 [P] Ensure repository can build locally: run `go build ./...` at repo root and verify `cmd/wt` builds (CI-ready check). Document failures to `specs/001-set-upstream-branch/notes/build-errors.md`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core plumbing required before user story implementation

- [ ] T004 Modify `internal/services/git_service.go` to add the following functions (implement using `exec.Command("git", ...)` and robust error handling):
  - `BranchExistsLocal(repoPath, branchName) (bool, error)` — checks `refs/heads/<branch>`
  - `BranchExistsRemote(repoPath, remoteName, branchName) (bool, error)` — checks `git ls-remote --heads <remote> <branch>`
  - `HasRemote(repoPath, remoteName) (bool, error)` — checks `git remote` list
  - `SetBranchUpstream(repoPath, branchName, remoteName) error` — runs `git branch --set-upstream-to=<remoteName>/<branchName> <branchName>` without pushing
  - `GetBranchUpstream(repoPath, branchName) (string, error)` — returns upstream like `origin/feature/foo` or empty

- [ ] T005 Update `internal/commands/add.go` (AddCommand) to accept new options and wire them through:
  - Add fields to `AddCommand` struct: `force bool`, `remoteOverride string`, `noUpstream bool`
  - Add setters: `SetForce(bool)`, `SetRemoteOverride(string)`, `SetNoUpstream(bool)`
  - Update `RunAddCommand` (internal/commands/add.go) signature to accept `force`, `remoteOverride`, `noUpstream` and pass them into the AddCommand instance

- [ ] T006 Update `internal/commands/add.go` execution path to perform pre-checks before creating a branch when `createBranch==true`:
  - Use `BranchExistsLocal` and `BranchExistsRemote` to detect conflicts
  - If branch exists (local or remote) AND `force==false`, return a clear error with suggested remedies (pick a different name, pass `--force`, or reuse existing branch explicitly)
  - If `force==true`, proceed but log that user requested override

---

## Phase 3: User Story 1 - Create new feature worktree (Priority: P1) 🎯 MVP

**Goal**: Create a new worktree with a new branch and configure the new branch's upstream locally so subsequent push/pull target the intended branch (no automatic push)

**Independent Test**: Create a repo with a remote, run `wt add -b feature/xyz master`, then verify `git rev-parse --abbrev-ref --symbolic-full-name @{u}` in the worktree returns the configured upstream (e.g., `origin/feature/xyz`).

### Implementation tasks

- [ ] T007 [US1] Update `internal/commands/add.go` to, after successful creation of a new branch/worktree, determine the remote to use for upstream:
  - If `remoteOverride` provided, use it
  - Else, if the base branch has a tracked remote, use that remote
  - Else, if remote named `origin` exists, use `origin`
  - Else, do not configure upstream

- [ ] T008 [US1] Call `SetBranchUpstream(repoPath, branchName, remoteName)` (from `internal/services/git_service.go`) when a remote was chosen and `noUpstream==false`; ensure this is a best-effort local config and does not push

- [ ] T009 [US1] Update `contracts/wt-add.md` (path: `specs/001-set-upstream-branch/contracts/wt-add.md`) to reflect the actual flag names and exit codes implemented (sync docs with code)

### Tests (Integration)

- [ ] T010 [US1] Create integration test `tests/integration/test_add_upstream_test.go` that:
  - Initializes a temporary git repo with a remote (bare) and a default branch
  - Runs `wt add -b feature/test <default-branch>` (invoke `cmd/wt` main or call into commands.RunAddCommand)
  - Verifies the created branch's upstream is configured locally to `origin/feature/test`

---

## Phase 4: User Story 2 - Create worktree for an existing remote branch (Priority: P2)

**Goal**: Create a worktree for an existing branch and preserve its existing upstream mapping

**Independent Test**: Given a branch `feature/x` that tracks `origin/feature/x`, run `wt add feature/x` and verify `GetBranchUpstream` returns `origin/feature/x` and that the CLI did not change it.

- [ ] T011 [US2] Ensure code path for creating a worktree from an existing branch (createBranch==false) does NOT alter upstream. Modify `internal/commands/add.go` to explicitly skip upstream configuration when branch already exists and has an upstream.
- [ ] T012 [US2] Add integration test `tests/integration/test_preserve_upstream_test.go` that creates a remote-tracking branch, runs `wt add <branch>`, and asserts upstream unchanged

---

## Phase 5: User Story 3 - Detached-start-point or commit-based worktree (Priority: P3)

**Goal**: When creating a worktree from a commit or tag without an explicit branch name, do not create or modify upstream configuration

**Independent Test**: Create a worktree from a commit hash using `wt add <commit>`, then verify no upstream was configured for the worktree's HEAD.

- [ ] T013 [US3] Verify `internal/commands/add.go` path for `createBranch==false` and `startPoint` being a commit does not invoke any `SetBranchUpstream` logic; add defensive check if necessary
- [ ] T014 [US3] Add integration test `tests/integration/test_detached_no_upstream_test.go` that creates a worktree from a commit and confirms no upstream exists

---

## Final Phase: Polish & Cross-Cutting Concerns

**Purpose**: Documentation, formatting, CI and final validation

- [ ] T015 [P] Update `specs/001-set-upstream-branch/quickstart.md` with explicit examples using new flags (`--force`, `--remote`, `--no-upstream`) and verification steps
- [ ] T016 [P] Run `gofmt` / `go vet` and fix linter issues in modified files: `cmd/wt/main.go`, `internal/commands/add.go`, `internal/services/git_service.go`
- [ ] T017 [P] Add user-facing error messages in `internal/commands/add.go` that match the contract (exit codes 2/3) and ensure messages include remediation steps
- [ ] T018 [P] Commit changes in a single feature commit on branch `001-set-upstream-branch` (include files changed) and push (if desired) — path: repo root (create commit message referencing spec)

---

## Dependencies & Execution Order

- Foundational (Phase 2) tasks T004-T006 must complete before User Story implementation (T007-T014)
- Phase 1 (Setup) tasks T001-T003 can run in parallel where marked [P]
- Phase 3 (US1) is the MVP and should be completed first (T007-T010). User Stories 2 and 3 can proceed after foundational tasks but may be worked in parallel by different developers

### Parallel opportunities

- T001 and T003 are parallelizable with other setup tasks [P]
- Documentation and formatting tasks (T009, T015, T016) are parallelizable [P]
- Integration test creation tasks (T010, T012, T014) can be written in parallel but rely on code changes to pass

## Implementation strategy

1. MVP-first: implement Phase 1 + Phase 2 + Phase 3 (US1) to deliver deterministic upstream configuration without auto-push
2. Validate US1 with integration test and quickstart before moving to US2/US3
3. Implement US2 and US3 and their integration tests
4. Polish docs, run `gofmt`/`go vet`, and finalize commit
