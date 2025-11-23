economic_security.md

(Canonical — Mbongo Chain Economic Security Model)

# Mbongo Chain — Economic Security Model  
Status: Canonical  
Version: v1.0

Mbongo Chain’s economic security model is built around three pillars:

1. **Fixed Supply Monetary Baseline** (31,536,000 MBO)  
2. **PoS + PoUW Hybrid Consensus**  
3. **AIDA-Regulated Economic Stability**  

This document defines how Mbongo Chain maintains security, prevents economic attacks, balances compute demand, and preserves long-term value.

---

# 1. Security Philosophy

Economic security in Mbongo Chain is based on:

- strong monetary predictability  
- verifiable compute rewards  
- bounded AI-regulation (AIDA)  
- strict governance controls  
- adversary-resistant economics  
- stable incentives for validators and compute providers  

The model ensures resistance to:

- spam  
- bribe attacks  
- economic capture  
- resource starvation  
- compute monopolization  
- supply shocks  
- governance takeover attempts  

---

# 2. Fixed Monetary Baseline

## 2.1 Immutable Supply

Total supply is permanently fixed at:

**31,536,000 MBO** (seconds in a year)

This provides:

- maximal predictability  
- anti-inflation guarantees  
- a mathematical foundation for scarcity  
- long-term value protection  

No actor — DAO, AIDA, Founder Council — can change this supply.

## 2.2 Emission Predictability

The emission curve:

- initial reward: 0.1 MBO/block  
- halving every 5 years  
- negligible inflation after ~20 years  

This results in:

- stable long-term staking yields  
- robust incentives for compute providers  
- protection against runaway inflation  

---

# 3. Hybrid Consensus Security (PoS + PoUW)

Mbongo Chain’s PoX model strengthens security by combining:

- **Proof-of-Stake (PoS)**  
- **Proof-of-Useful-Work (PoUW)**  
- **Proof-of-Compute (PoC)** for hardware validation  

### Benefits:

- attacks require both **stake** and **compute**  
- resistance to raw hashpower takeover  
- GPU compute is applied to **useful tasks**, not wasteful hashing  
- compute providers earn verifiable rewards  
- consensus receives real value from network activity  

---

## 3.1 PoS Security

PoS validators:

- secure block production  
- verify PoUW receipts  
- participate in finality  
- commit economic value through staking  
- are subject to **slashing** for misbehavior

Slashing conditions:

- double signing  
- invalid PoUW receipt acceptance  
- censorship  
- liveness failures  

Staking provides:

- skin-in-the-game  
- cost for attacking  
- penalty for malicious behavior  

---

## 3.2 PoUW Security

PoUW compute nodes:

- execute AI/compute tasks  
- produce verifiable task receipts  
- subject to redundancy validation  
- subject to fraud proofs  

Economic incentives:

- PoUW nodes earn **50% of block rewards**  
- additional compute fees from job submitters  
- conditioned on honest behavior and task success  

Security contributions:

- verifiable compute protects against malicious outputs  
- fraud-proof system catches incorrect task executions  
- redundant task execution detects inconsistencies  
- PoC hardware attestation prevents fake compute nodes  

---

# 4. Price Stability & Fee Security

Economic stability is provided by:

- **AIDA-regulated burn**  
- **base fee multiplier**  
- **priority fee constraints**  
- **PoUW pricing multiplier**

These mechanisms prevent:

- spam  
- fee volatility  
- compute-price manipulation  
- congestion-based attacks  
- sudden cost explosions for users  

---

# 5. AIDA: Bounded Economic Regulator

AIDA is a **bounded, supervised, deterministic regulator**  
(see `aida_regulator.md`).

AIDA adjusts:

- burn rate (0%–30%)  
- base fee multiplier (0.5–3.0)  
- PoUW multiplier (0.8–1.2)  

AIDA provides:

- fee stability  
- compute market equilibrium  
- demand-based burn  
- risk forecasting  
- anti-volatility behavior  

AIDA cannot:

- mint tokens  
- modify supply cap  
- touch emission schedule  
- override consensus  

Governance protections:

- DAO supervises normal operations  
- Founder Council oversees critical changes (10 years)

This prevents AIDA from being used as an attack vector.

---

# 6. Attack Resistance Model

Mbongo Chain is designed to resist:

---

## 6.1 Economic Capture Attacks

Attacker goal:  
Gain disproportionate influence by purchasing tokens or compute resources.

Mitigations:

- fixed supply → no infinite dilution  
- vesting locks → prevents sudden founder exit  
- Founder Council → veto protection for 10 years  
- AIDA bounds → prevent price manipulation via fees  
- PoUW redundancy → prevents compute-based takeover  
- slashing → cost-increasing factor for misconduct  

---

## 6.2 Governance Takeover Attempts

Attempt:
Acquire tokens to dominate votes.

Defense:

- quadratic lock-weighted voting  
- Founder Council veto on critical protocol changes  
- 90-day Safety Review Window for all upgrades  
- AIDA cannot influence governance  
- vesting tokens count toward long-term stability  
- DAO supermajority required for sensitive changes  

---

## 6.3 PoUW Manipulation Attacks

Attempts:

- fake compute  
- fabricated results  
- collusion between compute nodes  
- subversion of PoUW market pricing  

Defense:

- PoC hardware verification  
- fraud proofs  
- redundant computation  
- AIDA cannot modify job rewards beyond 0.8–1.2x  
- deterministic verification rules  
- validator-level checks  

---

## 6.4 Censorship Attacks

Mitigations:

- VRF-based leader selection  
- distributed validator set  
- PoUW contributions from diverse hardware  
- slashing for censorship  
- alternate routing via libp2p  

---

## 6.5 Fee Manipulation / Burn Manipulation

Defense:

- AIDA bounded ranges  
- Founder Council oversight  
- priority fee upper limits  
- deterministic fee model  
- circuit breaker for abnormal changes  
- rollback protection  

---

# 7. Treasury Security

The treasury is controlled by:

- DAO (primary)  
- Foundation (execution)  
- time-locked smart contract  

Treasury cannot:

- change token supply  
- influence AIDA burn  
- mint new MBO  
- override governance  

AIDA may **advise** treasury allocation, but not enforce actions.

---

# 8. Long-Term Security

Over decades:

- emission declines (like Bitcoin)  
- fees + PoUW payments become dominant  
- AIDA ensures stable economic operation  
- governance matures as Founder Council expires at year 10  
- PoUW grows with global compute demand  
- staking rewards remain stable due to predictable halving  

This allows a **century-scale economic model** similar to Bitcoin, but enhanced by:

- PoUW  
- AI compute economy  
- dynamic burn  
- governance protections  

---

# 9. Summary

Mbongo Chain’s economic security is grounded in:

- immutable supply  
- predictable emission  
- hybrid PoS + PoUW security  
- bounded AI regulation (AIDA)  
- governance protections (Founder Council 10 years)  
- redundancy & fraud-proof compute  
- anti-volatility fee mechanics  
- long-term equilibrium  

The system is resistant to:

- economic capture  
- compute-level manipulation  
- governance attacks  
- volatility-based instability  
- fee abuse  

- In practice, ~99% of the MBG supply will be issued after ~35 years, and ~99.9% after ~50 years from genesis. After that point, the network is primarily secured by fees and compute rewards.


Mbongo Chain is designed to maintain **robust, long-term economic security** while powering a global marketplace of verifiable compute.