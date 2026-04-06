# Implementation Plan: Provider Abstraction

**Branch**: `[00004-provider-abstraction]` | **Date**: 2026-04-06 | **Spec**: [spec.md](C:/Endava/Endava/EndevLocal/Repos/weather-cli-demo-2/specs/00004-provider-abstraction/spec.md)

## Summary

**Goal**: Isolate Open-Meteo specifics behind a provider-neutral internal contract without changing the public CLI behavior.  
**Approach**: Introduce provider-neutral result and error types, update the Open-Meteo client to map raw payloads internally, and keep contract compatibility locked through existing fixture-backed tests.  
**Key Constraint**: Public JSON and exit-code semantics from E003 must remain unchanged.

## Technical Context

**Language/Version**: Go 1.26.1  
**Primary Dependencies**: Cobra, Go standard library, existing Open-Meteo adapter  
**Storage**: N/A  
**Testing**: `go test`, fixture-backed compatibility tests  
**Target Platform**: macOS, Windows, Linux  
**Project Mode**: greenfield  
**Constraints**: No public contract changes, Open-Meteo remains default, all code under `/src`

## Instructions Check

| Principle | Status | Notes |
|-----------|--------|-------|
| Simplicity First | PASS | Refactor stays internal and does not add a second provider |
| Contract Stability | PASS | Existing JSON and exit-code fixtures remain the compatibility guardrail |
| Testable Reliability | PASS | Compatibility coverage is part of the plan |
| Release Automation Early | PASS | No workflow changes are required |
| Source Code Layout | PASS | All changes remain under `/src` |

## Architecture Decisions

| ID | Decision | Chosen | Rationale |
|----|----------|--------|-----------|
| AD-001 | Shared provider contract location | `src/internal/provider` | Keeps downstream packages independent from concrete adapters |
| AD-002 | Adapter return type | Provider-neutral weather result | Removes Open-Meteo response coupling from the service layer |
| AD-003 | Error abstraction | Provider-neutral error type | Lets failure classification avoid concrete adapter imports |

## Requirement Coverage Map

| Req ID | Component(s) | File Path(s) |
|--------|--------------|--------------|
| TR-001 | Provider contract | `src/internal/provider/types.go` |
| TR-002 | Open-Meteo adapter | `src/internal/provider/openmeteo/client.go` |
| TR-003 | Provider-neutral errors | `src/internal/provider/types.go`, `src/internal/output/failure.go` |
| TR-004 | Compatibility preservation | `src/cmd/weathercli/root.go`, `src/internal/output/*.go` |
| TR-005 | Compatibility tests | `src/cmd/weathercli/root_test.go`, `src/internal/weather/service_test.go`, `src/internal/provider/openmeteo/client_test.go` |

## Implementation Hints

- Keep raw Open-Meteo response structs private to the adapter package where possible.
- Reuse existing contract fixtures instead of adding parallel versions.
- Remove direct `openmeteo` imports from command and service layers.
