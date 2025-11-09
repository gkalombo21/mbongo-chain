# 🌍 Mbongo-Chain

**A hybrid blockchain for decentralized AI compute and digital banking.**

---

## 🧠 Overview

Mbongo-Chain is an open-source hybrid blockchain combining **Proof of Stake (PoS)** and **Proof of Useful Work (PoUW)** to reward AI-driven computation instead of wasting GPU energy on meaningless hashes.

### Core ideas:

- **Validators (PoS)** secure the network via staking.  
- **Compute nodes (PoUW)** contribute GPU power to perform real AI and ML workloads.  
- **Rewards** are distributed in **MBG tokens** proportional to useful work done.

---

## ⚙️ Architecture

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- Git
- Podman (recommended) or Docker

### Installation

```bash
git clone https://github.com/gkalombo21/mbongo-chain.git
cd mbongo-chain
go mod tidy
go run ./cmd/mbongo-chain
```

---

## 🧑‍💻 Developer Workflow

- All source code lives in the `internal/` directory.
- Follow idiomatic Go style conventions.
- Use feature branches (e.g. `feature/ai-networking`).
- Submit pull requests for review.

---

## 🧩 Project Modules

- **AI Compute Engine:** manages distributed GPU compute jobs.
- **Blockchain Core:** manages consensus and validator logic.
- **Banking Module:** enables tokenized accounts, transactions, and staking.
- **User Module:** manages identity and access permissions.

---

## 🧱 Future Goals

- Multi-chain interoperability (Cosmos SDK bridge).
- On-chain marketplace for AI inference jobs.
- Node deployment tools for GPU hosts.

---

## 👥 Community & Contribution

Please read the [`CONTRIBUTING.md`](CONTRIBUTING.md) and [`CODE_OF_CONDUCT.md`](docs/community/CODE_OF_CONDUCT.md) before contributing.

Discussions, PRs, and feature suggestions are always welcome.

---

## 📄 License

Mbongo-Chain is open-source under the MIT License.
