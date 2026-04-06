---
feature_branch: "[00004-provider-abstraction]"
created: "2026-04-06"
input: "E004 Provider Abstraction"
spec_type: "technical"
spec_maturity: "clarified"
epic_id: "E004"
epic_sources: "{PRD:CAP-005}{SAD:ADR-002}"
---

# Feature Specification: Provider Abstraction

**Feature Branch**: `[00004-provider-abstraction]`  
**Created**: 2026-04-06  
**Status**: Clarified  
**Spec Type**: technical  
**Spec Maturity**: clarified  
**Epic ID**: E004  
**Epic Sources**: {PRD:CAP-005}{SAD:ADR-002}  
**Product Document**: specs/prd.md

## Problem Statement *(mandatory)*

The CLI currently works, but the weather service and failure handling are still coupled directly to Open-Meteo response and error types. Maintainers are affected because provider-specific shapes leak beyond the adapter boundary, making future provider changes riskier than they need to be. If this boundary is not refined now, any later provider evolution will require broader internal rewiring.

## Scope *(mandatory)*

### Included

- Introduce provider-neutral interfaces and data types in the internal provider layer
- Move Open-Meteo specific payload mapping fully inside the Open-Meteo package
- Keep failure classification compatible while shifting command and service layers to depend on provider-neutral errors
- Add compatibility tests proving the public contract fixtures remain unchanged after the abstraction split

### Excluded

- Supporting a second live provider
- Adding user-facing provider selection
- Changing the public success or failure JSON contract

### Edge Cases & Boundaries

- The abstraction must not change existing exit codes or JSON fixtures
- Open-Meteo remains the default provider implementation after the refactor
- The command layer must not import provider-specific packages directly once the abstraction is complete

## Technical Objectives *(mandatory for technical specs only)*

### Objective 1 - Isolate Provider-Specific Logic (Priority: P1)

Create a provider-neutral contract so the weather service depends on generic provider types instead of Open-Meteo response structs.

**Why this priority**: This is the core maintainability goal of the epic.

**Deliverables**:
- Provider-neutral interface and data types
- Open-Meteo adapter implementing that interface

**Validation Criteria**:
1. **Given** the weather service is compiled after the refactor, **When** it is wired to the default provider, **Then** it depends only on provider-neutral contracts.

### Objective 2 - Preserve Public Compatibility (Priority: P1)

Keep the public CLI contract unchanged while the internal provider boundary is refactored.

**Why this priority**: The epic must not undo the stable contract work from E003.

**Deliverables**:
- Compatibility tests proving success and failure fixtures remain unchanged

**Validation Criteria**:
1. **Given** the CLI is run through existing contract tests, **When** the abstraction is complete, **Then** all public JSON fixtures and exit-code expectations still pass.

### Technical Constraints

- All implementation remains under `/src`
- Open-Meteo stays as the default runtime provider
- No public contract drift is allowed

## Integration Points *(mandatory for technical specs)*

- **IP-001**: Consumes the existing command and weather service layers from E001 and E003
- **IP-002**: Preserves the public JSON fixtures and exit-code contract from E003
- **IP-003**: Builds on the release baseline from E002 without requiring workflow changes

## Requirements *(mandatory)*

### Technical Requirements *(technical specs only)*

- **TR-001**: System MUST define a provider-neutral internal interface for current weather retrieval.
- **TR-002**: System MUST keep Open-Meteo transport and raw response mapping inside the Open-Meteo package.
- **TR-003**: System MUST expose provider-neutral error types to downstream internal packages.
- **TR-004**: System MUST preserve the public CLI success and failure contract unchanged after the refactor.
- **TR-005**: System MUST include compatibility tests proving the public contract remains stable.

### Key Entities *(include for product or technical specs if feature involves data)*

- **Provider Adapter**: The concrete implementation that talks to Open-Meteo.
- **Provider Weather Result**: The provider-neutral weather data returned to the service layer.
- **Provider Error**: The provider-neutral failure type used outside the adapter package.

## Assumptions & Risks *(mandatory)*

### Assumptions

- The single-provider runtime is still enough for MVP after the abstraction is introduced.
- Internal refactoring can be validated fully with automated tests and existing fixtures.

### Risks

- **Hidden coupling** *(likelihood: medium, impact: medium)*: Provider-specific assumptions may still be embedded in tests or command wiring.
- **Compatibility regression** *(likelihood: low, impact: high)*: Public fixtures could drift if mapping behavior changes during the refactor.

## Implementation Signals *(mandatory)*

- **NEW-API** — Introduces an internal provider boundary contract.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001 [OBJ1]**: The weather service and command layer compile without importing Open-Meteo-specific response types.
- **SC-002 [OBJ2]**: Existing public JSON fixtures and exit-code tests continue to pass unchanged.
