# CLI Contract: wt add --exec

## Command Interface

### Command Signature
```
wt add [--exec COMMAND] [--exec COMMAND ...] [-b BRANCH] [START-POINT]
```

### Flags

#### `--exec`, `-e` (string)
- **Description**: Command to execute in the worktree after creation
- **Type**: String (command with arguments)
- **Required**: No
- **Multiple**: Yes (can be specified multiple times)
- **Default**: None
- **Example**: `--exec "npm install"`, `-e "make test"`

#### `-b`, `--branch` (string)
- **Description**: Create a new branch for the worktree
- **Type**: String (branch name)
- **Required**: No (unless no START-POINT provided)
- **Multiple**: No
- **Default**: None

### Arguments

#### `START-POINT` (string)
- **Description**: Branch, commit, or tag to create worktree from
- **Type**: String (Git reference)
- **Required**: No (if -b is provided)
- **Multiple**: No
- **Default**: None

### Behavior Contract

#### Success Conditions
1. Worktree is successfully created at the specified location
2. All provided commands execute in the worktree directory
3. Commands execute sequentially in the order provided
4. Each command has access to the worktree's environment
5. Command output is displayed to the user

#### Error Conditions
1. **WorktreeCreationFailed**: Git worktree creation fails
   - Exit code: 1
   - Error message: "Failed to create worktree: {reason}"
   - Behavior: No commands are executed

2. **CommandExecutionFailed**: One or more commands fail to execute
   - Exit code: 2 (if worktree created successfully)
   - Error message: "Command failed: {command} - {reason}"
   - Behavior: Continue with remaining commands, preserve worktree

3. **InvalidCommand**: Command is not executable or invalid
   - Exit code: 3
   - Error message: "Invalid command: {command} - {reason}"
   - Behavior: Skip invalid command, continue with others

4. **TimeoutExceeded**: Command execution exceeds timeout
   - Exit code: 4
   - Error message: "Command timed out: {command} (after {duration})"
   - Behavior: Terminate command, continue with remaining commands

5. **NoCommandProvided**: --exec flag provided but no command specified
   - Exit code: 5
   - Error message: "No command provided for --exec flag"
   - Behavior: Worktree still created, no commands executed

### Output Contract

#### Standard Output
- Worktree creation status: "Created worktree at {path}"
- Command execution: "Executing: {command}"
- Command success: "✓ {command} completed successfully"
- Command failure: "✗ {command} failed: {reason}"

#### Exit Codes
- `0`: Success (worktree created, all commands executed successfully)
- `1`: Worktree creation failed
- `2`: One or more commands failed (worktree created)
- `3`: Invalid command specified
- `4`: Command timeout exceeded
- `5`: No command provided with --exec flag

### Examples

#### Basic Usage
```bash
# Create worktree and run a single command
wt add --exec "npm install" main

# Create worktree on new branch and run command
wt add -b feature/new --exec "make setup" main

# Create worktree and run multiple commands
wt add --exec "npm install" --exec "npm run build" main
```

#### Error Handling
```bash
# Invalid command
wt add --exec "nonexistent-command" main
# Output: ✗ nonexistent-command failed: command not found
# Exit code: 3

# Command timeout
wt add --exec "sleep 600" main  # With 5 minute default timeout
# Output: ✗ sleep 600 timed out: after 5m0s
# Exit code: 4
```

### Environment Contract

#### Working Directory
- All commands execute in the worktree directory
- Working directory is set explicitly before command execution

#### Environment Variables
- Inherits parent process environment
- Adds worktree-specific variables:
  - `WT_WORKTREE`: Path to worktree directory
  - `WT_BRANCH`: Branch name
  - `WT_REPO`: Path to main repository

#### Platform Compatibility
- Works consistently on Linux, macOS, Windows
- Handles platform-specific command differences
- Uses appropriate path separators for each platform

### Backward Compatibility

#### Existing Behavior Preserved
- `wt add` without `--exec` works exactly as before
- All existing flags and arguments remain unchanged
- Worktree creation logic is unchanged

#### New Functionality
- `--exec` flag is purely additive
- No impact on existing workflows
- Optional feature that doesn't affect core functionality

### Testing Contract

#### Test Coverage Requirements
- Unit tests for command parsing and validation
- Integration tests for actual command execution
- Cross-platform compatibility tests
- Error condition tests for all exit codes
- Performance tests for command execution overhead

#### Test Success Criteria
- All commands execute in correct worktree directory
- Exit codes match specified behavior
- Error messages are clear and actionable
- Performance overhead < 2 seconds for typical commands
- Works consistently across all supported platforms