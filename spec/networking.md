
# **networking.md — Mbongo Chain Networking Layer Specification (Final Version)**

---

# **Mbongo Chain — Networking Layer Specification**

*Final, cleaned, consolidated version.*

---

# 1. **Overview**

The networking layer enables Mbongo Chain nodes to:

* Discover and connect to peers
* Exchange blocks and transactions
* Broadcast PoUW compute proofs
* Maintain real-time synchronization
* Ensure fast finality (1s block time)
* Operate globally, even under high churn or unreliable connections

The protocol is built on **libp2p**, chosen for:

* Modularity
* Battle-tested performance (ETH2, Polkadot)
* NAT traversal
* Peer discovery
* Pub/sub messaging
* Secure channels

---

# 2. **Core Networking Responsibilities**

The networking layer manages:

* Peer-to-peer connectivity
* Gossip and broadcast protocols
* Block propagation
* Mempool synchronization
* Compute proof exchange
* Peer scoring
* Anti-spam protection
* Latency-optimized routing

It is optimized for:

* **sub-second propagation**
* **global distribution**
* **high throughput workloads**

---

# 3. **Network Architecture**

The network consists of:

* **Validator nodes**
* **Full nodes**
* **Compute nodes (PoUW)**
* **Light clients**

All nodes use libp2p but with different capabilities and bandwidth priorities.

### Node Capabilities Table

| Node Type    | Block Sync | Mempool  | PoS | PoUW     | Gossip  | Storage |
| ------------ | ---------- | -------- | --- | -------- | ------- | ------- |
| Validator    | ✔          | ✔        | ✔   | Optional | High    | Full    |
| Full Node    | ✔          | ✔        | ✘   | Optional | Medium  | Full    |
| Compute Node | Optional   | Optional | ✘   | ✔        | Low     | Minimal |
| Light Client | Partial    | ✘        | ✘   | ✘        | Minimal | None    |

---

# 4. **Protocols Used**

Mbongo Chain uses the following libp2p components:

### ✔ **GossipSub v1.1**

For:

* block broadcast
* transaction gossip
* PoUW proof distribution

### ✔ **mDNS + Kademlia DHT**

For:

* peer discovery
* bootstrapping

### ✔ **Noise Protocol**

For:

* encrypted channels
* authentication

### ✔ **QUIC (preferred) / TCP fallback**

For:

* low-latency data transmission

### ✔ **Request/Response Protocol**

For:

* block syncing
* state queries
* compute-task fetch

---

# 5. **Messaging Channels**

The networking system uses pub/sub topics:

| Topic               | Purpose                       |
| ------------------- | ----------------------------- |
| `/mbongo/blocks`    | block propagation             |
| `/mbongo/tx`        | transaction gossip            |
| `/mbongo/pouw`      | compute proof distribution    |
| `/mbongo/consensus` | PoS coordination              |
| `/mbongo/alerts`    | slashing & misbehavior alerts |

All messages include:

* signature
* timestamp
* anti-replay nonce

---

# 6. **Block Propagation Pipeline**

1. Node receives new block from proposer
2. Validate header and signatures
3. Validate hash-chain
4. Broadcast to peers via GossipSub
5. Peers validate block
6. Block enters local fork-choice
7. Block passed to state machine for execution

Propagation target:

* **< 350 ms worldwide**

---

# 7. **Mempool Synchronization**

Mempools synchronize through:

* gossip-based announcements
* direct transaction pull
* collision detection
* duplicate filtering

Rules include:

* max mempool size
* replacement fee policy
* anti-spam scoring

---

# 8. **Peer Scoring System**

Nodes score peers based on:

* uptime
* valid blocks relayed
* invalid blocks relayed
* bandwidth usage
* compute-proof reputation (for compute nodes)
* latency

Low-scoring peers are:

* throttled
* deprioritized
* eventually dropped

---

# 9. **Compute Node Networking (PoUW)**

Compute nodes have a simplified networking profile:

* Low gossip priority
* No block production
* Specialized channel for submitting VWP proofs
* Challenge-response round for anti-simulation checks

Compute nodes may be globally distributed, including:

* consumer GPUs
* cloud GPUs
* TPU/NPU devices
* IoT / edge compute clusters

Networking is optimized to support **millions of compute nodes**.

---

# 10. **Synchronization Modes**

Mbongo Chain supports:

### **Fast Sync**

Downloads recent blocks + state snapshot
~ recommended for new nodes

### **Full Sync**

Replays all blocks since genesis
~ recommended for auditors/archives

### **Light Sync**

Downloads only block headers + merkle proofs
~ recommended for mobile/wallets

---

# 11. **Anti-Spam & Security**

Networking includes:

* rate limiting
* proto-flood control
* invalid-block penalty
* signature verification at ingress
* duplicate proof filtering
* DoS protection
* IP/subnet scoring
* challenge-based verification for PoUW tasks

---

# 12. **Bootstrapping**

New nodes use:

* trusted bootstrap peers
* DHT crawling
* gossip warm-up
* peer reputation seeding

Bootstrap nodes:

* run redundant infrastructure
* geographically distributed
* signed by the Foundation initially
* replaced gradually by community nodes

---

# 13. **Network Upgradeability**

Changes to networking must:

* maintain backward compatibility
* preserve topology
* avoid partition risks
* go through governance
* provide fallback shims

Runtime handles protocol version tags.

---

# 14. **Testing Requirements**

Each part of the networking stack must include:

* unit tests
* performance benchmarks
* adversarial network simulation
* outage simulation
* long-distance latency tests
* gossip storm stress tests
* PoUW flood resistance tests

Goal:
**robustness under real-world global conditions**

---

# 15. **Summary**

The Mbongo Chain networking layer is built for:

* High-throughput block propagation
* Low-latency compute proof distribution
* Millions of compute nodes
* Strong security and anti-spam
* Seamless global scalability

It forms the backbone that keeps consensus synchronized across the world.

---
