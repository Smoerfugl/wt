package models

import (
	"fmt"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

// VersionInfo represents version information for the application
type VersionInfo struct {
	Version   string
	BuildDate string
	GitCommit string
	GoVersion string
	Platform  string
}

// NewVersionInfo creates a new VersionInfo instance with default values
func NewVersionInfo() *VersionInfo {
	return &VersionInfo{
		Version:   "0.0.1-dev", // Default fallback version
		GoVersion: getGoVersion(),
		Platform:  getPlatform(),
	}
}

// GetVersionInfo retrieves version information from build info or uses defaults
func GetVersionInfo() *VersionInfo {
	info := NewVersionInfo()

	// Try to get version from build info
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		if buildInfo.Main.Version != "" && buildInfo.Main.Version != "(devel)" {
			info.Version = buildInfo.Main.Version
		}

		// Extract build settings
		for _, setting := range buildInfo.Settings {
			switch setting.Key {
			case "vcs.revision":
				info.GitCommit = setting.Value
			case "vcs.time":
				if t, err := time.Parse(time.RFC3339, setting.Value); err == nil {
					info.BuildDate = t.Format("2006-01-02T15:04:05Z")
				}
			}
		}
	}

	return info
}

// Validate checks if the version information is valid
func (v *VersionInfo) Validate() error {
	if v.Version == "" {
		return fmt.Errorf("version cannot be empty")
	}

	if !isValidSemver(v.Version) {
		return fmt.Errorf("version %s is not valid semantic versioning", v.Version)
	}

	return nil
}

// isValidSemver checks if a version string follows semantic versioning
func isValidSemver(version string) bool {
	// Semantic versioning regex pattern
	pattern := `^v?(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`

	regex := regexp.MustCompile(pattern)
	return regex.MatchString(version)
}

// getGoVersion returns the current Go version
func getGoVersion() string {
	return "go" + strings.TrimPrefix(runtime.Version(), "go")
}

// getPlatform returns the current platform information
func getPlatform() string {
	return fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
}
