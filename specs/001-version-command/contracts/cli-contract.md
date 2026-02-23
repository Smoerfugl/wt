# CLI Contract: Version Command

**Date**: 2026-02-23
**Feature**: Version Command
**Interface Type**: CLI Command

## Command Specification

### Command Signature

```bash
wt version [--json] [--help]
```

### Arguments

| Argument | Type | Required | Description | Default |
|----------|------|----------|-------------|---------|
| `--json` | flag | optional | Output version information in JSON format | false |
| `--help` | flag | optional | Show help for version command | false |

### Exit Codes

| Code | Meaning | Description |
|------|--------|-------------|
| 0 | Success | Version information displayed successfully |
| 1 | Error | Invalid arguments or formatting error |
| 2 | Help | Help text displayed (when --help used) |

## Output Contracts

### Text Format (Default)

**Structure**:
```
wt version <VERSION> [BUILD_INFO]
```

**Example**:
```
wt version 1.0.0 (built with go1.21.0 on 2026-02-23, git commit abc1234)
```

**Fields**:
- `VERSION`: Semantic version (MAJOR.MINOR.PATCH)
- `BUILD_INFO`: Optional build information in parentheses

**Validation**:
- VERSION must match semantic versioning regex
- Output must be single line
- Must be human-readable

### JSON Format

**Structure**:
```json
{
  "version": "<SEMVER>",
  "buildDate": "<TIMESTAMP>",
  "gitCommit": "<COMMIT_HASH>",
  "goVersion": "<GO_VERSION>",
  "platform": "<PLATFORM>"
}
```

**Example**:
```json
{
  "version": "1.0.0",
  "buildDate": "2026-02-23T14:30:00Z",
  "gitCommit": "abc1234",
  "goVersion": "go1.21.0",
  "platform": "linux/amd64"
}
```

**Validation**:
- Must be valid JSON
- `version` field is required
- Other fields are optional (can be empty strings or null)
- Must be machine-parsable

## Help Contract

### Command Help

**Command**: `wt version --help`

**Output Structure**:
```
Display version information for wt

Usage:
  wt version [flags]

Flags:
  -h, --help   help for version
      --json    Output version information in JSON format
```

**Requirements**:
- Must show command usage
- Must list all available flags
- Must include brief description
- Must match existing CLI help format

### Main Help Integration

**Command**: `wt help`

**Requirements**:
- Version command must appear in command list
- Must have brief description: "Display version information"
- Must be alphabetically ordered with other commands

## Compatibility Contract

### Backward Compatibility
- ✅ No breaking changes (new command)
- ✅ Existing commands unaffected
- ✅ No changes to existing flags or behavior

### Forward Compatibility
- ✅ JSON output format stable
- ✅ Text output format may extend but not break
- ✅ New optional fields may be added to JSON

## Testing Contract

### Test Coverage Requirements

| Test Type | Coverage Requirement | Description |
|-----------|---------------------|-------------|
| Unit | 90%+ | Core formatting logic |
| Integration | 100% | CLI command execution |
| Contract | 100% | Output format validation |

### Test Scenarios

1. **Default execution**: `wt version` → text output
2. **JSON format**: `wt version --json` → JSON output
3. **Help**: `wt version --help` → help text
4. **Main help**: `wt help` → includes version command
5. **Error handling**: Invalid flags → appropriate error

## Performance Contract

| Metric | Requirement | Measurement |
|--------|-------------|-------------|
| Execution time | <100ms | From command invocation to output |
| Memory usage | <5MB | Peak memory during execution |
| Startup time | <50ms | Time to initialize command |

## Security Contract

- ✅ No network access required
- ✅ No file system writes
- ✅ No environment variable dependencies
- ✅ Safe to run in any directory
- ✅ No privilege escalation

## Documentation Contract

### Required Documentation

1. **CLI Help**: Integrated into `wt help` and `wt version --help`
2. **Usage Examples**: In command help text
3. **Output Format**: Documented in this contract
4. **Error Handling**: Documented exit codes and messages

### User-Facing Documentation

- Command must appear in main README
- Usage examples in documentation
- Version information in changelog format