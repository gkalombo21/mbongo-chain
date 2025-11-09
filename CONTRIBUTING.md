# Contributing to Mbongo-Chain

Welcome to Mbongo-Chain! This guide outlines how to get started, collaborate effectively, and follow the project’s engineering standards.

## 1. Development Setup

1. Install **Go 1.21+** (`go version`).
2. Fork the repository, then clone your fork:
   ```bash
   git clone https://github.com/<your-username>/mbongo-chain.git
   cd mbongo-chain
   ```
3. Synchronise dependencies and run the project once to confirm everything works:
   ```bash
   go mod tidy
   go run ./cmd/mbongo-chain
   ```
4. Prefer **Cursor AI** (recommended) or **VS Code** as your IDE and read `docs/architecture/MBONGO_SYSTEM_OVERVIEW.md` plus `docs/ai/AI_JOB_EXECUTION_FLOW.md` for project context.

## 2. Branching & Workflow

1. Update your fork from upstream and create a feature branch:
   ```bash
   git checkout -b feature/<feature-name>
   ```
2. Make focused changes. Keep commits atomic and descriptive.
3. Format code and run the full test suite before committing:
   ```bash
   go fmt ./...
   go test ./...
   ```
4. Commit with a clear message (English, imperative mood):
   ```bash
   git commit -m "Add feature <name>"
   ```
5. Push your branch and open a Pull Request:
   ```bash
   git push origin feature/<feature-name>
   ```

## 3. Coding Guidelines

- Use English for all identifiers, comments, documentation, and commit messages.
- Keep functions small and modular (target < 50 lines) with meaningful names.
- Avoid global variables; prefer dependency injection.
- Document every exported type or function with a concise Go comment.
- Add unit tests alongside new modules and follow table-driven patterns when possible.
- Do not introduce new dependencies without approval from the core maintainers.

## 4. Pull Request Rules

- One PR should represent one logical change.
- Provide a clear description, referencing related GitHub issues when applicable.
- Ensure `go fmt ./...` and `go test ./...` pass before submission.
- Highlight any AI-assisted changes in the PR description.

## 5. Core Team Workflow

- Core maintainers review and approve all PRs.
- Changes touching critical areas (`internal/ai`, `internal/blockchain`) require **two approvals**.
- Feature requests or substantial design changes should be proposed via GitHub Issues before implementation.

## 6. Community Conduct & Security

- Follow the behaviour guidelines in `CODE_OF_CONDUCT.md`.
- Never commit secrets. Use environment variables or secret managers.
- Report vulnerabilities responsibly to `security@mbongo-chain.org`.

---

Together we’re building Mbongo-Chain—the global bridge between decentralized AI and finance. Thank you for contributing! 🌍🚀

