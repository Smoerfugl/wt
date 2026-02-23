# Quickstart: Version Command Implementation

**Date**: 2026-02-23
**Feature**: Version Command
**Status**: Ready for Implementation

## Implementation Steps

### 1. Set Up Version Information

**File**: `internal/version/version.go`

```go
package version

import (
    "runtime/debug"
)

type VersionInfo struct {
    Version   string
    BuildDate string
    GitCommit string
    GoVersion string
    Platform  string
}

func GetVersionInfo() *VersionInfo {
    info := &VersionInfo{
        Version:   getVersionFromBuildInfo(),
        GoVersion: getGoVersion(),
        Platform:  getPlatform(),
    }
    
    // Try to get build info from debug
    if buildInfo, ok := debug.ReadBuildInfo(); ok {
        info.BuildDate = getBuildDate(buildInfo)
        info.GitCommit = getGitCommit(buildInfo)
    }
    
    return info
}

func getVersionFromBuildInfo() string {
    if buildInfo, ok := debug.ReadBuildInfo(); ok {
        return buildInfo.Main.Version
    }
    return "dev"
}

// Implement helper functions...
```

### 2. Create Version Command

**File**: `cmd/version.go`

```go
package cmd

import (
    "encoding/json"
    "fmt"
    "github.com/spf13/cobra"
    "worktree-manager/internal/version"
)

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Display version information",
    Long:  "Display version information for wt",
    RunE: func(cmd *cobra.Command, args []string) error {
        jsonFormat, _ := cmd.Flags().GetBool("json")
        return runVersion(jsonFormat)
    },
}

func init() {
    rootCmd.AddCommand(versionCmd)
    versionCmd.Flags().Bool("json", false, "Output version information in JSON format")
}

func runVersion(jsonFormat bool) error {
    info := version.GetVersionInfo()
    
    if jsonFormat {
        output, err := json.MarshalIndent(info, "", "  ")
        if err != nil {
            return fmt.Errorf("failed to format version info: %w", err)
        }
        fmt.Println(string(output))
    } else {
        fmt.Printf("wt version %s", info.Version)
        if info.BuildDate != "" || info.GitCommit != "" {
            fmt.Printf(" (built with %s on %s, git commit %s)", 
                info.GoVersion, info.BuildDate, info.GitCommit)
        }
        fmt.Println()
    }
    
    return nil
}
```

### 3. Update Main Command

**File**: `cmd/root.go`

```go
// Add to init() function:
rootCmd.AddCommand(versionCmd)
```

### 4. Add Build Information (Optional)

**Build Script**: Update your build command to include version info

```bash
go build -o wt -ldflags "\
-X 'worktree-manager/internal/version.Version=$(git describe --tags)' \
-X 'worktree-manager/internal/version.BuildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ)' \
-X 'worktree-manager/internal/version.GitCommit=$(git rev-parse HEAD)' \
"
```

## Testing Implementation

### Unit Tests

**File**: `internal/version/version_test.go`

```go
package version

import (
    "testing"
)

func TestGetVersionInfo(t *testing.T) {
    info := GetVersionInfo()
    
    if info.Version == "" {
        t.Error("Version should not be empty")
    }
    
    // Test semantic version format
    if !isValidSemver(info.Version) {
        t.Errorf("Version %s is not valid semver", info.Version)
    }
}

func isValidSemver(version string) bool {
    // Implement semver regex validation
    // ...
}
```

### Integration Tests

**File**: `cmd/version_test.go`

```go
package cmd

import (
    "testing"
    "bytes"
    "os"
)

func TestVersionCommand(t *testing.T) {
    // Test text output
    cmd := NewRootCmd()
    cmd.SetArgs([]string{"version"})
    
    old := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w
    
    err := cmd.Execute()
    if err != nil {
        t.Fatalf("Failed to execute version command: %v", err)
    }
    
    w.Close()
    os.Stdout = old
    
    var buf bytes.Buffer
    buf.ReadFrom(r)
    output := buf.String()
    
    if !containsVersion(output) {
        t.Error("Output should contain version information")
    }
}

func containsVersion(output string) bool {
    // Check for version pattern
    // ...
}
```

## Build and Test

```bash
# Build the application
go build -o wt .

# Test version command
./wt version
./wt version --json
./wt version --help

# Run tests
go test ./...

# Check help integration
./wt help
```

## Implementation Notes

1. **Error Handling**: Ensure all errors are properly handled and user-friendly messages are displayed
2. **Performance**: Keep execution time under 100ms
3. **Cross-platform**: Test on Linux, macOS, and Windows
4. **Documentation**: Update help text and examples
5. **Constitution Compliance**: Verify all principles are followed

## Troubleshooting

**Common Issues**:

1. **Version not found**: Ensure build info is properly embedded or use ldflags
2. **JSON formatting errors**: Validate the VersionInfo struct fields
3. **Help not showing**: Verify command is properly added to root command
4. **Slow execution**: Profile and optimize version info retrieval

**Debugging**:
```bash
# Verbose output (if available)
./wt version -v

# Check build info
go version -m ./wt
```