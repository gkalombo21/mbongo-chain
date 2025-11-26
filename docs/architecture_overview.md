# Mbongo Chain — Architecture Overview

> **Document Type:** Master Index  
> **Last Updated:** November 2025  
> **Status:** Canonical Reference

---

## 1. What This Document Is

This document is the **high-level architecture summary** of Mbongo Chain. It provides a concise overview of all major subsystems and serves as the central navigation hub for the technical documentation.

For detailed specifications, follow the links in the [Documentation Map](#6-documentation-map) below.

---

## 2. High-Level Diagram

```
┌─────────────────────────────────────────────────────────────────────────────────┐
│                          MBONGO CHAIN ARCHITECTURE                              │
├─────────────────────────────────────────────────────────────────────────────────┤
│                                                                                 │
│   ┌───────────────────────────────────────────────────────────────────────┐    │
│   │                          NODE LAYER                                   │    │
│   │   ┌───────────┐   ┌───────────┐   ┌───────────┐   ┌───────────┐      │    │
│   │   │   FULL    │   │ VALIDATOR │   │ GUARDIAN  │   │   LIGHT   │      │    │
│   │   │   NODE    │   │   NODE    │   │   NODE    │   │   NODE    │      │    │
│   │   └───────────┘   └───────────┘   └───────────┘   └───────────┘      │    │
│   └───────────────────────────────────┬───────────────────────────────────┘    │
│                                       │                                        │
│   ┌───────────────────────────────────▼───────────────────────────────────┐    │
│   │                        CONSENSUS LAYER                                │    │
│   │              ┌─────────────┐     ┌─────────────┐                      │    │
│   │              │     PoS     │◀───▶│    PoUW    │                      │    │
│   │              │   Engine    │     │   Engine   │                      │    │
│   │              └─────────────┘     └─────────────┘                      │    │
│   └───────────────────────────────────┬───────────────────────────────────┘    │
│                                       │                                        │
│   ┌───────────────┬───────────────────┼───────────────────┬───────────────┐    │
│   │               │                   │                   │               │    │
│   │   ┌───────────▼───────────┐   ┌───▼───────────┐   ┌───▼───────────┐  │    │
│   │   │   EXECUTION ENGINE    │   │    MEMPOOL    │   │COMPUTE ENGINE │  │    │
│   │   │   ┌───────────────┐   │   │               │   │  (GPU/PoUW)   │  │    │
│   │   │   │ STATE MACHINE │   │   │  Pending Txs  │   │               │  │    │
│   │   │   │   S' = F(S,T) │   │   │  Priority Q   │   │  Task Queue   │  │    │
│   │   │   └───────────────┘   │   │               │   │  Verification │  │    │
│   │   └───────────────────────┘   └───────────────┘   └───────────────┘  │    │
│   │                                                                       │    │
│   └───────────────────────────────────┬───────────────────────────────────┘    │
│                                       │                                        │
│   ┌───────────────────────────────────▼───────────────────────────────────┐    │
│   │                         STORAGE LAYER                                 │    │
│   │   ┌─────────┐   ┌─────────┐   ┌─────────┐   ┌─────────────────────┐  │    │
│   │   │ Blocks  │   │  State  │   │ Indexes │   │    Checkpoints      │  │    │
│   │   │         │   │  Trie   │   │         │   │                     │  │    │
│   │   └─────────┘   └─────────┘   └─────────┘   └─────────────────────┘  │    │
│   └───────────────────────────────────┬───────────────────────────────────┘    │
│                                       │                                        │
│   ┌───────────────────────────────────▼───────────────────────────────────┐    │
│   │                        NETWORKING LAYER                               │    │
│   │        libp2p  ──▶  Gossip  ──▶  Sync  ──▶  Peer Discovery           │    │
│   └───────────────────────────────────────────────────────────────────────┘    │
│                                                                                 │
└─────────────────────────────────────────────────────────────────────────────────┘
```

---

## 3. Core Architecture Layers

### Networking Layer
Handles all peer-to-peer communication using libp2p. Provides gossip-based message propagation, DHT peer discovery, and request/response protocols for chain synchronization.

### Mempool Layer
Stages pending transactions before block inclusion. Validates signatures and nonces, prioritizes by fee, and evicts low-priority transactions under memory pressure.

### Consensus Layer (PoS + PoUW)
Implements hybrid Proof-of-Stake and Proof-of-Useful-Work consensus. Leader election combines stake weight, PoUW scores, and VRF randomness. Fork choice follows the heaviest chain rule.

### Execution Layer
Processes transactions through a deterministic state transition function. The runtime validates inputs, executes state changes, and generates receipts. All execution is isolated and reproducible.

### Compute Layer (GPU/PoUW)
Coordinates useful GPU workloads for Proof-of-Useful-Work. Manages task assignment, result verification, and reward distribution. Compute receipts are integrated into block proposals.

### Storage Layer
Persists blockchain data using a key-value store (RocksDB). Maintains block store, Merkle state trie, indexes, and periodic checkpoints for fast synchronization.

### Node Architecture
Supports multiple node types: Full Nodes (validate and relay), Validator Nodes (propose blocks), Guardian Nodes (coordinate GPU compute), and Light Nodes (header verification only).

---

## 4. Block Flow Summary

| Step | Phase | Description |
|------|-------|-------------|
| **1** | Receive | Transaction arrives via network or RPC |
| **2** | Validate | Mempool verifies signature, nonce, balance |
| **3** | Queue | Transaction added to priority queue |
| **4** | Select | Block proposer selects transactions |
| **5** | Propose | Leader creates block with PoUW receipts |
| **6** | Consensus | Validators verify and attest to block |
| **7** | Execute | Runtime processes transactions, updates state |
| **8** | Commit | Block and state persisted, broadcast to network |

---

## 5. Security Principles

### Determinism & Validity
All state transitions are fully deterministic—given the same inputs, every node computes identical outputs. Block validity requires correct state roots, valid signatures, and proper proposer eligibility. Invalid blocks are rejected and may result in proposer slashing.

### Adversarial Model
The protocol assumes an adversarial network environment where up to 1/3 of stake may be controlled by malicious actors. Safety is maintained through cryptographic verification, economic penalties for misbehavior, and finality checkpoints that prevent long-range attacks.

### Compute & Network Security
GPU compute receipts undergo verification through replicated execution or probabilistic sampling. Invalid results trigger slashing. The networking layer employs peer scoring, connection limits, and rate limiting to defend against eclipse attacks, Sybil attacks, and denial-of-service attempts.

---

## 6. Documentation Map

| Document | Description |
|----------|-------------|
| `getting_started.md` | Quick start guide for new developers |
| `developer_introduction.md` | Development environment and workflow |
| `architecture_overview.md` | This document — master index |
| `full_system_overview.md` | Comprehensive system architecture |
| `architecture_master_overview.md` | Detailed layer-by-layer breakdown |
| `consensus_overview.md` | PoS + PoUW consensus mechanism |
| `execution_engine_overview.md` | Runtime and state machine |
| `compute_engine_overview.md` | GPU coordination and PoUW |
| `block_validation_pipeline.md` | Block verification process |
| `mempool_overview.md` | Transaction pool management |
| `network_overview.md` | P2P networking protocols |
| `node_architecture.md` | Node types and roles |
| `state_machine_validation.md` | State transition rules |
| `storage_overview.md` | Persistence and indexing |
| `security_model.md` | Threat analysis and mitigations |
| `final_architecture_overview.md` | Complete architecture reference |

---

## 7. Reading Order for Developers

**New to Mbongo Chain? Follow this order:**

| Order | Document | Purpose |
|-------|----------|---------|
| 1 | `getting_started.md` | Set up your environment |
| 2 | `developer_introduction.md` | Understand development workflow |
| 3 | `architecture_overview.md` | Get the big picture (this doc) |
| 4 | `full_system_overview.md` | Deep dive into architecture |
| 5 | `execution_engine_overview.md` | Understand runtime and state |
| 6 | `consensus_overview.md` | Learn consensus mechanism |
| 7 | `compute_engine_overview.md` | Explore GPU/PoUW system |
| 8 | `network_overview.md` | Study P2P protocols |
| 9 | `storage_overview.md` | Review persistence layer |
| 10 | `security_model.md` | Understand security design |

**For specific tasks:**
- *Implementing consensus changes* → Start with `consensus_overview.md`
- *Working on runtime* → Start with `execution_engine_overview.md`
- *GPU provider development* → Start with `compute_engine_overview.md`
- *Node operations* → Start with `node_architecture.md`

---

## 8. Future Extensions

### Parallel Execution
Transaction dependency analysis will enable parallel execution of non-conflicting transactions. Target: 10x throughput improvement with deterministic merge ordering.

### ZK Integration
Zero-knowledge proofs will provide succinct state transition verification. Light clients will sync via ZK proofs instead of full execution. Cross-chain bridges will use ZK for trust-minimized verification.

### GPU Marketplace
A global compute marketplace where external chains and applications can submit tasks. Features include price discovery, reputation systems, and cross-chain settlement.

### VM Design
A deterministic WASM execution engine for smart contracts. Includes instruction-level gas metering, sandboxed execution, and native PoUW integration for compute-heavy contracts.

---

## Quick Links

| Resource | Location |
|----------|----------|
| **Source Code** | `crypto/`, `network/`, `runtime/`, `pow/`, `node/`, `cli/` |
| **Tests** | `*/tests/` in each module |
| **Scripts** | `scripts/check.ps1`, `scripts/check.sh` |
| **Config** | `Cargo.toml`, `rustfmt.toml`, `.clippy.toml` |

---

*This document is the master index for Mbongo Chain technical documentation. Last updated November 2025.*
