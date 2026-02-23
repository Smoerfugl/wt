# Research: Add --exec Option to Worktree Add Command

## Technical Decisions

### Decision: Command Execution Approach
**Rationale**: Use Go's `os/exec` package to execute commands in the worktree directory. This provides cross-platform compatibility and proper process management.
**Alternatives considered**: 
- Shelling out to system shell (rejected due to security concerns and platform inconsistencies)
- Custom command parser (rejected as unnecessary complexity)

### Decision: Multiple Command Handling
**Rationale**: Execute multiple --exec commands sequentially in the order provided. This provides predictable behavior and easier debugging.
**Alternatives considered**:
- Parallel execution (rejected due to potential resource conflicts and output mixing)
- Single command only (rejected as too restrictive)

### Decision: Timeout Implementation
**Rationale**: Use a 5-minute default timeout with context.WithTimeout for command execution. This prevents hanging processes while allowing flexibility.
**Alternatives considered**:
- No timeout (rejected as risky for production use)
- Fixed short timeout (rejected as too restrictive for some use cases)

### Decision: Error Handling Strategy
**Rationale**: Continue worktree creation even if command execution fails, but report command errors clearly. This ensures users don't lose their worktree due to command issues.
**Alternatives considered**:
- Rollback worktree on command failure (rejected as too aggressive)
- Silent failure (rejected as poor user experience)

### Decision: Interactive Command Support
**Rationale**: Allow interactive commands but display warnings. This provides flexibility while acknowledging limitations.
**Alternatives considered**:
- Block interactive commands entirely (rejected as too restrictive)
- Automatic TTY allocation (rejected as complex and platform-dependent)

## Best Practices

### Command Execution Security
- Use `exec.CommandContext` for proper process lifecycle management
- Avoid shell interpretation when possible to prevent injection vulnerabilities
- Validate command arguments before execution
- Set appropriate working directory explicitly

### Cross-Platform Considerations
- Use filepath.Join for path construction
- Handle Windows vs Unix command differences gracefully
- Test on all target platforms (Linux, macOS, Windows)

### Testing Strategy
- Unit tests for command parsing and validation
- Integration tests for actual command execution
- Edge case testing for special characters, long commands, etc.
- Cross-platform test coverage

### Performance Optimization
- Minimize overhead between worktree creation and command execution
- Use efficient string handling for command arguments
- Avoid unnecessary file system operations

## Integration Points

### Existing Codebase Integration
- Extend `commands/add.go` to accept --exec flag
- Modify `services/worktree.go` to execute commands after creation
- Add utility functions in `utils/exec.go` for command execution

### Git Worktree Compatibility
- Ensure commands execute after Git worktree creation completes
- Preserve all existing Git worktree functionality
- Handle Git errors appropriately before command execution

### CLI Interface Design
- Add `--exec` and `-e` flags to add command
- Provide clear help text and examples
- Ensure consistent flag naming with existing commands

## Research Findings Summary

The implementation will use Go's standard library for command execution, extend existing command and service layers, and follow established patterns in the worktree-manager codebase. All technical decisions align with the project's constitution and maintain cross-platform compatibility.