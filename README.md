# GitWave

Effortlessly clone, sync, and ride the latest changes across all your repositories.

GitWave clones and pulls the latest data from a list of GitHub repositories, which can be helpful to run and detect a particular code pattern in the entire codebase for example quick verification of some code patterns or hardcoded keys with a custom SemGrep rule.

## Getting started

> Ensure $(go env GOPATH)/bin is in your $PATH

```bash
go install github.com/c0d3G33k/GitWave@main
```

## Usage

Setup GitHub personal access token (PAT) as the environment variable

```bash
export GITHUB_PAT=YOUR_GITHUB_PAT
```

Sync list of GitHub repositories

```bash
GitWave sync --input /path/to/csv/file --directory /path/to/work/dir/
```

## Input file CSV format

Create a CSV file with a list of GitHub repositories e.g.

```csv
https://github.com/boringtools/git-alerts
https://github.com/boringtools/polgate
```
