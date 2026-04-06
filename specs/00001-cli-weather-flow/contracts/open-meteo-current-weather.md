# Contract: Open-Meteo Current Weather

## Outbound Request

- **Method**: `GET`
- **Endpoint**: `/v1/forecast`
- **Purpose**: Retrieve the minimum Open-Meteo current weather fields needed for the MVP CLI payload
- **Auth**: None for MVP baseline
- **Query Parameters**:
  - `latitude` — required decimal
  - `longitude` — required decimal
  - `current` — required fixed field list covering temperature, wind speed, wind direction, weather code

## Successful Upstream Response

- **Required Response Fields**:
  - `latitude`
  - `longitude`
  - `current.time`
  - `current.temperature_2m`
  - `current.wind_speed_10m`
  - `current.wind_direction_10m`
  - `current.weather_code`
- **Mapping Rule**:
  - The CLI maps these upstream fields into the MVP payload fields `coordinates`, `current.temperature`, `current.wind_speed`, `current.wind_direction`, `current.weather_code`, and `observation_timestamp`

## Failure Conditions

- **Transport Failure**: timeout, DNS, or unreachable service returns a machine-readable failure result
- **Provider Data Failure**: missing or unusable required fields returns a machine-readable failure result
- **Validation Boundary**: invalid coordinates never produce this outbound request
