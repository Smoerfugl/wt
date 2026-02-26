# Quickstart: Creating a worktree with upstream configured

1. Create a new branch and worktree off the default branch (master/main):

```bash
# create and switch to a new worktree with branch `feature/foo`
wt add -b feature/foo master
```

2. Verify upstream is configured locally:

```bash
git -C <path-to-worktree> rev-parse --abbrev-ref --symbolic-full-name @{u}
# Expected: origin/feature/foo  (or the chosen remote)
```

3. Push the new branch explicitly:

```bash
git -C <path-to-worktree> push -u origin feature/foo
```

Notes:
- The CLI configures the branch to track the chosen remote locally but does not perform an automatic push. Use `git push -u` to create the remote branch if desired.
- If the requested branch name exists, the command will fail with an error describing options. Use `--force` only when you intentionally want to override.
