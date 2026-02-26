package commands

import (
	"os/exec"
	"strings"
	"testing"
)

// fakeGS is a test double implementing services.GitServiceIface for unit tests.
type fakeGS struct {
	setUpstreamCalled    bool
	setUpstreamArgs      struct{ branch, remoteBranch string }
	hasRemoteVal         bool
	getBranchUpstreamVal string
}

func (f *fakeGS) EnsureWorktreesDir(repoPath string) error                                { return nil }
func (f *fakeGS) AddWorktree(worktreePath, ref string) error                              { return nil }
func (f *fakeGS) AddWorktreeWithBranch(worktreePath, branchName, startPoint string) error { return nil }
func (f *fakeGS) GetDefaultRef(repoPath string) (string, error)                           { return "origin/main", nil }
func (f *fakeGS) BranchExistsLocal(repoPath, branch string) (bool, error)                 { return false, nil }
func (f *fakeGS) BranchExistsRemote(repoPath, remote, branch string) (bool, error)        { return false, nil }
func (f *fakeGS) HasRemote(repoPath, remote string) (bool, error)                         { return f.hasRemoteVal, nil }
func (f *fakeGS) SetBranchUpstream(repoPath, branch, remoteBranch string) error {
	f.setUpstreamCalled = true
	f.setUpstreamArgs.branch = branch
	f.setUpstreamArgs.remoteBranch = remoteBranch
	return nil
}
func (f *fakeGS) GetBranchUpstream(repoPath, branch string) (string, error) {
	return f.getBranchUpstreamVal, nil
}
func (f *fakeGS) GetGitVersion() (string, error) { return "test-git", nil }

func runGit(t *testing.T, dir string, args ...string) {
	t.Helper()
	cmd := exec.Command("git", args...)
	if dir != "" {
		cmd.Dir = dir
	}
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git %v failed: %v\noutput: %s", args, err, string(out))
	}
}

func gitOut(t *testing.T, dir string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", args...)
	if dir != "" {
		cmd.Dir = dir
	}
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("git %v failed: %v", args, err)
	}
	return strings.TrimSpace(string(out))
}

func TestRunAddCommand_UpstreamBehavior(t *testing.T) {
	// prepare AddCommand with fake service
	f := &fakeGS{hasRemoteVal: true, getBranchUpstreamVal: "origin/main"}
	ac := NewAddCommand(f)
	ac.SetCreateBranch(true)
	ac.SetBranchName("feature-x")
	ac.SetStartPoint("main")
	ac.SetNoUpstream(false)
	ac.SetRemote("")
	ac.SetForce(false)

	// Execute should call SetBranchUpstream using origin/feature-x
	if err := ac.Execute("/tmp/repo"); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
	if !f.setUpstreamCalled {
		t.Fatalf("expected SetBranchUpstream to be called")
	}
	if f.setUpstreamArgs.remoteBranch != "origin/feature-x" {
		t.Fatalf("unexpected remoteBranch: %s", f.setUpstreamArgs.remoteBranch)
	}
}

func TestRunAddCommand_BranchConflictWithoutForce(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not installed in test environment")
	}

	repo := t.TempDir()
	runGit(t, repo, "init")
	runGit(t, repo, "config", "user.email", "test@example.com")
	runGit(t, repo, "config", "user.name", "tester")
	runGit(t, repo, "commit", "--allow-empty", "-m", "initial")

	// create branch 'exists'
	runGit(t, repo, "checkout", "-b", "exists")

	// attempt to create same branch via RunAddCommand without force
	err := RunAddCommand(repo, "git", true, false, "exists", "", nil, false, "", false, "")
	if err == nil {
		t.Fatalf("expected error when creating existing branch without --force")
	}
	if !strings.Contains(err.Error(), "already exists locally") {
		t.Fatalf("unexpected error message: %v", err)
	}
}
