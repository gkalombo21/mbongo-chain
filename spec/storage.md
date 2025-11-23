# Mbongo Chain — Storage Layer, Merkle Structures & State Snapshots

This document defines how Mbongo Chain stores, indexes, and verifies all ledger data.  
The focus is on performance, deterministic state transitions, fast proofs, and scalable snapshotting.

---

# 1. Overview

Mbongo Chain uses a **modular storage architecture** composed of:

- **State Merkle Tree** (main world state)
- **Block Store**
- **Transaction Index Store**
- **PoUW Work Store**
- **Epoch & Validator Store**
- **Snapshots (prunable)**

Storage is embedded directly in the node and optimized for 1-second block times.

---

# 2. Storage Technology

### Backend
For MVP, storage is implemented using:

- **RocksDB** (default)
- Fully replaceable through a Rust trait abstraction

### Merkle Structure
Mbongo Chain uses a **Sparse Merkle Tree (SMT)**:

- Supports fast proof verification  
- Very compact for large empty state spaces  
- Used by Ethereum 2.0, Aptos, Sui  

Each leaf is keyed by a hashed address.

---

# 3. High-Level Layout

storage/
state/
accounts/
staking/
pouw/
gov/
blocks/
tx_index/
proofs/
epochs/
snapshots/


---

# 4. State Merkle Tree Specification

The state tree stores:

| Module | Data stored |
|--------|-------------|
| Accounts | balance, nonce, stake, reputation |
| Validators | voting power, keys, slashing state |
| PoUW Tasks | assigned tasks, difficulty, metadata |
| PoUW Proofs | submitted proofs, verification state |
| Governance | proposals, votes, outcomes |
| Epoch Data | randomness seeds, validator lists |

### State Root
Each block includes a **state root**, computed after applying all transactions.

state_root = merkle_hash(full_state)


Any node can verify correctness by replaying block transactions.

---

# 5. Block Store

Blocks are stored sequentially:


Any node can verify correctness by replaying block transactions.

---

# 5. Block Store

Blocks are stored sequentially:

blocks/
0.json
1.json
2.json
…


Each block includes:

- Header  
- Body (TXs + PoUW proofs)  
- Commit signatures  
- State root  

Blocks are immutable once finalized.

---

# 6. Transaction Index

To support explorers and fast RPC:

tx_index/
<tx_hash> → { block_height, index, status }


This allows:

- Fast retrieval by hash  
- Efficient wallet queries  
- Explorer indexing  

---

# 7. PoUW Work Store

Stores all relevant PoUW data:

proofs/
<proof_id>/
input.json
output.json
proof.bin
metadata.json


This provides:

- Proof persistence  
- Re-verification  
- Audits  
- Reputation scoring  

---

# 8. Epoch & Validator Store

epochs/
current.json
next.json
randomness/


Stores:

- Active validator set  
- Epoch boundaries  
- VRF randomness  
- Committee assignments  

Used for block proposer selection and PoS mechanics.

---

# 9. State Snapshots

Snapshots allow fast bootstrapping:

### Every X blocks (default 10,000):

- The full state is serialized  
- Merkle root is saved  
- Nodes can sync from snapshot instead of genesis  

Snapshot layout:
snapshots/
snapshot_10000/
state.bin
state_root
metadata.json


Snapshots are optional but strongly recommended for mainnet.

---

# 10. Pruning Strategy

To minimize disk usage:

### **Prune old states**  
Keep only the last N states (default 256 blocks).

### **Keep all blocks**  
Historical blocks are compressed but kept for audit.

### **Prune PoUW proofs after timeout**  
Unless flagged for “long-term archival”.

---

# 11. RocksDB Column Families

A node uses multiple column families for parallel IO:

default → key-value
state_accounts → account leaves
state_staking → staking leaves
state_pouw → PoUW state
tx_index → tx → block reference
block_headers → block metadata
block_bodies → block payload
snapshots → serialized states


---

# 12. Rust Traits for Storage Abstraction

```rust
pub trait KVStore {
    fn get(&self, key: &[u8]) -> Option<Vec<u8>>;
    fn put(&self, key: &[u8], value: &[u8]);
    fn delete(&self, key: &[u8]);
}

pub trait StateStore {
    fn get_account(&self, address: &Address) -> Option<Account>;
    fn put_account(&mut self, address: &Address, account: &Account);

    fn get_staking(&self, key: &StakingKey) -> Option<StakingEntry>;
    fn put_staking(&mut self, key: &StakingKey, value: &StakingEntry);

    // … same for PoUW, governance, epochs
}

13. State Root Verification

Every node independently verifies:

Replay all transactions

Apply state transitions

Compute Merkle root

Compare with block header

If mismatch → node rejects block.

This ensures:

No malicious proposer

No tampering

No invalid PoUW rewards

14. Summary

The Storage Layer is designed to provide:

High throughput for 1-second blocks

Deterministic and verifiable state

Fast synchronization via snapshots

Efficient indexing for explorers and wallets

Full support for PoUW data persistence

This component ensures the reliability and long-term integrity of Mbongo Chain.
