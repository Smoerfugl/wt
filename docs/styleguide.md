# Worktree Manager Style Guide

## Project Structure

The Worktree Manager project follows standard Go project structure conventions:

```
worktree-manager/
├── cmd/                  # Main applications for this project
│   └── wt/               # Worktree Manager CLI application
├── internal/             # Private application code (not importable)
│   ├── commands/         # CLI command implementations
│   ├── models/           # Data models and business logic
│   ├── services/         # Core services and business logic
│   └── utils/            # Utility functions and helpers
├── pkg/                  # Public library code (if needed)
├── docs/                 # Documentation
├── specs/                # Feature specifications
├── tests/                # Test files
├── integration/          # Integration tests
├── go.mod                # Go module definition
├── go.sum                # Go dependency checksums
└── README.md             # Project documentation
```

## Package Organization

### `cmd/` Directory
- Contains the main application entry points
- Each subdirectory represents a separate executable
- For Worktree Manager, this contains the `wt` CLI tool

### `internal/` Directory
- Contains private application code that should not be imported by other projects
- Organized by functional area:
  - `commands/`: CLI command implementations
  - `models/`: Data structures and domain models
  - `services/`: Core business logic and services
  - `utils/`: Utility functions and helpers

### `pkg/` Directory (if needed)
- Contains library code that can be imported by other projects
- Should be used sparingly - prefer keeping code in `internal/` unless explicitly designed for external use

## Code Style Guidelines

### Go Standards
- Follow official Go formatting guidelines (`go fmt`)
- Use `go vet` for static analysis
- Follow the single responsibility principle for functions
- Use explicit error handling
- Write clear, descriptive variable and function names

### Testing Requirements
- **Unit Tests**: Cover core functionality and edge cases
- **Integration Tests**: Verify Git worktree operations
- **Test Coverage**: Minimum 80% coverage required
- **Cross-Platform**: Tests must pass on Linux, macOS, Windows

### Documentation Standards
- **Go Doc Comments**: Required for all public functions/types
- **CLI Help**: Comprehensive usage examples and error messages
- **Changelog**: Document all user-facing changes

## Naming Conventions

### Files
- Use lowercase with underscores for file names: `git_service.go`
- Test files should end with `_test.go`: `git_service_test.go`

### Functions
- Use camelCase for function names: `getWorktreePath`
- Public functions should start with capital letter: `GetWorktreePath`
- Private functions should start with lowercase letter: `validateBranchName`

### Variables
- Use camelCase for variable names: `worktreePath`
- Constants should be in ALL_CAPS: `MAX_WORKTREES`
- Public variables should start with capital letter: `DefaultBranch`
- Private variables should start with lowercase letter: `currentBranch`

### Types
- Use PascalCase for type names: `WorktreeManager`
- Interface names should end with "er": `WorktreeManager`

## Error Handling

- Always handle errors explicitly
- Use descriptive error messages
- Consider using custom error types for domain-specific errors
- Follow the pattern of returning errors as the last return value

## Logging

- Use structured logging where appropriate
- Log levels should be appropriate for the message importance
- Avoid logging sensitive information
- Use consistent log message formats

## Testing

- Write tests before implementation (TDD approach)
- Test both happy paths and error cases
- Use table-driven tests for similar test cases
- Keep tests focused and fast
- Use test helpers to reduce boilerplate

## Code Organization

- Keep functions small and focused
- Group related functions together
- Use comments sparingly - code should be self-documenting
- Organize imports in groups (standard library, third-party, local)

## Git Workflow

- Use feature branches for new development
- Write meaningful commit messages
- Follow semantic versioning for releases
- Keep main branch stable and deployable

## Constitution Compliance

All changes must comply with the Project Constitution:
- **CLI-First Design**: All features accessible via command-line
- **Git Integration Excellence**: Use native Git worktree functionality
- **Safety and Predictability**: Safe operations with data integrity
- **Cross-Platform Compatibility**: Consistent behavior across OS
- **Observability and Debugging**: Clear output and error messages
- **Simplicity and Minimalism**: Focused feature set
