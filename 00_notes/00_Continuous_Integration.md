# Continuous Integration

## Running a Repo

Fork a Repo to our remote repository

Clone the Repo to our local

## Branching

```
git branch
git switch -c addtests
git push origin addtests
```

add, commit, push

## WorkFlows

GitHub Actions

```
mkdir .github
mkdir .github/workflows
touch .github/workflows/ci.yml
```

Sample that forces a failure

Typical script:

- Check out code
- install dependencies
- run tests

Script below:

- Set up Go
- Check out code
- Force failure of the CI job

```
name: ci

on:
  pull_request:
    branches: [main]

jobs:
  tests:
    name: Tests
    runs-on: ubuntu-latest

    steps:
      - name: Check out code
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.23.0'

      - name: Force Failure
        run: (exit 1)
```

Workflow > Job > Step

uses: external pre-defined Github action

run: custom command
