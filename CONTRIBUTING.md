# 🤝 Contributing to Mbongo-Chain

Mbongo-Chain fuses decentralized AI compute with digital banking rails to deliver a programmable, energy-efficient financial ecosystem. We welcome contributors who share our mission and can support code, documentation, AI/ML modules, research, security analysis, and community operations.

This guide codifies the standards we expect for a production-grade blockchain network. Please read it fully before submitting work.

## 1️⃣ Introduction

- **Mission**: Deploy a hybrid Proof of Stake + Proof of Useful Work network where validators secure the chain and GPU operators earn MBG for real AI workloads.
- **Contribution Streams**:
  - Core protocol engineering (consensus, staking, networking, cryptography under `internal/`).
  - AI compute schedulers, proofs of useful work, and runtime orchestration.
  - Wallets, SDKs, APIs, and application tooling (`app/`).
  - Documentation, governance proposals, tokenomics studies, and security research.
  - DevOps, CI/CD automation, observability, and blockchain infrastructure.
- **Collaboration Principles**: Transparency, reproducibility, respectful debate, and a security-first mindset.

## 2️⃣ Development Setup

1. **Prerequisites**
   - Go 1.21+ (`go version`)
   - Git (latest stable release)
   - Podman (preferred) or Docker
   - Optional: Cursor or VS Code with Go tooling
2. **Clone the Repository**
   ```bash
   git clone https://github.com/<your-username>/mbongo-chain.git
   cd mbongo-chain
   git remote add upstream https://github.com/gkalombo21/mbongo-chain.git
   ```
3. **Bootstrap Dependencies**
   ```bash
   go mod tidy
   go run ./cmd/mbongo-chain
   ```
4. **Context Reading**
   - `docs/architecture/MBONGO_SYSTEM_OVERVIEW.md`
   - `docs/ai/AI_JOB_EXECUTION_FLOW.md`
   - `docs/quantum-strategy.md`
   - `docs/tokenomics.yaml`

## 3️⃣ Branching & Workflow

1. Sync with upstream before coding:
   ```bash
   git fetch upstream
   git checkout main
   git pull upstream main
   ```
2. Create a focused branch:
   ```bash
   git checkout -b feature/<scope>
   ```
3. Keep commits atomic using imperative tense and issue references:
   ```bash
   git commit -m "Implement PoUW reward accounting"
   ```
4. Rebase on `upstream/main` before opening a PR:
   ```bash
   git fetch upstream
   git rebase upstream/main
   ```
5. Push and open a pull request once ready:
   ```bash
   git push origin feature/<scope>
   ```

## 4️⃣ Code Quality Standards

- **Language & Style**: `gofmt` is mandatory; respect idiomatic Go practices and follow `golangci-lint` recommendations.
- **Modularity**: Keep functions under ~50 lines; prefer interfaces and dependency injection; avoid global state.
- **Documentation**: Comment every exported package, struct, and function using GoDoc conventions.
- **Concurrency**: Guard shared resources explicitly; prefer channels or scoped mutexes; document locking assumptions.
- **Dependencies**: Justify new modules in PRs; update `go.mod` via `go get` plus `go mod tidy`.
- **Observability**: Reuse logging/metrics helpers in `internal/utils`; avoid ad-hoc logging formats.

## 5️⃣ Testing & Verification

- Unit tests are required for new logic; adopt table-driven patterns and cover failure cases.
- Integration tests live under `internal/<domain>/tests` with deterministic fixtures.
- Run baseline checks before commit:
  ```bash
  go fmt ./...
  go test ./...
  ```
- Document manual validation steps or scripts in the PR when automation is not feasible.
- Include benchmarks or profiling output for performance-sensitive changes.

## 6️⃣ Documentation & Research Contributions

- Update relevant guides (`README.md`, `docs/`, `ARCHITECTURE.md`, `TOKENOMICS.md`) alongside code.
- Place long-form analyses inside the appropriate `docs/` subdirectory and cross-reference existing strategy papers.
- Governance or monetary policy updates must align with `GOVERNANCE.md` and `docs/tokenomics.yaml`.
- Cite external research, standards (e.g., NIST PQC), or industry whitepapers when influencing design decisions.

## 7️⃣ Pull Request Checklist

Before requesting review confirm:
- [ ] Branch rebased on `upstream/main`
- [ ] `go fmt ./...` and `go test ./...` succeed locally
- [ ] CI pipeline passes or failing checks are justified
- [ ] PR description covers context, design decisions, and testing evidence
- [ ] Screenshots/logs for user-facing or operational changes included
- [ ] AI-assisted tooling disclosed in the PR summary

## 8️⃣ Review & Governance Expectations

- Core maintainers aim to review within three business days.
- Changes to `internal/ai`, `internal/blockchain`, consensus, or cryptography require **two approvals** plus a security note.
- Major architectural shifts begin as GitHub Issues or Mbongo Improvement Proposals (MIPs) for community deliberation.
- Disagreements are resolved via technical merit, performance data, and roadmap alignment.

## 9️⃣ Security & Responsible Disclosure

- Never commit secrets, private keys, or production configuration data.
- Use `.env.example` patterns and document secrets via secure deployment guides.
- Report vulnerabilities privately to `security@mbongo-chain.org` with reproduction steps and impact assessment.
- Expect coordinated disclosure timelines for high-severity findings.

## 🔟 Community Conduct

- Adhere to `CODE_OF_CONDUCT.md` across all community spaces.
- Encourage inclusive, respectful dialogue; no harassment, discrimination, or spam.
- Share learnings via documentation, community calls, or Discord to uplift builders globally.

---

Mbongo-Chain thrives on open collaboration between AI researchers, financial engineers, and builders worldwide. Thank you for helping us deliver trustworthy decentralized AI banking. 🌍🚀

