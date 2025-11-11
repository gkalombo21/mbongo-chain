# Mbongo-Chain Developer Onboarding Guide

## 1. Introduction
Welcome to the Mbongo-Chain project—a decentralized AI and financial blockchain implemented in Go and powered by the Cosmos SDK. This guide walks new developers through local setup, building, testing, and contributing effectively. Mbongo-Chain is open-source and welcomes collaborators from every background.

## 2. Prerequisites
Install and configure the following tools before you begin:
- **Go 1.21+**
- **Git**
- **Podman** *(recommended)* or **Docker** *(compatible)*
- **Cursor** *(AI-assisted development)*
- **GitHub account** with SSH access for cloning and pull requests

> 💡 *Tip:* Podman offers a rootless, lightweight workflow. Docker remains fully supported and aligns with existing CI pipelines.

## 3. Setting Up Your Environment
```bash
# Clone the repository
git clone https://github.com/gkalombo21/mbongo-chain.git
cd mbongo-chain

# Install dependencies
go mod tidy

# Copy environment file
cp .env.example .env
```
- Update `.env` with local credentials or overrides as needed.
- Open the project in Cursor or your preferred editor.
- Review `.cursor/rules`, `.cursor_style_rules.json`, and `.cursor_validate_language.json` to understand automation policies.

## 4. Running Mbongo-Chain Locally
```bash
# Start the node application
go run ./cmd/mbongo-chain
```
You should see demo ledger output showing account creation, deposits, withdrawals, and AI engine initialization. Modify code and rerun the command to iterate quickly.

## 5. Container Workflow with Podman or Docker
- Ensure Podman (preferred) or Docker is installed and running.
- Translate any `docker` commands in scripts to `podman` if using Podman.
- Create additional containers (databases, tooling) as required for integration tests.
- Use `podman ps` or `docker ps` to validate container status.

## 6. Testing and Quality Checks
```bash
# Run all unit tests
go test ./...

# Verbose output
go test -v ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```
- Run tests before every commit and pull request.
- GitHub Actions executes `go test ./...` automatically; aim for green pipelines.
- Cursor validation rules enforce Markdown structure, Go formatting, and English-only content.

## 7. Cursor and Automation Rules
- Cursor applies language validation prior to commits, builds, and Markdown saves.
- Style rules auto-correct Markdown headings, spacing, and GoDoc comments.
- Resolve all warnings flagged by Cursor before creating pull requests.

## 8. Git Workflow
1. Create a feature branch: `git checkout -b feature/<short-description>`.
2. Make code or documentation changes.
3. Run `go test ./...` to confirm stability.
4. Stage and commit: `git add . && git commit -m "Add <summary>"`.
5. Rebase with the latest main: `git pull origin main --rebase`.
6. Push your branch: `git push origin feature/<short-description>`.
7. Open a pull request with a summary, testing notes, and any screenshots or logs.

## 9. Common Issues and Fixes
- **act command errors:** install `act` with Chocolatey (`choco install act-cli`) or confirm it is present in your `PATH`.
- **Docker connection error:** switch to Podman, update scripts accordingly, and confirm the Podman service is running.
- **Permission issues on Windows:** run PowerShell or your terminal as Administrator when installing dependencies or interacting with containers.

## 10. Additional Resources
- [`README.md`](README.md) — project overview and quick links
- [`STYLE_GUIDE_FOR_CONTRIBUTORS.md`](STYLE_GUIDE_FOR_CONTRIBUTORS.md) — writing and coding standards
- [`CONTRIBUTING.md`](CONTRIBUTING.md) — contribution process
- [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md) — community expectations
- [`SECURITY.md`](SECURITY.md) — vulnerability reporting policy
- [`WHITEPAPER.md`](WHITEPAPER.md) — architecture and tokenomics deep dive
- Support: open a GitHub Issue or email `team@mbongo.io`

Welcome aboard! Your ideas and expertise help drive Mbongo-Chain toward a decentralized, AI-powered future. Reach out through GitHub Discussions or issues whenever you need assistance.

