— Canonical Version v1.0
# Mbongo Chain — Monorepo Architecture  
Status: Canonical  
Version: v1.0

Mbongo Chain uses a **Rust workspace monorepo** as its development architecture.  
This design ensures consistency, safety, modularity, and high performance across all components of the protocol.

The monorepo structure is a deliberate architectural choice, aligned with the requirements of a compute-native, PoS + PoUW blockchain.

---

# 1. Why a Monorepo?

A monorepo means:

- **one repository**  
- **one workspace**  
- **many crates/modules** managed together  

This approach is ideal for a blockchain because:

### 1.1 All modules progress together  
Consensus, runtime, PoUW engine, networking, cryptography and SDKs evolve **in sync**.

This prevents:

- compute engine mismatches  
- runtime version incompatibilities  
- node–SDK discrepancies  

Large multi-module blockchains like Solana, Sui, ICP, Near, and Polkadot also use monorepos for the same reason.

---

### 1.2 Developer onboarding becomes easier  
Instead of navigating multiple repos:

- everything is in one place  
- architecture becomes immediately understandable  
- contributing is faster  
- dependency management is cleaner  

A monorepo drastically lowers the barrier to entry for new contributors.

---

### 1.3 Simpler CI/CD and Testing  
A single command:



cargo build --workspace
cargo test --workspace


ensures:

- all modules compile together  
- all code paths are tested together  
- no hidden version drift  
- atomic upgrades  

Blockchain ecosystems with many tightly-coupled subsystems benefit enormously from unified testing.

---

### 1.4 Atomic versioning and safe upgrades  
A blockchain is a distributed system —  
**any incompatible update can break consensus.**

A monorepo guarantees:

- synchronized releases  
- consistent module versions  
- deterministic upgrade paths  
- zero fragmentation across nodes  

Every commit touching runtime + consensus + PoUW is tested as a whole.

---

### 1.5 Ideal for Rust Workspaces  
Rust supports monorepos natively via:

- workspaces  
- crates  
- shared dependencies  
- unified profiles  
- unified linting and formatting  

This makes the architecture clean and safe.

---

# 2. Mbongo Chain Monorepo Layout

The canonical structure:



mbongo-chain/
node/ # Core L1 node: networking, mempool, consensus
runtime/ # State machine, modules, execution engine
pouw/ # Compute engine, verification, VWP proofs
crypto/ # VRF, signatures, hashing, SMT
wallet/ # Key management, signing, CLI wallet
sdk/ # TypeScript & Rust SDKs for developers
infra/ # Deployment scripts, containers, monitoring
docs/ # Vision, specs, whitepaper, governance
tests/ # Integration tests & scenario simulations
Cargo.toml # Workspace definition


## 2.1 `node/`
Implements:

- networking (libp2p)
- block production  
- mempool  
- consensus (PoS + PoUW verification)  
- peer scoring  
- syncing  

---

## 2.2 `runtime/`
Implements:

- deterministic execution engine  
- state machine  
- E-Gas & C-Gas accounting  
- staking  
- governance  
- balance transfers  
- PoUW module  
- AIDA module (read-only economic parameters)

Everything inside `runtime/` is fully deterministic and reproducible.

---

## 2.3 `pouw/`
Implements:

- GPU/TPU/NPU compute engine  
- job scheduler  
- VWP proof generation  
- fraud proof submission  
- redundant execution  
- compute node reputation  

This crate interfaces with CUDA, ROCm, Vulkan, or CPU fallback modes.

---

## 2.4 `crypto/`
Provides:

- BLAKE3 hashing  
- Ed25519 signatures  
- VRF (Verifiable Random Functions)  
- Sparse Merkle Trees (SMT)  
- PoC (Proof-of-Compute) attestations  

All cryptographic primitives are isolated for safety.

---

## 2.5 `wallet/`
Includes:

- CLI key manager  
- secure signing  
- offline mode support  
- AIDA state queries  

---

## 2.6 `sdk/`
Contains:

- TypeScript SDK  
- Rust SDK  
- RPC abstractions  
- WASM bindings (future)  

---

## 2.7 `infra/`
Contains all infra tooling:

- docker compose for devnet  
- monitoring (Prometheus/Grafana)  
- deployment scripts  
- systemd templates  
- logging infrastructure  

---

## 2.8 `tests/`
Scenario-based tests:

- staking flow  
- PoUW fraud handling  
- AIDA economic adjustments  
- state transitions  
- execution engine tests  

A blockchain requires high-integrity integration testing — the monorepo makes this safe.

---

# 3. Advantages Over Multi-Repository Architecture

### 3.1 No version drift  
Multi-repo ecosystems suffer from:

- mismatched versions  
- incompatible APIs  
- out-of-sync SDKs  
- confusion for contributors  

Monorepo → everything moves together.

---

### 3.2 Better for Consensus & Runtime  
Consensus and runtime must evolve synchronously.  
Monorepo ensures:

- atomic updates  
- shared commit history  
- consistent versioning  

No unexpected forks.

---

### 3.3 Better for PoUW & Compute Logic  
Because PoUW and PoS interact, only a monorepo can guarantee:

- the compute engine matches the runtime logic  
- verification rules stay synchronized  
- AIDA economics integrate cleanly  

---

### 3.4 Much easier for audits  
Security firms prefer monorepos:

- one repository to scan  
- one dependency graph  
- one versioned codebase  

It reduces risk and audit cost.

---

# 4. CI/CD Model

Monorepo enables:

- one CI pipeline  
- unified linting (cargo fmt + clippy)  
- unified tests  
- workspace-level caching  
- multi-crate build optimization  

Example pipeline:



cargo fmt --all
cargo clippy --workspace
cargo test --workspace
cargo build --release --workspace


---

# 5. Summary

Mbongo Chain’s monorepo architecture provides:

- technical clarity  
- developer friendliness  
- synchronization across modules  
- safer consensus/runtime upgrades  
- unified CI/CD  
- optimal Rust workspace structure  
- robustness for PoUW compute verification  
- long-term maintainability  

A monorepo is the **only architecture** suitable for a compute-first blockchain where consensus, compute, economics, and execution must evolve together.
