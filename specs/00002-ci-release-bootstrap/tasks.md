---

description: "Task list for CI Release Bootstrap"
---

# Tasks: CI Release Bootstrap

**Input**: Design documents from `specs/00002-ci-release-bootstrap/`
**Prerequisites**: `plan.md` (required), `spec.md` (required), `research.md`, `data-model.md`

**Tests**: Include validation tasks because the technical spec requires real CI and snapshot packaging verification.

## Project Mode

`Mixed`

## Epic / Capability Map *(OPTIONAL)*

- `[OBJ1]` → CI validation workflow
- `[OBJ2]` → Cross-platform release packaging
- `[OBJ3]` → Release readability and trigger clarity

## Phase 1: Setup (Repository / Workspace Delta)

- [X] T001 Create release packaging configuration in .goreleaser.yml

---

## Phase 2: Foundational (Cross-Work-Item Blockers)

- [X] T002 Create CI validation workflow in .github/workflows/ci.yml

---

## Phase 3: Work Item 1 - Validate Go Changes (Priority: P1) 🎯 MVP

- [X] T003 [OBJ1] {TR-001,TR-002} Add Go build, test, coverage, and govulncheck steps to .github/workflows/ci.yml
- [X] T004 [OBJ1] {TR-001,TR-002} Verify repository commands locally against .github/workflows/ci.yml

---

## Phase 4: Work Item 2 - Package Cross-Platform Releases (Priority: P1) 🎯 MVP

- [X] T005 [OBJ2] {TR-003,TR-004} Configure GoReleaser builds and archives in .goreleaser.yml
- [X] T006 [OBJ2] {TR-004,TR-006} Add release workflow with snapshot validation and tagged release path in .github/workflows/release.yml
- [X] T007 [OBJ2] {TR-006} Validate snapshot packaging locally with GoReleaser using the repository entrypoint

---

## Phase 5: Work Item 3 - Preserve Release Readability (Priority: P2)

- [X] T008 [OBJ3] {TR-005} Add clear trigger and artifact comments in .github/workflows/release.yml
- [X] T009 [OBJ3] {TR-005} Add matching usage comments in .goreleaser.yml
