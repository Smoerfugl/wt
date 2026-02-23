package models

import (
	"testing"
)

func TestGetVersionInfo(t *testing.T) {
	info := GetVersionInfo()

	if info == nil {
		t.Fatal("GetVersionInfo() returned nil")
	}

	if info.Version == "" {
		t.Error("Version should not be empty")
	}

	if info.GoVersion == "" {
		t.Error("GoVersion should not be empty")
	}

	if info.Platform == "" {
		t.Error("Platform should not be empty")
	}

	// Test validation
	err := info.Validate()
	if err != nil {
		t.Errorf("VersionInfo.Validate() returned error: %v", err)
	}
}

func TestIsValidSemver(t *testing.T) {
	tests := []struct {
		version string
		want    bool
	}{
		{"1.0.0", true},
		{"v1.0.0", true},
		{"1.0.0-alpha", true},
		{"1.0.0-alpha.1", true},
		{"1.0.0+build.123", true},
		{"1.0.0-alpha.1+build.123", true},
		{"0.0.1-dev", true},
		{"invalid", false},
		{"", false},
		{"1.0", false},
		{"1.0.0.0", false},
	}

	for _, tt := range tests {
		t.Run(tt.version, func(t *testing.T) {
			got := isValidSemver(tt.version)
			if got != tt.want {
				t.Errorf("isValidSemver(%q) = %v, want %v", tt.version, got, tt.want)
			}
		})
	}
}
