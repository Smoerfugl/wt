<!--
SYNC IMPACT REPORT
Version change: 0.0.0 → 1.0.0 (Initial ratification)
- Added 6 core principles (CLI-First Design, Git Integration Excellence, Safety and Predictability, Cross-Platform Compatibility, Observability and Debugging, Simplicity and Minimalism)
- Added Development Standards section with code quality, testing, and documentation requirements
- Added Governance section with amendment process, versioning policy, and compliance review
- No sections removed

Templates requiring updates:
✅ .specify/templates/plan-template.md - Updated Constitution Check section with specific principle compliance checklist
⚠ .specify/templates/spec-template.md - No changes needed (already aligned with testing and documentation principles)
✅ .specify/templates/tasks-template.md - No changes needed (already follows CLI-first and testing principles)

Follow-up TODOs: None
-->

# Worktree Manager Constitution

## Core Principles

### I. CLI-First Design
Every feature must be accessible via the command-line interface. The CLI must follow Unix philosophy: do one thing well, support text input/output, and be composable with other tools. Commands must accept arguments via flags and provide both human-readable and machine-parsable output formats.

### II. Git Integration Excellence
The tool must work seamlessly with Git's native worktree functionality. It should never bypass or override Git's core mechanisms but rather provide a user-friendly wrapper that maintains compatibility with standard Git operations and workflows.

### III. Safety and Predictability
All operations must be safe by default. Destructive operations (like worktree removal) must require explicit confirmation or provide clear undo mechanisms. The tool must maintain data integrity and never corrupt Git repositories.

### IV. Cross-Platform Compatibility
The CLI must work consistently across different operating systems (Linux, macOS, Windows) and Git versions. Platform-specific behaviors must be clearly documented and minimized.

### V. Observability and Debugging
All operations must provide clear, actionable output. Error messages must be specific and include remediation guidance. The tool should support verbose/dry-run modes for debugging and learning purposes.

### VI. Simplicity and Minimalism
Follow the principle of least surprise. The tool should have a small, focused feature set that does worktree management exceptionally well. Avoid feature creep and maintain a clean, intuitive interface.

## Development Standards

### Code Quality
- All Go code must follow the official Go formatting guidelines (`go fmt`)
- Code must pass `go vet` static analysis without warnings
- Functions should be small, focused, and follow the single responsibility principle
- Error handling must be explicit and comprehensive

### Testing Requirements
- Unit tests must cover core functionality and edge cases
- Integration tests must verify Git worktree operations work correctly
- Tests must run successfully on the target platforms (Linux, macOS, Windows)
- Test coverage should be maintained at a minimum of 80%

### Documentation Standards
- All public functions and types must have Go doc comments
- CLI commands must have comprehensive help text and usage examples
- Changes to user-facing behavior must be documented in a changelog

## Governance

The Constitution supersedes all other development practices and guidelines. All contributions must comply with these principles.

### Amendment Process
1. Proposals for constitutional changes must be documented in a GitHub issue
2. Changes require approval from at least two maintainers
3. Major principle changes require a 7-day comment period
4. All amendments must include a clear rationale and impact analysis

### Versioning Policy
- **MAJOR**: Backward-incompatible changes to core principles or CLI interface
- **MINOR**: New principles added or significant expansions to existing ones
- **PATCH**: Clarifications, wording improvements, and non-substantive changes

### Compliance Review
- All pull requests must include a Constitution Check section verifying compliance
- Maintainers must review changes for adherence to principles before merging
- Complexity that violates principles must be explicitly justified

**Version**: 1.0.0 | **Ratified**: 2026-02-23 | **Last Amended**: 2026-02-23
