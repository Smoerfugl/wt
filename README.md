# wt — Worktree Manager

A CLI tool for managing Git worktrees. `wt` wraps `git worktree` with a friendlier interface, opinionated worktree placement, interactive selection menus, and multiple output formats.

**Zero external dependencies** — pure Go standard library.

## Installation

### From source

```bash
git clone https://github.com/smoerfugl/wt
cd wt
make install
```

### Build only

```bash
make build
# produces ./wt
```

### As a Go module

```bash
go install github.com/smoerfugl/wt@latest
```

Requires Go 1.20 or later.

## Usage

### List worktrees

```bash
wt list              # basic table output
wt list -v           # verbose: name, path, branch, commit, status, current
wt list -j           # JSON output
wt list -f <name>    # filter by name (case-insensitive substring)
wt list -b <branch>  # filter by exact branch name
```

### Add a worktree

```bash
wt add <branch>          # add worktree for an existing branch
wt add -b <new-branch>   # create a new branch and add its worktree
wt add -b <new-branch> <start-point>  # branch from a specific ref or commit
```

Worktrees are placed at `../worktrees/<repo-name>/<branch>` relative to the repository root. For example, adding a `feature-x` worktree to a repo at `/home/user/Projects/myapp` creates:

```
/home/user/Projects/worktrees/myapp/feature-x
```

When no `<start-point>` is given with `-b`, `wt` resolves the default ref from `origin/HEAD`, `origin`'s default branch, or the current branch.

### Remove a worktree

```bash
wt remove <path>   # remove by path
wt remove          # interactive numbered menu (excludes main worktree)
```

Enter `q` to cancel the interactive prompt.

### Prune stale worktrees

```bash
wt prune
```

Delegates to `git worktree prune`.

### Run a command in a worktree

```bash
wt exec <command> [args...]
```

Interactively select a non-main worktree, then run the given command in that directory with stdin/stdout/stderr forwarded. Enter `q` to cancel.

### Help

```bash
wt help
wt -h
wt --help
```

## Development

```bash
make fmt      # go fmt ./...
make vet      # go vet ./...
make test     # go test ./...
make build    # go build -o wt ./cmd/wt
```

### Project layout

```
cmd/wt/             # CLI entry point
internal/
  commands/         # Command structs (ListCommand, ...)
  models/           # Domain types (Worktree, Repository)
  services/         # GitService — shells out to git
  utils/            # Output formatters (basic, verbose, JSON)
integration/        # Black-box integration tests
specs/              # Per-command feature specifications
docs/               # Style guide and other documentation
```

### Testing

```bash
go test ./...           # all tests (unit + integration)
go test -cover ./...    # with coverage (80% minimum required)
go test -v -run <Name>  # single test
```

The integration tests build the real binary and exercise the full `add → list → remove → prune` lifecycle against a temporary git repository.

## Changelog

See [CHANGELOG.md](CHANGELOG.md).

## License

See repository for license details.
