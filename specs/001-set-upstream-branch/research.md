# research.md

Decision: Configure upstream locally to track chosen remote/branch without auto-push

Rationale: Safer default for developer workflows — avoids accidental remote writes while ensuring subsequent push/pull target the intended branch. Matches common Git workflows where branch tracking is set locally and pushes are explicit.

Alternatives considered:
- Defer upstream until first push: simpler but increases chance of accidental pushes to default ref if users forget to set upstream.
- Auto-push when configuring upstream: convenient but risk of unexpected remote branch creation and write.
- Prompt interactively: good UX for interactive sessions, but fragile in scripting/CI.

Decision: Fail on branch name conflicts and provide clear remediation guidance

Rationale: Prevents accidental overwrites or silent reuse. Encourages explicit user intent.

Alternatives considered:
- Reuse existing branch automatically: risky for destructive changes.
- Auto-rename and proceed: surprising; may confuse collaborators.

Decision: Prefer remote named `origin` if base branch has no tracked remote

Rationale: `origin` is the de facto default for most repositories and is deterministic.

Alternatives considered:
- Prompt user to choose remote: interactive and robust but not suitable for scripting.
- Use first listed remote: arbitrary and surprising.
