# Mbongo Chain — Roadmap

This document outlines the development roadmap for Mbongo Chain, organized by quarter with specific deliverables, milestones, and research directions.

---

## Overview

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                        MBONGO CHAIN ROADMAP                                 │
└─────────────────────────────────────────────────────────────────────────────┘

  2025                                              2026
  ────                                              ────
  
  Q1              Q2              Q3              Q4              Q1
   │               │               │               │               │
   ▼               ▼               ▼               ▼               ▼
┌─────┐         ┌─────┐         ┌─────┐         ┌─────┐         ┌─────┐
│FOUND│─────────│CORE │─────────│CONS │─────────│NODE │─────────│TEST │
│ATION│         │PROTO│         │ENSUS│         │SYNC │         │ NET │
└─────┘         └─────┘         └─────┘         └─────┘         └─────┘
   │               │               │               │               │
   │               │               │               │               │
   ▼               ▼               ▼               ▼               ▼
 Runtime        Execution       PoW/PoS         Full Node       Devnet
 Network        Block Format    Finality        Guardian        Explorer
 Mempool        Signatures      Audits          Fast Sync       Telemetry
 Specs          Gossip          Fork Choice     Monitoring      Faucet
```

---

## Q1 2025 — Protocol Foundations

### Objective

Establish the foundational architecture, specifications, and development infrastructure for Mbongo Chain.

### Deliverables

| Category | Deliverable | Status |
|----------|-------------|--------|
| **Runtime** | State machine scaffold | ✓ Complete |
| **Runtime** | Transaction format definition | ✓ Complete |
| **Runtime** | Account model implementation | ✓ Complete |
| **Networking** | P2P layer scaffold | ✓ Complete |
| **Networking** | Message type definitions | ✓ Complete |
| **Mempool** | Basic transaction pool | ✓ Complete |
| **Mempool** | Priority queue structure | ✓ Complete |
| **Specifications** | Runtime spec | ✓ Complete |
| **Specifications** | Networking spec | ✓ Complete |
| **Specifications** | Consensus spec (draft) | ✓ Complete |
| **Developer Tooling** | Workspace structure | ✓ Complete |
| **Developer Tooling** | CLI scaffold | ✓ Complete |
| **Developer Tooling** | Documentation framework | ✓ Complete |
| **Repository** | Module organization | ✓ Complete |
| **Repository** | Dependency management | ✓ Complete |
| **CI Pipeline** | Build automation | ✓ Complete |
| **CI Pipeline** | Lint checks (clippy, rustfmt) | ✓ Complete |
| **CI Pipeline** | Test automation | ✓ Complete |

### Milestones

```
Week 1-4:   Runtime and crypto module scaffolds
Week 5-8:   Networking layer foundation
Week 9-12:  Mempool, specs, CI pipeline
```

---

## Q2 2025 — Core Protocol Implementation

### Objective

Implement the core protocol components required for block production and transaction processing.

### Deliverables

| Category | Deliverable | Status |
|----------|-------------|--------|
| **Execution Engine** | Transaction execution | 🔄 In Progress |
| **Execution Engine** | State transition logic | 🔄 In Progress |
| **Execution Engine** | Gas metering scaffold | ⏳ Planned |
| **Execution Engine** | Receipt generation | ⏳ Planned |
| **Block Format** | Block header structure | 🔄 In Progress |
| **Block Format** | Block body encoding | ⏳ Planned |
| **Block Format** | Merkle tree construction | ⏳ Planned |
| **Cryptography** | Hash function integration (Blake3) | 🔄 In Progress |
| **Cryptography** | Ed25519 signatures | 🔄 In Progress |
| **Cryptography** | Keypair management | ⏳ Planned |
| **Cryptography** | Merkle proof generation | ⏳ Planned |
| **Gossip Networking** | Block announcement | ⏳ Planned |
| **Gossip Networking** | Transaction propagation | ⏳ Planned |
| **Gossip Networking** | Peer discovery protocol | ⏳ Planned |
| **Mempool Validation** | Signature verification | ⏳ Planned |
| **Mempool Validation** | Nonce ordering | ⏳ Planned |
| **Mempool Validation** | Balance checks | ⏳ Planned |

### Milestones

```
Week 1-4:   Execution engine core
Week 5-8:   Block format and cryptography
Week 9-12:  Gossip networking and mempool validation
```

### Technical Goals

- [ ] Execute 1,000 TPS in isolated benchmarks
- [ ] Block production latency < 100ms
- [ ] Deterministic execution verified across platforms

---

## Q3 2025 — Consensus & Security

### Objective

Implement consensus mechanisms and conduct security audits to prepare for testnet deployment.

### Deliverables

| Category | Deliverable | Status |
|----------|-------------|--------|
| **PoUW Prototype** | Compute proof format | ⏳ Planned |
| **PoUW Prototype** | Proof verification logic | ⏳ Planned |
| **PoUW Prototype** | Receipt integration | ⏳ Planned |
| **Validator Set** | Validator registration | ⏳ Planned |
| **Validator Set** | Stake tracking | ⏳ Planned |
| **Validator Set** | Proposer selection | ⏳ Planned |
| **Fork Choice** | Chain weight calculation | ⏳ Planned |
| **Fork Choice** | Reorg handling | ⏳ Planned |
| **Fork Choice** | Best chain selection | ⏳ Planned |
| **Finality** | Checkpoint structure | ⏳ Planned |
| **Finality** | Attestation aggregation | ⏳ Planned |
| **Finality** | Finality gadget (basic) | ⏳ Planned |
| **Security** | Internal code audit | ⏳ Planned |
| **Security** | External audit (Phase 1) | ⏳ Planned |
| **Security** | Penetration testing | ⏳ Planned |
| **Security** | Bug bounty program design | ⏳ Planned |

### Milestones

```
Week 1-4:   PoUW prototype and validator set
Week 5-8:   Fork choice and finality
Week 9-12:  Security audits and hardening
```

### Security Checklist

- [ ] Cryptographic review (signatures, hashing)
- [ ] Consensus safety analysis
- [ ] DoS resistance testing
- [ ] State machine invariant verification

---

## Q4 2025 — Nodes & Sync

### Objective

Build production-ready node implementations with efficient synchronization mechanisms.

### Deliverables

| Category | Deliverable | Status |
|----------|-------------|--------|
| **Full Node** | Complete node binary | ⏳ Planned |
| **Full Node** | Configuration system | ⏳ Planned |
| **Full Node** | Logging and metrics | ⏳ Planned |
| **Full Node** | RPC API | ⏳ Planned |
| **Guardian Node** | Header-only node | ⏳ Planned |
| **Guardian Node** | Checkpoint verification | ⏳ Planned |
| **Guardian Node** | Light client serving | ⏳ Planned |
| **Sync Pipeline** | Full sync implementation | ⏳ Planned |
| **Sync Pipeline** | Header-first download | ⏳ Planned |
| **Sync Pipeline** | Parallel body fetch | ⏳ Planned |
| **Fast Sync** | State snapshot format | ⏳ Planned |
| **Fast Sync** | Snapshot download | ⏳ Planned |
| **Fast Sync** | Checkpoint-based sync | ⏳ Planned |
| **Monitoring** | Prometheus metrics | ⏳ Planned |
| **Monitoring** | Health check endpoints | ⏳ Planned |
| **Monitoring** | Grafana dashboards | ⏳ Planned |

### Milestones

```
Week 1-4:   Full node implementation
Week 5-8:   Guardian node and sync pipeline
Week 9-12:  Fast sync and monitoring
```

### Performance Targets

| Metric | Target |
|--------|--------|
| Full sync speed | > 1,000 blocks/sec |
| Fast sync time | < 1 hour (from snapshot) |
| Node startup | < 30 seconds |
| Memory usage | < 4 GB (full node) |

---

## Q1 2026 — Testnet Alpha

### Objective

Launch the first public testnet with supporting infrastructure for developers and validators.

### Deliverables

| Category | Deliverable | Status |
|----------|-------------|--------|
| **Docker Devnet** | Single-node container | ⏳ Planned |
| **Docker Devnet** | Multi-node compose | ⏳ Planned |
| **Docker Devnet** | Development scripts | ⏳ Planned |
| **Cluster Tooling** | Node deployment automation | ⏳ Planned |
| **Cluster Tooling** | Validator onboarding | ⏳ Planned |
| **Cluster Tooling** | Network bootstrap | ⏳ Planned |
| **Faucet** | Token distribution service | ⏳ Planned |
| **Faucet** | Rate limiting | ⏳ Planned |
| **Faucet** | Web interface | ⏳ Planned |
| **Explorer** | Block explorer prototype | ⏳ Planned |
| **Explorer** | Transaction search | ⏳ Planned |
| **Explorer** | Account view | ⏳ Planned |
| **Telemetry** | Node telemetry collection | ⏳ Planned |
| **Telemetry** | Network visualization | ⏳ Planned |
| **Telemetry** | Performance dashboards | ⏳ Planned |
| **PoS/PoUW** | Integrated consensus | ⏳ Planned |
| **PoS/PoUW** | Validator rewards | ⏳ Planned |
| **PoS/PoUW** | Compute proof rewards | ⏳ Planned |

### Milestones

```
Week 1-4:   Docker devnet and cluster tooling
Week 5-8:   Faucet, explorer, telemetry
Week 9-12:  PoS/PoUW integration and testnet launch
```

### Testnet Launch Criteria

- [ ] 10+ validator nodes operational
- [ ] 99.9% uptime over 7 days
- [ ] Successful chain finalization
- [ ] Public faucet operational
- [ ] Explorer accessible
- [ ] Documentation complete

---

## Open Research Topics

### ZK Light Clients

Zero-knowledge proofs for ultra-light client verification:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     ZK LIGHT CLIENT RESEARCH                                │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Goals:                                                                     │
│  • Succinct state proofs (< 1 KB)                                          │
│  • Constant-time verification                                               │
│  • Browser-compatible verification                                          │
│                                                                             │
│  Approaches:                                                                │
│  • SNARK-based state commitments                                           │
│  • Recursive proof composition                                              │
│  • Aggregated signature verification                                        │
│                                                                             │
│  Status: Early research                                                     │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### GPU Verification

Hardware-accelerated proof verification for PoUW:

| Research Area | Description |
|---------------|-------------|
| CUDA/OpenCL integration | GPU-accelerated verification |
| Batch verification | Parallel proof checking |
| Hardware attestation | Trusted execution verification |
| Power efficiency | Verification cost optimization |

### Decentralized Compute Marketplace

Economic layer for compute task distribution:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                   COMPUTE MARKETPLACE RESEARCH                              │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Components:                                                                │
│  ┌─────────────┐    ┌─────────────┐    ┌─────────────┐                     │
│  │   Task      │───▶│   Matching  │───▶│   Settle-   │                     │
│  │   Posting   │    │   Engine    │    │   ment      │                     │
│  └─────────────┘    └─────────────┘    └─────────────┘                     │
│                                                                             │
│  Research Questions:                                                        │
│  • Optimal pricing mechanisms                                               │
│  • Provider reputation systems                                              │
│  • SLA enforcement on-chain                                                 │
│  • Dispute resolution                                                       │
│                                                                             │
│  Status: Conceptual design                                                  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Additional Research Directions

| Topic | Priority | Timeline |
|-------|----------|----------|
| Sharding | Medium | 2026+ |
| Cross-chain bridges | High | Q3 2026 |
| WASM smart contracts | High | Q2 2026 |
| Data availability sampling | Medium | 2027+ |
| Post-quantum signatures | Low | Research |

---

## Community & Governance

### Developer Onboarding

Building a strong developer community:

| Initiative | Description | Timeline |
|------------|-------------|----------|
| Documentation | Comprehensive guides and API docs | Ongoing |
| Tutorials | Step-by-step development tutorials | Q2 2025 |
| Examples | Reference implementations | Q2 2025 |
| Office Hours | Weekly community calls | Q3 2025 |
| Hackathons | Developer competitions | Q1 2026 |

### Security Program

Establishing security practices and incentives:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     SECURITY PROGRAM                                        │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Bug Bounty Program                                                         │
│  ─────────────────                                                          │
│  • Critical vulnerabilities: Up to $100,000                                │
│  • High severity: Up to $25,000                                            │
│  • Medium severity: Up to $5,000                                           │
│  • Low severity: Up to $1,000                                              │
│                                                                             │
│  Responsible Disclosure                                                     │
│  ─────────────────────                                                      │
│  • Contact: security@mbongo.money                                          │
│  • 90-day disclosure window                                                 │
│  • Public acknowledgment (optional)                                         │
│                                                                             │
│  Audit Schedule                                                             │
│  ──────────────                                                             │
│  • Q3 2025: Consensus and cryptography                                     │
│  • Q4 2025: Networking and sync                                            │
│  • Q1 2026: Full protocol audit                                            │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Governance RFCs

Establishing on-chain governance:

| RFC | Topic | Status |
|-----|-------|--------|
| RFC-001 | Governance framework | Draft |
| RFC-002 | Token economics | Planned |
| RFC-003 | Validator requirements | Planned |
| RFC-004 | Upgrade process | Planned |
| RFC-005 | Treasury management | Planned |

### Governance Timeline

```
Q2 2025:  RFC process established
Q3 2025:  Community feedback collection
Q4 2025:  Governance contracts design
Q1 2026:  Testnet governance trials
Q2 2026:  Mainnet governance activation
```

---

## Summary

The Mbongo Chain roadmap focuses on building a secure, performant, and developer-friendly blockchain platform. Each quarter builds upon the previous, progressing from foundational work to a fully operational testnet.

### Key Dates

| Milestone | Target Date |
|-----------|-------------|
| Core Protocol Complete | End Q2 2025 |
| Security Audits Complete | End Q3 2025 |
| Full Node Release | End Q4 2025 |
| Testnet Alpha Launch | Q1 2026 |

### Contributing

We welcome contributions at every stage. See [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines.

---

**Mbongo Chain** — Compute-first blockchain infrastructure for the global future.
