package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/smoerfugl/wt/internal/commands"
	"github.com/smoerfugl/wt/internal/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func runNewListCommand(verbose, jsonOutput bool, filter, branch string) error {
	// Get the repository path
	repoPath, err := gitTop()
	if err != nil {
		return err
	}
	return commands.RunListCommand(repoPath, "git", verbose, jsonOutput, filter, branch)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	cmd := os.Args[1]
	switch cmd {
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		verbose := listCmd.Bool("v", false, "show verbose output")
		jsonOutput := listCmd.Bool("j", false, "output in JSON format")
		filter := listCmd.String("f", "", "filter worktrees by name pattern")
		branch := listCmd.String("b", "", "filter worktrees by branch name")
		if err := listCmd.Parse(os.Args[2:]); err != nil {
			fatal(err)
		}
		if err := ensureRepo(); err != nil {
			fatal(err)
		}
		if err := runNewListCommand(*verbose, *jsonOutput, *filter, *branch); err != nil {
			fatal(err)
		}
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		createBranch := addCmd.Bool("b", false, "create a new branch with -b")
		addCmd.String("exec", "", "command to execute in the worktree after creation")
		addCmd.String("e", "", "command to execute in the worktree after creation (shorthand)")
		addCmd.Parse(os.Args[2:])
		args := addCmd.Args()

		// Parse multiple --exec/-e flags manually
		var execCommands []*utils.Command
		for i := 0; i < len(os.Args[2:]); i++ {
			if os.Args[2+i] == "--exec" || os.Args[2+i] == "-e" {
				if i+1 < len(os.Args[2:]) {
					cmdStr := os.Args[2+i+1]
					cmd := utils.NewCommand("sh", []string{"-c", cmdStr})
					if err := cmd.Validate(); err != nil {
						fatal(fmt.Errorf("invalid --exec command: %w", err))
					}
					execCommands = append(execCommands, cmd)
					i++ // Skip the command argument
				}
			}
		}

		if *createBranch {
			// wt add -b <new-branch> [<start-point>]
			if len(args) < 1 {
				fmt.Fprintln(os.Stderr, "usage: wt add -b <new-branch> [<start-point>]")
				os.Exit(2)
			}
		} else {
			// wt add <branch|commit>
			if len(args) < 1 {
				fmt.Fprintln(os.Stderr, "usage: wt add <branch|commit>")
				os.Exit(2)
			}
		}

		if err := ensureRepo(); err != nil {
			fatal(err)
		}

		// Get repository path
		repoPath, err := gitTop()
		if err != nil {
			fatal(err)
		}

		// Determine branch name and start point
		branchName := args[0]
		startPoint := ""
		if *createBranch && len(args) >= 2 {
			startPoint = args[1]
		}

		// Execute the add command
		if err := commands.RunAddCommand(repoPath, "git", *createBranch, false, branchName, startPoint, execCommands); err != nil {
			fatal(err)
		}
	case "remove":
		removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
		removeCmd.Parse(os.Args[2:])
		args := removeCmd.Args()

		if len(args) < 1 {
			// Show interactive selection if no path provided
			if err := interactiveRemove(); err != nil {
				fatal(err)
			}
			return
		}

		if err := ensureRepo(); err != nil {
			fatal(err)
		}
		if err := runGit("worktree", "remove", args[0]); err != nil {
			fatal(err)
		}
	case "prune":
		if err := ensureRepo(); err != nil {
			fatal(err)
		}
		if err := runGit("worktree", "prune"); err != nil {
			fatal(err)
		}
	case "exec":
		execCmd := flag.NewFlagSet("exec", flag.ExitOnError)
		execCmd.Parse(os.Args[2:])
		args := execCmd.Args()
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "usage: wt exec <command> [<args>...]")
			os.Exit(2)
		}
		if err := ensureRepo(); err != nil {
			fatal(err)
		}
		if err := interactiveExec(args); err != nil {
			fatal(err)
		}
	case "version":
		versionCmd := flag.NewFlagSet("version", flag.ExitOnError)
		jsonOutput := versionCmd.Bool("j", false, "Output version information in JSON format")
		help := versionCmd.Bool("h", false, "Show help for version command")
		if err := versionCmd.Parse(os.Args[2:]); err != nil {
			fatal(err)
		}
		if *help {
			commands.PrintVersionHelp()
			return
		}
		if err := commands.RunVersionCommand(versionCmd.Args(), *jsonOutput); err != nil {
			fatal(err)
		}
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", cmd)
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Fprint(os.Stdout, `wt - simple git worktree helper

Usage:
  wt list [flags]          List worktrees in the repository
    -v, --verbose        Show detailed information including branch, commit, and status
    -j, --json           Output in JSON format for programmatic use
    -f, --filter <name>  Filter worktrees by name pattern
    -b, --branch <name>  Filter worktrees by branch name
  wt add [-b] [--exec <command>] <branch>   Add a worktree (use -b to create a new branch)
                          Worktrees are created in ../worktrees/<branchname>
                          Use --exec to run commands in the new worktree
  wt remove [path]        Remove a worktree (interactive if no path specified)
  wt exec <command>       Execute a command in a selected worktree
  wt prune               Prune stale worktrees
  wt version             Display version information
  wt help                Show this help
`)
}

func fatal(err error) {
	fmt.Fprint(os.Stderr, "error: ", err, "\n")
	os.Exit(1)
}

func runGit(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func ensureRepo() error {
	_, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return errors.New("not a git repository (or any of the parent directories)")
	}
	return nil
}

type worktreeEntry struct {
	Path   string
	Head   string
	Branch string
}

func listWorktrees() error {
	out, err := exec.Command("git", "worktree", "list", "--porcelain").Output()
	if err != nil {
		return err
	}
	entries, err := parsePorcelain(out)
	if err != nil {
		return err
	}
	// print simple table
	fmt.Printf("%-40s %-12s %s\n", "PATH", "HEAD", "BRANCH")
	for _, e := range entries {
		branch := e.Branch
		if branch == "" {
			branch = "(detached)"
		}
		head := e.Head
		if head == "" {
			head = "(none)"
		} else if len(head) > 12 {
			head = head[:12]
		}
		fmt.Printf("%-40s %-12s %s\n", e.Path, head, branch)
	}
	return nil
}

func parsePorcelain(b []byte) ([]worktreeEntry, error) {
	scanner := bufio.NewScanner(bytes.NewReader(b))
	var entries []worktreeEntry
	var cur worktreeEntry
	seen := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "worktree ") {
			if seen {
				entries = append(entries, cur)
				cur = worktreeEntry{}
			}
			seen = true
			cur.Path = strings.TrimPrefix(line, "worktree ")
			// make path absolute-ish
			if !filepath.IsAbs(cur.Path) {
				// git outputs paths relative to repo; make them absolute to be clearer
				if top, err := gitTop(); err == nil {
					cur.Path = filepath.Join(top, cur.Path)
				}
			}
		} else if strings.HasPrefix(line, "HEAD ") {
			cur.Head = strings.TrimPrefix(line, "HEAD ")
		} else if strings.HasPrefix(line, "branch ") {
			cur.Branch = strings.TrimPrefix(line, "branch ")
		}
	}
	if seen {
		entries = append(entries, cur)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}

func gitTop() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func interactiveRemove() error {
	entries, err := getWorktreeEntries()
	if err != nil {
		return err
	}

	// Filter out main working tree
	var removableEntries []worktreeEntry
	for _, e := range entries {
		// Skip main working tree (typically the repo root)
		topLevel, err := gitTop()
		if err != nil {
			continue
		}
		if e.Path == topLevel {
			continue
		}
		removableEntries = append(removableEntries, e)
	}

	if len(removableEntries) == 0 {
		fmt.Println("No removable worktrees found.")
		return nil
	}

	// Display worktrees with numbers for selection
	fmt.Println("Available worktrees to remove:")
	for i, e := range removableEntries {
		branch := e.Branch
		if branch == "" {
			branch = "(detached)"
		}
		fmt.Printf("%d: %s (%s)\n", i+1, e.Path, branch)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter number to remove (or 'q' to quit): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	input = strings.TrimSpace(input)
	if input == "q" || input == "quit" {
		fmt.Println("Cancelled.")
		return nil
	}

	// Parse selection
	var selected int
	_, err = fmt.Sscanf(input, "%d", &selected)
	if err != nil || selected < 1 || selected > len(removableEntries) {
		return fmt.Errorf("invalid selection")
	}

	selectedEntry := removableEntries[selected-1]
	fmt.Printf("Removing worktree: %s\n", selectedEntry.Path)

	return runGit("worktree", "remove", selectedEntry.Path)
}

func interactiveExec(commandArgs []string) error {
	entries, err := getWorktreeEntries()
	if err != nil {
		return err
	}

	// Filter out main working tree for exec (usually you want to exec in worktrees, not main)
	var execEntries []worktreeEntry
	for _, e := range entries {
		// Skip main working tree (typically the repo root)
		topLevel, err := gitTop()
		if err != nil {
			continue
		}
		if e.Path == topLevel {
			continue
		}
		execEntries = append(execEntries, e)
	}

	if len(execEntries) == 0 {
		fmt.Println("No worktrees found to execute command in.")
		return nil
	}

	// Display worktrees with numbers for selection
	fmt.Println("Available worktrees to execute command:")
	for i, e := range execEntries {
		branch := e.Branch
		if branch == "" {
			branch = "(detached)"
		}
		fmt.Printf("%d: %s (%s)\n", i+1, e.Path, branch)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter number to select worktree (or 'q' to quit): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	input = strings.TrimSpace(input)
	if input == "q" || input == "quit" {
		fmt.Println("Cancelled.")
		return nil
	}

	// Parse selection
	var selected int
	_, err = fmt.Sscanf(input, "%d", &selected)
	if err != nil || selected < 1 || selected > len(execEntries) {
		return fmt.Errorf("invalid selection")
	}

	selectedEntry := execEntries[selected-1]
	fmt.Printf("Executing command in worktree: %s\n", selectedEntry.Path)

	// Change to the worktree directory and execute the command
	cmd := exec.Command(commandArgs[0], commandArgs[1:]...)
	cmd.Dir = selectedEntry.Path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func getWorktreeEntries() ([]worktreeEntry, error) {
	out, err := exec.Command("git", "worktree", "list", "--porcelain").Output()
	if err != nil {
		return nil, err
	}
	return parsePorcelain(out)
}

// defaultRef returns the repository's default branch ref (e.g. "origin/main" or "main").
// It first tries to read symbolic-ref "refs/remotes/origin/HEAD" and fallbacks to trying
// to parse `git remote show origin` or finally uses the current branch.
func defaultRef() (string, error) {
	// Try: git symbolic-ref refs/remotes/origin/HEAD
	out, err := exec.Command("git", "symbolic-ref", "refs/remotes/origin/HEAD").Output()
	if err == nil {
		ref := strings.TrimSpace(string(out))
		// refs/remotes/origin/HEAD -> refs/remotes/origin/main => use origin/main
		if strings.HasPrefix(ref, "refs/remotes/") {
			return strings.TrimPrefix(ref, "refs/remotes/"), nil
		}
		return ref, nil
	}

	// As a fallback, try `git remote show origin` and parse "HEAD branch: <name>"
	out2, err2 := exec.Command("git", "remote", "show", "origin").Output()
	if err2 == nil {
		scanner := bufio.NewScanner(bytes.NewReader(out2))
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(line, "HEAD branch:") {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) == 2 {
					branch := strings.TrimSpace(parts[1])
					// prefer origin/<branch>
					return "origin/" + branch, nil
				}
			}
		}
	}

	// Try to get current branch
	out3, err3 := exec.Command("git", "symbolic-ref", "--short", "HEAD").Output()
	if err3 == nil {
		branch := strings.TrimSpace(string(out3))
		return branch, nil
	}

	// Last resort: use 'main' or 'master' if they exist
	for _, cand := range []string{"main", "master"} {
		if err := exec.Command("git", "show-ref", "--verify", "refs/heads/"+cand).Run(); err == nil {
			return cand, nil
		}
	}

	return "", errors.New("could not determine default branch")
}
