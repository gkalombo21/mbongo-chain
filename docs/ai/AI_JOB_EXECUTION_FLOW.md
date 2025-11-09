# ⚙️ Mbongo Chain — AI Job Execution Flow

## 1. Overview

This document explains how AI tasks are created, dispatched, verified, and rewarded inside the Mbongo Chain network under the **Proof of Useful Work (PoUW)** consensus.

Each job represents a meaningful computation — for example:

- Model training  
- Image processing  
- Data classification  
- AI inference requests

---

## 2. Flow Summary

---

## 3. Step-by-Step Execution

### Step 1 — Job Creation

A user or partner app sends an AI task request via the API endpoint:

```json
POST /api/v1/ai/job
{
 "user": "kaayu-client-001",
 "model": "gpt-lite",
 "data_url": "https://kaayu.com/dataset-001.json",
 "reward_per_unit": 0.5
}
```

### Step 2 — Smart Contract Registration

The smart contract:

- Hashes the task metadata  
- Locks the corresponding reward pool in MBG tokens  
- Broadcasts an AIJobCreated event to the network  

Validators store the event in the blockchain ledger.

### Step 3 — Job Dispatch

The Task Pool analyzes available compute nodes (gpu and cpu types).

It selects a node based on:

- Performance score (reputation)  
- Latency  
- Region match or energy cost  

Once selected, the job is sent to the node via gRPC.

### Step 4 — Execution on Compute Node

The node downloads the dataset and runs the computation in a sandboxed environment.

Output example:

```
result_hash: "aeff99a1b2..."
accuracy: 0.93
runtime_ms: 3450
```

### Step 5 — Proof Submission

The node submits the result to the blockchain:

```json
POST /api/v1/ai/proof
{
 "job_id": "JOB-2025-0009",
 "node_id": "gpu-validator-4",
 "proof": "0xabcdef12345...",
 "hash": "aeff99a1b2...",
 "accuracy": 0.93
}
```

### Step 6 — Reward Distribution

Once verified, the contract distributes the reward:

```
Reward = WorkUnits × Accuracy × RewardRate
```

| Error | Cause | Resolution |

> |--------|--------|-------------|
> | Invalid proof | Hash mismatch | Node must re-submit ZKP |
> | Timeout | Node offline | Job is redispatched |
> | Low accuracy | Model failed | Partial reward applied |

| Event Name | Triggered When | Logged In |

> |-------------|----------------|------------|
> | AIJobCreated | New job registered | blockchain ledger |
> | AIProofSubmitted | Proof received | ledger and API logs |
> | AIRewardDistributed | Tokens paid | bank module |

---

## 6. Integration with Kaayu Platform

Kaayu users can submit AI jobs directly from the Kaayu dashboard.  
The platform will use Mbongo’s API gateway to dispatch jobs securely to the PoUW network.

Example use cases:

- HR AI recommendation engine  
- Resume parsing via language models  
- Automated data analysis for clients

---

| Feature | Description | ETA |

> |----------|--------------|-----|
> | Batch Job Scheduler | Group multiple jobs per epoch | Q2 2026 |
> | Federated Compute | Decentralized training across nodes | Q3 2026 |
> | AI Marketplace | User-submitted models for rent | Q4 2026 |
