Add new semantic version tag (patch)

Purpose
-------
Small, reproducible instructions and scripts to find the latest semantic `vMAJOR.MINOR.PATCH` tag
in a repository, bump the patch number, create an annotated tag, and push it to the remote.

Quick checklist
---------------
- Ensure local tags are up-to-date: `git fetch --tags`
- Find the latest semver-style tag (prefixed with `v`) and fall back to `v0.0.0` if none exist
- Bump the patch (Z in `vX.Y.Z`) to produce `vX.Y.(Z+1)`
- Create an annotated tag and push it: `git tag -a ...` then `git push origin <tag>`

Commands you can run interactively
---------------------------------

1) Fetch tags from remote (always do this first):

```bash
git fetch --tags
```

2) Get the latest tag (recommended):

```bash
# prefer a semver-sorted git listing (newest first)
git tag --list 'v[0-9]*' --sort=-v:refname | head -n1

# or, when annotated tags are used and reachable, git describe will return the nearest tag
git describe --tags --abbrev=0 2>/dev/null || git tag --list 'v[0-9]*' --sort=-v:refname | head -n1
```

Script: bump the patch and push the new tag
-----------------------------------------

Save the script below as `scripts/bump-patch.sh` (or run it inline). It:
- fetches tags
- finds the latest `vX.Y.Z` tag (falls back to `v0.0.0`)
- increments the patch number
- creates an annotated tag and pushes it

```bash
#!/usr/bin/env bash
set -euo pipefail

# Ensure we have up-to-date tags
git fetch --tags

# Try to get the latest semver tag. If none found, default to v0.0.0
latest=$(git describe --tags --abbrev=0 2>/dev/null || true)
if [ -z "$latest" ]; then
  latest=$(git tag --list 'v[0-9]*' --sort=-v:refname | head -n1 || true)
fi
if [ -z "$latest" ]; then
  latest="v0.0.0"
fi

# Strip leading 'v' and split into components
ver=${latest#v}
IFS='.' read -r major minor patch <<< "$ver"
major=${major:-0}
minor=${minor:-0}
patch=${patch:-0}

# Increment patch
patch=$((patch + 1))
new_tag="v${major}.${minor}.${patch}"

echo "Latest tag: $latest -> New tag: $new_tag"

# Create an annotated tag (preferred) and push it
git tag -a "$new_tag" -m "Bump version to $new_tag"
git push origin "$new_tag"

echo "Created and pushed $new_tag"
```

One-liner (for interactive use)
--------------------------------

```bash
git fetch --tags && latest=$(git tag --list 'v[0-9]*' --sort=-v:refname | head -n1) && [ -z "$latest" ] && latest=v0.0.0 && v=${latest#v} && IFS=. read -r MAJ MIN PAT <<< "$v" && NEW=v$MAJ.$MIN.$((PAT+1)) && git tag -a "$NEW" -m "Bump version to $NEW" && git push origin "$NEW"
```

Notes & edge-cases
-------------------
- This assumes tags follow `vMAJOR.MINOR.PATCH`. Tags that don't match (eg `release-1.2.3` or `1.2.3`) will be ignored by the default filter.
- If a tag only has two components (e.g. `v1.2`) the script treats missing parts as `0` (so `v1.2` -> `v1.2.1`).
- Use `git tag -s` to GPG-sign tags (`-s` instead of `-a`) if your project requires signed tags.
- Annotated tags (`-a`) are recommended because they carry a message and metadata; lightweight tags are created without `-a`.
- If you maintain a changelog, update it before tagging or include the changelog entry in the tag message.

Automation / CI tips
--------------------
- Only create and push tags from a trusted CI job or a developer machine with the right permissions.
- Protect the branch used to produce releases and ensure the CI job runs tests before tagging.
- Consider creating tags from a dedicated release workflow that also generates release notes.

Path
----
This skill lives at `.opencode/skills/add-new-version-tag/SKILL.md` in the repo.
