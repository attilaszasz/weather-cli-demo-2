# Data Model: CLI Weather Flow

## Entity: CoordinateInput

- **Purpose**: Represents the latitude and longitude pair supplied to the CLI for one weather lookup.
- **Fields**:
  - `latitude` — decimal, required, valid range `-90` to `90`
  - `longitude` — decimal, required, valid range `-180` to `180`
- **Validation Rules**:
  - Both fields are required together
  - Values must parse from CLI input before provider access
  - Any invalid value stops execution before network I/O
- **Relationships**:
  - Produces one `ProviderRequest`

## Entity: ProviderRequest

- **Purpose**: Canonical outbound request sent to Open-Meteo for current weather.
- **Fields**:
  - `latitude` — copied from `CoordinateInput`
  - `longitude` — copied from `CoordinateInput`
  - `current_fields` — fixed set required for MVP output
- **Validation Rules**:
  - Generated only from validated `CoordinateInput`
  - Must request the provider fields needed for the MVP payload
- **Relationships**:
  - Returns one `ProviderResponse` on success or one `ProviderFailure` on failure

## Entity: ProviderResponse

- **Purpose**: Upstream Open-Meteo response used to derive the CLI result.
- **Fields**:
  - `latitude`
  - `longitude`
  - `current.temperature_2m`
  - `current.wind_speed_10m`
  - `current.wind_direction_10m`
  - `current.weather_code`
  - `current.time`
- **Validation Rules**:
  - Required MVP fields must all be present and usable
  - Missing required fields convert the flow into a failure path
- **Relationships**:
  - Maps to one `MVPWeatherPayload`

## Entity: MVPWeatherPayload

- **Purpose**: Successful machine-readable CLI output for the first epic.
- **Fields**:
  - `coordinates.latitude`
  - `coordinates.longitude`
  - `current.temperature`
  - `current.wind_speed`
  - `current.wind_direction`
  - `current.weather_code`
  - `observation_timestamp`
- **Validation Rules**:
  - Includes only the approved MVP output fields
  - Emits only when provider data is valid and complete
- **Relationships**:
  - Produced from one `ProviderResponse`

## Entity: ProviderFailure

- **Purpose**: Lightweight machine-readable failure result for invalid provider or transport outcomes.
- **Fields**:
  - `error_type`
  - `message`
- **Validation Rules**:
  - Must not include weather payload fields
  - Used for upstream timeout, unreachable service, or unusable provider data
- **Relationships**:
  - Alternative outcome of `ProviderRequest`
