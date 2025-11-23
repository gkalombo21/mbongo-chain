
Version: Canonical
Status: Final
Author: Mbongo Foundation

Mbongo Chain uses a hybrid consensus architecture called PoX — Proof-of-Everything Useful, combining:

PoS (Proof-of-Stake) for block production, finality, and governance

PoUW (Proof-of-Useful-Work) for verifiable compute contributions

PoC (Proof-of-Compute) for hardware capability attestation

This consensus design ensures:

economic security

verifiable compute integrity

hardware diversity

long-term sustainability

AI-native scalability

1. Consensus Goals

Mbongo Chain consensus is designed to:

Secure the blockchain using stake-based validation

Integrate verifiable compute into block reward distribution

Incentivize heterogeneous hardware (GPU/TPU/NPU/ASIC)

Provide high throughput with 1-second blocks

Enable AI-native workloads as first-class citizens

Ensure fairness, transparency, and decentralization

Support future hardware generations and compute advances

2. Consensus Architecture Overview

Consensus relies on three integrated layers:

2.1 Proof-of-Stake (PoS)

PoS determines:

validator set

block proposers

finality

slashing

governance voting

transaction ordering

2.2 Proof-of-Useful-Work (PoUW)

PoUW determines:

compute task assignment

proof verification

compute reward distribution

compute reputation scoring

PoUW is not used for block production — only for economic rewards and compute validation.

2.3 Proof-of-Compute (PoC)

PoC provides:

hardware attestation

performance benchmarking

reward weighting

fairness and anti-centralization

Rewards scale with PoC score, not raw hashrate.

3. Block Production Process

Block production occurs every 1 second, in four phases:

3.1 Phase 1 — Validator Selection (VRF + Stake Weight)

Validators are selected using:

Verifiable Random Function (VRF)

Weighted by effective stake

Ensuring fairness and unpredictability

Each slot has:

1 proposer

committee validators for validation

PoUW compute receipts from previous tasks

3.2 Phase 2 — Block Assembly

Block includes:

transactions

state transitions

compute receipts

PoUW proof commitments

validator signatures

state root

randomness (VRF)

AIDA-logged economic parameters

3.3 Phase 3 — Block Validation (Committee)

Committee performs:

signature verification

state transition validity

gas limit checks

PoUW proof verification

cross-checking compute receipts

slashing detection

3.4 Phase 4 — Finality

Finality is achieved using:

multi-round BFT-style signatures

slashing for missed commitments

deterministic state root agreement

4. Reward Distribution — Synchronized With Emission
Reward per block:
R(t) = 0.1 MBO / 2^n  (n = number of 5-year halvings)

Distribution:
PoS validators = 50% of R(t)
PoUW compute nodes = 50% of R(t)


This reward split is:

hard-coded

approved by DAO + Founder Council for 10 years

never controlled by AIDA

never modified automatically

5. Proof-of-Stake (PoS) in Detail
5.1 Staking Requirements

minimum stake: DAO-defined

full-slash for double-sign

partial slash for downtime

stake unlocking delay: 21 days (suggested)

5.2 Validator Responsibilities

produce blocks

validate committee blocks

validate PoUW proofs

vote on governance

maintain uptime and HSM security

5.3 Slashing Conditions
Condition	Penalty
Double-sign	5% stake
Downtime	1% stake
Fraud detection	10% stake
Repeated violations	Jail or permanent removal

Funds go to:

treasury

or burn (if AIDA determines burn period active)

6. Proof-of-Useful-Work (PoUW) in Detail

PoUW allows compute providers to participate in consensus economically, not structurally.

6.1 Compute Task Lifecycle

Tasks published

Nodes bid or pull

Execution on hardware

Proof generation

Submission

Cross-check by PoS validators

Reward distribution

6.2 PoUW Verification Techniques

deterministic execution (when possible)

redundancy (R ≥ 2)

probabilistic verification

ZK-proofs (future upgrade)

reputation-based scaling

6.3 Fraud Penalties

Invalid proofs lead to:

forfeiture of current reward

reputation loss

potential slashing

permanent ban for repeated fraud

7. Proof-of-Compute (PoC) in Detail

PoC provides a hardware score to ensure fairness.

Measured attributes:

FLOPs

VRAM bandwidth

reliability

uptime

job success rate

PoC determines:

task eligibility

compute weight

reward multiplier

8. AIDA Integration (Bounded Role)

AIDA participates only in fee regulation, NOT in consensus.

AIDA CANNOT:

influence proposer selection

modify PoS/PoUW reward split

adjust stake weight

alter slashing

override validators

modify emission

AIDA CAN:

adjust burn rate of fees

adjust gas multiplier

adjust compute fee multiplier

provide risk simulation reports

AIDA increases economic resilience without centralizing consensus.

9. Consensus Security Model Summary

Mbongo Chain secures consensus through:

✔ Stake-based security (PoS)

Capturing >50% of staked MBO is economically unattainable.

✔ Compute-based security (PoUW)

Fraud proofs punish bad actors economically and reputationally.

✔ Hardware diversity (PoC)

Prevents specialized hardware takeovers.

✔ Fixed-supply scarcity

Resists inflationary pressure.

✔ Vesting locks

Prevent governance-buyouts in early years.

✔ Founder Council (10 years)

Prevents adversarial takeovers during network growth.

✔ AIDA bounded regulator

Ensures stability against volatility and compute-demand shock.

✔ Redundancy + slashing

Punishes malicious compute submissions.

✔ Decentralized committee validation

Reduces collusion risk.

10. Final Summary

Mbongo Chain consensus is a hybrid PoS + PoUW + PoC model designed for:

high security

high throughput

compute-first architecture

long-term decentralization

economic and compute integrity

AI-native workloads

multi-decade sustainability

This PoX model ensures that useful compute strengthens security, not wasteful hashing.
