package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/smoerfugl/wt/internal/models"
)

// RunVersionCommand executes the version command
func RunVersionCommand(args []string, jsonOutput bool) error {
	// Get version information
	info := models.GetVersionInfo()

	if err := info.Validate(); err != nil {
		return fmt.Errorf("invalid version information: %w", err)
	}

	// Format and output version information
	if jsonOutput {
		return outputJSONVersion(info)
	}

	return outputTextVersion(info)
}

// PrintVersionHelp displays help information for the version command
func PrintVersionHelp() {
	fmt.Fprint(os.Stdout, `Usage:
  wt version [flags]

Flags:
  -h, --help   Show help for version command
  -j, --json    Output version information in JSON format

Examples:
  wt version                    # Display version in text format
  wt version -j                 # Display version in JSON format
`)
}

// outputTextVersion outputs version information in text format
func outputTextVersion(info *models.VersionInfo) error {
	fmt.Printf("wt version %s", info.Version)

	// Add build information if available
	if info.BuildDate != "" || info.GitCommit != "" {
		fmt.Printf(" (built with %s on %s, git commit %s)",
			info.GoVersion, info.BuildDate, info.GitCommit)
	}
	fmt.Println()

	return nil
}

// outputJSONVersion outputs version information in JSON format
func outputJSONVersion(info *models.VersionInfo) error {
	output, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format version info as JSON: %w", err)
	}

	fmt.Println(string(output))
	return nil
}
