# Mbongo Chain — Technology Stack
Status: Canonical  
Version: v1.0

This document specifies the **official technology stack** for Mbongo Chain.  
It defines the languages, frameworks, libraries, and architectural patterns used across:

- the core node
- the runtime
- the PoUW compute layer
- the cryptography layer
- SDKs and tooling
- frontend interfaces
- observability and DevOps

The stack is designed to support:

- high-performance consensus
- verifiable compute (PoUW)
- GPU/TPU/NPU offload
- long-term maintainability
- strong developer ergonomics

---

## 1. Repository Structure & Monorepo Model

Mbongo Chain uses a **single Rust workspace monorepo**:

```text
mbongo-chain/
  node/               # L1 node (consensus, mempool, networking)
  runtime/            # State machine, modules, WASM runtime
  pouw/               # PoUW compute engine (workers + verifier)
  crypto/             # Cryptographic primitives and utilities
  wallet/             # CLI wallet / key management
  sdk/                # SDKs (Rust + TypeScript bindings)
  infra/              # DevOps scripts, deployment, monitoring
  docs/               # Specifications and documentation
  tests/              # Integration and scenario tests
  Cargo.toml          # Workspace configuration


This layout ensures:

a single cargo build / cargo test for the entire project

synchronized versions across all modules

easier onboarding for contributors

coherent CI/CD pipelines

2. Core Node (L1)
2.1 Language

Rust (stable toolchain)

Rationale:

memory safety without a garbage collector

high performance (close to C/C++)

excellent ecosystem for networking, crypto, and systems code

proven in multiple modern L1s (Solana, Polkadot, Near, Aptos, Sui, ICP)

2.2 Responsibilities

The node/ crate implements:

block production and validation

PoS + PoUW hybrid consensus orchestration

mempool and transaction routing

networking (P2P)

RPC interfaces

state persistence

light client support (future)

3. Runtime & Smart Contracts
3.1 Execution Model

Primary runtime:

State machine in Rust (deterministic logic)

modular architecture (accounts, staking, PoUW, governance, etc.)

Smart contracts (optional, later phase):

executed inside WASM (WebAssembly)

authored in Rust, compiled to WASM

3.2 WASM Runtime

Deterministic execution

Sandbox isolation

Bounded resource usage (gas metering)

Suitable for:

protocol-level modules

user-defined contracts (future)

potential ZK-friendly execution environments

3.3 Language for Contracts

Rust (via no_std + WASM target)

Rationale:

strong typing and safety

mature tooling (Cargo, testing, linting)

well-supported for WASM compilation

shared language across node, runtime, and contracts

4. PoUW Compute Engine

The pouw/ crate implements the Proof-of-Useful-Work layer.

4.1 Language & Runtime

Rust as the primary implementation language

Bindings to hardware APIs:

CUDA (NVIDIA GPUs)

ROCm (AMD GPUs)

Vulkan / WebGPU (future)

Optional Python integration:

for AI/ML workloads via PyTorch / TensorFlow

orchestrated and validated by a Rust controller

4.2 Responsibilities

job scheduling, assignment, and dispatch

execution of compute tasks (AI inference, rendering, ZK proving, etc.)

result validation (redundancy, fraud proofs, consensus with PoS)

hardware capability benchmarking (PoC)

production of PoUW receipts for inclusion in blocks

4.3 Design Goals

CPU-only support for MVP (fallback)

GPU support as first-class once stable

modular job definitions (pluggable compute types)

verifiable and reproducible results

5. Networking
5.1 P2P Layer

libp2p (Rust implementation) or equivalent P2P framework

Responsibilities:

peer discovery

secure channels (noise / TLS)

block and transaction gossip

PoUW proof propagation

support for future subnets / sharded compute domains

5.2 RPC & External Interfaces

JSON-RPC over HTTP

for wallets, dApps, exchanges, indexers

WebSocket

real-time subscriptions (new blocks, PoUW receipts, events)

gRPC

high-performance service communication (compute nodes, indexers, infra services)

All RPC surface is specified in rpc_api.md.

6. Storage & State
6.1 State Database

RocksDB (primary key-value store)

or equivalent Rust bindings (e.g., rust-rocksdb)

Future alternative: ParityDB (optional)

Responsibilities:

store canonical state (accounts, balances, staking, PoUW metadata)

store block metadata and indices

store PoUW job metadata and receipts

6.2 State Tree

Sparse Merkle Tree (SMT) or similar structure for:

state root computation

efficient proofs for light clients

verifiable state transitions

7. Cryptography

Implemented primarily in the crypto/ crate.

7.1 Signatures

Ed25519

Properties:

strong security guarantees

widely used in modern blockchains

efficient verification

compatible with WASM environments

7.2 Hashing

SHA3-256 as the primary hash function

Option for BLAKE2b in specific modules if needed

Uses:

transaction IDs

block IDs

Merkle roots

PoUW proof commitments

randomness beacons (input to VRFs)

7.3 VRF

Ed25519-based VRF (or curve25519-compatible VRF library)

Uses:

leader election

randomness injection into consensus

fair task assignment in PoUW

8. SDKs & Client Libraries

Located in sdk/.

8.1 TypeScript SDK

TypeScript-first design

Bundled for Node.js and browser environments

Responsibilities:

connect to JSON-RPC and WebSocket endpoints

handle account management (client-side)

sign transactions

query state

interact with PoUW marketplace APIs

used in:

dApps

explorers

monitoring dashboards

integration scripts

8.2 Rust SDK

Rust crate for:

build tooling

custom modules / off-chain workers

Rust services interacting with the chain

Uses:

build smart contracts (WASM)

interact with node (RPC/gRPC)

integrate with external services in Rust

9. Frontend Technology

Mbongo Chain does not impose a mandatory frontend framework,
but recommends a canonical stack for:

block explorers

staking dashboards

PoUW compute dashboards

governance UIs

monitoring / admin consoles

9.1 Recommended Stack

React.js + TypeScript

Next.js (v14+, App Router)

Reasons:

de-facto standard for Web3 dApps and dashboards

built-in support for:

Server-Side Rendering (SSR)

Static Site Generation (SSG/ISR)

API routes (backend + frontend in one project)

fast data fetching and streaming

seamless integration with the TypeScript SDK

widely understood by modern frontend engineers

9.2 Usage

Typical structure for an official Mbongo frontend:

mbongo-frontend/
  app/
    explorer/
    staking/
    compute/
    governance/
    api/
  lib/
    mbongo-sdk/      # TS SDK wrapper
  components/
    charts/
    layout/
    widgets/

10. Tooling, Testing & Observability
10.1 Build & Workspace

cargo workspaces for all Rust crates

single build and test pipeline:

cargo build --workspace

cargo test --workspace

10.2 Testing

Unit tests (crate-level)

Integration tests (tests/ directory)

Scenario tests:

multiple nodes

PoS + PoUW interactions

chain reorg scenarios

PoUW fraud attempts

10.3 Observability

Metrics:

Prometheus-compatible metrics endpoints

Node health, latency, block time, PoUW throughput, etc.

Logging:

structured JSON logs

correlation IDs for tracing

Tracing:

OpenTelemetry integration (future)

11. DevOps & Deployment

High-level responsibilities in infra/:

containerization (Docker images)

deployment scripts (Kubernetes / bare metal)

configuration templates:

validators

PoUW compute nodes

indexers

RPC nodes

CI/CD:

GitHub Actions (or equivalent)

build, test, lint, format, security checks

testnet deployment pipelines

12. Design Goals

The chosen stack is guided by these goals:

Performance

Rust from top to bottom for consensus, runtime, and compute.

Safety

memory safety by design

strong typing

secure crypto primitives

AI & Compute Native

PoUW engine optimized for GPU/TPU/NPU from day one

stack ready for verifiable AI workloads.

Developer Experience

Rust workspace + TypeScript SDK

standardized frontend (React / Next.js)

clear SDKs and tooling.

Long-Term Maintainability

minimal number of primary languages (Rust + TypeScript)

monorepo structure

modular design.

This stack is the canonical baseline for Mbongo Chain and must be considered the reference for all future modules, documentation, and ecosystem tooling.