# Implementation Plan: Remove Command

**Branch**: `003-remove-command` | **Date**: 2026-02-23 | **Spec**: specs/003-remove-command/spec.md

## Summary

Implement `wt remove [path]` supporting both non-interactive removal by path and interactive selection when no path is given. Use `git worktree remove <path>` to perform removal and ensure the main worktree is never removable.

## Technical Context

- Language: Go
- Key files: `cmd/wt/main.go` has interactive removal (`interactiveRemove()`) and porcelain parsing helpers. Ensure safety checks and error messages.

## Plan

1. Harden interactiveRemove to confirm deletions and refuse removal of repository top-level.
2. Ensure `wt remove <path>` calls `git worktree remove <path>` and surfaces errors (already present in main.go).
3. Add tests for interactive and non-interactive flows.
4. Update quickstart and spec documentation.
