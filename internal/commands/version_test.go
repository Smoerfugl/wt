package commands

import (
	"bytes"
	"os"
	"testing"

	"github.com/smoerfugl/wt/internal/models"
)

func TestRunVersionCommand(t *testing.T) {
	tests := []struct {
		name       string
		jsonOutput bool
		wantPrefix string
	}{
		{
			name:       "Text output",
			jsonOutput: false,
			wantPrefix: "wt version ",
		},
		{
			name:       "JSON output",
			jsonOutput: true,
			wantPrefix: "{",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := RunVersionCommand([]string{}, tt.jsonOutput)

			w.Close()
			os.Stdout = old

			if err != nil {
				t.Fatalf("RunVersionCommand() error = %v, wantErr = nil", err)
			}

			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := buf.String()

			if len(output) < len(tt.wantPrefix) || output[:len(tt.wantPrefix)] != tt.wantPrefix {
				t.Errorf("RunVersionCommand() output = %q, want prefix %q", output, tt.wantPrefix)
			}
		})
	}
}

func TestVersionInfoValidation(t *testing.T) {
	tests := []struct {
		name    string
		version string
		wantErr bool
	}{
		{
			name:    "Valid semantic version",
			version: "1.0.0",
			wantErr: false,
		},
		{
			name:    "Valid semantic version with v prefix",
			version: "v1.0.0",
			wantErr: false,
		},
		{
			name:    "Valid semantic version with prerelease",
			version: "1.0.0-alpha.1",
			wantErr: false,
		},
		{
			name:    "Valid semantic version with build metadata",
			version: "1.0.0+build.123",
			wantErr: false,
		},
		{
			name:    "Invalid semantic version",
			version: "invalid",
			wantErr: true,
		},
		{
			name:    "Empty version",
			version: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := &models.VersionInfo{
				Version: tt.version,
			}
			err := info.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("VersionInfo.Validate() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
