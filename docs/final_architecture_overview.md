# Mbongo Chain — Final Architecture Overview

This document provides the complete end-to-end architecture overview of Mbongo Chain, serving as the master reference for understanding the system design, component interactions, and operational flows.

---

## 1. High-Level System Overview

### System Layers

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     MBONGO CHAIN ARCHITECTURE                               │
└─────────────────────────────────────────────────────────────────────────────┘

  ┌─────────────────────────────────────────────────────────────────────────┐
  │                           USER INTERFACE                                │
  │                                                                         │
  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐    │
  │  │    CLI      │  │   Wallet    │  │   DApps     │  │   Explorer  │    │
  │  │  (mbongo)   │  │    Apps     │  │             │  │             │    │
  │  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘    │
  │         │                │                │                │           │
  └─────────┼────────────────┼────────────────┼────────────────┼───────────┘
            │                │                │                │
            └────────────────┴────────────────┴────────────────┘
                                    │
                                    ▼
  ┌─────────────────────────────────────────────────────────────────────────┐
  │                           RPC / API LAYER                               │
  │                                                                         │
  │  ┌─────────────────────────────────────────────────────────────────┐   │
  │  │  JSON-RPC  │  WebSocket  │  GraphQL (future)  │  REST (future)  │   │
  │  └─────────────────────────────────────────────────────────────────┘   │
  └────────────────────────────────────┬────────────────────────────────────┘
                                       │
                                       ▼
  ┌─────────────────────────────────────────────────────────────────────────┐
  │                           NODE LAYER                                    │
  │                                                                         │
  │  ┌─────────────────────────────────────────────────────────────────┐   │
  │  │                      NODE ORCHESTRATOR                          │   │
  │  │  • Lifecycle management   • Configuration   • Event routing     │   │
  │  └─────────────────────────────────────────────────────────────────┘   │
  │         │                │                │                │           │
  │         ▼                ▼                ▼                ▼           │
  │  ┌───────────┐    ┌───────────┐    ┌───────────┐    ┌───────────┐     │
  │  │NETWORKING │    │  MEMPOOL  │    │ CONSENSUS │    │ EXECUTION │     │
  │  │   LAYER   │◀──▶│   LAYER   │◀──▶│   LAYER   │◀──▶│   LAYER   │     │
  │  └─────┬─────┘    └─────┬─────┘    └─────┬─────┘    └─────┬─────┘     │
  │        │                │                │                │           │
  └────────┼────────────────┼────────────────┼────────────────┼───────────┘
           │                │                │                │
           ▼                ▼                ▼                ▼
  ┌─────────────────────────────────────────────────────────────────────────┐
  │                           STORAGE LAYER                                 │
  │                                                                         │
  │  ┌───────────┐    ┌───────────┐    ┌───────────┐    ┌───────────┐     │
  │  │   STATE   │    │   BLOCK   │    │ RECEIPTS  │    │CHECKPOINTS│     │
  │  │   TRIE    │    │   STORE   │    │   STORE   │    │   STORE   │     │
  │  └───────────┘    └───────────┘    └───────────┘    └───────────┘     │
  │                                                                         │
  │  ┌─────────────────────────────────────────────────────────────────┐   │
  │  │                     DATABASE (RocksDB / LevelDB)                │   │
  │  └─────────────────────────────────────────────────────────────────┘   │
  └─────────────────────────────────────────────────────────────────────────┘
```

### Node Roles

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           NODE ROLE HIERARCHY                               │
└─────────────────────────────────────────────────────────────────────────────┘

                          ┌─────────────────────┐
                          │   VALIDATOR NODE    │
                          │                     │
                          │  • Full state       │
                          │  • Block production │
                          │  • Consensus voting │
                          │  • PoUW verification│
                          └──────────┬──────────┘
                                     │
              ┌──────────────────────┼──────────────────────┐
              │                      │                      │
              ▼                      ▼                      ▼
   ┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
   │   FULL NODE     │    │  GUARDIAN NODE  │    │  COMPUTE NODE   │
   │                 │    │                 │    │                 │
   │  • Full state   │    │  • Headers only │    │  • Light state  │
   │  • Validation   │    │  • Checkpoints  │    │  • PoUW tasks   │
   │  • Tx relay     │    │  • Light serving│    │  • Proof submit │
   │  • Block relay  │    │  • Peer scoring │    │  • Task execute │
   └────────┬────────┘    └────────┬────────┘    └────────┬────────┘
            │                      │                      │
            └──────────────────────┼──────────────────────┘
                                   │
                                   ▼
                        ┌─────────────────┐
                        │   LIGHT NODE    │
                        │   (Future)      │
                        │                 │
                        │  • Headers only │
                        │  • Proof verify │
                        │  • No relay     │
                        └─────────────────┘
```

### Layer Responsibilities

| Layer | Responsibility | Components |
|-------|----------------|------------|
| **Networking** | P2P communication, peer management | Gossip, sync, discovery |
| **Mempool** | Transaction queuing, prioritization | Pool, eviction, broadcast |
| **Consensus** | Block production, finality | PoS, PoUW, checkpoints |
| **Execution** | State transitions, validation | VM, gas metering, trie |
| **Storage** | Persistence, indexing | Blocks, state, receipts |

---

## 2. Module-Level Architecture

### Module Dependency Graph

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           MODULE DEPENDENCIES                               │
└─────────────────────────────────────────────────────────────────────────────┘

                              ┌─────────────┐
                              │     cli     │
                              │  (binary)   │
                              └──────┬──────┘
                                     │
                                     ▼
                              ┌─────────────┐
                              │    node     │
                              │(orchestrator)│
                              └──────┬──────┘
                                     │
         ┌───────────────────────────┼───────────────────────────┐
         │                           │                           │
         ▼                           ▼                           ▼
  ┌─────────────┐            ┌─────────────┐            ┌─────────────┐
  │   network   │            │   runtime   │            │     pow     │
  │   (p2p)     │            │ (execution) │            │  (compute)  │
  └─────────────┘            └──────┬──────┘            └──────┬──────┘
                                    │                          │
                                    └────────────┬─────────────┘
                                                 │
                                                 ▼
                                          ┌─────────────┐
                                          │   crypto    │
                                          │(primitives) │
                                          └─────────────┘
```

### crypto/ Module

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  MODULE: crypto/                                                            │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  RESPONSIBILITIES:                                                          │
│  • Cryptographic primitives (hashing, signatures)                          │
│  • Key generation and management                                           │
│  • Merkle tree operations                                                  │
│  • Address derivation                                                      │
│                                                                             │
│  DEPENDENCIES:                                                              │
│  • None (standalone)                                                       │
│                                                                             │
│  INPUTS:                                                                    │
│  • Raw data for hashing                                                    │
│  • Messages for signing                                                    │
│  • Keys for cryptographic operations                                       │
│                                                                             │
│  OUTPUTS:                                                                   │
│  • Hash digests (Blake3, SHA256)                                           │
│  • Signatures (Ed25519, ECDSA)                                             │
│  • Public keys, addresses                                                  │
│  • Merkle proofs                                                           │
│                                                                             │
│  KEY TYPES:                                                                 │
│  • Hash, Signature, PublicKey, PrivateKey, Address                         │
│  • MerkleTree, MerkleProof                                                 │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### pow/ Module

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  MODULE: pow/                                                               │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  RESPONSIBILITIES:                                                          │
│  • Proof of Useful Work verification                                       │
│  • Compute task management                                                 │
│  • Receipt generation and validation                                       │
│  • Work scoring and reward calculation                                     │
│                                                                             │
│  DEPENDENCIES:                                                              │
│  • crypto (hashing, signatures)                                            │
│                                                                             │
│  INPUTS:                                                                    │
│  • Compute tasks (from requesters)                                         │
│  • Execution proofs (from providers)                                       │
│  • Block context (for verification)                                        │
│                                                                             │
│  OUTPUTS:                                                                   │
│  • Verified compute receipts                                               │
│  • Work scores                                                             │
│  • Provider rewards                                                        │
│                                                                             │
│  KEY TYPES:                                                                 │
│  • ComputeTask, ComputeProof, ComputeReceipt                              │
│  • WorkScore, ProviderReward                                               │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### runtime/ Module

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  MODULE: runtime/                                                           │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  RESPONSIBILITIES:                                                          │
│  • State machine implementation                                            │
│  • Transaction execution                                                   │
│  • State trie management                                                   │
│  • Gas metering and accounting                                             │
│  • Receipt generation                                                      │
│                                                                             │
│  DEPENDENCIES:                                                              │
│  • crypto (hashing, signatures, merkle)                                    │
│                                                                             │
│  INPUTS:                                                                    │
│  • Transactions (validated)                                                │
│  • Current state                                                           │
│  • Block context                                                           │
│                                                                             │
│  OUTPUTS:                                                                   │
│  • New state root                                                          │
│  • Execution receipts                                                      │
│  • Logs and events                                                         │
│  • Gas consumption                                                         │
│                                                                             │
│  KEY TYPES:                                                                 │
│  • State, Account, Transaction, Receipt                                    │
│  • BlockContext, ExecutionResult                                           │
│  • StateTrie, StorageValue                                                 │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### network/ Module

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  MODULE: network/                                                           │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  RESPONSIBILITIES:                                                          │
│  • P2P networking                                                          │
│  • Peer discovery and management                                           │
│  • Message routing and gossip                                              │
│  • Block and transaction propagation                                       │
│  • Sync protocol                                                           │
│                                                                             │
│  DEPENDENCIES:                                                              │
│  • None (standalone)                                                       │
│                                                                             │
│  INPUTS:                                                                    │
│  • Peer connections                                                        │
│  • Messages from peers                                                     │
│  • Data to broadcast                                                       │
│                                                                             │
│  OUTPUTS:                                                                   │
│  • Received messages                                                       │
│  • Peer status                                                             │
│  • Network events                                                          │
│                                                                             │
│  KEY TYPES:                                                                 │
│  • Peer, PeerInfo, PeerScore                                              │
│  • Message, MessageType                                                    │
│  • NetworkEvent, SyncStatus                                                │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### node/ Module

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  MODULE: node/                                                              │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  RESPONSIBILITIES:                                                          │
│  • Node orchestration                                                      │
│  • Component coordination                                                  │
│  • Configuration management                                                │
│  • Lifecycle management (start, stop, restart)                             │
│  • Metrics and monitoring                                                  │
│                                                                             │
│  DEPENDENCIES:                                                              │
│  • runtime (execution)                                                     │
│  • network (p2p)                                                           │
│  • crypto (primitives)                                                     │
│  • pow (compute verification)                                              │
│                                                                             │
│  INPUTS:                                                                    │
│  • Configuration files                                                     │
│  • Network events                                                          │
│  • User commands                                                           │
│                                                                             │
│  OUTPUTS:                                                                   │
│  • Running node instance                                                   │
│  • Metrics data                                                            │
│  • RPC responses                                                           │
│                                                                             │
│  KEY TYPES:                                                                 │
│  • Node, NodeConfig, NodeStatus                                            │
│  • BlockProducer, ChainSync                                                │
│  • Mempool, ValidatorSet                                                   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### cli/ Module

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  MODULE: cli/                                                               │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  RESPONSIBILITIES:                                                          │
│  • Command-line interface                                                  │
│  • User interaction                                                        │
│  • Node management commands                                                │
│  • Key management                                                          │
│                                                                             │
│  DEPENDENCIES:                                                              │
│  • node (via RPC or direct)                                                │
│                                                                             │
│  INPUTS:                                                                    │
│  • Command-line arguments                                                  │
│  • User input                                                              │
│  • Configuration files                                                     │
│                                                                             │
│  OUTPUTS:                                                                   │
│  • Console output                                                          │
│  • Exit codes                                                              │
│  • Generated files (keys, config)                                          │
│                                                                             │
│  COMMANDS:                                                                  │
│  • mbongo run     - Start node                                             │
│  • mbongo version - Show version                                           │
│  • mbongo info    - Show node info                                         │
│  • mbongo key     - Key management                                         │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 3. Block Lifecycle (End-to-End)

### Complete Transaction to Block Flow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     BLOCK LIFECYCLE: END-TO-END                             │
└─────────────────────────────────────────────────────────────────────────────┘

  USER                                                                  NETWORK
    │                                                                      │
    │  1. SUBMIT TRANSACTION                                               │
    │  ═══════════════════════                                             │
    ▼                                                                      │
  ┌─────────────────┐                                                      │
  │   Transaction   │                                                      │
  │   Submitted     │                                                      │
  │   (via RPC)     │                                                      │
  └────────┬────────┘                                                      │
           │                                                               │
           ▼                                                               │
  ┌─────────────────┐                                                      │
  │  2. RECEIVE TX  │                                                      │
  │  ══════════════ │                                                      │
  │                 │                                                      │
  │  • Decode format│                                                      │
  │  • Check size   │                                                      │
  └────────┬────────┘                                                      │
           │                                                               │
           ▼                                                               │
  ┌─────────────────┐                                                      │
  │  3. VERIFY TX   │                                                      │
  │  ════════════   │                                                      │
  │                 │                                                      │
  │  • Signature    │                                                      │
  │  • Nonce        │                                                      │
  │  • Balance      │                                                      │
  │  • Gas limit    │                                                      │
  └────────┬────────┘                                                      │
           │                                                               │
           ▼                                                               │
  ┌─────────────────┐         ┌─────────────────┐                         │
  │  4. MEMPOOL     │────────▶│  5. GOSSIP TX   │────────────────────────▶│
  │  CLASSIFICATION │         │  ═════════════  │                         │
  │  ══════════════ │         │                 │                         │
  │                 │         │  • Announce hash│                         │
  │  • Ready queue  │         │  • Send on req  │                         │
  │  • Future queue │         └─────────────────┘                         │
  │  • Prioritize   │                                                      │
  └────────┬────────┘                                                      │
           │                                                               │
           │  (wait for slot assignment)                                   │
           ▼                                                               │
  ┌─────────────────┐                                                      │
  │  6. PROPOSAL    │                                                      │
  │  ═══════════    │                                                      │
  │                 │                                                      │
  │  • Select txs   │                                                      │
  │  • Include PoUW │                                                      │
  │  • Build header │                                                      │
  │  • Sign block   │                                                      │
  └────────┬────────┘                                                      │
           │                                                               │
           ▼                                                               │
  ┌─────────────────┐         ┌─────────────────────────────────────────┐ │
  │  7. GOSSIP BLOCK│────────▶│                   PEER NETWORK          │◀┘
  │  ══════════════ │         │                                         │
  │                 │         │  ┌─────────┐  ┌─────────┐  ┌─────────┐ │
  │  • Announce     │         │  │ Peer A  │  │ Peer B  │  │ Peer C  │ │
  │  • Serve body   │         │  └────┬────┘  └────┬────┘  └────┬────┘ │
  └────────┬────────┘         │       │            │            │      │
           │                  │       ▼            ▼            ▼      │
           │                  │  ┌─────────────────────────────────┐   │
           │                  │  │        VALIDATION               │   │
           │                  │  │  • Header check                 │   │
           │                  │  │  • PoUW verify                  │   │
           │                  │  │  • Tx validation                │   │
           │                  │  └─────────────────────────────────┘   │
           │                  └─────────────────────────────────────────┘
           │
           ▼
  ┌─────────────────┐
  │  8. VALIDATION  │
  │  ═════════════  │
  │                 │
  │  • Header valid │
  │  • PoUW valid   │
  │  • Txs valid    │
  │  • Parent known │
  └────────┬────────┘
           │
           ▼
  ┌─────────────────┐
  │  9. EXECUTION   │
  │  ══════════════ │
  │                 │
  │  • Apply txs    │
  │  • Update state │
  │  • Generate     │
  │    receipts     │
  │  • Compute root │
  └────────┬────────┘
           │
           ▼
  ┌─────────────────┐
  │  10. COMMIT     │
  │  ════════════   │
  │                 │
  │  • Verify root  │
  │  • Persist block│
  │  • Update head  │
  │  • Prune mempool│
  └────────┬────────┘
           │
           ▼
  ┌─────────────────┐         ┌─────────────────────────────────────────┐
  │  11. SYNC       │────────▶│              PEER SYNC                  │
  │  ══════════     │         │                                         │
  │                 │         │  • Announce commit                      │
  │  • Broadcast    │         │  • Serve to syncing peers               │
  │  • Update peers │         │  • Update attestations                  │
  └─────────────────┘         └─────────────────────────────────────────┘
```

### Lifecycle Timing

| Stage | Typical Duration | Parallelizable |
|-------|------------------|----------------|
| Receive TX | < 1ms | Yes |
| Verify TX | 1-5ms | Yes (per tx) |
| Mempool Add | < 1ms | No |
| Gossip TX | 50-200ms | Yes |
| Proposal | 50-100ms | No |
| Gossip Block | 100-300ms | Yes |
| Validation | 50-100ms | Partial |
| Execution | 100-500ms | No |
| Commit | 10-50ms | No |
| Sync | 50-200ms | Yes |

---

## 4. Execution Pipeline Overview

### Deterministic Execution Rules

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     DETERMINISTIC EXECUTION                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  FUNDAMENTAL RULES:                                                         │
│  ══════════════════                                                         │
│                                                                             │
│  1. Same inputs → Same outputs (always)                                    │
│  2. No external state (time, random, network)                              │
│  3. Fixed arithmetic (checked, no float)                                   │
│  4. Ordered execution (strict tx sequence)                                 │
│  5. Bounded computation (gas limits)                                       │
│                                                                             │
│  FORBIDDEN OPERATIONS:                                                      │
│  ═════════════════════                                                      │
│                                                                             │
│  ❌ System time / clock                                                    │
│  ❌ Random number generation                                               │
│  ❌ File system access                                                     │
│  ❌ Network calls                                                          │
│  ❌ Thread spawning                                                        │
│  ❌ Floating point arithmetic                                              │
│  ❌ HashMap iteration (non-deterministic order)                            │
│                                                                             │
│  ALLOWED "RANDOMNESS":                                                      │
│  ═════════════════════                                                      │
│                                                                             │
│  ✓ Block hash (previous block)                                            │
│  ✓ VRF output (verifiable)                                                │
│  ✓ Transaction hash combinations                                          │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### State Transition Phases

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     STATE TRANSITION PHASES                                 │
└─────────────────────────────────────────────────────────────────────────────┘

  PRE-STATE                 INTRA-STATE                 POST-STATE
  ═════════                 ═══════════                 ══════════

  ┌─────────┐              ┌─────────────┐              ┌─────────┐
  │         │              │             │              │         │
  │ State_N │──────────────│  Execution  │──────────────│State_N+1│
  │         │              │             │              │         │
  │ root_N  │              │ • Validate  │              │ root_N+1│
  │         │              │ • Execute   │              │         │
  └─────────┘              │ • Metering  │              └─────────┘
                           │ • Logging   │
       │                   └──────┬──────┘                   │
       │                          │                          │
       ▼                          ▼                          ▼
  ┌─────────┐              ┌─────────────┐              ┌─────────┐
  │Accounts │              │  Checkpoint │              │Accounts │
  │ Nonces  │              │  + Rollback │              │ Nonces  │
  │Balances │              │             │              │Balances │
  │ Storage │              │  On failure:│              │ Storage │
  └─────────┘              │  revert to  │              └─────────┘
                           │  checkpoint │
                           └─────────────┘


  On Success:              On Revert:                 On OOG:
  ───────────              ──────────                 ────────
  • All changes committed  • Changes rolled back      • All gas consumed
  • Logs preserved         • Nonce incremented        • Changes rolled back
  • Gas refunded           • Gas consumed             • Nonce incremented
```

### Commit Invariants

| Invariant | Description | Enforcement |
|-----------|-------------|-------------|
| State Root Match | Computed root == header root | Block rejection |
| Receipts Root Match | Computed root == header root | Block rejection |
| Gas Used Match | Total gas == header gas | Block rejection |
| Supply Conservation | Pre + rewards == Post | Assertion |
| Nonce Monotonicity | Nonce always increases | State check |

### Security Checks

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     EXECUTION SECURITY CHECKS                               │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  PRE-EXECUTION:                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  □ Signature valid                                                  │   │
│  │  □ Nonce == expected                                                │   │
│  │  □ Balance >= max_cost                                              │   │
│  │  □ Gas limit >= intrinsic                                           │   │
│  │  □ Chain ID matches                                                 │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  DURING EXECUTION:                                                          │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  □ Gas consumed <= limit                                            │   │
│  │  □ Stack depth <= 1024                                              │   │
│  │  □ Memory bounds checked                                            │   │
│  │  □ Arithmetic overflow checked                                      │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  POST-EXECUTION:                                                            │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  □ State root computed correctly                                    │   │
│  │  □ Receipts root computed correctly                                 │   │
│  │  □ Refund calculated correctly                                      │   │
│  │  □ Proposer paid correctly                                          │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 5. Consensus Overview

### Hybrid PoS + PoUW Architecture

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     HYBRID CONSENSUS                                        │
└─────────────────────────────────────────────────────────────────────────────┘

                    ┌─────────────────────────────────┐
                    │         CONSENSUS LAYER         │
                    └────────────────┬────────────────┘
                                     │
              ┌──────────────────────┼──────────────────────┐
              │                      │                      │
              ▼                      ▼                      ▼
     ┌────────────────┐     ┌────────────────┐     ┌────────────────┐
     │  PROOF OF      │     │  PROOF OF      │     │   FINALITY     │
     │  STAKE (PoS)   │     │ USEFUL WORK    │     │   GADGET       │
     │                │     │   (PoUW)       │     │                │
     │ • Validator    │     │ • Compute      │     │ • Checkpoints  │
     │   selection    │     │   verification │     │ • Attestations │
     │ • Block sign   │     │ • Work scoring │     │ • 2/3 threshold│
     │ • Slashing     │     │ • Rewards      │     │                │
     └───────┬────────┘     └───────┬────────┘     └───────┬────────┘
             │                      │                      │
             └──────────────────────┼──────────────────────┘
                                    │
                                    ▼
                         ┌─────────────────────┐
                         │    CHAIN WEIGHT     │
                         │                     │
                         │  W = 0.7×Stake +    │
                         │      0.3×PoUW       │
                         └─────────────────────┘
```

### Stake-Weighted Leader Selection

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     LEADER SELECTION                                        │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  INPUTS:                                                                    │
│  • Slot number (S)                                                         │
│  • Epoch randomness (R) - from previous epoch                              │
│  • Validator set (V) - active validators with stakes                       │
│                                                                             │
│  ALGORITHM:                                                                 │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  seed = hash(S || R)                                                │   │
│  │  total_stake = sum(v.stake for v in V)                              │   │
│  │  selection_point = seed mod total_stake                             │   │
│  │                                                                     │   │
│  │  cumulative = 0                                                     │   │
│  │  for validator in V (sorted):                                       │   │
│  │      cumulative += validator.stake                                  │   │
│  │      if cumulative > selection_point:                               │   │
│  │          return validator  // Leader for slot S                     │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  PROPERTIES:                                                                │
│  • Deterministic (same inputs → same leader)                               │
│  • Proportional (higher stake → higher probability)                        │
│  • Unpredictable (randomness from previous epoch)                          │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Fork-Choice Rule

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     FORK-CHOICE RULE                                        │
└─────────────────────────────────────────────────────────────────────────────┘

  Chain Weight Calculation:
  ═════════════════════════

  Weight = Σ(block_weight) for all blocks

  block_weight = stake_attestations × 0.7 + pouw_score × 0.3

  Where:
  • stake_attestations = Σ(attester_stake) for block attestations
  • pouw_score = Σ(compute_score) for verified receipts in block


  Fork Resolution:
  ═════════════════

       Finalized                Fork A                 Fork B
          │                       │                       │
          ▼                       ▼                       ▼
       ┌─────┐                ┌─────┐                ┌─────┐
       │  F  │                │ A1  │                │ B1  │
       │     │                │W:100│                │W:150│
       └──┬──┘                └──┬──┘                └──┬──┘
          │                      │                      │
          │        ┌─────────────┴───┐                  │
          │        │                 │                  │
          ▼        ▼                 │                  ▼
       ┌─────┐  ┌─────┐              │              ┌─────┐
       │ ... │──│ A2  │              │              │ B2  │◀── Winner
       │     │  │W:80 │              │              │W:200│    (highest
       └─────┘  └─────┘              │              └─────┘     weight)
                                     │
                   Total: 180        │        Total: 350
                                     │
                              Canonical: Fork B
```

### Block Finality Model

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     FINALITY MODEL                                          │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  CHECKPOINT FINALIZATION:                                                   │
│  ════════════════════════                                                   │
│                                                                             │
│  Epoch N-1              Epoch N                Epoch N+1                   │
│  ────────               ───────                ─────────                   │
│                                                                             │
│  ┌─────┐              ┌─────┐                ┌─────┐                       │
│  │ CP  │──────────────│ CP  │────────────────│ CP  │                       │
│  │ N-1 │              │  N  │                │ N+1 │                       │
│  └──┬──┘              └──┬──┘                └──┬──┘                       │
│     │                    │                      │                          │
│     │   ◄─ Finalized     │   ◄─ Justified       │   ◄─ Pending            │
│     │      (>2/3 on N)   │      (>2/3 votes)    │      (collecting)       │
│     │                    │                      │                          │
│                                                                             │
│  FINALITY RULE:                                                             │
│  ──────────────                                                             │
│  A checkpoint is FINALIZED when:                                            │
│  1. It is JUSTIFIED (>2/3 stake attestation)                               │
│  2. Its parent checkpoint is also JUSTIFIED                                 │
│  3. Epochs are consecutive (N, N+1)                                        │
│                                                                             │
│  TIME TO FINALITY: ~2 epochs (~12-13 minutes at 1s blocks)                 │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 6. Networking Overview

### P2P Model

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     P2P NETWORKING MODEL                                    │
└─────────────────────────────────────────────────────────────────────────────┘

  BOOTSTRAP:
  ══════════

  New Node                    Seed Nodes                    Network
      │                           │                            │
      │─── Connect ──────────────▶│                            │
      │                           │                            │
      │◀── Peer List ─────────────│                            │
      │                           │                            │
      │─── Connect to Peers ──────────────────────────────────▶│
      │                           │                            │
      │◀── Status Exchange ───────────────────────────────────▶│
      │                           │                            │


  PEER SCORING:
  ═════════════

  ┌────────────────────────────────────────────────────────────────────────┐
  │  Score = 100 (initial)                                                 │
  │                                                                        │
  │  Positive:                     Negative:                               │
  │  • Valid response: +1          • Timeout: -5                           │
  │  • Fast response: +2           • Invalid data: -30                     │
  │  • New valid block: +5         • Bad signature: -100 (ban)             │
  │                                • Protocol violation: -100 (ban)        │
  │                                                                        │
  │  Action at score 0: Disconnect and ban                                 │
  └────────────────────────────────────────────────────────────────────────┘


  GOSSIP:
  ═══════

  Origin Node                 Peer A                    Peer B
      │                         │                         │
      │── NewTxHash ───────────▶│                         │
      │                         │── NewTxHash ───────────▶│
      │                         │                         │
      │◀── GetTx ───────────────│                         │
      │                         │                         │
      │── Tx ──────────────────▶│                         │
      │                         │── Tx ──────────────────▶│
```

### Message Types and Flow

| Category | Message | Direction | Purpose |
|----------|---------|-----------|---------|
| **Handshake** | Hello | Bidirectional | Protocol version |
| | Status | Bidirectional | Chain head, height |
| | Disconnect | Unidirectional | Graceful close |
| **Liveness** | Ping | Request | Check alive |
| | Pong | Response | Confirm alive |
| **Blocks** | NewBlockHash | Broadcast | Announce block |
| | GetBlock | Request | Request block |
| | Block | Response | Block data |
| **Transactions** | NewTxHash | Broadcast | Announce tx |
| | GetTx | Request | Request tx |
| | Tx | Response | Transaction data |
| **Sync** | GetHeaders | Request | Header range |
| | Headers | Response | Header data |
| | GetBodies | Request | Block bodies |
| | Bodies | Response | Body data |

### Network Topology

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     NETWORK TOPOLOGY                                        │
└─────────────────────────────────────────────────────────────────────────────┘


                         VALIDATOR MESH
              ┌────────────────────────────────┐
              │                                │
              │    ┌────┐        ┌────┐       │
              │    │ V1 │◀──────▶│ V2 │       │
              │    └─┬──┘        └──┬─┘       │
              │      │              │         │
              │      │    ┌────┐    │         │
              │      └───▶│ V3 │◀───┘         │
              │           └─┬──┘              │
              │             │                 │
              └─────────────┼─────────────────┘
                            │
              ┌─────────────┼─────────────────┐
              │             │                 │
              │             ▼                 │
              │          FULL NODES           │
              │                               │
              │  ┌────┐  ┌────┐  ┌────┐      │
              │  │ F1 │──│ F2 │──│ F3 │      │
              │  └─┬──┘  └──┬─┘  └──┬─┘      │
              │    │        │       │         │
              └────┼────────┼───────┼─────────┘
                   │        │       │
              ┌────┼────────┼───────┼─────────┐
              │    │        │       │         │
              │    ▼        ▼       ▼         │
              │       GUARDIAN NODES          │
              │                               │
              │     ┌────┐     ┌────┐        │
              │     │ G1 │     │ G2 │        │
              │     └─┬──┘     └──┬─┘        │
              │       │           │           │
              └───────┼───────────┼───────────┘
                      │           │
              ┌───────┼───────────┼───────────┐
              │       │           │           │
              │       ▼           ▼           │
              │        LIGHT CLIENTS          │
              │                               │
              │   ┌────┐ ┌────┐ ┌────┐       │
              │   │ L1 │ │ L2 │ │ L3 │       │
              │   └────┘ └────┘ └────┘       │
              │                               │
              └───────────────────────────────┘
```

---

## 7. Storage Layer Overview

### Storage Architecture

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     STORAGE ARCHITECTURE                                    │
└─────────────────────────────────────────────────────────────────────────────┘

  ┌─────────────────────────────────────────────────────────────────────────┐
  │                           STORAGE LAYER                                 │
  ├─────────────────────────────────────────────────────────────────────────┤
  │                                                                         │
  │  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐         │
  │  │   STATE TRIE    │  │   BLOCK STORE   │  │  RECEIPT STORE  │         │
  │  │                 │  │                 │  │                 │         │
  │  │  • Accounts     │  │  • Headers      │  │  • Tx receipts  │         │
  │  │  • Storage      │  │  • Bodies       │  │  • Logs         │         │
  │  │  • Code         │  │  • Indices      │  │  • Bloom        │         │
  │  └────────┬────────┘  └────────┬────────┘  └────────┬────────┘         │
  │           │                    │                    │                   │
  │           └────────────────────┼────────────────────┘                   │
  │                                │                                        │
  │  ┌─────────────────┐  ┌────────▼────────┐  ┌─────────────────┐         │
  │  │ MEMPOOL STORE   │  │  CHECKPOINT     │  │   INDEX STORE   │         │
  │  │                 │  │     STORE       │  │                 │         │
  │  │  • Pending txs  │  │                 │  │  • Height→Hash  │         │
  │  │  • Tx metadata  │  │  • Finalized    │  │  • TxHash→Block │         │
  │  │  • Priority     │  │  • Justified    │  │  • Address→Txs  │         │
  │  └────────┬────────┘  └────────┬────────┘  └────────┬────────┘         │
  │           │                    │                    │                   │
  │           └────────────────────┼────────────────────┘                   │
  │                                │                                        │
  │  ┌─────────────────────────────▼─────────────────────────────────────┐ │
  │  │                       DATABASE BACKEND                            │ │
  │  │                                                                   │ │
  │  │  ┌─────────────────┐                   ┌─────────────────┐       │ │
  │  │  │    RocksDB      │       OR          │    LevelDB      │       │ │
  │  │  │  (production)   │                   │  (development)  │       │ │
  │  │  └─────────────────┘                   └─────────────────┘       │ │
  │  └───────────────────────────────────────────────────────────────────┘ │
  │                                                                         │
  └─────────────────────────────────────────────────────────────────────────┘
```

### State Trie (Placeholder)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     STATE TRIE STRUCTURE                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Implementation: Merkle Patricia Trie (placeholder)                        │
│                                                                             │
│  Structure:                                                                 │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │                         State Root                                  │   │
│  │                             │                                       │   │
│  │              ┌──────────────┼──────────────┐                       │   │
│  │              │              │              │                       │   │
│  │           Account        Account        Account                    │   │
│  │           0x123...       0x456...       0x789...                   │   │
│  │              │                              │                       │   │
│  │          Storage                        Storage                    │   │
│  │           Trie                           Trie                      │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  Key Features:                                                              │
│  • O(log n) lookups                                                        │
│  • Merkle proofs for light clients                                         │
│  • Efficient diff computation                                              │
│  • Pruning support                                                         │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Storage Data Types

| Store | Key | Value | Size Estimate |
|-------|-----|-------|---------------|
| **State Trie** | Address | Account (88 bytes) | ~1GB per 10M accounts |
| **Block Headers** | Hash | Header (~500 bytes) | ~500MB per 1M blocks |
| **Block Bodies** | Hash | Body (variable) | ~50GB per 1M blocks |
| **Receipts** | Hash + Index | Receipt (variable) | ~20GB per 1M blocks |
| **Checkpoints** | Epoch | Checkpoint (~2KB) | ~20MB per 10K epochs |

---

## 8. Node Roles & Responsibilities

### Role Comparison

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     NODE ROLE COMPARISON                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌────────────────┬────────────┬────────────┬────────────┬────────────┐   │
│  │   Capability   │   Full     │ Validator  │  Guardian  │   Light    │   │
│  ├────────────────┼────────────┼────────────┼────────────┼────────────┤   │
│  │ Full state     │     ✓      │     ✓      │     ✗      │     ✗      │   │
│  │ Execute blocks │     ✓      │     ✓      │     ✗      │     ✗      │   │
│  │ Produce blocks │     ✗      │     ✓      │     ✗      │     ✗      │   │
│  │ Validate blocks│     ✓      │     ✓      │  Headers   │     ✗      │   │
│  │ Verify PoUW    │     ✓      │     ✓      │  Optional  │     ✗      │   │
│  │ Relay blocks   │     ✓      │     ✓      │  Headers   │     ✗      │   │
│  │ Relay txs      │     ✓      │     ✓      │     ✗      │     ✗      │   │
│  │ Serve proofs   │     ✓      │     ✓      │     ✗      │     ✗      │   │
│  │ Attestations   │     ✗      │     ✓      │     ✗      │     ✗      │   │
│  │ Checkpoints    │     ✓      │     ✓      │     ✓      │     ✓      │   │
│  └────────────────┴────────────┴────────────┴────────────┴────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Full Node

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  FULL NODE                                                                  │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  HARDWARE REQUIREMENTS:                                                     │
│  • CPU: 4+ cores                                                           │
│  • RAM: 8+ GB                                                              │
│  • Storage: 500+ GB SSD                                                    │
│  • Network: 100+ Mbps                                                      │
│                                                                             │
│  RESPONSIBILITIES:                                                          │
│  • Maintain full blockchain state                                          │
│  • Validate all blocks and transactions                                    │
│  • Relay blocks and transactions to peers                                  │
│  • Serve data to light clients and guardians                               │
│  • Participate in gossip network                                           │
│                                                                             │
│  CONSENSUS INTERACTION:                                                     │
│  • Follows fork-choice rule                                                │
│  • Tracks finalized checkpoints                                            │
│  • Does NOT produce blocks or vote                                         │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Validator Node

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  VALIDATOR NODE                                                             │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  HARDWARE REQUIREMENTS:                                                     │
│  • CPU: 8+ cores (16 recommended)                                          │
│  • RAM: 16+ GB (32 recommended)                                            │
│  • Storage: 1+ TB NVMe SSD                                                 │
│  • Network: 1+ Gbps, low latency                                           │
│  • Uptime: 99.9%+                                                          │
│                                                                             │
│  RESPONSIBILITIES:                                                          │
│  • All full node responsibilities                                          │
│  • Produce blocks when selected                                            │
│  • Sign attestations for blocks                                            │
│  • Participate in finality protocol                                        │
│  • Maintain validator key security                                         │
│                                                                             │
│  CONSENSUS INTERACTION:                                                     │
│  • Leader election participant                                             │
│  • Block producer (when assigned)                                          │
│  • Attestation signer                                                      │
│  • Subject to slashing                                                     │
│                                                                             │
│  STAKING:                                                                   │
│  • Minimum stake: TBD                                                      │
│  • Lock period: TBD                                                        │
│  • Slashing risk: Double-sign, inactivity                                  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Guardian Node

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  GUARDIAN NODE                                                              │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  HARDWARE REQUIREMENTS:                                                     │
│  • CPU: 2+ cores                                                           │
│  • RAM: 4+ GB                                                              │
│  • Storage: 50+ GB SSD                                                     │
│  • Network: 50+ Mbps                                                       │
│                                                                             │
│  RESPONSIBILITIES:                                                          │
│  • Validate block headers only                                             │
│  • Verify checkpoint signatures                                            │
│  • Relay headers to light clients                                          │
│  • Maintain peer scores                                                    │
│  • Reject invalid headers                                                  │
│                                                                             │
│  CONSENSUS INTERACTION:                                                     │
│  • Follows header chain                                                    │
│  • Verifies checkpoint attestations                                        │
│  • Does NOT execute transactions                                           │
│  • Does NOT maintain full state                                            │
│                                                                             │
│  USE CASES:                                                                 │
│  • Light client server                                                     │
│  • Bridge relay                                                            │
│  • Network monitoring                                                      │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Light Node (Future)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  LIGHT NODE (FUTURE)                                                        │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  HARDWARE REQUIREMENTS:                                                     │
│  • CPU: 1+ cores                                                           │
│  • RAM: 512+ MB                                                            │
│  • Storage: 1+ GB                                                          │
│  • Network: 10+ Mbps                                                       │
│                                                                             │
│  RESPONSIBILITIES:                                                          │
│  • Sync headers only                                                       │
│  • Request proofs on demand                                                │
│  • Submit transactions via full nodes                                      │
│  • Verify inclusion proofs                                                 │
│                                                                             │
│  CONSENSUS INTERACTION:                                                     │
│  • Trust checkpoints                                                       │
│  • Verify Merkle proofs                                                    │
│  • No participation in consensus                                           │
│                                                                             │
│  USE CASES:                                                                 │
│  • Mobile wallets                                                          │
│  • Browser applications                                                    │
│  • IoT devices                                                             │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 9. Security Model Summary

### Threat Model

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     THREAT MODEL                                            │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ADVERSARY CAPABILITIES:                                                    │
│  ────────────────────────                                                   │
│  • Control up to f < n/3 validators (safety)                               │
│  • Control network between honest nodes (bounded delay)                    │
│  • Observe all network traffic                                             │
│  • Unlimited compute resources                                             │
│  • Adaptive corruption over time                                           │
│                                                                             │
│  ADVERSARY LIMITATIONS:                                                     │
│  ─────────────────────                                                      │
│  • Cannot break cryptographic primitives                                   │
│  • Cannot control >2/3 stake instantaneously                              │
│  • Cannot prevent message delivery indefinitely                            │
│                                                                             │
│  SECURITY GOALS:                                                            │
│  ───────────────                                                            │
│  • Safety: No conflicting finalizations                                    │
│  • Liveness: Transactions eventually included                              │
│  • Censorship resistance: Cannot exclude transactions forever              │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Security Summary Table

| Property | Mechanism | Assumption |
|----------|-----------|------------|
| **Gas Safety** | Metering, limits | Correct implementation |
| **Replay Protection** | Nonce, chain ID | Per-account ordering |
| **Invalid Block** | Validation pipeline | Deterministic execution |
| **Double Spend** | Finality, checkpoints | Honest majority |
| **Long-Range Attack** | Weak subjectivity | Recent checkpoint |
| **Sybil Resistance** | Stake requirement | Economic cost |
| **Eclipse Attack** | Peer diversity | Outbound connections |
| **DoS Protection** | Rate limiting | Resource bounds |

### Memory Isolation

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     MEMORY ISOLATION                                        │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  TRANSACTION ISOLATION:                                                     │
│  • Each transaction executes in isolated context                           │
│  • No shared mutable state between transactions                            │
│  • Checkpoint/rollback for atomic execution                                │
│                                                                             │
│  CONTRACT ISOLATION:                                                        │
│  • Separate storage namespace per contract                                 │
│  • No direct memory access between contracts                               │
│  • Cross-contract calls via message passing                                │
│                                                                             │
│  NODE ISOLATION:                                                            │
│  • Execution sandboxed from host system                                    │
│  • No file/network access during execution                                 │
│  • Resource limits enforced                                                │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 10. GPU Offload Opportunities

### GPU-Friendly Workloads

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     GPU OFFLOAD ANALYSIS                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  HIGH GPU POTENTIAL:                                                        │
│  ═══════════════════                                                        │
│                                                                             │
│  ┌────────────────────────────────────────────────────────────────────┐    │
│  │  Signature Verification (batch)           Speedup: ~100x          │    │
│  │  • ECDSA recovery parallelizable                                  │    │
│  │  • Batch all block signatures                                     │    │
│  └────────────────────────────────────────────────────────────────────┘    │
│                                                                             │
│  ┌────────────────────────────────────────────────────────────────────┐    │
│  │  SNARK Proof Verification                 Speedup: ~50x           │    │
│  │  • MSM operations                                                 │    │
│  │  • Pairing computations                                           │    │
│  └────────────────────────────────────────────────────────────────────┘    │
│                                                                             │
│  ┌────────────────────────────────────────────────────────────────────┐    │
│  │  Hash Computation (bulk)                  Speedup: ~30x           │    │
│  │  • Merkle tree construction                                       │    │
│  │  • State root computation                                         │    │
│  └────────────────────────────────────────────────────────────────────┘    │
│                                                                             │
│  LOW GPU POTENTIAL (Sequential):                                            │
│  ═══════════════════════════════                                            │
│                                                                             │
│  • Transaction execution (state dependencies)                              │
│  • Nonce validation (per-account ordering)                                 │
│  • Balance updates (order-dependent)                                       │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### PoUW Compute Integration

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     PoUW COMPUTE WORKLOADS                                  │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  CURRENT (Placeholder):                                                     │
│  • Basic proof format                                                      │
│  • Signature verification                                                  │
│  • Score calculation                                                       │
│                                                                             │
│  PLANNED:                                                                   │
│  ┌────────────────────────────────────────────────────────────────────┐    │
│  │  AI/ML Workloads                                                  │    │
│  │  • Model inference verification                                   │    │
│  │  • Gradient computation                                           │    │
│  │  • Matrix operations                                              │    │
│  └────────────────────────────────────────────────────────────────────┘    │
│                                                                             │
│  ┌────────────────────────────────────────────────────────────────────┐    │
│  │  Cryptographic Workloads                                          │    │
│  │  • ZK proof generation                                            │    │
│  │  • VDF computation                                                │    │
│  │  • Randomness beacon                                              │    │
│  └────────────────────────────────────────────────────────────────────┘    │
│                                                                             │
│  ┌────────────────────────────────────────────────────────────────────┐    │
│  │  Data Workloads                                                   │    │
│  │  • Large dataset processing                                       │    │
│  │  • Merkle proof generation                                        │    │
│  │  • Data availability sampling                                     │    │
│  └────────────────────────────────────────────────────────────────────┘    │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Future WASM VM GPU Execution

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     WASM VM + GPU (FUTURE)                                  │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ARCHITECTURE:                                                              │
│                                                                             │
│  ┌─────────────────┐                                                       │
│  │   WASM Module   │                                                       │
│  │  (Smart Contract)│                                                       │
│  └────────┬────────┘                                                       │
│           │                                                                 │
│           ▼                                                                 │
│  ┌─────────────────┐     ┌─────────────────┐                               │
│  │  WASM Runtime   │────▶│  GPU Precompiles│                               │
│  │  (Wasmer/Wasmtime)│    │                 │                               │
│  └────────┬────────┘     │  • Matrix mult  │                               │
│           │              │  • Hash batch   │                               │
│           ▼              │  • Sig verify   │                               │
│  ┌─────────────────┐     └────────┬────────┘                               │
│  │   Host Functions│              │                                        │
│  │                 │◀─────────────┘                                        │
│  │  • State access │                                                       │
│  │  • Crypto ops   │                                                       │
│  └─────────────────┘                                                       │
│                                                                             │
│  DETERMINISM:                                                               │
│  • Fixed-function GPU operations only                                      │
│  • No floating point in consensus path                                     │
│  • Reproducible across GPU vendors                                         │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 11. Roadmap Links

### Core Documentation

| Document | Description | Path |
|----------|-------------|------|
| **Roadmap** | Development timeline and milestones | [roadmap.md](roadmap.md) |
| **Runtime Architecture** | Execution engine design | [runtime_architecture.md](runtime_architecture.md) |
| **Networking Overview** | P2P layer specification | [networking_overview.md](networking_overview.md) |
| **Mempool Overview** | Transaction pool design | [mempool_overview.md](mempool_overview.md) |
| **Sync Validation** | Chain synchronization | [sync_validation.md](sync_validation.md) |
| **Block Validation Pipeline** | Block processing flow | [block_validation_pipeline.md](block_validation_pipeline.md) |

### Consensus & Validation

| Document | Description | Path |
|----------|-------------|------|
| **Consensus Overview** | PoS + PoUW hybrid | [consensus_overview.md](consensus_overview.md) |
| **Consensus Validation** | Validation rules | [consensus_validation.md](consensus_validation.md) |
| **State Machine Validation** | Execution validation | [state_machine_validation.md](state_machine_validation.md) |

### Node & Infrastructure

| Document | Description | Path |
|----------|-------------|------|
| **Node Architecture** | Node internals | [node_architecture.md](node_architecture.md) |
| **Guardian Status** | Light node design | [guardian_status.md](guardian_status.md) |
| **Setup Validation** | Environment setup | [setup_validation.md](setup_validation.md) |

### Getting Started

| Document | Description | Path |
|----------|-------------|------|
| **Getting Started** | Quick start guide | [getting_started.md](getting_started.md) |
| **Developer Guide** | Full development guide | [developer_guide.md](developer_guide.md) |
| **Architecture Overview** | High-level design | [architecture_overview.md](architecture_overview.md) |

---

## Summary

Mbongo Chain is a modular, Rust-native blockchain designed for compute-first applications. The architecture combines proven consensus mechanisms (PoS) with novel compute validation (PoUW) to create a secure, performant, and useful network.

Key architectural decisions:

1. **Modular Design**: Clear separation between crypto, networking, execution, and consensus
2. **Deterministic Execution**: Bit-for-bit reproducible state transitions
3. **Hybrid Consensus**: Economic security (PoS) + compute utility (PoUW)
4. **Scalable Roles**: Full nodes, validators, guardians, and light clients
5. **GPU-Ready**: Architecture supports future GPU acceleration

---

**Mbongo Chain** — Compute-first blockchain infrastructure for the global future.

