# Quickstart: Using --exec with wt add

## Basic Usage

The `--exec` option allows you to automatically execute commands in a worktree immediately after creation.

### Single Command Execution

```bash
# Create a worktree and run npm install
wt add --exec "npm install" main

# Create a worktree on a new branch and run tests
wt add -b feature/new --exec "make test" main
```

### Multiple Commands

```bash
# Run multiple commands sequentially
wt add --exec "npm install" --exec "npm run build" main

# Mix with other flags
wt add -b hotfix --exec "bundle install" --exec "rake db:migrate" main
```

## Common Patterns

### Dependency Installation

```bash
# Node.js projects
wt add --exec "npm install" main
wt add --exec "yarn install" main

# Ruby projects  
wt add --exec "bundle install" main

# Python projects
wt add --exec "pip install -r requirements.txt" main
```

### Build and Test

```bash
# Run build process
wt add --exec "make build" main

# Run tests
wt add --exec "npm test" main
wt add --exec "go test ./..." main

# Build and test sequence
wt add --exec "make build" --exec "make test" main
```

### Development Setup

```bash
# Setup development environment
wt add -b feature/ui --exec "npm install && npm run dev" main

# Database setup
wt add --exec "bundle exec rails db:create db:migrate" main
```

## Advanced Usage

### Interactive Commands

```bash
# Interactive commands work but may require user input
wt add --exec "rails console" main
# Note: You'll see a warning about interactive commands
```

### Long-running Commands

```bash
# Commands are subject to 5-minute default timeout
wt add --exec "npm run dev" main
# For longer operations, the timeout can be configured (future enhancement)
```

### Error Handling

```bash
# If a command fails, the worktree is still created
wt add --exec "nonexistent-command" main
# Output: ✗ nonexistent-command failed: command not found
# Worktree is still available at the specified path

# Check exit codes
wt add --exec "false" main  # Exit code 2 (command failed)
wt add --exec "true" main   # Exit code 0 (success)
```

## Best Practices

### Command Ordering

Commands execute in the order specified:
```bash
# This runs 'install' before 'build'
wt add --exec "npm install" --exec "npm run build" main
```

### Error Recovery

If a command fails, subsequent commands still execute:
```bash
# Even if install fails, build will still attempt to run
wt add --exec "npm install" --exec "npm run build" main
```

### Worktree Preservation

The worktree is created even if commands fail:
```bash
# Worktree created successfully, command failed
wt add --exec "invalid-command" main
# You can still cd into the worktree and fix issues manually
```

## Troubleshooting

### Command Not Found

```bash
# Error: command not found
wt add --exec "my-custom-command" main

# Solution: Use full path or ensure command is in PATH
wt add --exec "/full/path/to/command" main
```

### Permission Issues

```bash
# Error: permission denied
wt add --exec "/protected/command" main

# Solution: Check file permissions or use sudo (not recommended)
```

### Timeout Issues

```bash
# Error: command timed out
wt add --exec "long-running-process" main

# Solution: Break into smaller commands or optimize the process
wt add --exec "quick-step1" --exec "quick-step2" main
```

## Platform Notes

### Windows

```bash
# Use double quotes for commands with spaces
wt add --exec "dir" main

# For complex commands, consider using PowerShell syntax
wt add --exec "powershell -Command Write-Host 'Hello'" main
```

### Linux/macOS

```bash
# Standard shell commands work as expected
wt add --exec "ls -la" main
wt add --exec "echo $PWD" main
```

## Examples by Use Case

### Frontend Development

```bash
# Create feature branch and start dev server
wt add -b feature/ui --exec "npm install && npm run dev" main

# Run tests in isolated worktree
wt add --exec "npm test" main
```

### Backend Development

```bash
# Ruby on Rails setup
wt add -b feature/api --exec "bundle install && rails db:migrate" main

# Go project setup
wt add --exec "go mod tidy && go build" main
```

### DevOps/CI

```bash
# Setup and test in one command
wt add --exec "make setup && make test" main

# Lint and format check
wt add --exec "npm run lint && npm run format:check" main
```

## Performance Tips

- Simple commands add minimal overhead (< 1 second)
- Complex commands may take longer but won't affect worktree creation
- Multiple commands run sequentially, so order matters for performance
- Use `--exec` for setup tasks to avoid manual repetition