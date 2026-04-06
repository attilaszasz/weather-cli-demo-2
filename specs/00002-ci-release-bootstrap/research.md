# Research: CI Release Bootstrap
> Feature | 2026-04-06 | Purpose: inform GitHub Actions and GoReleaser setup for cross-platform CLI releases

## GitHub Actions for Go
- **Decision**: Use GitHub Actions as the canonical CI and release orchestrator for the Go CLI.
- **Rationale**: It aligns with the project technical context and supports build, test, and release automation directly in the repository.
- **Rejected**: Manual release scripting, because it conflicts with the project governance requirement for early release automation.
- **Pitfalls**: Workflow steps must pin stable actions and use repository-native commands to avoid drift between local and CI behavior.
- **Sources**: https://docs.github.com/actions, https://go.dev/doc/tutorial/create-module

## GoReleaser Packaging
- **Decision**: Use GoReleaser to build and package binaries for macOS, Windows, and Linux from a single config file.
- **Rationale**: GoReleaser is the project-approved packaging path and reduces bespoke per-platform release logic.
- **Rejected**: Hand-authored per-OS archive scripts, because they increase maintenance and artifact inconsistency risk.
- **Pitfalls**: Binary path, main package location, archive naming, and snapshot validation need to match the repository’s `/src/cmd/weathercli` layout.
- **Sources**: https://goreleaser.com/, https://goreleaser.com/ci/actions/

## Summary
| Topic | Decision | Rationale |
|-------|----------|-----------|
| CI platform | GitHub Actions | Matches project architecture and governance |
| Packaging | GoReleaser | Standard cross-platform release path for Go CLI |

## Sources Index
| URL | Topic | Fetched |
|-----|-------|---------|
| https://docs.github.com/actions | GitHub Actions for Go | 2026-04-06 |
| https://go.dev/doc/tutorial/create-module | GitHub Actions for Go | 2026-04-06 |
| https://goreleaser.com/ | GoReleaser Packaging | 2026-04-06 |
| https://goreleaser.com/ci/actions/ | GoReleaser Packaging | 2026-04-06 |
