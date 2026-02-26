# Specification Quality Checklist: Unit Tests for Requirements (Requirements Quality)

**Purpose**: Validate the specification's clarity, completeness, consistency, and measurability for the feature "Set upstream branch on worktree checkout" (intended for PR reviewers).
**Created**: 2026-02-26
**Feature**: ../spec.md

## Requirement Completeness

- [ ] CHK001 - Are all primary functional requirements present (branch creation, upstream configuration, preservation rules)? [Completeness, Spec §FR-001/FR-002/FR-003]
- [ ] CHK002 - Are acceptance criteria for each user story explicitly listed and tied to requirements (US1-US3)? [Completeness, Spec §User Scenarios]
- [ ] CHK003 - Are failure and error flows documented for each primary operation (e.g., remote unreachable, name conflict)? [Completeness, Spec §FR-005] [Gap]

## Requirement Clarity

- [ ] CHK004 - Is the term "configure the branch's upstream" precisely defined (what exact git ref is written, local vs remote, and when)? [Clarity, Spec §FR-001]
- [ ] CHK005 - Is the chosen remote-selection heuristic described unambiguously (order: base branch tracked remote → `origin` → unset)? [Clarity, Spec §FR-004]
- [ ] CHK006 - Are the meanings of flags like `--force`, `--no-upstream`, and `--remote` specified in requirements (semantics and constraints)? [Clarity, Spec §contracts/wt-add.md] [Gap]

## Requirement Consistency

- [ ] CHK007 - Do requirements about preserving existing upstreams (FR-002) and auto-configuring upstreams (FR-001) have no conflict when a branch exists remotely? [Consistency, Spec §FR-001/FR-002]
- [ ] CHK008 - Are messages and behaviors described for name-conflict handling consistent across User Stories, Contracts, and Error handling sections? [Consistency, Spec §Edge Cases; contracts/wt-add.md]

## Acceptance Criteria Quality

- [ ] CHK009 - Are success criteria measurable and testable without implementation detail (e.g., how to verify upstream is set) and do they reference specific checks? [Measurability, Spec §SC-001]
- [ ] CHK010 - Do acceptance criteria include both positive and negative outcomes (e.g., remote present vs remote absent) for each story? [Completeness, Spec §SC-001/SC-003]

## Scenario Coverage

- [ ] CHK011 - Are primary, alternate, exception, and recovery scenarios enumerated for worktree creation (new-branch, existing-branch, detached commit)? [Coverage, Spec §User Scenarios]
- [ ] CHK012 - Are cross-repository and multi-remote scenarios covered (multiple remotes, remote selection ambiguity)? [Coverage, Spec §Edge Cases] [Gap]
- [ ] CHK013 - Are scripting/CI use-cases explicitly stated (non-interactive behavior and non-prompting flows)? [Coverage, Spec §Notes] [Gap]

## Edge Case Coverage

- [ ] CHK014 - Are branch-name collision behaviors explicitly required (fail, prompt, auto-rename, or force) and recorded as a single canonical requirement? [Edge Case, Spec §FR-006]
- [ ] CHK015 - Are behavior and requirements defined for repositories with no remotes or for unreachable remotes? [Edge Case, Spec §User Scenarios, FR-005]
- [ ] CHK016 - Are requirements for unusual default-refs (repos without `origin` or with non-standard default branch names) documented? [Edge Case, Spec §Assumptions]

## Non-Functional Requirements (NFRs)

- [ ] CHK017 - Are accessibility requirements required for any CLI output that is consumed by assistive tooling or scripts (machine-parsable JSON vs human text)? [Completeness, NFR]
- [ ] CHK018 - Are observability and debugging requirements specified (verbosity levels, error codes, machine-parsable output) with measurable expectations? [Clarity, Spec §Constitution - Observability]
- [ ] CHK019 - Are security considerations documented for operations that might write remote branches (e.g., permission assumptions, safe-by-default policy)? [Coverage, Spec §Notes]
- [ ] CHK020 - Are performance expectations or acceptable latencies defined where they matter for developer UX (e.g., typical worktree creation time target)? [Clarity, Spec §Plan - Performance Goals] [Gap]

## Dependencies & Assumptions

- [ ] CHK021 - Are external dependencies (presence of `git` binary, remote reachability) and their required versions/assumptions documented? [Completeness, Spec §Assumptions]
- [ ] CHK022 - Are assumptions about default remote names and branch naming conventions recorded and justified (e.g., prefer `origin`)? [Clarity, Spec §Assumptions]

## Ambiguities & Conflicts

- [ ] CHK023 - Are any ambiguous terms ("configure upstream", "default ref", "sensible default") identified and replaced with measurable definitions or [Gap] markers? [Ambiguity, Spec §FR-001] [Gap]
- [ ] CHK024 - Do any requirements conflict across sections (e.g., plan says no auto-push while contracts imply sync) and are conflicts annotated with resolution guidance? [Conflict, Spec §contracts/wt-add.md]

## Traceability

- [ ] CHK025 - Does the spec include an ID scheme or explicit section references for ≥80% of requirements and acceptance criteria, or are missing IDs marked with [Gap]? [Traceability, Spec §All] [Gap]
- [ ] CHK026 - Are mappings from each acceptance criterion back to a specific requirement (FR-###) present or requested as remediation? [Traceability, Spec §SC-001/SC-004]

## Summary Checks

- [ ] CHK027 - Is there a short remediation plan in the spec for gaps identified by this checklist (who fixes, by when)? [Completeness, Spec §Notes] [Gap]
- [ ] CHK028 - Has the spec been annotated to indicate which items are intentionally out-of-scope (explicit exclusions)? [Clarity, Spec §Notes]
