
Mbongo Chain — Mempool Specification

Version: Final Canonical Specification

The mempool is a critical component of Mbongo Chain.
It maintains the set of pending transactions and PoUW proofs that are waiting to be included in the next block.

Because Mbongo Chain integrates AI compute workloads, the mempool includes two queues:

Transaction Mempool (TxPool)

Compute Proof Mempool (ProofPool)

This document defines structure, validation rules, ordering, prioritization, spam resistance, and block-building integration.

1. Objectives

The mempool must ensure:

deterministic ordering rules

resistance to spam and DoS

support for 1-second block times

parallel acceptance of PoUW proofs

fair ordering across users and compute nodes

compatibility with future WASM smart contracts

lightweight operation for validator nodes

redundancy for compute proofs

It must allow validators to build blocks quickly, fairly, and efficiently.

2. High-Level Architecture
         ┌──────────────────────────────┐
         │      Network (libp2p)        │
         └───────────────▲──────────────┘
                         │ gossip
         ┌───────────────┴──────────────┐
         │        Mempool Service        │
         └──────────────────────────────┘
                │                 │
                ▼                 ▼
        TxPool (transactions)   ProofPool (PoUW proofs)


Validator nodes run both pools.
Compute nodes only send proofs; they do not store the full mempool.

3. Transaction Mempool (TxPool)
3.1 Structure

The TxPool contains:

a hashmap indexed by transaction hash

one nonce-sorted queue per account

a priority heap sorted by effective gas price

Data structures:

TxPool
├── by_hash: HashMap<tx_hash, Transaction>
├── by_account: HashMap<Address, VecDeque<Transaction>>
├── priority_heap: BinaryHeap<PoolEntry>
└── size_bytes: u64


PoolEntry includes:

{
  effective_gas_price: u64,
  arrival_time: u64,
  tx_hash: H256
}

3.2 Admission Rules (TX MUST satisfy all)

A transaction is accepted into the pool ONLY if:

Signature is valid

Nonce == next_expected_nonce(account)

Balance >= fee_max

Gas limit <= max_block_gas

Size <= MAX_TX_SIZE (64 KB)

Module/method exists

Sender not banned for spam

Pool is not full

If pool is full:

transactions with lowest effective gas price are evicted first

Effective Gas Price
effective_gas_price = base_fee + priority_fee

3.3 Ordering Rules

Transactions are ordered by:

Priority (effective_gas_price descending)

Arrival time (FIFO)

Nonce order per account

This ordering is deterministic across all validators.

3.4 Local Revalidation

Before each block:

check balances again

re-check gas price >= required base_fee

drop transactions with outdated nonces

drop transactions whose senders were slashed or jailed

4. PoUW Proof Mempool (ProofPool)

PoUW submissions are large, diverse, and arrive from GPU/TPU/NPU nodes.
They must be handled separately.

4.1 Structure
ProofPool
├── by_id: HashMap<proof_id, ProofEntry>
├── by_task: HashMap<task_id, Vec<proof_id>>
├── hardware_queue: HashMap<hardware_class, Vec<proof_id>>
└── priority_heap: BinaryHeap<ProofPriority>


Each proof entry includes:

{
  task_id: u64,
  node_id: H256,
  proof_bytes: Vec<u8>,
  metadata: {...},
  quality_score: f32,
  latency_ms: u32,
  hardware_class: GPU/TPU/NPU/CPU/ASIC,
  arrival_time: u64
}

4.2 Admission Rules

A PoUW proof is accepted ONLY if:

task_id exists and is active

proof size <= MAX_POUW_PROOF_SIZE (128 MB default)

node identity signature is valid

node reputation >= minimum

proof matches compute commitment

proof not expired (deadline)

ProofPool not full

Invalid proofs trigger reputation slashing.

4.3 Prioritization Rules

Proofs are ordered by:

Quality score (higher = better)

Latency (lower = better)

Hardware class weighting

Arrival time (FIFO)

Priority formula
priority = α * quality_score − β * latency_ms + γ * hardware_weight


Values α, β, γ are governance-defined.

4.4 Deduplication

For each task:

only the best proof (highest priority) is included

others remain available for redundancy checks

malicious duplicates → reputation penalty

5. Gossip & Networking Rules
5.1 Gossip Topics

TxPool listens on:

/mbongo/txs/1.0.0

ProofPool listens on:

/mbongo/pouw/1.0.0

5.2 Gossip Controls

To prevent spam:

rate limits per peer

reputation scoring

invalid TX/Proof blacklisting

bloom filter for duplicates

size caps for proofs/timestamps

6. Block Builder Integration

When a validator builds a block:

6.1 Transaction Selection

Pull from priority heap

Respect per-account nonce order

Stop when:

block gas limit reached

block size reached

max TX count reached

6.2 PoUW Proof Selection

Validator selects:

1 winning proof per task

optional redundant proofs for verification

enforce:

max proofs per block

max bytes per block

Proof selection is deterministic:

highest priority

lowest latency

earliest arrival time

7. Eviction & Garbage Collection
Transactions are evicted when:

nonce becomes invalid

fee too low

sender balance insufficient

pool full → lowest priority evicted

expired after X blocks (default 300)

Proofs are evicted when:

task completed

proof invalidated

expired (deadline passed)

pool full → lowest priority evicted

8. Limits
Parameter	Default
Max TX size	64 KB
Max PoUW proof size	128 MB
Max TXs per block	10,000
Max proofs per block	128
Max mempool size	1 GB (configurable)
Max proofs in ProofPool	2,000
Max TXs per account	64 pending

Governance may update these parameters.

9. Security & Anti-Spam Rules

proof and TX signatures verified before propagation

stake-based rate limits for validators

reputation-based rate limits for compute nodes

misbehavior → slashing / temp ban

invalid TX or proof → peer reputation penalty

large-proof DoS prevention via size caps

10. Future Extensions

mempool-level ZK proofs of ordering

prioritized QoS lanes for enterprise compute providers

AI-driven congestion control

encrypted mempool (research phase)

parallel mempool sharding for high-load scenarios

11. Summary

The Mbongo Chain mempool:

maintains a fast, deterministic transaction pipeline

integrates PoUW proofs in parallel

is resistant to spam and large-proof attacks

supports 1-second block times

ensures fair ordering via priority-based selection

provides a robust foundation for hybrid PoS + PoUW execution

It acts as the gateway between the network and the execution engine.
