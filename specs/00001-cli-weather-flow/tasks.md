---

description: "Task list for CLI Weather Flow"
---

# Tasks: CLI Weather Flow

**Input**: Design documents from `specs/00001-cli-weather-flow/`
**Prerequisites**: `plan.md` (required), `spec.md` (required), `research.md`, `data-model.md`, `contracts/`

**Tests**: Include test tasks because the specification explicitly requires automated verification of validation and provider-response mapping behavior.

**Organization**: Tasks are grouped by user story so each flow stays independently testable while shared CLI scaffolding lands once in the earliest blocking phases.

## Project Mode

`Greenfield`

## Epic / Capability Map *(OPTIONAL)*

- `[US1]` → Coordinate-based successful current-weather lookup
- `[US2]` → Invalid coordinate rejection before provider calls
- `[US3]` → Safe handling of upstream transport and unusable-response failures

## Phase 1: Setup (Repository / Workspace Delta)

- [ ] T001 Initialize Go module metadata in go.mod
- [ ] T002 [P] Create CLI entrypoint scaffold in src/cmd/weathercli/main.go
- [ ] T003 [P] Add Cobra root command scaffold in src/cmd/weathercli/root.go

---

## Phase 2: Foundational (Cross-Work-Item Blockers)

- [ ] T004 Create shared weather domain types in src/internal/weather/types.go
- [ ] T005 [P] Create Open-Meteo transport and payload types in src/internal/provider/openmeteo/types.go
- [ ] T006 [P] Create JSON payload structs in src/internal/output/payload.go and failure structs in src/internal/output/failure.go

---

## Phase 3: Work Item 1 - Get Current Weather (Priority: P1) 🎯 MVP

- [ ] T007 [P] [US1] {FR-001,FR-002} Implement coordinate parsing helpers in src/internal/validation/coordinates.go
- [ ] T008 [P] [US1] {FR-008} Add validation unit tests in src/internal/validation/coordinates_test.go
- [ ] T009 [US1] {FR-003} Implement Open-Meteo client request flow in src/internal/provider/openmeteo/client.go
- [ ] T010 [P] [US1] {FR-008} Add mocked provider success tests in src/internal/provider/openmeteo/client_test.go
- [ ] T011 [US1] {FR-003,FR-004,FR-005} Implement weather orchestration service in src/internal/weather/service.go
- [ ] T012 [US1] {FR-004,FR-005} Implement approved MVP success payload builder in src/internal/output/payload.go
- [ ] T013 [US1] {FR-001,FR-003,FR-004,FR-005} Wire the CLI command success flow in src/cmd/weathercli/root.go
- [ ] T014 [US1] {FR-008} Add service-level success path tests in src/internal/weather/service_test.go
- [ ] T015 [US1] {FR-008} Add provider success fixtures in src/tests/testdata/openmeteo-success.json

---

## Phase 4: Work Item 2 - Reject Invalid Coordinates (Priority: P1) 🎯 MVP

- [ ] T016 [US2] {FR-002,FR-006} Enforce missing, malformed, and range validation in src/internal/validation/coordinates.go
- [ ] T017 [P] [US2] {FR-006,FR-008} Expand invalid-input validation tests in src/internal/validation/coordinates_test.go
- [ ] T018 [US2] {FR-006} Add lightweight command failure handling for invalid input in src/internal/output/failure.go
- [ ] T019 [US2] {FR-006} Connect invalid-input failure flow in src/cmd/weathercli/root.go

---

## Phase 5: Work Item 3 - Handle Upstream Failure Safely (Priority: P2)

- [ ] T020 [US3] {FR-007} Implement transport and unusable-response failure mapping in src/internal/provider/openmeteo/client.go
- [ ] T021 [P] [US3] {FR-007,FR-008} Add mocked upstream failure tests in src/internal/provider/openmeteo/client_test.go
- [ ] T022 [US3] {FR-007} Implement machine-readable provider failure payloads in src/internal/output/failure.go
- [ ] T023 [US3] {FR-007} Propagate provider failure outcomes through src/internal/weather/service.go
- [ ] T024 [US3] {FR-008} Add service-level failure path tests in src/internal/weather/service_test.go
- [ ] T025 [US3] {FR-008} Add unusable-response fixture in src/tests/testdata/openmeteo-invalid.json

---

## Phase 6: Polish & Cross-Cutting Concerns *(OPTIONAL)*

- [ ] T026 [P] Update CLI usage and coordinate-only behavior notes in src/cmd/weathercli/root.go
- [ ] T027 [P] Run gofmt-oriented cleanup across src/cmd/weathercli/main.go
- [ ] T028 [P] Run gofmt-oriented cleanup across src/internal/provider/openmeteo/client.go
- [ ] T029 [P] Run gofmt-oriented cleanup across src/internal/weather/service.go
