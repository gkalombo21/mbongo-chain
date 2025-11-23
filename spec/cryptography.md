# Mbongo Chain — Cryptography, Keys, Signatures & Randomness

This document defines the cryptographic primitives used by Mbongo Chain.  
The design is optimized for security, performance, and long-term decentralization.

---

# 1. Overview of Cryptographic Philosophy

Mbongo Chain uses **modern, battle-tested primitives** with the following priorities:

- **Security** (first-class priority)
- **Post-quantum awareness**
- **Fast verification** for 1-second blocks
- **Compatibility** with hardware (GPU/TPU/ASIC)
- **Deterministic randomness** for consensus

The system avoids outdated mechanisms such as ECDSA in favor of faster and more secure options.

---

# 2. Key Cryptographic Components

| Component | Algorithm | Purpose |
|----------|-----------|---------|
| Signature Scheme | **Ed25519** | Fast block signatures & TX signatures |
| Hash Function | **BLAKE3** | Merkle trees, hashing, commitments |
| VRF | **Ristretto VRF** | Validator selection, PoS randomness |
| PoUW Proof Hashing | **BLAKE3 / Poseidon** | Hashing inside SNARK circuits |
| Address Derivation | **SHA3-256 → bech32** | Human-friendly, checksum-safe |
| ZK-Friendly Circuits | **Poseidon + PLONK** | Optional proofs inside PoUW |

---

# 3. Signature System

Mbongo Chain uses **Ed25519** for all signatures.

Benefits:
- Very fast
- Small signatures
- Battle-tested (Solana, Polkadot, Cosmos SDK, Tendermint, etc.)
- Resistant to common side-channel attacks
- Easy to implement in Rust (`ed25519-dalek`)

### Signature Format

signature = Ed25519.sign(private_key, message)


Messages are always hashed before signing:

message_hash = BLAKE3(message)
signature = Ed25519.sign(private_key, message_hash)


---

# 4. Key Types

## 4.1 Account Keys
- Used for sending transactions
- Ed25519 private/public pairs
- Derived using SHA3 + bech32 encoding

## 4.2 Validator Keys
Each validator holds:

| Key | Purpose |
|-----|---------|
| **Consensus Key** | Signing block proposals |
| **Network Key** | libp2p networking |
| **VRF Key** | Randomness and leader selection |

Validators must securely store these keys offline or in HSM if possible.

---

# 5. Hashing System

### Primary Hash Function: **BLAKE3**

Reasons:

- Extremely fast (GPU-friendly)
- Low CPU load → ideal for 1-second blocks
- Perfect for Merkle trees
- Secure and widely adopted

### Hash format:

hash = BLAKE3(input_bytes)


Used for:
- Block headers
- Transaction hashes
- Merkle proofs
- PoUW proof commitments
- State root hashing

---

# 6. Merkle Tree Hashing

Mbongo Chain uses **Sparse Merkle Trees** with:

- **BLAKE3** as node hash
- Deterministic sparse-layout
- Proof compression for efficient RPC

Leaf hashing:

leaf = BLAKE3(address || value)


Internal node hashing:

node = BLAKE3(left_child || right_child)


---

# 7. VRF (Verifiable Random Function)

Mbongo Chain uses **Ristretto255 VRF**.

Uses:
- Randomness for epochs
- Validator slot selection
- PoUW challenge randomness

Benefits:
- Non-inflatable randomness
- Bias-resistant
- Verifiable on-chain

### VRF Output Format:

vrf_output = VRF.prove(secret_key, input)
vrf_proof = VRF.proof
vrf_verify(public_key, input, proof)


---

# 8. Randomness System

Every epoch generates a **Randomness Beacon** using:

1. Aggregated VRF outputs  
2. BLAKE3 hashing  
3. Anti-bias protection  

Formula:

epoch_randomness = BLAKE3(vrf_output_1 || vrf_output_2 || ... || vrf_output_n)


If a validator withholds VRF → slashing event.

This ensures:
- Unpredictable randomness
- No validator can influence selection
- Civil-resistance for PoUW tasks

---

# 9. Address Format

Addresses follow:

address = bech32_encode("mbo", SHA3-256(pubkey)[:20])


Benefits:
- Human-readable
- Error-detection built in
- Same style as Cosmos / Celestia

---

# 10. PoUW Cryptography

PoUW tasks follow two possibilities:

## 10.1 Non-ZK-secured tasks (MVP)
Compute nodes deliver:

- input hash  
- output hash  
- execution metadata  

Validated deterministically by the verification subsystem.

## 10.2 ZK-secured tasks (future upgrade)
For tasks that require cryptographic verification:

### Supported ZK systems:
- **PLONK**
- **Groth16**
- **Halo2**
- **STARKs** (in future)

### Hash inside circuits:
- **Poseidon Hash** → ZK-friendly
- **BLAKE3** used on-chain

This approach balances performance and future-proofing.

---

# 11. Security Rules

### Signature Rules
- All signatures must be Ed25519
- Nonces used to prevent replay
- Blocks require aggregated signatures (BLS optional future)

### Hashing Rules
- All state must use BLAKE3
- No SHA-256 unless necessary for compatibility
- Inputs must be canonicalized before hashing

### VRF Rules
- VRF output must be used as-is
- No custom transformations (prevents bias)
- Validators must reveal VRF output each slot

---

# 12. Quantum Considerations

Mbongo Chain is **quantum-aware**, future-ready:

Today:
- Ed25519 (practical and performant)

Future upgrade path:
- Ed448
- Dilithium (PQ signature)
- Falcon2 (PQ fast verification)

A flag for PQ migration exists in the protocol upgrade rules.

---

# 13. Summary

The cryptographic stack of Mbongo Chain is built for:

- High throughput  
- Strong security  
- Scalable verification  
- Fast hashing (BLAKE3)  
- Robust randomness (VRF-based)  
- Optional ZK-friendly computation  

This provides a future-proof foundation for the entire ecosystem, from consensus to PoUW compute markets.
