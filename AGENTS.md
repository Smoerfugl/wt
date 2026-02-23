# AGENTS.md

This file provides guidance for AI coding agents working on the Worktree Manager project.

## Project Overview

**Worktree Manager** is a Go-based CLI tool for managing Git worktrees. It provides a user-friendly interface for creating, listing, removing, and executing commands in Git worktrees while maintaining compatibility with native Git functionality.

## Development Philosophy

### Test-Driven Development (TDD)

We follow strict TDD principles:
1. **Red**: Write failing tests first
2. **Green**: Implement minimal code to pass tests
3. **Refactor**: Improve code while keeping tests passing

### Spec-Driven Development

All features must be specified before implementation:
1. Create feature specification in `.specify/` directory
2. Define user stories and acceptance criteria
3. Get specification approval before coding

## Setup Commands

```bash
# Install Go dependencies
go mod tidy

# Build the CLI
go build -o wt ./cmd/wt

# Run tests
go test ./...

# Format code
go fmt ./...

# Run static analysis
go vet ./...
```

## Code Style

### Go Standards
- Follow official Go formatting (`go fmt`)
- Use `go vet` for static analysis
- Single responsibility principle for functions
- Explicit error handling

### Testing Requirements
- **Unit Tests**: Cover core functionality and edge cases
- **Integration Tests**: Verify Git worktree operations
- **Test Coverage**: Minimum 80% coverage required
- **Cross-Platform**: Tests must pass on Linux, macOS, Windows

### Documentation Standards
- **Go Doc Comments**: Required for all public functions/types
- **CLI Help**: Comprehensive usage examples and error messages
- **Changelog**: Document all user-facing changes

## Testing Instructions

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestFunctionName

# Test in verbose mode
go test -v ./...
```

## Constitution Compliance

All changes must comply with the [Project Constitution](.specify/memory/constitution.md):
- **CLI-First Design**: All features accessible via command-line
- **Git Integration Excellence**: Use native Git worktree functionality
- **Safety and Predictability**: Safe operations with data integrity
- **Cross-Platform Compatibility**: Consistent behavior across OS
- **Observability and Debugging**: Clear output and error messages
- **Simplicity and Minimalism**: Focused feature set

## Workflow

### Feature Development
1. **Specify**: Create specification in `.specify/` directory
2. **Plan**: Create implementation plan
3. **Test**: Write failing tests first
4. **Implement**: Write minimal code to pass tests
5. **Refactor**: Improve code quality
6. **Document**: Update documentation and help text

### Pull Request Requirements
- Include Constitution Check section
- Verify all tests pass
- Maintain minimum 80% test coverage
- Follow Go formatting standards
- Include comprehensive documentation

## Agent-Specific Configuration

### OpenCode AI
Agents are configured in `.opencode/command/` directory with Markdown files defining behavior.

### Worktree Management Commands
```bash
# List all worktrees
wt list

# Add a new worktree
wt add <branch|commit>

# Add worktree with new branch
wt add -b <new-branch> [<start-point>]

# Remove a worktree
wt remove [path]

# Prune stale worktrees
wt prune

# Execute command in worktree
wt exec <command> [<args>...]
```

### Constitution Compliance Commands
```bash
# Check constitution compliance
./speckit.constitution

# Update constitution
./speckit.constitution update

# Validate against constitution
./speckit.analyze
```

## Security Considerations

- Never commit secrets or sensitive data
- Validate all user input
- Use secure coding practices
- Follow Go security guidelines

## Deployment

```bash
# Build release binary
go build -ldflags="-s -w" -o wt ./cmd/wt

# Install to GOPATH
go install ./cmd/wt
```

## Maintenance

- Update AGENTS.md when workflows change
- Keep constitution aligned with project needs
- Review and update agent configurations regularly
- Maintain comprehensive test coverage
