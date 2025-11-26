# Mbongo Chain — Architecture Overview

This document provides a high-level overview of the Mbongo Chain architecture, designed for developers, researchers, and contributors who want to understand how the system works.

---

## 1. Introduction

**Mbongo Chain** is a compute-first Layer 1 blockchain built entirely in Rust. It combines **Proof of Stake (PoS)** for economic security with **Proof of Useful Work (PoUW)** for verifiable compute validation.

The protocol is optimized for:

- AI inference and training workloads
- High-performance computing (HPC) tasks
- Parallelizable batch processing
- Decentralized GPU coordination

Mbongo Chain achieves 1-second block times with deterministic execution, enabling real-time compute markets at global scale.

---

## 2. Design Philosophy

### Core Principles

| Principle | Description |
|-----------|-------------|
| **Rust-Native** | Built from scratch in Rust with no legacy code or compatibility layers |
| **Secure-by-Default** | Memory safety, type safety, and explicit error handling throughout |
| **Modular** | Clear separation between networking, consensus, execution, and storage |
| **Developer-Friendly** | Clean APIs, comprehensive documentation, and intuitive tooling |
| **Compute-First** | Designed around verifiable compute rather than simple token transfers |

### Why Compute-First?

Traditional blockchains optimize for financial transactions. Mbongo Chain optimizes for **compute verification** — proving that specific computations were performed correctly. This enables:

- Trustless AI model inference
- Verifiable scientific computations
- Decentralized rendering and simulation
- Provably fair resource allocation

---

## 3. High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────────┐
│                              CLI                                    │
│                    (Developer Interface)                            │
└───────────────────────────┬─────────────────────────────────────────┘
                            │
┌───────────────────────────▼─────────────────────────────────────────┐
│                             NODE                                    │
│              (Orchestration & Coordination Layer)                   │
├─────────────┬─────────────┬─────────────┬─────────────┬─────────────┤
│             │             │             │             │             │
│  ┌──────────▼──────────┐  │  ┌──────────▼──────────┐  │             │
│  │     NETWORKING      │  │  │       MEMPOOL       │  │             │
│  │   (P2P Protocol)    │  │  │  (Tx Queue/Relay)   │  │             │
│  └──────────┬──────────┘  │  └──────────┬──────────┘  │             │
│             │             │             │             │             │
│             └─────────────┼─────────────┘             │             │
│                           │                           │             │
│             ┌─────────────▼─────────────┐             │             │
│             │         RUNTIME           │             │             │
│             │   (State Machine Logic)   │             │             │
│             └─────────────┬─────────────┘             │             │
│                           │                           │             │
│             ┌─────────────▼─────────────┐             │             │
│             │    EXECUTION ENGINE       │             │             │
│             │  (Transaction Processing) │             │             │
│             └─────────────┬─────────────┘             │             │
│                           │                           │             │
│  ┌────────────────────────┼────────────────────────┐  │             │
│  │                        │                        │  │             │
│  │  ┌─────────▼─────────┐ │ ┌─────────▼─────────┐  │  │             │
│  │  │      STORAGE      │ │ │     CONSENSUS     │  │  │             │
│  │  │   (State/Blocks)  │ │ │   (PoS + PoUW)    │  │  │             │
│  │  └───────────────────┘ │ └───────────────────┘  │  │             │
│  └────────────────────────┴────────────────────────┘  │             │
│                                                       │             │
└───────────────────────────────────────────────────────┴─────────────┘
                            │
              ┌─────────────▼─────────────┐
              │          CRYPTO           │
              │   (Primitives & Proofs)   │
              └───────────────────────────┘
```

---

## 4. Core Components

### 4.1 Node

The **Node** module is the orchestration layer that coordinates all other components.

**Responsibilities:**
- Lifecycle management (startup, shutdown, restart)
- Configuration loading and validation
- Inter-module communication
- Event routing and dispatch
- Health monitoring and metrics

**Location:** `/node/src/`

---

### 4.2 Networking

The **Networking** module handles all peer-to-peer communication.

**Responsibilities:**
- Peer discovery and connection management
- Message serialization and transport
- Block and transaction gossip
- Network topology management
- DoS protection and rate limiting

**Protocol:** libp2p-based with custom Mbongo extensions

**Location:** `/network/src/`

---

### 4.3 Crypto

The **Crypto** module provides cryptographic primitives used throughout the system.

**Responsibilities:**
- Hash functions (block hashes, transaction IDs)
- Digital signatures (Ed25519, future BLS support)
- Keypair generation and management
- Merkle tree construction
- Proof serialization

**Location:** `/crypto/src/`

---

### 4.4 Mempool

The **Mempool** manages pending transactions before they are included in blocks.

**Responsibilities:**
- Transaction ingestion and validation
- Priority ordering (fee-based, compute-weighted)
- Duplicate detection and rejection
- Expiration and eviction policies
- Transaction relay to peers

**Location:** `/runtime/src/` (integrated with runtime)

---

### 4.5 Runtime

The **Runtime** module implements the blockchain's state machine.

**Responsibilities:**
- State representation and transitions
- Block validation rules
- Transaction format enforcement
- Account and balance management
- System parameter governance

**Location:** `/runtime/src/`

---

### 4.6 Execution Engine

The **Execution Engine** processes transactions and produces state changes.

**Responsibilities:**
- Transaction execution (apply operations to state)
- Gas/compute metering
- Deterministic execution guarantees
- State root computation
- Receipt generation

**Location:** `/runtime/src/` (integrated with runtime)

---

### 4.7 PoW Module (Compute Verification)

The **PoW** module implements Proof of Useful Work for compute verification.

**Responsibilities:**
- Compute task definition and distribution
- Proof generation interface
- Proof verification logic
- Reward calculation for compute providers
- Integration with consensus layer

**Location:** `/pow/src/`

---

### 4.8 CLI

The **CLI** module provides developer and operator tooling.

**Responsibilities:**
- Node management commands
- Key generation and wallet operations
- Transaction submission
- Chain inspection and debugging
- Configuration management

**Location:** `/cli/src/`

---

## 5. Block Lifecycle

A block progresses through the following stages:

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│ Transaction │────▶│  Validation │────▶│   Mempool   │
│  Received   │     │   (Format)  │     │   (Queue)   │
└─────────────┘     └─────────────┘     └──────┬──────┘
                                               │
                    ┌──────────────────────────┘
                    │
              ┌─────▼─────┐     ┌─────────────┐     ┌─────────────┐
              │   Block   │────▶│   Runtime   │────▶│    State    │
              │  Assembly │     │  Execution  │     │ Transition  │
              └───────────┘     └─────────────┘     └──────┬──────┘
                                                          │
                    ┌─────────────────────────────────────┘
                    │
              ┌─────▼─────┐     ┌─────────────┐     ┌─────────────┐
              │   Block   │────▶│  Consensus  │────▶│   Network   │
              │  Sealing  │     │  Finality   │     │   Gossip    │
              └───────────┘     └─────────────┘     └─────────────┘
```

### Detailed Flow

1. **Transaction Received** — Node receives transaction from user or peer
2. **Validation** — Format, signature, and basic validity checks
3. **Mempool** — Valid transactions queued by priority
4. **Block Assembly** — Block producer selects transactions for inclusion
5. **Runtime Execution** — Transactions executed against current state
6. **State Transition** — New state root computed from execution results
7. **Block Sealing** — Block header finalized with state root and signatures
8. **Consensus Finality** — PoS validators attest to block validity
9. **Network Gossip** — Finalized block propagated to all peers

---

## 6. Consensus Overview

Mbongo Chain uses a **hybrid consensus model** combining Proof of Stake with Proof of Useful Work.

### Proof of Stake (PoS)

**Role:** Economic security and block finality

- Validators stake tokens to participate in consensus
- Block producers are selected based on stake weight
- Attestations from validators finalize blocks
- Slashing penalties discourage misbehavior

### Proof of Useful Work (PoUW)

**Role:** Compute verification and resource allocation

- Compute providers prove they performed useful work
- Proofs are verified on-chain with minimal overhead
- Valid proofs earn compute rewards
- Work includes AI inference, scientific computation, rendering

### Hybrid Model Benefits

| Aspect | PoS Contribution | PoUW Contribution |
|--------|------------------|-------------------|
| Security | Economic finality | Compute integrity |
| Incentives | Staking rewards | Compute rewards |
| Participation | Token holders | Compute providers |
| Attack Cost | Capital requirements | Hardware + computation |

The hybrid model ensures that both capital and computation are required to attack the network, significantly raising the cost of malicious behavior.

---

## 7. Storage & State

### State Model

Mbongo Chain maintains the following state:

| State Type | Description |
|------------|-------------|
| **Account State** | Balances, nonces, metadata for all accounts |
| **System State** | Chain parameters, validator set, epoch data |
| **Compute State** | Active compute tasks, proofs, rewards |
| **Block State** | Block headers, transaction receipts, state roots |

### Data Flow

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│ Transaction │────▶│  Execution  │────▶│    State    │
│   Input     │     │   Engine    │     │   Change    │
└─────────────┘     └─────────────┘     └──────┬──────┘
                                               │
                                        ┌──────▼──────┐
                                        │   Merkle    │
                                        │    Root     │
                                        └──────┬──────┘
                                               │
                                        ┌──────▼──────┐
                                        │   Storage   │
                                        │  (Persist)  │
                                        └─────────────┘
```

### Storage Backend

- **In-Memory:** Development and testing (current implementation)
- **RocksDB:** Production persistent storage (planned)
- **Archive Mode:** Full historical state (planned)

---

## 8. Node Responsibilities

A Mbongo Chain node performs the following duties:

### Network Operations
- **Peer Discovery** — Find and connect to other nodes
- **Block Propagation** — Receive and relay new blocks
- **Transaction Relay** — Gossip pending transactions
- **State Sync** — Download chain state from peers

### Consensus Operations
- **Block Production** — Assemble and propose new blocks (if validator)
- **Attestation** — Vote on block validity (if validator)
- **Finalization** — Track and apply finalized blocks

### Execution Operations
- **Transaction Processing** — Execute transactions in blocks
- **State Management** — Maintain and update chain state
- **Receipt Generation** — Produce execution receipts

### Storage Operations
- **Block Storage** — Persist blocks and headers
- **State Storage** — Store account and system state
- **Pruning** — Remove old data per retention policy

---

## 9. Security Model

### Safety Assumptions

Mbongo Chain's security relies on:

| Assumption | Requirement |
|------------|-------------|
| **Honest Majority** | >2/3 of staked tokens controlled by honest validators |
| **Network Synchrony** | Messages delivered within known time bounds |
| **Cryptographic Hardness** | Hash functions and signatures remain secure |

### Validator Security

- **Slashing:** Validators lose stake for provable misbehavior
  - Double signing (signing conflicting blocks)
  - Surround voting (conflicting attestations)
  - Inactivity (extended offline periods)

- **Key Management:** Validators must secure signing keys
  - Hot keys for attestations
  - Cold keys for withdrawals (planned)

### Compute Security

- **Proof Verification:** All compute proofs verified on-chain
- **Task Integrity:** Compute tasks cryptographically committed
- **Result Validation:** Multiple providers can verify results

### Attack Resistance

| Attack Vector | Mitigation |
|---------------|------------|
| 51% Attack | High capital + compute requirements |
| Long-Range Attack | Checkpointing and weak subjectivity |
| DoS | Rate limiting and peer scoring |
| Eclipse | Diverse peer connections |

---

## 10. Future Extensions

### Smart Contract Layer

*Status: Planned*

- WebAssembly (WASM) execution environment
- Deterministic gas metering
- Contract deployment and invocation
- Standard token interfaces

### Cross-Chain Interoperability

*Status: Planned*

- Light client bridges to other chains
- Asset transfer protocols
- Message passing interfaces
- Trustless verification

### Off-Chain Compute Marketplace

*Status: In Development*

- Compute task bidding and matching
- Provider reputation system
- Automated pricing and settlement
- SLA enforcement

### Additional Roadmap Items

- [ ] Hardware security module (HSM) support
- [ ] Threshold signatures for validators
- [ ] Zero-knowledge proof integration
- [ ] Sharding for horizontal scaling
- [ ] Data availability sampling

---

## Summary

Mbongo Chain is a modular, Rust-native blockchain designed for verifiable compute at scale. Its hybrid PoS + PoUW consensus provides both economic security and compute integrity, enabling a new class of decentralized applications.

For implementation details, see the [Developer Guide](developer_guide.md).

For contribution guidelines, see [CONTRIBUTING.md](../CONTRIBUTING.md).

---

**Mbongo Chain** — Compute-first blockchain infrastructure for the global future.

