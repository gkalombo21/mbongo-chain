(Canonical — Mbongo Chain Economic Regulator Specification)

# AIDA Regulator Specification
Status: Canonical  
Version: v1.0

AIDA (Autonomous Intelligent Dynamic Adjuster) is the AI-assisted economic regulator of Mbongo Chain.  
It governs *fee burn*, *network economic stability*, *PoUW marketplace parameters*, and *long-term monetary balance* — but always under **strict governance limits** and **DAO oversight**.

AIDA **cannot** modify consensus, mint tokens, or change the supply cap.  
It is a bounded, supervised intelligence module — *not* an autonomous economic agent.

This specification defines AIDA’s responsibilities, constraints, algorithms, safety guarantees, and governance boundaries.

---

# 1. Purpose

AIDA exists to:

- maintain long-term network economic health  
- adjust burn rate dynamically based on demand  
- stabilize fees during high-load periods  
- discourage spam and abusive compute requests  
- align PoUW rewards with actual supply/demand  
- detect anomalies (economic or traffic-related)  
- provide risk forecasts for governance proposals  
- run economic simulations for major upgrades  
- protect the network against governance attacks  
- ensure predictable monetary behavior  

AIDA is an **advisory + regulatory module**, not a monetary controller.

---

# 2. Design Principles

AIDA follows these principles:

### **2.1 Bounded Autonomy**
AIDA can adjust parameters only within strict min/max ranges.

### **2.2 Governance Supervision**
DAO and Founder Council must approve any expansion of AIDA’s control envelope.

### **2.3 No Monetary Issuance**
AIDA can never:
- mint tokens  
- increase supply  
- accelerate emission  
- cancel vesting  

### **2.4 Transparency**
All AIDA actions are:
- logged
- auditable
- signed by the AIDA module
- published on-chain

### **2.5 Determinism**
Although AIDA uses AI-based predictions,  
**all on-chain actions must be deterministic**.

This means:

- off-chain ML models may advise  
- on-chain enforcement uses rule-based logic  
- final parameter updates follow deterministic formulas

---

# 3. Responsibilities

AIDA is responsible for four economic domains:

---

## 3.1 Dynamic Fee Burn Regulation

AIDA adjusts the **burn_rate** within the governance-approved range:



0% ≤ burn_rate ≤ 30%


Purpose:

- burn more when network demand is high  
- burn less when network slows down  
- encourage economic equilibrium  
- maintain scarcity aligned with activity  

AIDA uses:

- block fill ratio  
- PoUW job queue depth  
- mempool pressure  
- transaction gas statistics  
- AI-predicted demand curves  

But the on-chain decision is deterministic.

Formula (simplified):

```text
if demand > high_threshold:
    burn_rate = min(burn_rate + α, max_burn_rate)
else if demand < low_threshold:
    burn_rate = max(burn_rate - β, min_burn_rate)


Where α and β are small step functions.

3.2 Fee Stabilization (Anti-Volatility)

AIDA stabilizes fees by modifying:

base_fee_multiplier

priority_fee_limit

PoUW job pricing coefficients

But must remain in bounds defined by DAO:

base_fee_multiplier ∈ [0.5 , 3.0]
priority_fee_limit  ∈ [0 , 2.0]


AIDA tries to:

reduce fee volatility

keep compute affordable

avoid spikes caused by one-off events

ensure consistent user experience

3.3 PoUW Marketplace Balancing

AIDA monitors:

compute supply (available GPUs)

compute demand (tasks pending)

completion rates

job failure patterns

hardware reliability

worker reputation scores

If supply < demand → increase PoUW reward multiplier
If supply > demand → reduce multiplier

Boundaries:

PoUW_multiplier ∈ [0.8 , 1.2]


This ensures:

stable compute pricing

fair market equilibrium

predictable earnings for compute providers

3.4 Economic Risk Forecasting

AIDA evaluates:

emission curve health

staking participation trends

PoUW profitability

DAO treasury viability

systemic risks (spam, overload, economic attacks)

AIDA provides:

forecasts

alerts

dashboards

simulation results

AIDA does not enforce decisions here — only informs DAO.

4. Governance Boundaries

AIDA is NOT autonomous.

AIDA requires DAO approval for:

expanding any parameter range

upgrading AIDA logic

deploying new AI models

modifying weight coefficients

adding new PoUW job classes

modifying marketplace rules

AIDA requires Founder Council approval for:

any change affecting:

burn limits

emission-related parameters

PoUW reward split allocations

consensus-critical fee formulas

macroeconomic primitives (gas system, resource pricing)

For 10 years, the Founder Council acts as:

a safety validator

a protection layer

an anti-attack mechanism

5. Safety Guarantees

AIDA includes strong guardrails:

5.1 Immutable Supply

AIDA cannot:

mint

burn supply outside fees

modify emission schedule

5.2 Deterministic Control

Only the decision logic is on-chain.
ML models only produce external advisory signals.

5.3 Rate Limiting

AIDA can adjust parameters at most once every N blocks:

N = 300 blocks (5 minutes)

5.4 Undo Protection

If an abnormal adjustment is detected:

AIDA auto-reverts to previous safe state

Validator nodes reject conflicting updates

5.5 Circuit Breaker

If AIDA behaves erratically:

the DAO can freeze AIDA adjustments

Founder Council can revoke adjustments for 10 years

an emergency safe-mode is activated

6. On-Chain Decision Mechanism

AIDA uses a deterministic decision engine:

Inputs:

block statistics

gas price oracle

PoUW queue metrics

mempool pressure

validator participation

stake ratio

historical volatility

compute supply/demand ratios

Output:

updated burn rate

updated base fee multiplier

updated PoUW multiplier

All updates are:

signed by AIDA

included in block headers

logged via AIDAUpdate event

7. Off-Chain Intelligence Layer (Optional)

AIDA’s “AI brain” may run off-chain:

large models

ML predictors

anomaly detectors

demand forecasting

compute load prediction

This layer is advisory only.

On-chain AIDA receives bounded signals, such as:

“expected demand = medium-high”

“risk = low”

“network pressure = rising”

But the final update is strictly rule-based.

8. Event Logs

AIDA emits:

AIDAUpdate(burn_rate, base_fee, pouw_multiplier, block)

AIDAWarn(message)

AIDARisk(level, category)

AIDARevert(previous_state)

AIDAFrozen()

AIDAResumed()

9. Future Extensions

per-job-class pricing

AI-powered block production optimization

cross-chain predictive pricing

ZK-verified simulation models

AIDA governance dashboard (Next.js + TS SDK)

long-term treasury optimization

All upgrades require DAO and/or Founder Council approval.

10. Summary

AIDA is a bounded, supervised, deterministic regulator that:

adjusts burn rate dynamically

stabilizes fees

balances PoUW market supply/demand

forecasts economic risks

protects against manipulation

operates under strict governance limits

cannot mint, print, or alter supply

AIDA ensures Mbongo Chain remains:

economically stable

predictable

resistant to attacks

aligned with real-world compute demand

sustainable for decades

This document represents the canonical specification of AIDA.