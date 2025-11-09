# 🏗️ Mbongo Chain System Architecture Overview

## 1. Introduction

The **Mbongo Chain** project combines decentralized blockchain technology, AI computation, and a hybrid consensus model known as **Proof of Useful Work (PoUW)**.  
Its goal is to power a sustainable, useful, and AI-driven ecosystem that rewards contributors for real computation.

---

## 2. System Layers

The Mbongo system is organized into **four primary layers**:

| Layer | Description | Example Modules |
|--------|--------------|----------------|
| **Core Blockchain Layer** | Manages blocks, transactions, and consensus. | `internal/blockchain`, `internal/bank` |
| **AI Compute Layer** | Executes AI jobs using GPU/CPU nodes under PoUW. | `ai/policy`, `ai/security` |
| **API Layer** | Provides REST/gRPC endpoints for users and apps. | `internal/api`, `routes.go`, `handlers.go` |
| **Application Layer** | Handles user logic, wallets, and Kaayu integrations. | `cmd/mbongo-chain`, `pkg/common` |

---

## 3. Module Overview

### 3.1 Core Blockchain (`internal/blockchain`)

- Responsible for block creation, validation, and storage.  
- Uses a **hybrid consensus** of Proof of Stake + Proof of Useful Work.  
- Core files include:  
  - `block.go` — defines the block data model  
  - `chain.go` — handles chain synchronization  
  - `consensus.go` — manages validator logic and staking  

---

### 3.2 Bank Module (`internal/bank`)

- Handles token balances, staking, and rewards.  
- Linked to PoUW smart contracts for payout distribution.  
- Files:  
  - `account.go` — manages wallet creation  
  - `transaction.go` — logs transfers and staking  
  - `ledger.go` — keeps immutable transaction history  

---

### 3.3 AI Module (`ai/policy` and `ai/security`)

- Provides AI compute coordination and security verification.  
- Files:  
  - `AI_COMPUTE_GUIDE.md` — explains the compute flow  
  - `AI_COMPUTE_SECURITY.md` — describes ZK-proof and encryption models  

---

### 3.4 API Module (`internal/api`)

- Exposes endpoints for external applications.  
- Supports JSON-RPC, REST, and WebSocket interfaces.  
- Future roadmap includes GraphQL and gRPC.  

---

### 3.5 Utility Module (`internal/utils`)

- Common utilities used across the system:  
  - Logging (`logger.go`)  
  - Configuration management (`config.go`)  
  - Cryptographic helpers (`crypto.go`)  

---

## 4. Consensus Model

Mbongo Chain runs on a **hybrid consensus**:

- **Proof of Stake (PoS):** Validators lock MBG tokens to secure the chain.  
- **Proof of Useful Work (PoUW):** GPU and CPU operators contribute real computation (AI training, data processing, etc.).  

Reward Distribution:

---

## 5. Data Flow Example

---

## 6. Network Scaling

For production, the network will include:

- **Validator Nodes (PoS)**  
- **Compute Nodes (AI PoUW)**  
- **API Nodes (Public access)**  

Each node communicates using gRPC and REST over secure TLS channels.
