# Implementation Plan: Exec Command

**Branch**: `005-exec-command` | **Date**: 2026-02-23 | **Spec**: specs/005-exec-command/spec.md

## Summary

Implement `wt exec <command> [args...]` to list non-main worktrees for selection and then execute the provided command in the selected worktree directory while attaching stdin/stdout/stderr.

## Technical Context

- `cmd/wt/main.go` already contains `interactiveExec()` and `exec` case wiring. Verify selection logic, command execution, and edge case handling.

## Plan

1. Add tests for executing commands in selected worktree and behavior when none are available.
2. Ensure commands that require stdin work by connecting `cmd.Stdin` to `os.Stdin` (already implemented).
3. Improve error messages when command is not found or returns non-zero.
