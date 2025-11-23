Mbongo Chain — State Machine Specification (Final, Canonical Version)
1. Overview

The State Machine defines how the global state of Mbongo Chain evolves from block to block.
It provides deterministic, cryptographically verifiable, and reproducible state transitions necessary for hybrid PoS + PoUW consensus.

Key properties:

Deterministic execution

Replayable transitions

Cryptographically committed root (MPT/SMT)

Modular + versioned structure

Optimized for 1-second block times

AI-native (PoUW + AI model integration)

2. Global State Architecture

Mbongo Chain uses a Merkle-Patricia Trie (MPT) or Sparse Merkle Tree (SMT) depending on module needs.
The state_root is included in every block header.

Global state layout:
state_root
│
├── accounts_root
├── staking_root
├── pouw_root
├── ai_registry_root
├── storage_root
└── system_root

2.1 Accounts Tree

Each account contains:

Field	Type	Description
balance	u128	Spendable MBO
nonce	u64	Replay protection
stake	u128	Staked MBO (PoS)
storage_root	H256	Contract storage (future WASM)
code_hash	Option<H256>	WASM contract code hash

Addresses are 32 bytes.

2.2 Staking / Validator State

Stores:

active validators

delegations

voting power

slashing flags

epoch committees

reward accumulators

VRF public keys

Updated at epoch boundaries.

2.3 PoUW (Proof-of-Useful-Work) State

Each compute node has a dedicated entry:

node_id: H256
hardware_class: enum(GPU/TPU/NPU/CPU/ASIC)
success_rate: f32
avg_latency_ms: u32
completed_tasks: u64
invalid_proofs: u32
reputation_score: i64


This governs:

scheduling probability

reward weighting

slashing penalties

2.4 AI Model Registry

A governance-controlled registry of approved AI models.

Each entry contains:

model_id
sha256_checksum
size_bytes
version
approved_by_governance
model_type (onnx / quantized / custom)
gas_multiplier


Purpose:

ensure deterministic inference

allow PoUW workloads referencing approved models

enable future on-chain inference verification

2.5 Contract Storage (WASM Phase)

Contracts have their own sub-trie:

hash(contract_address, storage_key) → value


Properties:

deterministic key-value structure

gas-charged writes

prefix iteration supported

proofs compatible with light clients

2.6 System State

Contains chain-level metadata:

epoch_number
slot_number
validator_set
randomness_seed
governance_flags
halving_epoch
fees_collected
protocol_version

3. State Transition Function (STF)

The STF is the deterministic function:

new_state = ApplyBlock(old_state, block)

The STF performs:

Header validation

PoS signature verification

PoUW proof commit (optional)

Sequential transaction processing

Module hooks execution

Reward distribution (PoS + PoUW)

Epoch transitions

Recompute Merkle roots

Commit new state_root

No nondeterministic operations are permitted.

4. Block Processing Pipeline
4.1 Step 1 — Block Header Validation

Checks:

parent hash

height

proposer signature

PoS committee signatures

timestamp tolerance

previous state root link

4.2 Step 2 — PoUW Proof Validation (If Present)

The state machine verifies the Validatable Work Proof (VWP):

task ID

compute node identity

commitment hashes

model checksum

deterministic checkpoints

redundancy checks

latency & hardware metadata

If valid:

update compute reputation

mark task completed

accumulate compute rewards

If invalid:

slash reputation

reject proof

4.3 Step 3 — Transaction Execution

Executed in order, deterministically.

For each transaction:

verify signature

check nonce

check balance

deduct base gas

execute payload → module

apply storage writes

update Merkle branches

generate events

accumulate gas used

If transaction fails:

state reverts

minimum gas fee is still charged

4.4 Step 4 — Block-Level Module Hooks

Modules may define:

on_block_start()

on_transaction()

on_compute_proof()

on_block_end()

Used for:

computing epoch transitions

updating staking rewards

issuing PoUW bonuses

auto-governance updates

4.5 Step 5 — Rewards & Slashing
Rewards:

50% → PoS validators

50% → PoUW compute nodes

gas fees → validators

compute marketplace fees → nodes + validators + treasury

Slashing:

double-signing

downtime

invalid PoUW proof

fraud attempts

repeated poor performance

4.6 Step 6 — Commit New State Root

The final output:

deterministic updated state

recomputed Merkle roots

new state_root in block header

Consensus finalization (2–4 seconds).

5. Gas Rules for State
Reads

constant cost

Writes

expensive (new node creation, trie expansion)

Deletes

refunded partially (griefing protected)

Governance can tune:

base_fee

storage_fee

compute_fee multipliers

6. Snapshots & Fast Sync

The state machine supports multi-level sync:

6.1 Snapshots

Generated every N blocks (default 10,000):

Includes:

full state

Merkle multiproof bundle

epoch metadata

version tagging

6.2 Sync Modes

Full Sync — replay all blocks

Fast Sync — apply snapshots then replay last blocks

Light Sync — verify state_root only

7. Determinism Guarantees

To avoid consensus splits:

no floating point

no randomness

no time access

no nondeterministic WASM ops

pure host functions

deterministic trie ordering

versioned modules

consistent encoding (SCALE/Borsh)

All honest nodes must compute the same state_root.

8. Error Handling
Transaction Errors

InvalidNonce

InvalidSignature

InsufficientBalance

InvalidModule

GasOverflow

Unauthorized

Proof Errors

InvalidWorkProof

InvalidModelChecksum

LatencyViolation

FraudAttempt

System Errors

EpochTransitionError

StorageError

VersionMismatch

All errors must be deterministic and testable.

9. Testing Requirements
Unit tests

account transitions

staking logic

PoUW reputation scoring

MPT/SMT consistency

Integration tests

multi-block execution

epoch transitions

parallel transaction scenarios

Fuzzing

invalid payloads

storage corruption attempts

adversarial PoUW proofs

10. Summary

The Mbongo Chain State Machine:

defines all global state and rules

enables PoS + PoUW hybrid operation

integrates AI model registry

tracks compute reputation

ensures deterministic, cryptographically secure transitions

supports fast sync, snapshots, and proofs

is the backbone of execution and consensus

It is the foundational layer enabling Mbongo Chain to operate as a verifiable compute-first blockchain.
