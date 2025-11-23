
# 1. Overview

Each Mbongo Chain node runs a unified executable composed of multiple modules:

- **Networking Layer (libp2p)**
- **Consensus Engine (PoS + PoUW merge)**
- **Block Builder**
- **Execution Engine (State Machine)**
- **Storage Engine**
- **Transaction Pool**
- **PoUW Compute Engine (optional)**
- **RPC & API Layer**
- **Telemetry + Metrics**

Nodes differ by activated modules:

| Node Type | Activated Modules |
|----------|-------------------|
| **Validator** | Consensus + Networking + Execution + Storage + Block Builder |
| **Full Node** | Networking + Execution + Storage + TxPool |
| **Compute Node** | Networking + PoUW Engine + Light Execution |

---

# 2. High-Level Architecture Diagram



+-------------------------------------------------------+
| Mbongo Node |
+-------------------------------------------------------+
| Networking (libp2p) |
| - GossipSub |
| - Peer Discovery (Kademlia) |
| - Encrypted Transport (Noise) |
+-------------------------------------------------------+
| Consensus Engine |
| - PoS Validator |
| - PoUW Verifier |
| - Slot scheduler |
+-------------------------------------------------------+
| Block Builder |
| - Tx selection |
| - Proof merging |
| - State root calculation |
+-------------------------------------------------------+
| Execution Engine |
| - WASM virtual machine |
| - Transaction execution |
| - Smart contract runtime |
+-------------------------------------------------------+
| PoUW Compute Engine (optional) |
| - GPU/CPU/TPU tasks |
| - Proof generator |
| - Challenge-response verifier |
+-------------------------------------------------------+
| Storage Engine |
| - Hot DB (RocksDB) |
| - Cold Archive storage (optional) |
+-------------------------------------------------------+
| RPC / API Layer |
| - JSON-RPC |
| - gRPC |
| - WebSocket |
+-------------------------------------------------------+
| Telemetry |
| - Prometheus metrics |
| - Logs & diagnostics |
+-------------------------------------------------------+


---

# 3. Node Types

## 3.1 Validator Node

Validators are responsible for:

- Block proposals  
- Voting in consensus  
- Verifying PoUW proofs  
- Maintaining full state  
- Participating in finality  

### Requirements:
- Full DB storage  
- High uptime (≥ 95%)  
- 250 MBG minimum stake  
- Secure key storage (HSM recommended)  

---

## 3.2 Full Node

Full nodes:

- Execute all transactions  
- Store full chain state  
- Relay gossip messages  
- Validate PoUW proofs (optional)  

### Characteristics:
- No staking required  
- Helps decentralization  
- Can serve RPC for users and dApps  

---

## 3.3 Compute Node (PoUW Worker)

Compute nodes:

- Pull AI/compute tasks  
- Execute GPU/CPU/TPU jobs  
- Produce **Validatable Work Proofs (VWP)**  
- Submit results to block producers  

### Features:

- Does **not** run full chain state  
- Uses **light client** mode  
- Requires compute hardware (GPU/TPU/etc.)  
- Reputation score impacts rewards  

---

# 4. Slot Timing & Responsibilities

Mbongo Chain uses **1-second blocks**.

Each second:

1. A validator is selected as block proposer  
2. A PoUW task is issued to the compute network  
3. Compute nodes return proofs  
4. Validator merges proofs and produces block  
5. Finality sub-round ensures confirmation (2–4 seconds)

---

# 5. Transaction Pipeline



User → TxPool → Block Builder → Execution Engine → Storage → Gossip


### Steps:

1. **Incoming Tx** validated cryptographically  
2. Added to **Transaction Pool**  
3. Block builder selects highest-fee or highest-weight tx  
4. Execution engine runs WASM contract  
5. State root calculated  
6. Block committed & propagated  

---

# 6. Execution Engine (WASM)

Mbongo Chain uses a **WASM-based VM** for smart contracts:

- Deterministic execution  
- Language agnostic  
- High performance  
- Sandboxed environment  

Supported languages (compiled to WASM):
- Rust  
- AssemblyScript  
- Go (via WASM compile)  

---

# 7. Storage Engine

Mbongo uses:

### **RocksDB** as Hot Storage  
Stores:
- state  
- receipts  
- PoUW proofs  
- block metadata  

### **Cold Storage** (optional)
- Archive nodes  
- Snapshots  
- Historical data  

---

# 8. RPC & API Layer

Nodes expose:

- **JSON-RPC** for wallets  
- **WebSocket** for real-time data  
- **gRPC** for high-performance apps  
- **REST API** for explorers  

Example endpoints:
- `/getBlock`  
- `/getBalance`  
- `/submitProof`  
- `/taskQueue`  

---

# 9. Telemetry, Monitoring & Observability

Metrics provided through:

### Prometheus:
- block time  
- slot misses  
- peer count  
- memory usage  
- PoUW proof rate  

### Logs:
- JSON structured logs  
- Rotating logs  
- Optional cloud export  

---

# 10. Security Considerations

- Key isolation (HSM/TPM recommended)  
- Anti-DoS checks on gossip messages  
- Rate limits for PoUW submissions  
- Validator private key never leaves memory  
- Encrypted P2P channels (Noise Protocol)  

---

# 11. Summary

The Mbongo Chain node is:

- modular  
- high-performance  
- secure  
- AI-native  
- optimized for global decentralized compute  

This architecture allows:
- Validators to secure the network  
- Full nodes to maintain decentralization  
- Compute nodes to provide useful AI work  
- dApps to interact through a clean API  

This completes the node-level specification needed for developers to understand and build Mbongo Chain.
