# Research: CLI Weather Flow
> Feature | 2026-04-06 | Purpose: inform product-scope requirements for the first runnable weather-cli flow

## Open-Meteo Coordinate Access
- **Decision**: Base the MVP feature scope on explicit latitude/longitude input and current-condition retrieval from Open-Meteo.
- **Rationale**: Open-Meteo documents latitude and longitude as required inputs and supports current-condition style weather retrieval aligned with the product direction.
- **Rejected**: City-name lookup or forecast-first scope, because they expand the MVP beyond the project plan and product document.
- **Pitfalls**: Provider field shape and model-specific availability can vary, so the feature should describe stable user outcomes rather than raw provider payload details.
- **Sources**: https://open-meteo.com/en/docs/gfs-api, https://open-meteo.com/en/docs/metno-api

## CLI Machine-Readable Behavior
- **Decision**: Prioritize machine-readable output and explicit failure behavior in the feature stories and requirements.
- **Rationale**: CLI best-practice guidance emphasizes predictable output and error semantics for scripting and automation workflows.
- **Rejected**: Human-formatted terminal output as a primary experience, because the project PRD centers automation-first value.
- **Pitfalls**: Mixing display-oriented output with machine-readable responses makes downstream scripting brittle.
- **Sources**: https://clig.dev/

## Summary
| Topic | Decision | Rationale |
|-------|----------|-----------|
| Provider access | Coordinate-only Open-Meteo current weather flow | Matches epic scope and MVP boundary |
| CLI behavior | Machine-readable output first | Supports automation and stable scripting |

## Sources Index
| URL | Topic | Fetched |
|-----|-------|---------|
| https://open-meteo.com/en/docs/gfs-api | Open-Meteo Coordinate Access | 2026-04-06 |
| https://open-meteo.com/en/docs/metno-api | Open-Meteo Coordinate Access | 2026-04-06 |
| https://clig.dev/ | CLI Machine-Readable Behavior | 2026-04-06 |
