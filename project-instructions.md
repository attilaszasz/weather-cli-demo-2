<!-- template-version: 2 -->
# weather-cli Project Instructions

## Core Principles

### I. Simplicity First

The project MUST optimize for the smallest implementation that delivers the current CLI value cleanly, and SHOULD reject speculative features beyond the documented MVP scope — a narrow, understandable codebase is the fastest path to a dependable standalone executable.

### II. Contract Stability

Public CLI behavior MUST remain stable once introduced, especially argument semantics, JSON response structure, and exit-code meaning — automation users depend on predictable interfaces more than on rapid feature expansion.

### III. Testable Reliability

All shipped behavior MUST be covered by automated checks appropriate to the risk, with special focus on validation, provider integration boundaries, JSON contracts, and release safety — reliability for a CLI product is only credible when failure paths and packaging are verifiable.

### IV. Release Automation Early

Build, test, lint, security scanning, and cross-platform release automation MUST be established early and kept green as the product evolves — distributable binaries are a core product outcome, not a late-stage packaging detail.

### V. Agent Output Style

All agent output MUST be concise and outcome-oriented. This principle supersedes any verbose defaults.

- **Progress reports**: Facts and outcomes only — no narration, no restating the task.
- **Artifacts**: Emit required sections only — no preamble paragraphs, no summary epilogues.
- **Reasoning**: Omit unless the user asks "why" or the decision is non-obvious.
- **Errors / blockers**: State the problem, the attempted fix, and the result — nothing else.
- **Phase-boundary reports**: ≤ 5 bullet points.
- **Preserve without compressing**: Artifact template structure and required sections; explicit decision / registration / validation guidance in shared skills; delegation constraints and sub-agent role definitions; existing size limits (spec ≤ 10 KB, research ≤ 4 KB, stories ≤ 200 words).

## Technology Stack

- **Language/Runtime**: Go 1.24
- **Frameworks**: Cobra, Go standard library, Open-Meteo integration client, GoReleaser
- **Storage**: none
- **Infrastructure**: local CLI runtime, GitHub Actions CI, GitHub Releases artifact distribution

## Testing & Quality Policy

- **Coverage Target**: 80%
- **Required QC Categories**: linting, security scanning, coverage, performance
- **Test Strategy**: Test-after with unit and integration coverage for validation, provider mapping, JSON contracts, and release-critical paths; benchmark critical command latency when behavior stabilizes
- **Linting / Formatting**: golangci-lint, `gofmt`, and `go test` must pass cleanly

## Source Code Layout

- **Policy**: ENFORCE_SRC_ROOT
- **Convention**: All source code lives under `/src`; repository-root files are limited to governance, configuration, release automation, and documentation; tests may live under `/src` packages or a dedicated `/tests` directory when justified

## Development Workflow

- **Branching**: Feature branches from main with squash merge
- **Commit Convention**: Conventional Commits
- **CI Requirements**: All pull requests and release candidates must pass `go test`, lint, security scanning, and cross-platform build validation before merge or publication

<!-- Optional: add additional sections below (Security Requirements, Performance Standards, Compliance, etc.) -->

## Governance

- Project instructions supersede all other documentation and practices.
- Amendments require a version bump with ISO-dated changelog entry.
- All implementations MUST pass the Instructions Check gate during planning.
- Complexity beyond these principles MUST be justified and documented.

- Bootstrap document registrations in `.github/sddp-config.md` for `specs/prd.md`, `specs/sad.md`, and `specs/project-plan.md` MUST be preserved unless explicitly replaced by the user.
- Changes that affect the CLI contract, provider boundary, or release automation MUST update the relevant bootstrap artifacts before downstream implementation continues.
- Source-code layout enforcement applies to implementation files only; workflow files, release configuration, and governance documents may remain at repository-standard locations outside `/src`.

**Version**: 1.0.0 | **Last Amended**: 2026-04-06
