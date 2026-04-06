# weather-cli

`weather-cli` is a standalone Go command-line tool that fetches current weather for a latitude and longitude and prints machine-readable JSON.

## What It Does

- Accepts explicit `--latitude` and `--longitude` parameters
- Fetches current conditions from Open-Meteo
- Returns structured JSON on both success and failure
- Uses stable non-zero exit codes for validation, network, provider, and internal failures

## Scope

Included:

- Current weather by coordinates
- JSON-first output for scripts and automation
- Cross-platform release artifacts for macOS, Linux, and Windows

Not included:

- Forecasts
- City-name lookup or geocoding
- Human-formatted output modes
- Caching or offline support

## Run Locally

Requirements:

- Go installed locally

Run directly from source:

```powershell
go run .\src\cmd\weathercli --latitude 44.4268 --longitude 26.1025
```

Build a local binary:

```powershell
go build -o .\bin\weathercli.exe .\src\cmd\weathercli
```

On macOS or Linux:

```bash
go build -o ./bin/weathercli ./src/cmd/weathercli
```

## Example Success Output

```json
{
  "status": "success",
  "timestamp": "2026-04-06T09:15:00Z",
  "location": {
    "latitude": 44.4268,
    "longitude": 26.1025
  },
  "current": {
    "temperature": 14.2,
    "wind_speed": 7.1,
    "wind_direction": 220,
    "weather_code": 3,
    "observation_timestamp": "2026-04-06T10:00"
  },
  "source": {
    "name": "open-meteo"
  }
}
```

## Example Failure Output

```json
{
  "status": "error",
  "timestamp": "2026-04-06T09:15:00Z",
  "error": {
    "code": "validation_error",
    "message": "latitude must be a valid decimal number",
    "retryable": false
  }
}
```

## Exit Codes

| Exit Code | Meaning |
|-----------|---------|
| `0` | Success |
| `2` | Validation failure |
| `3` | Network or transport failure |
| `4` | Provider data failure |
| `5` | Internal failure |

## Releases

Tagged releases are published from GitHub Actions when a tag matching `v*` is pushed. GoReleaser builds archives for:

- macOS
- Linux
- Windows

Archive names follow this pattern:

```text
weathercli_<version>_<os>_<arch>
```

Examples:

- `weathercli_v0.0.1_linux_amd64.tar.gz`
- `weathercli_v0.0.1_windows_amd64.zip`

Snapshot packaging can also be validated through the release workflow dispatch path.

## Quality Checks

Repository validation uses:

- `go build ./...`
- `go vet ./...`
- `go test ./...`
- `govulncheck ./...`

## Project Layout

```text
src/
  cmd/weathercli/        CLI entrypoint
  internal/validation/   coordinate parsing and validation
  internal/provider/     provider contracts and Open-Meteo adapter
  internal/weather/      weather service orchestration
  internal/output/       success and failure JSON shaping
```
