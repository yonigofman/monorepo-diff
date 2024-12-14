# MonorepoDiff

**MonorepoDiff** is a powerful tool designed to detect changes in applications within an Nx monorepo, based on Git commit differences. It helps developers and CI/CD pipelines quickly identify which applications have been modified between two commits, enabling efficient workflows and targeted deployments.

## Features

- Detects changed applications in an Nx monorepo.
- Compares the differences between two Git commits.
- Provides a list of applications that have changed based on the files modified.
- Supports customizable commit range comparison.
- Can be integrated into CI/CD pipelines to optimize build and deployment processes.

## Requirements

- Go 1.18+ (for building and running the script).
- Git installed and properly configured.
- Nx monorepo structure (for identifying applications in a specific folder).

## Installation

1. Clone the repository to your local machine:
    ```bash
    git clone https://github.com/yonigofman/monorepo-diff.git
    ```

2. Navigate into the project folder:
    ```bash
    cd monorepo-diff
    ```

3. Build the Go script:
    ```bash
    go build -o monorepo-diff main.go
    ```

4. (Optional) Move the binary to a location in your `PATH` for easy access.

## Usage

### Command-line Arguments

- `--base` (default: `HEAD^`): The base commit to compare against (e.g., the previous commit or a specific commit hash).
- `--head` (default: `HEAD`): The head commit to compare to (typically the current commit).
- `--projects` (default: `apps`): The folder in your Nx monorepo containing the applications to check for changes.

### Example 1: Detect changes in applications between the current commit and the previous commit

```bash
./monorepo-diff --base HEAD^ --head HEAD --projects apps
