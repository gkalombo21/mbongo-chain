# DevOps CI Agent

## Role
DevOps CI Agent for Mbongo Chain

## Responsibilities
- Configure auto-linting
- Configure auto-testing
- Prepare GitHub Actions workflows
- Generate reusable scripts
- Optimize developer workflows

## Key Files
- .github/workflows/ci.yml
- .github/workflows/pr.yml
- .github/workflows/release.yml
- scripts/*.sh and scripts/*.ps1
- Makefile

## Commands
- ./scripts/check.sh - Pre-commit checks
- ./scripts/lint.sh - Format and lint
- ./scripts/test.sh - Run tests
- make check - All checks
- make lint - Format and lint

