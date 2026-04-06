---
feature_branch: "[00003-json-contract-hardening]"
created: "2026-04-06"
input: "E003 JSON Contract Hardening"
spec_type: "product"
spec_maturity: "clarified"
epic_id: "E003"
epic_sources: "{PRD:CAP-003,CAP-004}{SAD:ADR-003,ADR-004}"
---

# Feature Specification: JSON Contract Hardening

**Feature Branch**: `[00003-json-contract-hardening]`  
**Created**: 2026-04-06  
**Status**: Clarified  
**Spec Type**: product  
**Spec Maturity**: clarified  
**Epic ID**: E003  
**Epic Sources**: {PRD:CAP-003,CAP-004}{SAD:ADR-003,ADR-004}  
**Product Document**: specs/prd.md

## Problem Statement *(mandatory)*

The CLI already returns machine-readable JSON, but its success and failure shapes are still lightweight MVP outputs without a locked public envelope or distinct shell exit semantics. Developers and automation users are affected because downstream scripts need a durable contract for both parsing and control flow. If this feature is not completed, later provider work will sit on top of an unstable public interface.

## Scope *(mandatory)*

### Included

- Replace the MVP success payload with a stable normalized top-level schema containing `status`, `timestamp`, `location`, `current`, and `source`
- Replace the lightweight failure payload with a structured error envelope that stays stable across validation, network, provider, and internal failures
- Map command failures to distinct non-zero exit codes for validation, network, provider, and internal error categories
- Add contract-focused tests and fixtures that lock representative success and failure outputs

### Excluded

- New weather fields beyond the current MVP weather data set
- Provider abstraction changes or alternate provider support
- Human-oriented output modes, debug flags, or stderr logging features

### Edge Cases & Boundaries

- Validation failures must produce structured JSON and a validation-specific exit code without making a provider request
- Upstream transport failures must produce a network-classified failure payload and a distinct exit code
- Unusable provider responses must produce a provider-classified failure payload and a distinct exit code
- Unexpected serialization or command-layer failures must fall back to an internal-classified failure payload or internal exit code

## User Scenarios & Testing *(mandatory for product specs only)*

### User Story 1 - Parse Stable Success Output (Priority: P1)

An automation user runs the CLI successfully and receives a stable normalized JSON envelope that is independent of raw provider shape details.

**Why this priority**: Stable success parsing is the core public contract promised by the product and technical context.

**Independent Test**: Run the CLI success flow and compare the JSON output to a locked fixture.

**Acceptance Scenarios**:

1. **Given** valid coordinates and a valid provider response, **When** the CLI succeeds, **Then** the output contains `status`, `timestamp`, `location`, `current`, and `source` in the normalized contract.
2. **Given** a downstream script consumes the success response, **When** provider-specific field names change internally, **Then** the public CLI success schema remains stable.

### User Story 2 - Distinguish Failure Categories Reliably (Priority: P1)

An automation user runs the CLI on failing inputs or provider conditions and can distinguish the failure category from the structured JSON payload.

**Why this priority**: Error parsing is as important as success parsing for script safety.

**Independent Test**: Exercise validation, network, provider-data, and internal failure mappings and compare them to the expected contract.

**Acceptance Scenarios**:

1. **Given** malformed coordinates, **When** the CLI fails, **Then** it returns a structured validation error payload and no success object.
2. **Given** an upstream transport failure, **When** the CLI fails, **Then** it returns a structured network error payload and no success object.
3. **Given** unusable upstream weather data, **When** the CLI fails, **Then** it returns a structured provider error payload and no success object.

### User Story 3 - Branch Script Flow by Exit Code (Priority: P1)

An automation user checks the command exit status and can branch logic without parsing free-form text.

**Why this priority**: Exit-code semantics are explicitly required by the SAD and are critical for shell automation.

**Independent Test**: Trigger each failure category and assert the command returns the expected non-zero exit code.

**Acceptance Scenarios**:

1. **Given** invalid input, **When** the CLI exits, **Then** it returns the validation exit code.
2. **Given** a network failure, **When** the CLI exits, **Then** it returns the network exit code.
3. **Given** provider-data or internal failures, **When** the CLI exits, **Then** it returns their distinct non-zero exit codes.

## Requirements *(mandatory)*

### Functional Requirements *(product specs only)*

- **FR-001**: System MUST emit a stable success JSON envelope with top-level `status`, `timestamp`, `location`, `current`, and `source` fields.
- **FR-002**: System MUST emit a structured failure JSON envelope with top-level `status`, `timestamp`, and `error` fields.
- **FR-003**: System MUST classify failures into validation, network, provider, or internal categories in the public error contract.
- **FR-004**: System MUST return exit code `0` for success.
- **FR-005**: System MUST return distinct non-zero exit codes for validation, network, provider, and internal failures.
- **FR-006**: System MUST keep public error payloads generic enough to avoid leaking raw upstream provider schema details.
- **FR-007**: System MUST protect representative success and failure outputs with automated contract tests and fixtures.

### Key Entities *(include for product or technical specs if feature involves data)*

- **Success Envelope**: The normalized JSON structure returned on successful weather retrieval.
- **Failure Envelope**: The normalized JSON structure returned on command failure.
- **Error Descriptor**: The public error object containing the stable failure code, message, and retryability hint.
- **Exit Code Map**: The numeric command outcome mapping used by shells and automation.

## Assumptions & Risks *(mandatory)*

### Assumptions

- The current weather field set from E001 remains sufficient for this contract-hardening slice.
- A generated response timestamp is acceptable in addition to the provider observation timestamp already present in weather data.
- Retryability is useful in the public error envelope for scripts deciding whether to retry.

### Risks

- **Contract churn** *(likelihood: medium, impact: high)*: Renaming or reshaping output fields can break early adopters if not locked with fixtures and tests.
- **Exit-code drift** *(likelihood: low, impact: high)*: Inconsistent numeric mappings across command paths would make automation unreliable.
- **Overexposed provider detail** *(likelihood: medium, impact: medium)*: Surfacing provider-specific error labels directly would weaken the normalized contract.

## Implementation Signals *(mandatory)*

- **NEW-API** — This epic finalizes the public CLI JSON and exit-code contract.
- **EXTERNAL-SERVICE** — Failure mapping still depends on Open-Meteo transport and data-validation outcomes.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001 [US1]**: The success path matches a stable normalized JSON fixture with top-level `status`, `timestamp`, `location`, `current`, and `source`.
- **SC-002 [US2]**: Validation, network, and provider-data failures each match structured failure fixtures with stable public error codes.
- **SC-003 [US3]**: The command returns distinct exit codes for validation, network, provider, and internal failures in automated tests.

## Glossary *(include when spec introduces 2+ domain-specific terms)*

| Term | Definition |
|------|------------|
| Success envelope | The top-level JSON object returned when the command succeeds. |
| Failure envelope | The top-level JSON object returned when the command fails. |
| Error descriptor | The nested JSON object that describes a public failure category. |
| Exit-code map | The stable numeric mapping from command outcome to shell exit status. |

## Clarifications

### Session 2026-04-06

- Q: Should autopilot ask follow-up questions for E003? -> A: No, because the project plan already carries the `skip_clarify` pipeline hint.
- Q: What success schema should this epic target? -> A: The SAD-recommended normalized top-level contract with `location`, `current`, `source`, and `timestamp`, plus a top-level `status`.
- Q: How should public failures classify provider problems? -> A: Use generic public categories for validation, network, provider, and internal failures rather than provider-specific labels.

## Compliance Check

### Instructions Check Report
**Target**: spec.md
**Status**: PASS

| Principle | Verdict | Notes |
|-----------|---------|-------|
| Simplicity First | PASS | Scope hardens the existing public contract without expanding weather features or provider capabilities. |
| Contract Stability | PASS | The feature explicitly locks both JSON envelopes and exit-code semantics. |
| Testable Reliability | PASS | Contract fixtures and exit-code assertions are required outcomes. |
| Release Automation Early | PASS | The work stays compatible with the existing CI and release pipeline. |
| Agent Output Style | N/A | This principle governs agent communication rather than feature behavior. |
