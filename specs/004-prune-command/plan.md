# Implementation Plan: Prune Command

**Branch**: `004-prune-command` | **Date**: 2026-02-23 | **Spec**: specs/004-prune-command/spec.md

## Summary

Implement `wt prune` which delegates to `git worktree prune` at repository root and reports success/failure to the user.

## Technical Context

- Location: `cmd/wt/main.go` already contains a `case "prune"` implementation that runs `runGit("worktree","prune")`. Verify error handling and messaging.

## Plan

1. Add tests to ensure `wt prune` invokes `git worktree prune` and handles no-op cases gracefully.
2. Improve messaging when no stale worktrees are found.
