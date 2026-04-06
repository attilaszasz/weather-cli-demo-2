---

description: "Task list for JSON Contract Hardening"
---

# Tasks: JSON Contract Hardening

**Input**: Design documents from `specs/00003-json-contract-hardening/`
**Prerequisites**: `plan.md` (required), `spec.md` (required)

**Tests**: Include test tasks because the specification explicitly requires fixture-backed contract assertions and exit-code verification.

**Organization**: Tasks are grouped by user story so the stable success contract, stable failure contract, and exit-code semantics remain independently testable.

## Project Mode

`Greenfield`

## Epic / Capability Map *(OPTIONAL)*

- `[US1]` → Stable success envelope
- `[US2]` → Structured failure envelope
- `[US3]` → Distinct shell exit codes

## Phase 1: Setup (Repository / Workspace Delta)

- [X] T001 Create autopilot-ready feature workspace artifacts in specs/00003-json-contract-hardening

---

## Phase 2: Foundational (Cross-Work-Item Blockers)

- [X] T002 [P] Create a reusable exit-code mapper in src/internal/exitcode/exitcode.go
- [X] T003 [P] Reshape public success and failure envelope types in src/internal/output/payload.go and src/internal/output/failure.go

---

## Phase 3: Work Item 1 - Parse Stable Success Output (Priority: P1) 🎯 MVP

- [X] T004 [US1] {FR-001} Implement the normalized success envelope in src/internal/output/payload.go
- [X] T005 [P] [US1] {FR-007} Add success contract fixture in src/tests/testdata/contract-success.json
- [X] T006 [US1] {FR-001,FR-007} Lock success contract tests in src/internal/output/payload_test.go and src/cmd/weathercli/root_test.go

---

## Phase 4: Work Item 2 - Distinguish Failure Categories Reliably (Priority: P1) 🎯 MVP

- [X] T007 [US2] {FR-002,FR-003,FR-006} Implement public failure classification in src/internal/output/failure.go
- [X] T008 [P] [US2] {FR-007} Add failure contract fixtures in src/tests/testdata/contract-failure-validation.json, src/tests/testdata/contract-failure-network.json, and src/tests/testdata/contract-failure-provider.json
- [X] T009 [US2] {FR-002,FR-003,FR-006,FR-007} Update command failure tests in src/cmd/weathercli/root_test.go

---

## Phase 5: Work Item 3 - Branch Script Flow by Exit Code (Priority: P1) 🎯 MVP

- [X] T010 [US3] {FR-004,FR-005} Wire real exit-code mapping through src/cmd/weathercli/main.go and src/cmd/weathercli/root.go
- [X] T011 [P] [US3] {FR-005,FR-007} Add exit-code unit coverage in src/cmd/weathercli/root_test.go and src/internal/exitcode/exitcode_test.go

---

## Phase 6: Polish & Cross-Cutting Concerns *(OPTIONAL)*

- [X] T012 [P] Re-run formatting and keep contract helper naming consistent across src/cmd/weathercli and src/internal/output
