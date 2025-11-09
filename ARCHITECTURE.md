# 🧠 Mbongo-Chain Architecture

## Overview

**Mbongo-Chain** is a hybrid blockchain platform designed for real-world useful computation. It combines:

- **Proof of Stake (PoS)** — validator-based security and governance
- **Proof of Useful Work (PoUW)** — GPU nodes rewarded for verifiable AI workloads

The project is implemented in **Go 1.21** with a modular layout inspired by the Cosmos SDK, ensuring clear separation between command-line tooling, blockchain logic, AI compute orchestration, and APIs.

---

## 🧱 1. Module Map

| Path | Purpose |
|------|---------|
| `cmd/` | Node entrypoints, CLI bootstrap, configuration loading |
| `internal/bank/` | Account, balance, and ledger management |
| `internal/blockchain/` | Blocks, consensus integration, state machine |
| `internal/powu/` | Proof of Useful Work (AI job scheduling and verification) |
| `internal/api/` | REST and gRPC interfaces for external clients |
| `internal/user/` | Identity, roles, validator/worker registration |
| `internal/utils/` | Shared helpers for config, logging, and errors |
| `pkg/` | Future shared libraries and SDK components |

---

## 🧠 3. Core Components

### 🪙 Bank Module
Handles wallet creation, balances, and transactions.

- `CreateAccount()` — generates a new wallet.
- `Deposit()` / `Withdraw()` — basic financial operations.
- `Transfer()` — moves funds between accounts.
- Ledger updates are persisted on-chain.

---

### ⛓️ Blockchain Module
Responsible for:

- Block creation and validation
- State management
- Consensus enforcement (PoS + PoUW hybrid)

It ensures that every AI computation and transaction is **verified** and **recorded immutably**.

---

### 🤖 AI Module (Proof of Useful Work)
Converts GPU compute tasks into blockchain rewards.

- `job.go` — defines AI job structure and submission
- `verifier.go` — validates computation proofs
- `rewards.go` — distributes MBG tokens for verified work

**Workflow example:**
1. A GPU node executes an AI task (e.g., image classification).
2. The node submits a cryptographic proof of work done.
3. The blockchain verifies the proof through consensus.
4. The node receives MBG tokens as reward.

---

### 🌐 API Module
Provides REST and gRPC endpoints for:

- Account creation and management
- Transaction submission
- Blockchain state queries
- AI job submissions and result verifications

Example route: `POST /api/v1/ai/submit`

---

### 🔐 User Module
Manages:

- User registration
- Authentication via private/public key pairs
- Access control for different roles (validator, miner, developer)

---

### ⚙️ Utils Module
Contains helper functions used across the system:

- `config.go` — load environment and network configuration
- `logger.go` — centralized logging handler

---

## 🧮 4. Data Flow Summary

1. **Submission** — User submits a transaction or AI job via REST/gRPC.
2. **Validation** — API layer validates inputs and forwards them to blockchain services.
3. **Processing** — Blockchain logic executes state transitions and updates the ledger.
4. **Consensus** — Validators stake MBG and finalize blocks via PoS.
5. **Useful Work** — If applicable, PoUW verifies AI computations and triggers rewards.
6. **Observation** — Events and state updates are exposed back through APIs or UI dashboards.

---

## 🔄 5. Consensus Overview

| Component | Purpose | Mechanism |
|-----------|---------|-----------|
| **Proof of Stake (PoS)** | Governance and security | Validators stake MBG tokens |
| **Proof of Useful Work (PoUW)** | Productive computation | Nodes perform AI/GPU tasks |
| **Hybrid Validation** | Combine both | PoS validators confirm PoUW results |

---

## 🧰 6. Dependencies

- **Go 1.21+**
- **Gin / Gorilla Mux** — HTTP routing
- **Tendermint RPC (planned)** — inter-chain communication
- **TensorFlow / PyTorch bindings (planned)** — AI computation layer
- **GORM / SQLite (optional)** — local state persistence for tooling

---

## 🚀 7. Future Extensions

- **Smart contracts** with sandboxed WASM runtime
- **AI dataset marketplace** for verified training jobs
- **Decentralized GPU marketplace** using PoUW nodes
- **Cross-chain PoUW interoperability** with other networks

---

## 📘 8. Summary

Mbongo-Chain is a hybrid blockchain designed for usefulness—combining secure staking, AI computation, and community-driven innovation. Its modular structure ensures scalability, transparency, and long-term growth.

---

*Maintained by the Mbongo-Chain Core Development Team*  
*Last updated: November 2025*
