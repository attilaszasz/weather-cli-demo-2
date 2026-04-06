# Data Model: CI Release Bootstrap

## Entity: CIWorkflow

- **Purpose**: Validates the repository with Go build, test, coverage, and vulnerability checks.
- **Fields**:
  - `trigger`
  - `go_version`
  - `validation_steps`
  - `status`
- **Relationships**:
  - Uses one `CLIEntrypoint`

## Entity: ReleaseWorkflow

- **Purpose**: Runs the release or snapshot packaging path in GitHub Actions.
- **Fields**:
  - `trigger`
  - `permissions`
  - `release_steps`
  - `artifact_targets`
- **Relationships**:
  - Uses one `GoReleaserConfig`

## Entity: GoReleaserConfig

- **Purpose**: Defines build targets, archive naming, and snapshot behavior.
- **Fields**:
  - `project_name`
  - `builds`
  - `archives`
  - `snapshot`
- **Relationships**:
  - Packages one `CLIEntrypoint`

## Entity: ReleaseArtifact

- **Purpose**: Packaged binary output for a target operating system.
- **Fields**:
  - `os`
  - `arch`
  - `archive_name`
- **Relationships**:
  - Produced by one `ReleaseWorkflow`
