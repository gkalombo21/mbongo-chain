
Mbongo Chain — Proof-of-Useful-Work (PoUW) Proof Specification

Validatable Work Proof (VWP) Protocol

1. Introduction

Mbongo Chain integrates verifiable, hardware-accelerated compute directly into consensus.
To securely reward useful work, the chain requires a cryptographically sound method to:

validate compute results

prevent cheating or shortcutting

support heterogeneous hardware (GPU / TPU / NPU / CPU / ASIC)

integrate compute outputs into state

provide deterministic reproducibility when possible

allow probabilistic verification for non-deterministic workloads

This system is called the Validatable Work Proof (VWP).

VWP proofs are included in blocks and validated by every block-producer.

2. PoUW Pipeline Overview
Task Spec → Compute Node → Execution → VWP Proof → Mempool → Validator → Block


More detailed:

A task is created on-chain

A compute node pulls the task

Node executes workload off-chain

Node generates a Validatable Work Proof (VWP)

The proof is submitted to the ProofPool

A validator includes the best proof per task in the block

Execution Engine verifies proof correctness

State Machine updates:

task completed

compute rewards

node reputation

3. VWP Proof Structure

A proof contains:

VWPProof {
    task_id: u64,
    node_id: H256,
    hardware_class: enum { GPU, TPU, NPU, CPU, ASIC },
    checkpoint_hash: H256,
    output_hash: H256,
    redundancy_index: u8,
    execution_time_ms: u32,
    memory_used_mb: u32,
    compute_signature: Signature,      // signs commitment
    metadata: Vec<u8>,                 // hw info, model id, etc.
    proof_bytes: Vec<u8>,              // optional ZK proof or logs
}


Key components:

3.1 output_hash

Hash of final output (tensor, rendering, ZK result, etc.)

3.2 checkpoint_hash

Intermediate hash checkpoints allow fraud detection without re-executing the entire task.

3.3 compute_signature

Compute node must sign:

H(task_id || output_hash || checkpoint_hash || hardware_metadata)

3.4 hardware_class

Determines reputation weighting and scheduling fairness.

3.5 redundancy_index

Used for cross-verification when multiple nodes compute the same task.

4. Task Commitment Scheme

To prevent cheating, Mbongo Chain uses deterministic commitment hashing:

commitment = H(
    input_hash ||
    model_hash ||
    parameters_hash ||
    expected_precision ||
    hardware_metadata
)


Each task created on-chain includes this commitment.
Nodes must generate outputs consistent with the commitment.

This prevents:

using different models

shortcutting with cached results

submitting random values

5. Verification Modes

Depending on workload type, the chain uses 3 verification modes.

5.1 Mode A — Deterministic Verification (AI Inference, ZK Tasks)

Workloads like ONNX inference can reproduce bit-by-bit identical outputs if:

fixed model

fixed input

fixed precision

deterministic kernels

Verification:

recompute small deterministic subsets
validate output_hash
validate checkpoint_hash


No full re-execution required.

5.2 Mode B — Probabilistic Verification (GPU Rendering, ML ops)

For tasks that cannot be fully deterministic:

validator reruns a subset:

random tiles

random batch indices

random rays / pixels

compare with submitted checkpoint

Probability of cheating → decreases exponentially with number of checks.

Example:

Probability of undetected cheating = (acceptable_error)^num_checks

5.3 Mode C — Redundant Multi-node Verification

For heavy tasks:

2–4 compute nodes do the same job

cross-validate each other

Rules:

if ≥ 2 nodes match → valid

if disagreement → run extra checks

malicious nodes → slashed reputation

6. Proof Validation Workflow (On Validator)

When a validator receives a proof:

Step 1 — Basic Validation

task exists & active

node is registered

node has valid reputation

proof format valid

size within bounds

Step 2 — Signature Check
verify(node.pubkey, compute_signature, commitment)

Step 3 — Commitment Consistency

input_hash matches task spec

model_hash approved in AI registry

precision matches

hardware metadata acceptable

Step 4 — Checkpoint Verification

validator re-executes ≤ 1% of task

matches checkpoint_hash

Step 5 — Output Hash Verification

Ensure output obeys expected patterns:

shape

type

quantization

range constraints

Step 6 — Redundancy Resolution (if needed)

If multiple proofs exist:

highest score proof wins

others used for verification

inconsistent proofs → slash offenders

7. Fraud Detection

If a compute node attempts fraud:

Types of Fraud

incorrect output

fake hardware

replayed proof

abandoned tasks

repeated failures

latency manipulation

Penalties

reduce reputation score

reduce scheduling chance

temporarily ban node

permanent ban (critical fraud)

Slashing affects reward eligibility and task priority.

8. Reputation System (PoUW Reputation Tree)

Reputation is stored on-chain under pouw_root.

Each successful proof increases:

reputation += α * quality_score


Each failed/invalid proof:

reputation -= β


Decay per epoch:

reputation *= (1 - decay_rate)


Defaults:

decay_rate = 0.01

α = 1.0

β = 2.0

Nodes with higher reputation:

receive more tasks

earn higher-quality rewards

get prioritized in block inclusion

Nodes with negative reputation are excluded.

9. Reward Calculation

Rewards depend on:

block reward pool

task difficulty

quality score

hardware class weighting

reputation multiplier

Reward formula:
reward = base_reward
       * difficulty
       * quality_factor
       * hardware_weight
       * rep_multiplier


Base reward comes from:

50% of block rewards

compute marketplace fees

10. Proof Size Constraints

PoUW proofs can be large.
Limits are enforced:

Class	Max Size
Small (ZK / logs)	1–5 MB
Medium (inference traces)	≤ 32 MB
Large (render tiles / tensors)	≤ 128 MB

Validators reject exceeding proofs.

11. Mempool Integration (Summary)

ProofPool enforces:

priority ranking

size caps

per-task best proof selection

spam protection

redundancy control

It works in parallel with the TxPool.

12. Block Inclusion Rules

A validator may include:

1 winning proof per task

optional redundant proofs

max 128 proofs per block

max 128 MB total proof data

Proof order is deterministic:

by task_id ascending

by priority descending

by arrival time

13. Future Upgrades

Zero-knowledge proofs for full verification

GPU kernel attestation via remote attestation frameworks

Confidential compute enclaves

Multi-step pipeline for training workloads

Decentralized model-weight shards

14. Summary

The VWP (Validatable Work Proof) system enables:

useful compute to secure consensus

verifiable AI inference

scalable heterogeneous hardware

fair reward distribution

fraud detection

redundancy and probabilistic validation

compatibility with future ZK verification

This system is the core innovation that differentiates Mbongo Chain from other compute networks.