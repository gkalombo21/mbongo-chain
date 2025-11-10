# ⚛️ Quantum Risk & Post-Quantum Strategy

## 1️⃣ Context — Why Prepare Now
- Quantum computers threaten classical cryptography (ECDSA, RSA, secp256k1) via Shor’s algorithm, enabling rapid discrete log and integer factorisation.
- Mbongo-Chain relies on digital signatures and hashing for consensus, staking, and banking modules; a quantum-capable adversary could forge validator identities or rewrite history.
- Anticipatory planning protects long-term token value, institutional trust, and regulatory alignment for AI-driven financial services.

## 2️⃣ Current Situation — Reality Check
- No publicly known quantum computer can break 256-bit elliptic curves or SHA-3 security today.
- Experts project viable cryptanalytic machines within 10–15 years, aligning with NIST and industry roadmaps.
- Adversaries may already harvest encrypted traffic for “store now, decrypt later” attacks, increasing urgency for migration.

## 3️⃣ Our Strategy — Three-Phase Migration
- **Phase I · Initial Hardening (Now–2026)**  
  - Prioritise Ed25519 adoption wherever compatible to reduce legacy RSA usage.  
  - Introduce hybrid signature modes (Ed25519 + PQC) in non-critical modules for experimentation.  
  - Build cryptographic abstraction layers in `internal/` to swap algorithms without core rewrites.
- **Phase II · Transition (2026–2029)**  
  - Pilot CRYSTALS-Dilithium and Falcon for validator keys and inter-node authentication.  
  - Offer SPHINCS+ as stateless fallback for resistant environments.  
  - Launch migration toolchain: dual-signature transactions, backwards-compatible wallets, and PQ-aware key ceremonies.  
  - Run incentivised testnets validating performance, storage, and network overhead.
- **Phase III · Post-Quantum Defaults (2029 onward)**  
  - Make PQC signatures mandatory for validators, compute nodes, and banking operations.  
  - Deprecate classical-only keys with automated revocation windows.  
  - Embed PQC in consensus protocol, ensuring finality and slashing rules remain deterministic under larger key sizes.

## 4️⃣ Governance & Upgrades — Coordinated Change
- Mbongo-Chain governance will ratify cryptographic upgrades through community votes requiring supermajority thresholds.  
- Emergency “quantum fork” procedures define rollback, snapshot, and replay policies if zero-day quantum breaks occur.  
- Upgrade milestones ship as versioned network proposals with on-chain documentation, security audits, and phased activation.  
- Dedicated Quantum Risk Council evaluates external intelligence, NIST updates, and cross-chain interoperability requirements.

## 5️⃣ Key Takeaways — Quantum Readiness
- Proactive roadmap shields validator integrity, AI workloads, and user assets from future quantum adversaries.  
- Hybrid signature support ensures compatibility today while easing migration costs tomorrow.  
- Governance-driven rollouts align technical resilience with community consensus and regulatory expectations.

## 6️⃣ References — Further Reading
- NIST Post-Quantum Cryptography Program (Round 3 Reports, 2023)  
- Deloitte Insights: “Decrypting the Post-Quantum Era”  
- Google Willow Team: “Practical Quantum-Resistant Cryptography Deployments”  
- Openware Research: “Hybrid PQC Strategies for Financial Blockchains”  
- Mbongo-Chain Architecture Docs & Roadmap (internal)

