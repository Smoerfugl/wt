# Implementation Plan: Add --exec Option to Worktree Add Command

**Branch**: `002-exec` | **Date**: 2026-02-23 | **Spec**: [specs/002-exec/spec.md](specs/002-exec/spec.md)
**Input**: Feature specification from `/specs/002-exec/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command. See `.specify/templates/plan-template.md` for the execution workflow.

## Summary

This feature adds a `--exec` option to the `wt add` command that allows users to execute commands automatically in newly created worktrees. The implementation will extend the existing worktree creation logic to accept an optional command parameter and execute it in the worktree's directory context after successful creation.

## Technical Context

<!--
  ACTION REQUIRED: Replace the content in this section with the technical details
  for the project. The structure here is presented in advisory capacity to guide
  the iteration process.
-->

**Language/Version**: Go 1.21+ (current project language)  
**Primary Dependencies**: Standard Go library only (no external dependencies needed)  
**Storage**: N/A (uses existing Git worktree storage)  
**Testing**: Go testing framework (`go test`)  
**Target Platform**: Linux, macOS, Windows (cross-platform CLI)  
**Project Type**: CLI tool for Git worktree management  
**Performance Goals**: Command execution adds <2 seconds overhead to worktree creation  
**Constraints**: Must maintain compatibility with native Git worktree functionality  
**Scale/Scope**: Single new command flag with execution logic

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

**Core Principles Compliance:**
- [x] CLI-First Design: All functionality accessible via command-line
- [x] Git Integration Excellence: Uses native Git worktree functionality
- [x] Safety and Predictability: Safe operations with clear confirmation
- [x] Cross-Platform Compatibility: Works consistently across platforms
- [x] Observability and Debugging: Clear output and error messages
- [x] Simplicity and Minimalism: Focused feature set, intuitive interface

**Development Standards Compliance:**
- [x] Code Quality: Follows Go formatting and vet standards
- [x] Testing Requirements: Unit and integration test coverage
- [x] Documentation Standards: Comprehensive Go docs and CLI help

## Project Structure

### Documentation (this feature)

```text
specs/[###-feature]/
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output (/speckit.plan command)
├── data-model.md        # Phase 1 output (/speckit.plan command)
├── quickstart.md        # Phase 1 output (/speckit.plan command)
├── contracts/           # Phase 1 output (/speckit.plan command)
└── tasks.md             # Phase 2 output (/speckit.tasks command - NOT created by /speckit.plan)
```

### Source Code (repository root)
<!--
  ACTION REQUIRED: Replace the placeholder tree below with the concrete layout
  for this feature. Delete unused options and expand the chosen structure with
  real paths (e.g., apps/admin, packages/something). The delivered plan must
  not include Option labels.
-->

```text
src/
├── commands/
│   └── add.go          # Extended with --exec flag handling
├── services/
│   └── worktree.go     # Worktree creation logic
├── models/
│   └── worktree.go     # Worktree data model
└── utils/
    └── exec.go         # Command execution utilities

tests/
├── unit/
│   ├── commands/
│   │   └── add_test.go # Unit tests for add command
│   └── services/
│       └── worktree_test.go # Unit tests for worktree service
└── integration/
    └── add_exec_test.go # Integration tests for --exec functionality
```

**Structure Decision**: Single project structure following existing worktree-manager conventions. The feature extends existing command and service layers without requiring new major components.

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| [e.g., 4th project] | [current need] | [why 3 projects insufficient] |
| [e.g., Repository pattern] | [specific problem] | [why direct DB access insufficient] |
