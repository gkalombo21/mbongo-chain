Mbongo Chain — Core Architecture Specification

Version finale, consolidée, sans doublons.

1. Overview

Mbongo Chain is a next-generation, AI-native blockchain designed for verifiable compute, high-throughput execution, and a hybrid consensus model (PoS + PoUW).
The architecture emphasizes:

High-speed block production (~1s)

Modular Rust implementation

Efficient state transitions

AI/GPU/TPU-compatible compute marketplace

Extensible execution environment

Deterministic, secure state machine

This document defines the core architecture of the system and how each component interacts.

2. High-Level Architecture Diagram (Conceptual)
 ┌───────────────────────────┐
 │        Networking         │ ← libp2p-based P2P
 └───────────────┬──────────┘
                 │
 ┌───────────────▼────────────────┐
 │        Consensus Engine        │ ← PoS + PoUW hybrid
 └───────────────┬────────────────┘
                 │
 ┌───────────────▼────────────────┐
 │     Block Builder / Executor    │
 └───────────────┬────────────────┘
                 │
 ┌───────────────▼────────────────┐
 │        Runtime System          │ ← Modular WASM-based execution
 └───────────────┬────────────────┘
                 │
 ┌───────────────▼────────────────┐
 │        State Machine           │ ← deterministic state transitions
 └───────────────┬────────────────┘
                 │
 ┌───────────────▼────────────────┐
 │          Ledger / Storage      │ ← Merkle/SMT, database, prunable
 └────────────────────────────────┘

3. Node Architecture
3.1 Node Types

Mbongo Chain supports multiple node roles:

Validator Node

Participates in PoS consensus

Proposes and votes on blocks

Maintains full ledger

Verifies PoUW proofs

Compute Node (PoUW Worker)

Performs useful compute (AI, render, math, ZK)

Produces Validatable Work Proofs (VWP)

Does not need full chain state

Full Node

Maintains full chain state

Participates in networking

Verifies all block data

Light Node

Verifies headers

Fetches state on demand

4. Consensus Layer (Structural Integration)

(Full rules defined in consensus_spec.md — here we document architecture-level integration.)

Mbongo Chain uses a hybrid consensus engine:

4.1 PoS Integration

Slot leader election via VRF

Committee-based attestation

Block proposal pipeline interacts with runtime and state machine

Provides security, finality, and fork-choice logic

4.2 PoUW Integration

Work Tasks dispatched through Networking

Compute nodes return Validatable Work Proofs

PoUW proofs inserted into block builder

Rewards split 50/50 with PoS validators

5. Block Lifecycle
5.1 Steps from Slot Start to Finalization

Slot begins (1s window)

VRF selects PoS proposer

Proposer builds a block:

transactions

PoUW proof bundle

state root

metadata

Validator committee signs block

Networking gossips block

State machine applies transitions

Block reaches economic finality (2–4s)

6. Runtime System Architecture

(Full details in runtime.md — here is the architectural overview.)

The runtime is:

Modular

WASM-capable

Versioned

Upgradable via governance

Responsibilities:

Transaction format, decoding, validation

Execution gas metering

Module routing (accounts, compute, contracts…)

Invoking state transitions

Exposing host functions to smart modules

Maintaining deterministic execution

7. State Machine Integration

The state machine is:

deterministic

transition-driven

cryptographically committed

Responsibilities:

Apply transactions in order

Update Merkle/SMT state

Manage account balances & nonces

Apply PoUW result commits

Produce state roots

Expose commit & rollback

Provide proofs for light clients

8. Ledger & Storage Layer

Mbongo Chain uses:

SMT (Sparse Merkle Tree) for state

Append-only block storage

Prunable historical data

Fast key-value store for runtime access

Optional:

Archive nodes

Storage shards

Distributed availability in later versions

9. Networking Layer (Structural Overview)

(Full rules in networking.md)

Built on libp2p, including:

GossipSub for blocks & transactions

WorkTask propagation

Peer scoring (spam/DoS protection)

Sync protocol (header-first)

WorkProof distribution

Node discovery via mDNS + DHT

10. Compute Marketplace (Integration Overview)

(Details in compute_marketplace.md)

Architecture elements:

Work dispatcher

Proof verifier

Compute reputation system

Market incentives

Task categories (AI, 3D, ZK, math…)

PoUW block proof integration

11. Module System

The system is split into Rust crates inside a monorepo:

/crates
   /consensus
   /networking
   /runtime
   /state
   /ledger
   /pouw
   /accounts
   /mempool
   /cli
   /node


Each module is:

independently testable

versioned

documented

compatible with fuzzing & simulation

12. Mempool Architecture

Handles:

tx validation

tx prioritization

tx gossip

PoUW proof mempool

anti-spam scoring

gas pricing for workloads

13. Security Architecture

Includes:

Slashing (PoS)

Redundant verification (PoUW)

Anti-spoofing workloads

Network-level peer scoring

Hash-based integrity on state roots

Deterministic runtime

Replay protection via nonces

14. Upgrade Path
Phase 1

Validator network

PoUW MVP (AI inference + math tasks)

Basic runtime

Phase 2

WASM smart contracts

3D rendering work types

ZK verification pipeline

Phase 3

Sharding (compute-centric)

Storage expansion

Advanced marketplace features

15. Summary

This architecture ensures:

High throughput

Deterministic state execution

AI-aligned workloads

Strong security via hybrid consensus

Modular maintainability

Long-term scalability

Mbongo Chain is built for global-scale compute & trust.
