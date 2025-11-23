# Mbongo Chain — Infrastructure & Deployment Guide
Status: Canonical  
Version: v1.0

This document describes how to deploy, operate, and maintain nodes in the Mbongo Chain network.  
It covers:

- validator nodes
- PoUW compute nodes
- RPC nodes
- indexers
- monitoring and observability
- recommended hardware profiles
- deployment workflows (devnet, testnet, mainnet)

The infrastructure assumes:

- a Rust-based monorepo
- libp2p networking
- RocksDB storage
- JSON-RPC + WebSocket + gRPC APIs
- PoS + PoUW hybrid consensus

---

# 1. Node Types

Mbongo Chain operates with four primary node types:

| Node Type | Purpose |
|-----------|---------|
| **Validator Node** | Produces blocks, participates in PoS, validates PoUW receipts |
| **PoUW Compute Node** | Executes GPU/CPU tasks, generates verifiable receipts |
| **RPC Node** | Exposes public RPC endpoints (JSON-RPC, WebSocket, gRPC) |
| **Indexer Node** | Processes chain data for explorers, dashboards, analytics |

Each node can run independently but benefits from shared infrastructure.

---

# 2. System Requirements

## 2.1 General Requirements (all nodes)

- **OS:** Linux (Ubuntu 22.04 recommended)  
- **CPU:** 4+ cores  
- **RAM:** 8–16 GB  
- **Storage:**  
  - NVMe SSD strongly recommended  
  - Minimum: 200 GB  
- **Network:**  
  - stable low-latency connection  
  - 10 Mbps minimum (up/down)

---

## 2.2 Validator Nodes — Recommended Hardware

Validators have strict uptime and performance requirements.

| Component | Recommendation |
|-----------|----------------|
| CPU | 8+ cores (Ryzen 7 / Xeon) |
| RAM | 32 GB |
| Disk | 1 TB NVMe SSD |
| Network | 100 Mbps symmetric |
| GPU | **Not required** |

Validators verify PoUW results but do *not* execute GPU tasks.

---

## 2.3 PoUW Compute Nodes — Recommended Hardware

PoUW compute nodes perform GPU-accelerated tasks.

### **Profile A — Standard GPU Node**
- GPU: NVIDIA RTX 3080 / 3090 / 4080  
- VRAM: 10–24 GB  
- CPU: 8 cores  
- RAM: 32 GB  
- Storage: 500 GB SSD  
- OS: Ubuntu + CUDA 12.x

### **Profile B — AI-Optimized Node**
- GPU: NVIDIA A100 / H100 or AMD MI210  
- VRAM: 40–80 GB  
- CPU: 16+ cores  
- RAM: 64–128 GB  
- Network: 1 Gbps  
- Use case: heavy AI inference, batch jobs, ZK proving

### **Profile C — CPU-Only Node** (MVP)
- No GPU required  
- Executes light PoUW jobs  
- Validates compute receipts  

Used for minimum-spec PoUW participation.

---

# 3. Installation & Build

All nodes are built from the monorepo:

```bash
git clone https://github.com/mbongo-chain/mbongo-chain
cd mbongo-chain
cargo build --release


Output binaries:

mbongo-node

mbongo-wallet

mbongo-pouw

mbongo-indexer

4. Validator Deployment
4.1 Generate Keys
./target/release/mbongo-wallet keygen --output validator.json

4.2 Create Validator Config
./target/release/mbongo-node init \
  --validator \
  --key validator.json \
  --data-dir ~/.mbongo/validator

4.3 Join Network
./target/release/mbongo-node start \
  --config ~/.mbongo/validator/config.toml


Validators:

participate in PoS

verify PoUW receipts

maintain consensus

produce blocks when selected via VRF

5. PoUW Compute Node Deployment
5.1 GPU Drivers & Dependencies

For NVIDIA:

sudo apt install nvidia-driver-550
sudo apt install cuda-toolkit-12-3


For AMD:

sudo apt install rocm-dev

5.2 Initialize
./target/release/mbongo-pouw init \
  --data-dir ~/.mbongo/pouw

5.3 Start Compute Worker
./target/release/mbongo-pouw start \
  --gpu \
  --jobs ai,rendering,zk \
  --stake <amount> \
  --rpc http://validator:8545


PoUW nodes:

request compute tasks

execute them

generate receipts

submit proofs to validators

6. RPC Node Deployment

RPC nodes expose public endpoints for:

wallets

dApps

explorers

indexers

Start an RPC node:

./target/release/mbongo-node start \
  --rpc \
  --ws \
  --grpc \
  --data-dir ~/.mbongo/rpc


External ports:

Protocol	Port	Description
HTTP JSON-RPC	8545	Standard RPC
WebSocket	8546	Realtime events
gRPC	50051	High-performance compute & infra
7. Indexer Deployment

Indexers ingest chain data and populate:

explorers

dashboards

compute stats

PoUW verification metrics

Start indexer:

./target/release/mbongo-indexer start \
  --rpc http://localhost:8545 \
  --db postgres://...


Indexer storage is external (PostgreSQL, ElasticSearch).

8. Monitoring & Observability
8.1 Metrics

All nodes expose Prometheus endpoints:

http://host:9600/metrics


Metrics include:

block time

peer count

PoUW throughput

CPU/GPU utilization

VRAM usage

consensus latency

8.2 Logging

Structured JSON logs:

./mbongo-node --log-format json


Recommended stack:

Loki

Grafana

Promtail

OpenTelemetry (future)

9. Deployment Environments
Network	Description
Devnet	Fast iteration for developers
Testnet	Public, resets allowed, faucets available
Mainnet	Production, immutable, secure
10. Best Practices

run validators on isolated dedicated servers

run PoUW compute nodes separately from validator nodes

keep GPU drivers and CUDA updated

monitor GPU thermals during heavy compute

use SSDs only (no HDD)

configure automatic restarts

protect RPC nodes behind rate-limiters

enable firewall rules on all nodes

11. Future Upgrades

auto-scaling PoUW clusters

multi-GPU orchestration

subnet-based compute sharding

remote attestation (TEE support)

ZK-enabled PoUW verification

Summary

This document defines:

how to deploy Mbongo validator nodes

how to deploy PoUW compute nodes

supported hardware configurations

system requirements

RPC/indexer guidelines

observability stack

security best practices

It is the canonical infrastructure reference for the Mbongo Chain ecosystem.