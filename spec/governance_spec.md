Mbongo Chain — Governance Specification

Version: Canonical
Status: Final

Mbongo Chain adopts a hybrid governance model designed to ensure long-term protocol stability, protect the architectural vision, prevent hostile governance takeovers, and gradually transition toward full decentralization. This specification defines the roles, processes, and constraints governing the evolution of the protocol.

1. Governance Overview

Mbongo Chain governance consists of four complementary components:

The DAO (Decentralized Autonomous Organization)
– Handles normal governance activities
– Controls treasury, grants, and operational decisions

The Founder Council (10-Year Strategic Oversight)
– Ensures protection against adversarial governance during the early evolution
– Approves only critical protocol changes
– 10-year limited mandate

The Mbongo Foundation (Non-Profit Entity)
– Supports ecosystem growth
– Provides research, security audits, and grant distribution
– Holds a permanent founder seat (non-revocable for 10 years)

AIDA Regulator (AI-Assisted Governance Module)
– Provides simulations, audits, and risk forecasting
– Adjusts burn-rate parameters within strict boundaries
– Guided and supervised by the DAO and the Founder Council

This structure ensures:

decentralization

stability

founder protection

transparency

long-term sustainability

2. Governance Principles

Governance in Mbongo Chain follows these principles:

2.1 Transparency

All governance actions are:

on-chain

signed

publicly visible

auditable

2.2 Stability First

Critical economic parameters cannot be modified without extensive review.

2.3 Checks & Balances

No single entity can unilaterally modify protocol rules.

2.4 Decentralization Roadmap

The system is designed to become progressively more decentralized over the first decade.

2.5 Skin-in-the-Game

Voting power tied to time-locked tokens ensures long-term alignment.

3. The DAO

The DAO governs:

treasury allocations

community grants

validator incentive parameters

PoUW marketplace fees

protocol upgrades (non-critical)

operational parameters

AIDA configuration boundaries

DAO votes follow:

One token = proportional voting power

Time-locked vesting tokens count as voting tokens

Delegation allowed

Vote types:

Type	Description	Requirements
Normal vote	Treasury, grants, fees, operational changes	Simple majority
Parameter vote	Non-critical system parameters	55% majority
Critical vote	Affects core protocol logic	DAO + Founder Council approval

The DAO cannot:

modify emission schedule

change maximum supply

alter consensus-critical rules

override the Founder Council during the 10-year window

4. Founder Council (10-Year Strategic Oversight)

The Founder Council is a multi-signature governance body with strictly limited authority over critical protocol decisions.

4.1 Duration & Mandate

The Founder Council operates for 10 years from genesis (Year 0 → Year 10).
After Year 10, powers automatically dissolve unless renewed by DAO supermajority.

Purpose:

Protect architectural integrity

Prevent governance attacks

Ensure safe evolution during early years

4.2 Structure (Multi-Sig, 2/3 Threshold)

The Founder Council is composed of:

Founder of Mbongo Chain (permanent seat for 10 years)

Independent Technical Expert (appointed by Foundation)

Core Contributor Representative (elected by DAO)

Approval threshold:

2 out of 3 signatures required.

4.3 Critical Scope

Founder Council approval is required for:

A. Economic Parameters

Emission schedule

Halving interval

Block reward formula

Maximum supply (31,536,000 MBG)

PoS/PoUW split ratio

AIDA regulator parameter range

B. Consensus & State Machine

Consensus rule changes

Finality rules

Slashing parameters

Hard forks

State machine invariants

C. Governance Layer

Core governance module upgrades

Critical governance rule changes

Emergency shutdown rules

4.4 Strict Limitations

The Founder Council cannot:

Mint tokens

Modify the total supply cap

Shorten vesting schedules

Override DAO votes

Control treasury spending

Freeze or seize any user funds

Influence marketplace allocations

Change PoUW reputation rules

This prevents abuse while ensuring protection.

5. Governance Safety Review Window

All Critical Protocol Changes undergo a 90-day public review process:

Publication on governance forums

Public discussion period

Foundation risk assessment

AIDA simulation report

DAO vote

Founder Council approval

On-chain timelock (14 days)

Activation

This guarantees:

transparency

informed decision-making

economic safety

6. Mbongo Foundation

The Mbongo Foundation is an independent non-profit supporting:

research

grants

developer programs

audits

security teams

ecosystem adoption

The founder holds a permanent board seat for the first 10 years.

The Foundation cannot override DAO decisions.

7. AIDA Regulator (AI-Assisted Governance)

AIDA is a bounded AI-assisted governance component.

7.1 Responsibilities

Analyze network health

Estimate economic trends

Simulate parameter changes

Provide risk assessments

Adjust burn-rate inside allowed limits:

min_burn_rate = 0%

max_burn_rate = 30%

Provide “early warning” alerts

7.2 Limitations

AIDA cannot:

mint or burn supply outside fees

bypass DAO or Founder Council

modify consensus rules

alter vesting

AIDA is an advisor + regulator under strict governance boundaries.

8. Upgrade Process
Step	Action
1	Proposal published
2	90-day public review
3	Foundation risk analysis
4	AIDA simulations
5	DAO vote
6	Founder Council multi-sig (if critical)
7	On-chain timelock
8	Activation

Critical upgrades require:

DAO supermajority (67%)

Founder Council approval (2/3)

Non-critical upgrades require:

simple DAO majority

9. Decentralization Roadmap
Phase	Years	Description
Phase 1	0–3	Founder-led stability, high oversight
Phase 2	3–7	DAO expansion, Foundation autonomy
Phase 3	7–10	Shared governance, maturity stage
Phase 4	10+	Full decentralization; Founder Council dissolves
10. Summary

Mbongo Chain governance is built to be:

secure

transparent

investor-friendly

technical

stable

decentralized over time

With:

DAO → daily governance

Founder Council → strategic protection (10 years)

Foundation → ecosystem development

AIDA → AI-assisted regulation

90-day safety window → transparency

Multi-sig checks → no centralization

This governance model is particularly well suited for a compute-first, AI-native L1.