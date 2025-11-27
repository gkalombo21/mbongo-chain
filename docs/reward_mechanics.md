# Reward Mechanics & Distribution Model

> **Document Type:** Reward Specification  
> **Last Updated:** November 2025  
> **Status:** Official Reference

---

## Table of Contents

1. [Overview](#1-overview)
2. [PoS Reward Model](#2-pos-reward-model)
3. [PoUW Reward Model](#3-pouw-reward-model)
4. [Fee Integration](#4-fee-integration)
5. [Reward Lifecycle Diagram](#5-reward-lifecycle-diagram)
6. [Example Reward Scenarios](#6-example-reward-scenarios)
7. [Security Considerations](#7-security-considerations)
8. [For Participants](#8-for-participants)

---

## 1. Overview

### Total Reward Budget

Each finalized block generates a **block reward** from the protocol's emission schedule. This reward, combined with transaction fees collected in the block, forms the total reward pool for that block.

```
┌─────────────────────────────────────────────────────────────┐
│                 TOTAL REWARD COMPOSITION                    │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   Total Block Reward = Base Emission + Transaction Fees    │
│                                                             │
│   Where:                                                    │
│   • Base Emission: Protocol-defined MBO per block          │
│   • Transaction Fees: Sum of all fees in block             │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Distribution Split

The total reward pool is distributed according to a fixed ratio:

```
┌─────────────────────────────────────────────────────────────┐
│                 REWARD DISTRIBUTION SPLIT                   │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   Recipient                    │  Share                    │
│   ─────────────────────────────┼──────────────────────────│
│   PoS Validators               │  70%                      │
│   PoUW Compute Providers       │  30%                      │
│   ─────────────────────────────┼──────────────────────────│
│   Total                        │  100%                     │
│                                                             │
│   This split applies to:                                   │
│   • Base block emission                                    │
│   • Transaction fee pool                                   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## 2. PoS Reward Model

### 2.1 Validator Reward Calculation

Validator rewards are calculated based on three factors:

```
┌─────────────────────────────────────────────────────────────┐
│              VALIDATOR REWARD FACTORS                       │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   1. STAKE WEIGHT                                          │
│      Proportion of total staked MBO                        │
│                                                             │
│   2. PERFORMANCE SCORE                                     │
│      Uptime and correctness metrics                        │
│                                                             │
│   3. BLOCK PRODUCTION                                      │
│      Proposer bonus for produced blocks                    │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Stake Weight Formula

A validator's stake weight determines their share of the PoS reward pool:

```
┌─────────────────────────────────────────────────────────────┐
│                 STAKE WEIGHT CALCULATION                    │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   Stake_Weight(v) = Staked_MBO(v) / Total_Staked_MBO       │
│                                                             │
│   Where:                                                    │
│   • Staked_MBO(v): Validator's own stake + delegations     │
│   • Total_Staked_MBO: Sum of all active validator stakes   │
│                                                             │
│   Example:                                                  │
│   ─────────                                                 │
│   Validator A stake: 100,000 MBO                           │
│   Total network stake: 1,000,000 MBO                       │
│   Stake_Weight(A) = 100,000 / 1,000,000 = 10%             │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.3 Performance Score

Validators earn a performance multiplier based on operational metrics:

| Metric | Weight | Description |
|--------|--------|-------------|
| **Uptime** | 50% | Percentage of expected attestations signed |
| **Correctness** | 30% | Percentage of valid (non-slashed) actions |
| **Latency** | 20% | Speed of vote propagation |

```
┌─────────────────────────────────────────────────────────────┐
│              PERFORMANCE SCORE CALCULATION                  │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   Performance_Score(v) = (0.50 × Uptime)                   │
│                        + (0.30 × Correctness)              │
│                        + (0.20 × Latency_Score)            │
│                                                             │
│   Score Range: 0.0 to 1.0                                  │
│                                                             │
│   Thresholds:                                               │
│   • ≥ 0.95: Excellent (full rewards)                       │
│   • 0.80 - 0.95: Good (minor reduction)                    │
│   • 0.50 - 0.80: Acceptable (moderate reduction)           │
│   • < 0.50: Poor (significant reduction)                   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.4 Final Validator Reward

```
Validator_Reward(v) = PoS_Pool × Stake_Weight(v) × Performance_Score(v)

Where:
  PoS_Pool = Total_Block_Reward × 0.70
```

### 2.5 Slashing Overview

Validators may lose stake for protocol violations:

| Violation | Penalty | Recovery |
|-----------|---------|----------|
| **Double Signing** | 100% of stake | None (permanent) |
| **Invalid Block** | 10% of stake | Rejoin after unbonding |
| **Extended Downtime** | 1% per day | Automatic after recovery |
| **Invalid Attestation** | 5% of stake | Rejoin after unbonding |

> **Note:** Detailed slashing mechanics are specified in `slashing_spec.md`.

---

## 3. PoUW Reward Model

### 3.1 Compute Task Verification Pipeline

PoUW rewards are distributed based on verified compute contributions:

```
┌─────────────────────────────────────────────────────────────┐
│           COMPUTE VERIFICATION PIPELINE                     │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   ┌─────────────┐                                          │
│   │  Task       │  Client submits compute task             │
│   │  Submission │  with input data and fee                 │
│   └──────┬──────┘                                          │
│          │                                                  │
│          ▼                                                  │
│   ┌─────────────┐                                          │
│   │  Task       │  Protocol assigns task to                │
│   │  Assignment │  eligible GPU provider                   │
│   └──────┬──────┘                                          │
│          │                                                  │
│          ▼                                                  │
│   ┌─────────────┐                                          │
│   │  GPU        │  Provider executes task                  │
│   │  Execution  │  in deterministic environment            │
│   └──────┬──────┘                                          │
│          │                                                  │
│          ▼                                                  │
│   ┌─────────────┐                                          │
│   │  Receipt    │  Provider generates signed               │
│   │  Generation │  compute receipt with output hash        │
│   └──────┬──────┘                                          │
│          │                                                  │
│          ▼                                                  │
│   ┌─────────────┐                                          │
│   │  Verification│  Protocol verifies receipt via         │
│   │  (Replicated │  replicated execution or sampling      │
│   │   or Sampled)│                                         │
│   └──────┬──────┘                                          │
│          │                                                  │
│          ▼                                                  │
│   ┌─────────────┐                                          │
│   │  Reward     │  Verified receipt included in block,    │
│   │  Distribution│  provider receives MBO reward          │
│   └─────────────┘                                          │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 3.2 Compute Receipt to MBO Payout

Each verified compute receipt generates a reward based on:

```
┌─────────────────────────────────────────────────────────────┐
│              RECEIPT REWARD FACTORS                         │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   1. TASK DIFFICULTY                                       │
│      Computational complexity of the task                  │
│                                                             │
│   2. RESOURCE USAGE                                        │
│      GPU time, memory, and operations consumed             │
│                                                             │
│   3. VERIFICATION STATUS                                   │
│      Verified (1.0) or Failed (0.0)                        │
│                                                             │
│   4. PROVIDER SCORE                                        │
│      Historical reliability and performance                │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 3.3 Work Units to Rewards

```
┌─────────────────────────────────────────────────────────────┐
│              WORK UNIT CONVERSION                           │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   Work_Units(task) = Base_Units × Difficulty_Multiplier    │
│                                                             │
│   Difficulty Tiers:                                         │
│   ─────────────────                                         │
│   • Tier 1 (Simple):     1x multiplier                     │
│   • Tier 2 (Standard):   2x multiplier                     │
│   • Tier 3 (Complex):    5x multiplier                     │
│   • Tier 4 (Intensive): 10x multiplier                     │
│                                                             │
│   Provider_Share = Work_Units(p) / Total_Work_Units        │
│                                                             │
│   Provider_Reward = PoUW_Pool × Provider_Share             │
│                                                             │
│   Where:                                                    │
│     PoUW_Pool = Total_Block_Reward × 0.30                  │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 3.4 GPU Scoring Model

Provider scores influence task assignment priority and reward multipliers:

| Factor | Weight | Description |
|--------|--------|-------------|
| **Success Rate** | 40% | Verified tasks / Total assigned |
| **Timeliness** | 25% | On-time completions |
| **Capacity** | 20% | Available compute resources |
| **History** | 15% | Long-term reliability |

```
┌─────────────────────────────────────────────────────────────┐
│                 GPU PROVIDER SCORE                          │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   Provider_Score = (0.40 × Success_Rate)                   │
│                  + (0.25 × Timeliness)                     │
│                  + (0.20 × Capacity_Score)                 │
│                  + (0.15 × History_Score)                  │
│                                                             │
│   Score Range: 0.0 to 1.0                                  │
│                                                             │
│   Score Impact:                                             │
│   • Higher score → More task assignments                   │
│   • Higher score → Priority in task queue                  │
│   • Score < 0.5 → Reduced assignments                      │
│   • Score < 0.3 → Temporary suspension                     │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

> **Note:** Detailed GPU scoring aligned with AIDA is specified in `aida_scoring.md`.

---

## 4. Fee Integration

### 4.1 Fee Types

```
┌─────────────────────────────────────────────────────────────┐
│                    FEE STRUCTURE                            │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   FEE TYPE           │ DESCRIPTION                         │
│   ───────────────────┼─────────────────────────────────────│
│   Gas Fee            │ Per-operation execution cost        │
│   Storage Fee        │ State storage allocation cost       │
│   Base Transaction   │ Minimum fee per transaction         │
│   Priority Fee       │ Optional tip for faster inclusion   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 4.2 Gas Fees (Execution)

Gas fees compensate for computational resources consumed during transaction execution:

```
Gas_Fee = Gas_Used × Gas_Price

Where:
  Gas_Used: Actual computation units consumed
  Gas_Price: MBO per gas unit (set by user)
```

### 4.3 Storage Fees

Storage fees compensate for permanent state storage:

```
Storage_Fee = Storage_Units × Storage_Price_Per_Unit

Where:
  Storage_Units: Bytes of new state created
  Storage_Price_Per_Unit: Protocol-defined rate
```

### 4.4 Base Transaction Fee

Every transaction pays a minimum base fee regardless of complexity:

```
Base_Fee = MINIMUM_TX_FEE (protocol constant)
```

### 4.5 Fee Redistribution

All collected fees are added to the block reward pool and distributed:

```
┌─────────────────────────────────────────────────────────────┐
│                 FEE REDISTRIBUTION                          │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   Total_Fees = Σ(Gas_Fees + Storage_Fees + Base_Fees)      │
│                                                             │
│   Fee Distribution:                                         │
│   ─────────────────                                         │
│   • 70% to PoS Validators (proposer + attesters)           │
│   • 30% to PoUW Providers (if receipts in block)           │
│                                                             │
│   If no PoUW receipts in block:                            │
│   • 100% to PoS Validators                                 │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## 5. Reward Lifecycle Diagram

```
┌─────────────────────────────────────────────────────────────────────────────────────────┐
│                           REWARD LIFECYCLE                                              │
└─────────────────────────────────────────────────────────────────────────────────────────┘

  ┌──────────────┐
  │    BLOCK     │
  │  FINALIZED   │
  └──────┬───────┘
         │
         │ Block contains:
         │ • Transactions (fees)
         │ • PoUW Receipts
         │
         ▼
  ┌──────────────────────────────────────────────────────────────────────────────────────┐
  │                              REWARD POOL                                             │
  │                                                                                      │
  │   ┌────────────────────────┐         ┌────────────────────────┐                     │
  │   │    BASE EMISSION       │    +    │   TRANSACTION FEES     │                     │
  │   │   (Protocol MBO)       │         │   (Gas + Storage)      │                     │
  │   └────────────────────────┘         └────────────────────────┘                     │
  │                                                                                      │
  │                              = TOTAL REWARD POOL                                    │
  │                                                                                      │
  └──────────────────────────────────────┬───────────────────────────────────────────────┘
                                         │
                    ┌────────────────────┴────────────────────┐
                    │                                         │
                    ▼                                         ▼
  ┌──────────────────────────────────┐      ┌──────────────────────────────────┐
  │         PoS POOL (70%)           │      │        PoUW POOL (30%)           │
  │                                  │      │                                  │
  │   Distribution by:               │      │   Distribution by:               │
  │   • Stake Weight                 │      │   • Work Units                   │
  │   • Performance Score            │      │   • Verification Status          │
  │   • Block Production             │      │   • Provider Score               │
  │                                  │      │                                  │
  └───────────────┬──────────────────┘      └───────────────┬──────────────────┘
                  │                                         │
       ┌──────────┴──────────┐               ┌──────────────┴──────────────┐
       │                     │               │                             │
       ▼                     ▼               ▼                             ▼
  ┌──────────┐         ┌──────────┐    ┌──────────┐               ┌──────────┐
  │ PROPOSER │         │ ATTESTERS│    │PROVIDER 1│               │PROVIDER N│
  │          │         │          │    │          │               │          │
  │ Higher   │         │ Stake-   │    │ Work-    │               │ Work-    │
  │ share    │         │ weighted │    │ weighted │               │ weighted │
  └──────────┘         └──────────┘    └──────────┘               └──────────┘
       │                     │               │                             │
       │                     │               │                             │
       ▼                     ▼               ▼                             ▼
  ┌──────────────────────────────────────────────────────────────────────────────────────┐
  │                              PARTICIPANT WALLETS                                     │
  │                                                                                      │
  │   MBO credited to validator and compute provider accounts                           │
  │                                                                                      │
  └──────────────────────────────────────────────────────────────────────────────────────┘
```

---

## 6. Example Reward Scenarios

### Scenario 1: Normal Validator

```
┌─────────────────────────────────────────────────────────────┐
│              SCENARIO: NORMAL VALIDATOR                     │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   Context:                                                  │
│   • Block Reward: 1.0 MBO                                  │
│   • Transaction Fees: 0.1 MBO                              │
│   • Total Pool: 1.1 MBO                                    │
│   • PoS Pool: 0.77 MBO (70%)                               │
│                                                             │
│   Validator Stats:                                          │
│   • Stake: 50,000 MBO                                      │
│   • Total Network Stake: 1,000,000 MBO                     │
│   • Stake Weight: 5%                                       │
│   • Performance Score: 0.95                                │
│                                                             │
│   Calculation:                                              │
│   ─────────────                                             │
│   Base Share = 0.77 × 0.05 = 0.0385 MBO                    │
│   With Performance = 0.0385 × 0.95 = 0.0366 MBO            │
│                                                             │
│   Per-Block Reward: ~0.037 MBO                             │
│   Daily Reward (43,200 blocks): ~1,598 MBO                 │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Scenario 2: High-Performance GPU Provider

```
┌─────────────────────────────────────────────────────────────┐
│         SCENARIO: HIGH-PERFORMANCE GPU PROVIDER             │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   Context:                                                  │
│   • Block Reward: 1.0 MBO                                  │
│   • Transaction Fees: 0.1 MBO                              │
│   • Total Pool: 1.1 MBO                                    │
│   • PoUW Pool: 0.33 MBO (30%)                              │
│                                                             │
│   Provider Stats:                                           │
│   • Tasks Completed (block): 5                             │
│   • Work Units: 50 (Tier 3 tasks)                          │
│   • Total Network Work Units: 200                          │
│   • Provider Score: 0.92                                   │
│                                                             │
│   Calculation:                                              │
│   ─────────────                                             │
│   Work Share = 50 / 200 = 25%                              │
│   Base Reward = 0.33 × 0.25 = 0.0825 MBO                   │
│                                                             │
│   Per-Block Reward: ~0.083 MBO                             │
│   Daily Reward (if consistent): ~3,586 MBO                 │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Scenario 3: Heavy-Block Epoch with High Fees

```
┌─────────────────────────────────────────────────────────────┐
│        SCENARIO: HEAVY-BLOCK EPOCH (HIGH FEES)              │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   Context:                                                  │
│   • Block Reward: 1.0 MBO                                  │
│   • Transaction Fees: 5.0 MBO (network congestion)         │
│   • Total Pool: 6.0 MBO                                    │
│   • PoS Pool: 4.2 MBO (70%)                                │
│   • PoUW Pool: 1.8 MBO (30%)                               │
│                                                             │
│   Block Proposer (10% stake, 0.98 performance):            │
│   ─────────────────────────────────────────────            │
│   Base Share = 4.2 × 0.10 = 0.42 MBO                       │
│   With Performance = 0.42 × 0.98 = 0.41 MBO                │
│   Proposer Bonus = 0.05 MBO                                │
│   Total: ~0.46 MBO                                         │
│                                                             │
│   Top GPU Provider (30% work share):                       │
│   ─────────────────────────────────                        │
│   Work Share = 1.8 × 0.30 = 0.54 MBO                       │
│                                                             │
│   Note: High-fee blocks significantly boost rewards        │
│   for all participants proportionally.                     │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## 7. Security Considerations

### 7.1 Preventing Reward Manipulation

| Attack Vector | Mitigation |
|---------------|------------|
| **Stake Splitting** | Minimum stake requirement prevents dust attacks |
| **Fake Compute** | Replicated verification catches fraudulent receipts |
| **Self-Dealing** | Task assignment is deterministic and verifiable |
| **Collusion** | Statistical analysis detects abnormal patterns |
| **Double Claiming** | Unique receipt IDs prevent duplicate rewards |

### 7.2 Economic Integrity with Fixed Supply

```
┌─────────────────────────────────────────────────────────────┐
│              ECONOMIC INTEGRITY GUARANTEES                  │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   ✓ Rewards never exceed emission schedule                 │
│   ✓ Fees are redistributed, not created                    │
│   ✓ Slashed stake is burned or redistributed               │
│   ✓ No mechanism can mint MBO beyond cap                   │
│   ✓ All rewards are auditable on-chain                     │
│                                                             │
│   Invariant:                                                │
│   ───────────                                               │
│   Total_MBO_Issued ≤ 31,536,000 MBO (always)               │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 7.3 Reward Verification

All reward distributions are verifiable:

- Block rewards computed from deterministic emission schedule
- Fee totals derived from executed transactions
- PoS shares calculated from stake registry
- PoUW shares calculated from verified receipts
- All distributions recorded on-chain

---

## 8. For Participants

### For Validators

As a validator, your MBO rewards depend on:

- **Stake Size:** Larger stake increases your share of the PoS pool
- **Performance:** Maintain high uptime and vote correctly to maximize rewards
- **Block Production:** Earn proposer bonuses when elected as block leader
- **Commission:** Set competitive rates to attract delegations while remaining profitable

Optimize by: Running reliable infrastructure, minimizing downtime, and building delegation trust.

### For Compute Providers

As a GPU provider, your MBO rewards depend on:

- **Work Volume:** Complete more verified tasks to increase your work unit share
- **Task Difficulty:** Higher-tier tasks generate more work units per completion
- **Provider Score:** Maintain high success rate and timeliness for priority assignment
- **Capacity:** More GPU resources enable handling larger task volumes

Optimize by: Ensuring deterministic execution, maintaining high availability, and investing in capable hardware.

### For Delegators (Future Extension)

As a delegator, you can earn passive MBO rewards by:

- **Selecting Validators:** Choose high-performance validators with fair commission rates
- **Diversifying:** Spread delegation across multiple validators to reduce risk
- **Monitoring:** Track validator performance and redelegate if necessary
- **Compounding:** Restake earned rewards to increase future returns

> **Note:** Delegation mechanics will be detailed in `delegation_guide.md`.

---

## 9. Related Documentation

| Document | Description |
|----------|-------------|
| `token_intro.md` | MBO token introduction |
| `monetary_policy.md` | Fixed supply and halving schedule |
| `slashing_spec.md` | Detailed slashing mechanics |
| `aida_scoring.md` | GPU scoring model |
| `delegation_guide.md` | Delegator instructions |

---

*This document defines the reward mechanics for Mbongo Chain. All reward calculations are deterministic and verifiable on-chain.*

