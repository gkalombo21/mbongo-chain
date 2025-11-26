# Mbongo Chain — Sync & Validation

This document describes the chain synchronization mechanisms and block validation pipeline in Mbongo Chain. It covers full sync, fast sync, state sync, and the security rules that govern peer interactions during synchronization.

---

## 1. Sync Pipeline Overview

### High-Level Flow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                        SYNC PIPELINE                                        │
└─────────────────────────────────────────────────────────────────────────────┘

  ┌─────────────┐     ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
  │    PEER     │     │   HEADER    │     │    BODY     │     │  VALIDATE   │
  │  DISCOVERY  │────▶│    SYNC     │────▶│    SYNC     │────▶│   & EXEC    │
  └─────────────┘     └─────────────┘     └─────────────┘     └─────────────┘
        │                   │                   │                   │
        ▼                   ▼                   ▼                   ▼
  ┌─────────────┐     ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
  │ Find peers  │     │ Download    │     │ Download    │     │ Execute     │
  │ with higher │     │ headers     │     │ block       │     │ txs, verify │
  │ chain head  │     │ sequentially│     │ bodies in   │     │ state root  │
  └─────────────┘     └─────────────┘     │ parallel    │     └──────┬──────┘
                                          └─────────────┘            │
                                                                     ▼
                                                               ┌─────────────┐
                                                               │   COMMIT    │
                                                               │   STATE     │
                                                               └─────────────┘
```

### Sync Modes

| Mode | Use Case | Speed | Trust Model |
|------|----------|-------|-------------|
| **Full Sync** | Archive nodes, maximum security | Slowest | Trustless |
| **Fast Sync** | Full nodes joining network | Fast | Checkpoint trust |
| **State Sync** | Validators, light clients | Fastest | Checkpoint trust |

---

## 2. Full Sync

Full sync downloads and verifies every block from genesis, providing the highest security guarantees.

### Header-First Workflow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     HEADER-FIRST SYNC                                       │
└─────────────────────────────────────────────────────────────────────────────┘

  Phase 1: Header Chain Construction
  ──────────────────────────────────

  Local Node                              Remote Peer
       │                                       │
       │─── GetHeaders(start: 0, count: 512) ─▶│
       │                                       │
       │◀── Headers [h0, h1, h2, ... h511] ────│
       │                                       │
       │    ┌────────────────────────────┐     │
       │    │ Verify header chain:       │     │
       │    │ • h[i].parent == h[i-1].hash│    │
       │    │ • h[i].height == h[i-1] + 1│    │
       │    │ • Signature valid (PoS)    │     │
       │    │ • Timestamp monotonic      │     │
       │    └────────────────────────────┘     │
       │                                       │
       │─── GetHeaders(start: 512, count: 512)▶│
       │                                       │
       │◀── Headers [h512, ... h1023] ─────────│
       │                                       │
       │    ... continue until tip ...         │
       │                                       │

  Phase 2: Body Download (Parallel)
  ─────────────────────────────────

       │                    Peer A              Peer B              Peer C
       │                      │                   │                   │
       │─ GetBodies [h0-h99] ▶│                   │                   │
       │─ GetBodies [h100-h199] ─────────────────▶│                   │
       │─ GetBodies [h200-h299] ───────────────────────────────────▶│
       │                      │                   │                   │
       │◀─ Bodies [b0-b99] ───│                   │                   │
       │◀─ Bodies [b100-b199] ────────────────────│                   │
       │◀─ Bodies [b200-b299] ─────────────────────────────────────│
       │                      │                   │                   │
```

### Block Body Download

Bodies are downloaded in parallel from multiple peers:

```rust
// Conceptual body download orchestration
struct BodyDownloader {
    /// Headers awaiting body download
    pending_headers: VecDeque<BlockHeader>,
    
    /// In-flight requests by peer
    in_flight: HashMap<PeerId, BodyRequest>,
    
    /// Completed bodies awaiting execution
    ready_bodies: BTreeMap<BlockHeight, Block>,
    
    /// Configuration
    max_concurrent_requests: usize,  // default: 64
    bodies_per_request: usize,       // default: 128
    request_timeout: Duration,       // default: 30s
}

impl BodyDownloader {
    fn schedule_requests(&mut self, peers: &[PeerId]) {
        for peer in peers {
            if self.in_flight.len() >= self.max_concurrent_requests {
                break;
            }
            if let Some(batch) = self.next_batch() {
                self.send_request(peer, batch);
            }
        }
    }
}
```

### Signature Verification

Each block header undergoes signature verification:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                   SIGNATURE VERIFICATION                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  For each block header:                                                     │
│                                                                             │
│  1. Extract proposer from slot assignment                                   │
│     ┌─────────────────────────────────────────────────────────────────┐    │
│     │ proposer = validator_set[slot % validator_count]                │    │
│     │ (weighted by stake for production implementation)               │    │
│     └─────────────────────────────────────────────────────────────────┘    │
│                                                                             │
│  2. Verify block signature                                                  │
│     ┌─────────────────────────────────────────────────────────────────┐    │
│     │ message = hash(header_without_signature)                        │    │
│     │ valid = verify(proposer.pubkey, message, header.signature)      │    │
│     └─────────────────────────────────────────────────────────────────┘    │
│                                                                             │
│  3. Verify attestations (if present)                                        │
│     ┌─────────────────────────────────────────────────────────────────┐    │
│     │ for attestation in block.attestations:                          │    │
│     │     verify(attestation.validator, block.hash, attestation.sig)  │    │
│     └─────────────────────────────────────────────────────────────────┘    │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### PoUW Receipt Verification

*Status: Placeholder for future implementation*

Compute proofs included in blocks are verified during sync:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                   PoUW RECEIPT VERIFICATION                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  For each compute_receipt in block:                                         │
│                                                                             │
│  1. Verify proof format                                                     │
│     • Task ID exists and is valid                                          │
│     • Proof structure matches expected schema                              │
│                                                                             │
│  2. Verify proof correctness                                                │
│     • Execute verification algorithm                                       │
│     • Compare result against commitment                                    │
│                                                                             │
│  3. Verify provider eligibility                                             │
│     • Provider registered at block height                                  │
│     • No duplicate submission for same task                                │
│                                                                             │
│  4. Accumulate rewards                                                      │
│     • Track verified compute for provider                                  │
│     • Update reward distribution state                                     │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Accumulated Stake Weight Verification

The canonical chain is determined by accumulated stake weight:

```rust
// Conceptual stake weight calculation
struct ChainWeight {
    /// Sum of attesting stake across all blocks
    total_attesting_stake: u128,
    
    /// Number of finalized checkpoints
    finalized_checkpoints: u64,
    
    /// Tip block height
    height: u64,
}

impl ChainWeight {
    fn compare(&self, other: &ChainWeight) -> Ordering {
        // 1. Prefer chain with more finalized checkpoints
        match self.finalized_checkpoints.cmp(&other.finalized_checkpoints) {
            Ordering::Equal => {}
            ord => return ord,
        }
        
        // 2. Prefer chain with more attesting stake
        match self.total_attesting_stake.cmp(&other.total_attesting_stake) {
            Ordering::Equal => {}
            ord => return ord,
        }
        
        // 3. Tie-breaker: prefer higher block (more work)
        self.height.cmp(&other.height)
    }
}
```

### Full Sync Validation Checklist

| Check | Stage | Failure Action |
|-------|-------|----------------|
| Header parent hash | Header sync | Reject header |
| Header signature | Header sync | Reject header, penalize peer |
| Transactions merkle root | Body validation | Reject body |
| State root after execution | Execution | Reject block, re-request |
| Receipts root | Execution | Reject block |
| PoUW proofs valid | Execution | Reject block |
| Attestations valid | Post-execution | Log warning (soft) |

---

## 3. Fast Sync (Planned)

Fast sync downloads a recent state snapshot and only executes recent blocks.

### State Snapshot Download

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     FAST SYNC WORKFLOW                                      │
└─────────────────────────────────────────────────────────────────────────────┘

  1. Find Pivot Block (recent finalized)
  ────────────────────────────────────────

       │                              Network
       │                                 │
       │─── GetFinalizedCheckpoint ─────▶│
       │                                 │
       │◀── Checkpoint(height: 1000000) ─│
       │                                 │
       │    pivot_block = 1000000        │
       │                                 │

  2. Download State Snapshot
  ────────────────────────────────────────

       │                              Peer A
       │                                 │
       │─── GetStateChunk(root, path: 0x00...) ──▶│
       │                                 │
       │◀── StateChunk(nodes: [...]) ────│
       │                                 │
       │─── GetStateChunk(root, path: 0x01...) ──▶│
       │                                 │
       │◀── StateChunk(nodes: [...]) ────│
       │                                 │
       │    ... continue for all paths ... │
       │                                 │

  3. Verify and Commit
  ────────────────────────────────────────

       │    ┌─────────────────────────────────────┐
       │    │ Reconstruct state trie              │
       │    │ Verify root == pivot_block.state_root│
       │    │ Commit to database                  │
       │    └─────────────────────────────────────┘

  4. Sync Recent Blocks (with execution)
  ────────────────────────────────────────

       │    Download and execute blocks from     │
       │    pivot_block to chain tip             │
       │                                         │
```

### Merkle Root Validation

State chunks are verified against the expected state root:

```rust
// Conceptual state verification
fn verify_state_chunk(
    chunk: &StateChunk,
    expected_root: Hash,
    path_prefix: &[u8],
) -> Result<(), SyncError> {
    // Verify merkle proof for chunk
    let computed_root = merkle_root_from_chunk(chunk, path_prefix)?;
    
    if computed_root != expected_root {
        return Err(SyncError::InvalidStateChunk {
            expected: expected_root,
            computed: computed_root,
        });
    }
    
    Ok(())
}
```

### Finality Checkpoint Verification

Fast sync relies on trusted checkpoints:

| Checkpoint Source | Trust Level | Use Case |
|-------------------|-------------|----------|
| Hard-coded | Highest | Initial sync |
| Finality gadget | High | Recent checkpoints |
| Peer consensus | Medium | Fallback |

### Skip Block Execution

Between genesis and pivot block:

- Headers are downloaded and verified (signatures only)
- Bodies are NOT downloaded
- Transactions are NOT executed
- State is obtained from snapshot

---

## 4. State Sync (Placeholder)

State sync provides the fastest synchronization by downloading only current state.

### Validator / Light Node Mode

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     STATE SYNC MODES                                        │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Validator Mode                                                             │
│  ──────────────                                                             │
│  • Download full current state                                              │
│  • Verify against finality checkpoint                                       │
│  • Begin validating immediately                                             │
│  • Backfill recent blocks in background                                     │
│                                                                             │
│  Light Client Mode                                                          │
│  ─────────────────                                                          │
│  • Download headers only                                                    │
│  • Request state proofs on demand                                           │
│  • No local state storage                                                   │
│  • Rely on full nodes for data                                              │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Snapshot Chunking

State is divided into chunks for parallel download:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     STATE CHUNKING                                          │
└─────────────────────────────────────────────────────────────────────────────┘

  State Trie (simplified)
  
                         Root
                          │
            ┌─────────────┼─────────────┐
            │             │             │
          0x0...        0x5...        0xA...
            │             │             │
       ┌────┴────┐   ┌────┴────┐   ┌────┴────┐
       │         │   │         │   │         │
     Chunk 0   Chunk 1  Chunk 2  Chunk 3  Chunk 4  Chunk 5
  
  
  Chunk Format:
  ┌─────────────────────────────────────────────────────────────────┐
  │  path_prefix: [u8; N]     // Trie path this chunk covers       │
  │  nodes: Vec<TrieNode>     // All nodes under this prefix       │
  │  proof: MerkleProof       // Proof connecting to root          │
  │  version: u64             // State version (block height)      │
  └─────────────────────────────────────────────────────────────────┘
```

### Versioned State

State snapshots are versioned by block height:

| Field | Description |
|-------|-------------|
| `block_height` | Block at which snapshot was taken |
| `block_hash` | Hash of the snapshot block |
| `state_root` | Merkle root of state trie |
| `chunk_count` | Number of chunks in snapshot |
| `total_size` | Total snapshot size in bytes |

---

## 5. Sync Failure Modes

### Failure Categories

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     SYNC FAILURE MODES                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  TIMEOUT                                                            │   │
│  │                                                                     │   │
│  │  Trigger: Peer doesn't respond within deadline                      │   │
│  │                                                                     │   │
│  │  Actions:                                                           │   │
│  │  • Mark request as failed                                           │   │
│  │  • Decrease peer score                                              │   │
│  │  • Re-assign request to different peer                              │   │
│  │  • If repeated: disconnect peer                                     │   │
│  │                                                                     │   │
│  │  Timeouts:                                                          │   │
│  │  • Header request: 10s                                              │   │
│  │  • Body request: 30s                                                │   │
│  │  • State chunk: 60s                                                 │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  INVALID BLOCK                                                      │   │
│  │                                                                     │   │
│  │  Trigger: Block fails validation                                    │   │
│  │                                                                     │   │
│  │  Sub-types:                                                         │   │
│  │  • Invalid signature → ban peer immediately                         │   │
│  │  • Invalid state root → re-request, then ban                        │   │
│  │  • Malformed data → re-request, decrease score                      │   │
│  │                                                                     │   │
│  │  Actions:                                                           │   │
│  │  • Discard block                                                    │   │
│  │  • Penalize sending peer                                            │   │
│  │  • Request from alternative peer                                    │   │
│  │  • Log for analysis                                                 │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  FORK SWITCH                                                        │   │
│  │                                                                     │   │
│  │  Trigger: Discovered chain with higher weight than current         │   │
│  │                                                                     │   │
│  │  Actions:                                                           │   │
│  │  • Identify common ancestor                                         │   │
│  │  • Revert to common ancestor                                        │   │
│  │  • Re-execute blocks on new fork                                    │   │
│  │  • Update canonical chain pointer                                   │   │
│  │                                                                     │   │
│  │  Constraints:                                                       │   │
│  │  • Cannot revert past finalized checkpoint                          │   │
│  │  • Maximum reorg depth: 64 blocks (configurable)                    │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  CHECKPOINT MISMATCH                                                │   │
│  │                                                                     │   │
│  │  Trigger: Peer's finalized checkpoint conflicts with ours          │   │
│  │                                                                     │   │
│  │  Implications:                                                      │   │
│  │  • Either we or peer is on wrong chain                              │   │
│  │  • Possible network partition                                       │   │
│  │  • Possible attack                                                  │   │
│  │                                                                     │   │
│  │  Actions:                                                           │   │
│  │  • Disconnect from peer                                             │   │
│  │  • Alert operator                                                   │   │
│  │  • Seek additional peers for consensus                              │   │
│  │  • If majority disagrees: halt and require manual intervention      │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Recovery Procedures

| Failure | Recovery | Max Retries |
|---------|----------|-------------|
| Timeout | Retry with different peer | 3 |
| Invalid data | Re-request, penalize | 2 |
| Fork detected | Reorg if within bounds | 1 |
| Checkpoint conflict | Manual intervention | 0 |

---

## 6. Sync Security Rules

### Finality Checkpoints

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     FINALITY CHECKPOINT RULES                               │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Hard-Coded Checkpoints                                                     │
│  ──────────────────────                                                     │
│  • Embedded in client binary                                                │
│  • Used for initial sync                                                    │
│  • Cannot be reverted                                                       │
│                                                                             │
│  Example:                                                                   │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  static CHECKPOINTS: &[(u64, Hash)] = &[                           │   │
│  │      (0, genesis_hash),                                            │   │
│  │      (100_000, hash!("0x1234...")),                                │   │
│  │      (500_000, hash!("0xabcd...")),                                │   │
│  │  ];                                                                │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  Dynamic Checkpoints (Finality Gadget)                                      │
│  ─────────────────────────────────────                                      │
│  • Created by consensus (>2/3 attestations)                                 │
│  • Stored in local database                                                 │
│  • Pruning boundary for state                                               │
│                                                                             │
│  Weak Subjectivity                                                          │
│  ──────────────────                                                         │
│  • New nodes must start from recent checkpoint                              │
│  • Prevents long-range attacks                                              │
│  • Checkpoint age limit: 2 weeks (configurable)                             │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Peer Ban Score

Peers accumulate penalty scores for misbehavior:

```rust
// Conceptual peer scoring
struct PeerScore {
    score: i32,  // Starts at 100, banned at 0
}

impl PeerScore {
    const INITIAL: i32 = 100;
    const BAN_THRESHOLD: i32 = 0;
    
    // Penalties
    const TIMEOUT_PENALTY: i32 = -5;
    const INVALID_HEADER_PENALTY: i32 = -50;
    const INVALID_BODY_PENALTY: i32 = -30;
    const INVALID_SIGNATURE_PENALTY: i32 = -100;  // Instant ban
    const CHECKPOINT_MISMATCH_PENALTY: i32 = -100;  // Instant ban
    
    // Rewards
    const VALID_RESPONSE_REWARD: i32 = 1;
    const FAST_RESPONSE_REWARD: i32 = 2;
    
    fn apply_penalty(&mut self, penalty: i32) -> bool {
        self.score += penalty;
        self.score <= Self::BAN_THRESHOLD
    }
}
```

### Penalty Schedule

| Offense | Penalty | Ban After |
|---------|---------|-----------|
| Request timeout | -5 | 20 timeouts |
| Empty response | -10 | 10 occurrences |
| Invalid header | -50 | 2 occurrences |
| Invalid body | -30 | 3 occurrences |
| Invalid signature | -100 | Immediate |
| Checkpoint mismatch | -100 | Immediate |
| Protocol violation | -100 | Immediate |

### DoS Protection

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     DoS PROTECTION                                          │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Rate Limiting                                                              │
│  ─────────────                                                              │
│  • Max header requests per peer: 10/second                                  │
│  • Max body requests per peer: 5/second                                     │
│  • Max state requests per peer: 2/second                                    │
│                                                                             │
│  Resource Limits                                                            │
│  ───────────────                                                            │
│  • Max concurrent sync peers: 8                                             │
│  • Max in-flight requests: 64                                               │
│  • Max pending headers: 10,000                                              │
│  • Max pending bodies: 1,000                                                │
│                                                                             │
│  Backpressure                                                               │
│  ────────────                                                               │
│  • Slow down requests if execution falls behind                             │
│  • Pause sync if disk I/O is saturated                                      │
│  • Limit memory usage for buffered data                                     │
│                                                                             │
│  Connection Diversity                                                       │
│  ────────────────────                                                       │
│  • Minimum outbound connections: 4                                          │
│  • Prefer geographically diverse peers                                      │
│  • Rotate slow peers periodically                                           │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 7. Implementation Notes

### Sync State Machine

```rust
// Conceptual sync state machine
enum SyncState {
    /// Discovering peers with higher chain
    PeerDiscovery,
    
    /// Downloading and verifying headers
    HeaderSync {
        target_height: u64,
        current_height: u64,
    },
    
    /// Downloading block bodies
    BodySync {
        target_height: u64,
        downloaded: u64,
        executed: u64,
    },
    
    /// Downloading state snapshot (fast sync)
    StateSync {
        pivot_block: u64,
        chunks_total: u32,
        chunks_received: u32,
    },
    
    /// Fully synced, following chain tip
    Synced,
    
    /// Sync failed, awaiting retry
    Failed {
        reason: SyncError,
        retry_at: Instant,
    },
}
```

### Progress Tracking

| Metric | Description |
|--------|-------------|
| `sync_height` | Current sync progress (block height) |
| `chain_tip` | Known highest block from peers |
| `sync_speed` | Blocks per second |
| `eta` | Estimated time to completion |
| `peer_count` | Active sync peers |

---

## Summary

Mbongo Chain provides multiple sync modes to accommodate different node types and security requirements. Full sync offers maximum security through complete verification, while fast sync and state sync enable rapid onboarding. The security rules protect against malicious peers and ensure nodes converge on the correct chain.

For networking details, see [Networking Overview](networking_overview.md).

For consensus details, see [Consensus Overview](consensus_overview.md).

---

**Mbongo Chain** — Compute-first blockchain infrastructure for the global future.

