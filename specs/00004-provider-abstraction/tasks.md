---

description: "Task list for Provider Abstraction"
---

# Tasks: Provider Abstraction

**Input**: Design documents from `specs/00004-provider-abstraction/`
**Prerequisites**: `plan.md` (required), `spec.md` (required)

## Project Mode

`Greenfield`

## Phase 1: Setup

- [X] T001 Create autopilot-ready feature workspace artifacts in specs/00004-provider-abstraction

## Phase 2: Foundational

- [X] T002 [P] {TR-001,TR-003} Create provider-neutral contracts in src/internal/provider/types.go
- [X] T003 [P] {TR-002} Refactor Open-Meteo adapter to return provider-neutral results in src/internal/provider/openmeteo/client.go

## Phase 3: Compatibility

- [X] T004 {TR-001,TR-003} Update the weather service to depend on provider-neutral contracts in src/internal/weather/service.go
- [X] T005 {TR-003,TR-004} Remove adapter-specific imports from failure classification in src/internal/output/failure.go and src/cmd/weathercli/root.go
- [X] T006 [P] {TR-005} Update provider, service, and command tests to prove compatibility
- [X] T007 [P] {TR-004,TR-005} Re-run public contract fixtures without changing them
