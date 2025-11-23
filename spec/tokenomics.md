(Canonical — Mbongo Chain Tokenomics Specification)

# Mbongo Chain — Tokenomics Specification  
Status: Canonical  
Version: v1.0

The MBO token is the native asset of Mbongo Chain.  
It secures the network through Proof-of-Stake (PoS), powers the Proof-of-Useful-Work (PoUW) economy, supports governance, and aligns long-term incentives across all ecosystem actors.

Mbongo Chain adopts a **non-inflationary, fixed-supply, time-based monetary model** designed for durability, stability, and verifiable compute.

---

# 1. Monetary Parameters

## 1.1 Total Fixed Supply

**Total Supply: 31,536,000 MBO**  
(= number of seconds in a year)

This establishes Mbongo Chain as a **time-backed, compute-native digital asset**, aligning scarcity with a universal physical constant.

Supply is immutable and enforced at protocol level.  
No governance vote, AIDA update, or upgrade can modify this number.

---

# 2. Emission Schedule (PoS + PoUW)

MBO follows a predictable, Bitcoin-style emission curve:

- **Block time:** 1 second  
- **Initial block reward:** **0.1 MBO/block**  
- **Halving interval:** every **157,680,000 blocks** (≈ 5 years)  
- **Asymptotic max supply:** 31,536,000 MBO

### 2.1 Reward Progression

| Years | Reward/block | Reward/year | Total emitted in period |
|-------|--------------|-------------|--------------------------|
| 0–5   | 0.1 MBO      | 3,153,600   | 15,768,000 MBO |
| 5–10  | 0.05 MBO     | 1,576,800   | 7,884,000 MBO |
| 10–15 | 0.025 MBO    | 788,400     | 3,942,000 MBO |
| 15–20 | 0.0125 MBO   | 394,200     | 1,971,000 MBO |
| 20–25 | 0.00625 MBO  | 197,100     | 985,500 MBO |
| ...   | ...          | ...         | ... |

### 2.2 Emission Completion

- ≈ **95%** emitted by year ~20  
- ≈ **99.9%** emitted by year ~50  
- asymptotic tail continues → emissions become negligible

This ensures:

- predictable scarcity  
- long-term security  
- sustainability for validators & PoUW compute nodes

---

# 3. Reward Split: PoS + PoUW

Block rewards are split between:

- **PoS Validators** — providing economic security  
- **PoUW Compute Nodes** — contributing useful compute

### Canonical Split:

- **50% → PoUW** (compute marketplace, AI tasks, GPU work)  
- **50% → PoS** (stake-based consensus)

This ensures:

- alignment between compute supply and economic security  
- predictable incentives for compute providers  
- stable staking returns during early network growth

Note: the DAO may adjust this split **within a safe range (40%–60%)**, with 10-year Founder Council oversight.

---

# 4. Transaction Fees, Gas, and Burn

Fees on Mbongo Chain serve three purposes:

1. Prevent spam  
2. Incentivize validators and compute nodes  
3. Enable **dynamic burn** via AIDA

## 4.1 Fee Components

- **Base Fee** — algorithmically adjusted  
- **Priority Fee** — optional tip  
- **Compute Fee** — PoUW job cost  
- **AIDA-regulated Burn** (0%–30%)

---

# 5. AIDA — Dynamic Burn Regulation

AIDA (Autonomous Intelligent Dynamic Adjuster) governs **fee burn**, **PoUW pricing multipliers**, and **base fee stability**.

AIDA respects strict boundaries:

### Allowed Ranges



burn_rate ∈ [0% , 30%]
base_fee_multiplier ∈ [0.5 , 3.0]
priority_fee_limit ∈ [0 , 2.0]
pouw_multiplier ∈ [0.8 , 1.2]


### AIDA Is Not Allowed To:

- mint new tokens  
- increase supply  
- touch emission curve  
- override consensus rules  
- modify vesting schedules  

AIDA operates **under DAO governance + Founder Council oversight (10 years)**.

---

# 6. Token Distribution (Genesis Allocation)

Mbongo Chain follows a **balanced, credibility-first** distribution model used by top L1s (Solana, Aptos, Sui, Avalanche).

### 6.1 Genesis Supply Breakdown

| Category | % of Supply | Amount (MBO) | Vesting |
|----------|--------------|---------------|----------|
| **Network Rewards (PoS + PoUW)** | **45%** | 14,190,000 | Emitted over ~50 years |
| **Foundation + Ecosystem** | **20%** | 6,307,200 | 4-year vesting, 1-year cliff |
| **Founders Allocation** | **10%** | 3,153,600 | 4-year vesting, 1-year cliff |
| **Early Core Contributors** | **5%** | 1,576,800 | 4-year vesting |
| **Community & Airdrops** | **10%** | 3,153,600 | unlocked gradually |
| **Public Sale / Strategic Partners** | **10%** | 3,153,600 | 12–24 month vesting |

Total = **31,536,000 MBO**

---

# 7. Long-Term Vesting Model

### 7.1 Founders Allocation
- **4-year vesting**
- **1-year cliff**
- **monthly unlock**
- ensures long-term alignment
- cannot be changed by AIDA or the DAO

### 7.2 Core Contributors
- vesting tied to continued contributions
- optional DAO revocation for inactive members

### 7.3 Foundation
- controls grant allocation  
- cannot sell > X% per quarter (DAO-regulated)

---

# 8. Governance Utility

MBO tokens grant:

- voting power in the DAO  
- right to propose upgrades  
- veto participation (through Founder Council until year 10)  
- voting weight **increases with lock duration**  
  (Quadratic Lock Power Model)

---

# 9. Economic Security Model

 MBO is used to:

- secure PoS validators  
- fund PoUW compute rewards  
- support treasury operations  
- fuel governance  
- maintain network safety via slashing  

This aligns incentives between:

- validators  
- compute providers  
- developers  
- token holders  
- Foundation  
- governance participants  

See `economic_security.md` for full details.

---

# 10. Fee Model Summary

| Component | Purpose | Controlled By |
|-----------|----------|----------------|
| Base Fee | Prevent spam | AIDA (bounded) |
| Priority Fee | User incentive | Market |
| Compute Fee | Pay for AI/compute jobs | PoUW marketplace |
| Burn | Reduce supply, balance demand | AIDA (0%–30%) |

AIDA stabilizes fees without compromising neutrality.

---

# 11. Deflation Model (AIDA + Fees)

Mbongo Chain uses a **controlled deflation** similar to EIP-1559 but enhanced via:

- PoUW multipliers  
- dynamic burn  
- bounded ML projections  
- controlled parameter ranges  
- Founder Council oversight  

This produces:

- increasing scarcity when demand is high  
- predictable cost during low demand  
- balanced economics for compute providers  

---

# 12. Summary

Mbongo Chain’s tokenomics provide:

- **fixed supply** (31.536M MBO)  
- **Bitcoin-style halving schedule**  
- **hybrid PoS + PoUW reward system**  
- **dynamic burn regulation via AIDA**  
- **strong founder protection (10 years)**  
- **balanced initial distribution**  
- **predictable long-term scarcity**  
- **incentives aligned with verifiable compute**  

- In practice, ~99% of the MBG supply will be issued after ~35 years, and ~99.9% after ~50 years from genesis. After that point, the network is primarily secured by fees and compute rewards.

This monetary model enables Mbongo Chain to become the foundational settlement layer for decentralized AI compute.
