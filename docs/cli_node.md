# Mbongo Chain — CLI Node Commands

> **Document Type:** CLI Reference  
> **Last Updated:** November 2025  
> **Status:** Official Reference  
> **Parent:** [cli_overview.md](./cli_overview.md)

---

## Table of Contents

1. [Purpose of Node Commands](#1-purpose-of-node-commands)
2. [Node Command Structure](#2-node-command-structure)
3. [Detailed Command Documentation](#3-detailed-command-documentation)
4. [Configuration Integration](#4-configuration-integration)
5. [Security Rules](#5-security-rules)
6. [Advanced Topics](#6-advanced-topics)
7. [ASCII Diagrams](#7-ascii-diagrams)

---

## 1. Purpose of Node Commands

### 1.1 What is a Mbongo Node?

A Mbongo Node is the core software component that participates in the Mbongo Chain network. The `mbongo node` CLI commands manage the lifecycle and operations of this node.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         NODE TYPES                                          │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   FULL NODE                                                                 │
│   ═════════                                                                 │
│   • Stores complete blockchain state                                       │
│   • Validates all blocks and transactions                                  │
│   • Serves RPC queries                                                     │
│   • Does NOT participate in consensus                                      │
│   • Command: mbongo node start                                             │
│                                                                             │
│   VALIDATOR NODE                                                            │
│   ══════════════                                                            │
│   • Full node + consensus participation                                    │
│   • Proposes and attests blocks                                            │
│   • Requires staked MBO (minimum 50,000)                                   │
│   • Subject to slashing                                                    │
│   • Command: mbongo node start --validator                                 │
│                                                                             │
│   GPU PROVIDER NODE (PoUW)                                                  │
│   ════════════════════════                                                  │
│   • Full node + compute task execution                                     │
│   • Submits compute receipts                                               │
│   • Requires GPU hardware                                                  │
│   • Earns PoUW rewards                                                     │
│   • Command: mbongo node start --compute-provider                          │
│                                                                             │
│   HYBRID NODE                                                               │
│   ═══════════                                                               │
│   • Validator + GPU Provider combined                                      │
│   • Maximum reward potential                                               │
│   • Command: mbongo node start --validator --compute-provider              │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 1.2 When to Use CLI vs Config Files vs RPC

| Method | Use Case | Example |
|--------|----------|---------|
| **CLI Flags** | Quick overrides, one-time changes | `--rpc-port 8546` |
| **Config File** | Persistent settings, production | `config.toml` |
| **RPC** | External queries, monitoring | `curl http://localhost:8545` |
| **Environment** | Secrets, CI/CD pipelines | `MBONGO_KEYSTORE_PASSWORD` |

**Priority Order:** CLI Flags > Environment Variables > Config File > Defaults

### 1.3 Required Environment

```bash
# Minimum requirements
- Rust binary: mbongo (compiled release build)
- Data directory: ~/.mbongo/ (or custom via --data-dir)
- Config file: ~/.mbongo/config.toml (optional)
- Port access: 30303 (P2P), 8545 (RPC), 9090 (metrics)

# Verify installation
mbongo --version
mbongo node --help
```

---

## 2. Node Command Structure

### 2.1 Syntax

```
mbongo node <command> [subcommand] [flags]
```

### 2.2 Subcommands

| Command | Description |
|---------|-------------|
| `start` | Start the node |
| `stop` | Stop the running node |
| `restart` | Restart the node |
| `status` | Show node status |
| `info` | Show node information |
| `peers` | List connected peers |
| `sync-status` | Show synchronization status |
| `purge-data` | Delete node data |
| `snapshot export` | Export state snapshot |
| `snapshot import` | Import state snapshot |
| `logs` | View/stream logs |
| `rpc` | RPC server management |
| `metrics` | Prometheus metrics |

---

## 3. Detailed Command Documentation

### 3.1 `mbongo node start`

**Description:** Start the Mbongo node daemon.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--config` | `-c` | No | `~/.mbongo/config.toml` | Config file path |
| `--data-dir` | `-d` | No | `~/.mbongo` | Data directory |
| `--validator` | | No | `false` | Enable validator mode |
| `--compute-provider` | | No | `false` | Enable GPU provider mode |
| `--keystore` | | No | `~/.mbongo/keystore` | Keystore directory |
| `--rpc-addr` | | No | `127.0.0.1` | RPC bind address |
| `--rpc-port` | | No | `8545` | RPC port |
| `--p2p-port` | | No | `30303` | P2P listen port |
| `--bootnodes` | | No | (built-in) | Bootstrap node addresses |
| `--log-level` | | No | `info` | Log verbosity |
| `--metrics` | | No | `false` | Enable Prometheus metrics |
| `--metrics-port` | | No | `9090` | Metrics port |

**Examples:**

```bash
# Basic start
mbongo node start

# Start as validator
mbongo node start --validator --keystore /secure/keystore

# Start with GPU provider
mbongo node start --compute-provider --gpu-device 0

# Start with custom config
mbongo node start -c /etc/mbongo/config.toml -d /var/lib/mbongo

# Start with verbose logging
mbongo node start --log-level debug

# Start with external RPC access (use with caution)
mbongo node start --rpc-addr 0.0.0.0 --rpc-port 8545
```

**Output (table):**

```
Node started successfully
────────────────────────────────────────────────
  Node ID      │ 16Uiu2HAmXyz...
  Mode         │ Full Node + Validator
  P2P Port     │ 30303
  RPC Endpoint │ http://127.0.0.1:8545
  Data Dir     │ /var/lib/mbongo
  PID          │ 12345
────────────────────────────────────────────────
```

---

### 3.2 `mbongo node stop`

**Description:** Gracefully stop the running node.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--force` | `-f` | No | `false` | Force kill (SIGKILL) |
| `--timeout` | | No | `30` | Shutdown timeout (seconds) |

**Examples:**

```bash
# Graceful stop
mbongo node stop

# Force stop
mbongo node stop --force

# Stop with custom timeout
mbongo node stop --timeout 60
```

**Output:**

```
Stopping node (PID: 12345)...
Flushing state to disk...
Disconnecting peers...
Node stopped successfully.
```

---

### 3.3 `mbongo node restart`

**Description:** Restart the node (stop + start).

**Flags:** Inherits all flags from `start` plus `stop` flags.

**Examples:**

```bash
# Simple restart
mbongo node restart

# Restart with new config
mbongo node restart -c /etc/mbongo/new-config.toml

# Force restart
mbongo node restart --force
```

---

### 3.4 `mbongo node status`

**Description:** Show current node status and health.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--output` | `-o` | No | `table` | Output format |

**Examples:**

```bash
mbongo node status
mbongo node status --output json
```

**Output (table):**

```
┌────────────────────────────────────────────────────────────┐
│ Node Status                                                │
├─────────────────────┬──────────────────────────────────────┤
│ Status              │ Running                              │
│ Uptime              │ 3d 14h 22m                           │
│ Block Height        │ 12,345,678                           │
│ Sync Status         │ Synced                               │
│ Peers               │ 42                                   │
│ Pending Txs         │ 156                                  │
│ Mode                │ Validator + GPU Provider             │
│ Validator Active    │ Yes                                  │
│ Compute Tasks       │ 3 in progress                        │
│ Memory Usage        │ 4.2 GB                               │
│ Disk Usage          │ 128 GB                               │
└─────────────────────┴──────────────────────────────────────┘
```

**Output (JSON):**

```json
{
  "status": "running",
  "uptime_seconds": 310932,
  "block_height": 12345678,
  "sync_status": "synced",
  "peer_count": 42,
  "pending_transactions": 156,
  "mode": ["validator", "compute_provider"],
  "validator_active": true,
  "compute_tasks_active": 3,
  "memory_bytes": 4509715456,
  "disk_bytes": 137438953472
}
```

---

### 3.5 `mbongo node info`

**Description:** Display detailed node information.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--output` | `-o` | No | `table` | Output format |

**Examples:**

```bash
mbongo node info
mbongo node info --output json
```

**Output:**

```
┌────────────────────────────────────────────────────────────┐
│ Node Information                                           │
├─────────────────────┬──────────────────────────────────────┤
│ Node ID             │ 16Uiu2HAmXyz...abc                   │
│ Version             │ 0.1.0-alpha                          │
│ Network             │ mainnet                              │
│ Chain ID            │ 1                                    │
│ Genesis Hash        │ 0xabc123...                          │
│ Protocol Version    │ 1                                    │
│ Client              │ mbongo/v0.1.0/rust-1.75              │
│ Data Directory      │ /var/lib/mbongo                      │
│ Config File         │ /etc/mbongo/config.toml              │
│ Listen Addresses    │ /ip4/0.0.0.0/tcp/30303               │
│ RPC Endpoint        │ http://127.0.0.1:8545                │
└─────────────────────┴──────────────────────────────────────┘
```

---

### 3.6 `mbongo node peers`

**Description:** List connected peers.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--output` | `-o` | No | `table` | Output format |
| `--limit` | `-l` | No | `50` | Maximum peers to show |
| `--sort` | | No | `latency` | Sort by (latency, id, connected) |

**Examples:**

```bash
mbongo node peers
mbongo node peers --limit 10 --sort latency
mbongo node peers --output json
```

**Output:**

```
Connected Peers: 42
┌──────────────────────────────┬────────────┬──────────┬───────────┐
│ Peer ID                      │ Address    │ Latency  │ Direction │
├──────────────────────────────┼────────────┼──────────┼───────────┤
│ 16Uiu2HAm...abc              │ 1.2.3.4    │ 12ms     │ outbound  │
│ 16Uiu2HAm...def              │ 5.6.7.8    │ 24ms     │ inbound   │
│ 16Uiu2HAm...ghi              │ 9.10.11.12 │ 45ms     │ outbound  │
└──────────────────────────────┴────────────┴──────────┴───────────┘
```

---

### 3.7 `mbongo node sync-status`

**Description:** Show blockchain synchronization status.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--output` | `-o` | No | `table` | Output format |
| `--watch` | `-w` | No | `false` | Continuously update |

**Examples:**

```bash
mbongo node sync-status
mbongo node sync-status --watch
```

**Output:**

```
Synchronization Status
────────────────────────────────────────────────
  Current Block   │ 12,340,000
  Highest Block   │ 12,345,678
  Sync Progress   │ 99.95%
  Blocks Behind   │ 5,678
  Sync Speed      │ ~1,200 blocks/sec
  ETA             │ ~5 seconds
  State           │ Syncing
────────────────────────────────────────────────
```

---

### 3.8 `mbongo node purge-data`

**Description:** Delete all node data (blockchain, state, keystore).

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--yes` | `-y` | No | `false` | Skip confirmation |
| `--keep-keystore` | | No | `false` | Preserve keystore |
| `--keep-config` | | No | `false` | Preserve config |

**Examples:**

```bash
# Interactive purge
mbongo node purge-data

# Non-interactive, keep keys
mbongo node purge-data --yes --keep-keystore
```

**Output:**

```
⚠️  WARNING: Data Purge
────────────────────────────────────────────────
This will permanently delete:
  • Blockchain data (128 GB)
  • State database (45 GB)
  • Transaction index (12 GB)

This will NOT delete:
  • Keystore (--keep-keystore)
  • Config file (--keep-config)

Type 'DELETE' to confirm: DELETE

Purging data...
  [████████████████████████████████] 100%
Data purged successfully.
```

---

### 3.9 `mbongo node snapshot export`

**Description:** Export a state snapshot for fast sync.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--output` | `-o` | Yes | — | Output file path |
| `--block` | | No | `latest` | Block height for snapshot |
| `--compress` | | No | `true` | Enable compression |

**Examples:**

```bash
mbongo node snapshot export --output ./snapshot.tar.zst
mbongo node snapshot export --output ./snapshot.tar.zst --block 12000000
```

---

### 3.10 `mbongo node snapshot import`

**Description:** Import a state snapshot.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--input` | `-i` | Yes | — | Snapshot file path |
| `--verify` | | No | `true` | Verify integrity |

**Examples:**

```bash
mbongo node snapshot import --input ./snapshot.tar.zst
mbongo node snapshot import --input ./snapshot.tar.zst --verify=false
```

---

### 3.11 `mbongo node logs`

**Description:** View or stream node logs.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--follow` | `-f` | No | `false` | Stream logs |
| `--lines` | `-n` | No | `100` | Number of lines |
| `--level` | | No | `all` | Filter by level |
| `--grep` | | No | — | Filter by pattern |

**Examples:**

```bash
mbongo node logs
mbongo node logs -f
mbongo node logs -n 500 --level error
mbongo node logs --grep "consensus"
```

---

### 3.12 `mbongo node rpc`

**Description:** Manage RPC server.

**Subcommands:**

| Subcommand | Description |
|------------|-------------|
| `status` | RPC server status |
| `enable` | Enable RPC server |
| `disable` | Disable RPC server |

**Examples:**

```bash
mbongo node rpc status
mbongo node rpc enable --addr 127.0.0.1 --port 8545
mbongo node rpc disable
```

---

### 3.13 `mbongo node metrics`

**Description:** Prometheus metrics endpoint management.

**Examples:**

```bash
mbongo node metrics status
mbongo node metrics enable --port 9090
curl http://localhost:9090/metrics
```

---

## 4. Configuration Integration

### 4.1 Default Paths

```
~/.mbongo/
├── config.toml          # Main configuration
├── keystore/            # Encrypted keys
│   ├── validator.json   # Validator key
│   └── node.json        # Node identity key
├── data/
│   ├── db/              # LevelDB/RocksDB
│   ├── state/           # State trie
│   └── blocks/          # Block storage
├── logs/                # Log files
└── snapshots/           # State snapshots
```

### 4.2 Config File Override

```bash
# Use custom config
mbongo node start --config /etc/mbongo/production.toml

# Config file example (config.toml)
[node]
data_dir = "/var/lib/mbongo"
log_level = "info"

[network]
listen_addr = "/ip4/0.0.0.0/tcp/30303"
bootnodes = ["/ip4/1.2.3.4/tcp/30303/p2p/16Uiu2HAm..."]

[rpc]
enabled = true
addr = "127.0.0.1"
port = 8545

[validator]
enabled = true
keystore = "/secure/keystore"

[compute]
enabled = true
gpu_device = 0
```

### 4.3 Node Identity

```bash
# Node identity is auto-generated on first start
# Stored in: ~/.mbongo/keystore/node.json

# View node ID (public only)
mbongo node info | grep "Node ID"

# ⚠️ Private keys are NEVER exposed via CLI
```

---

## 5. Security Rules

### 5.1 Local-Only Commands

Commands that require local access (no RPC):

```
mbongo node start       # Process control
mbongo node stop        # Process control
mbongo node restart     # Process control
mbongo node purge-data  # Data deletion
mbongo node snapshot *  # File I/O
```

### 5.2 RPC Exposure Warnings

```
⚠️  PRODUCTION SECURITY CHECKLIST
────────────────────────────────────────────────
✗ Never expose RPC to 0.0.0.0 without firewall
✗ Never enable admin RPC methods publicly
✗ Never run node as root user
✓ Use reverse proxy (nginx) with TLS
✓ Implement IP whitelisting
✓ Use separate user with limited permissions
✓ Enable rate limiting
```

### 5.3 Production Recommendations

```bash
# Create dedicated user
sudo useradd -r -s /bin/false mbongo

# Set permissions
sudo chown -R mbongo:mbongo /var/lib/mbongo
sudo chmod 700 /var/lib/mbongo/keystore

# Run with restricted permissions
sudo -u mbongo mbongo node start -c /etc/mbongo/config.toml
```

### 5.4 Logging Best Practices

```bash
# Production: info level, file output
mbongo node start --log-level info

# Debugging: debug level, console
mbongo node start --log-level debug

# Never log sensitive data
# Logs automatically redact:
# - Private keys
# - Passwords
# - Session tokens
```

---

## 6. Advanced Topics

### 6.1 GPU Provider Mode

```bash
# Start with GPU provider
mbongo node start --compute-provider \
  --gpu-device 0 \
  --compute-threads 4 \
  --max-tasks 10

# Check GPU status
mbongo compute status

# Multiple GPUs
mbongo node start --compute-provider --gpu-device 0,1,2,3
```

### 6.2 Validator Mode via CLI

```bash
# Register validator (requires 50,000 MBO staked)
mbongo validator register \
  --keystore ~/.mbongo/keystore \
  --stake 50000

# Start node as validator
mbongo node start --validator

# Monitor validator
mbongo validator status
```

### 6.3 Developer Mode

```bash
# Fast sync, skip verification (UNSAFE)
mbongo node start --dev \
  --fast-sync \
  --skip-verification \
  --unsafe-rpc

# Local devnet
mbongo node start --dev --chain dev --mine
```

### 6.4 Container Deployment

```dockerfile
# Dockerfile
FROM rust:1.75-slim
COPY mbongo /usr/local/bin/
EXPOSE 30303 8545 9090
ENTRYPOINT ["mbongo", "node", "start"]
CMD ["--config", "/config/config.toml"]
```

```bash
# Docker run
docker run -d \
  -v /data/mbongo:/var/lib/mbongo \
  -v /config:/config \
  -p 30303:30303 \
  -p 8545:8545 \
  mbongo-node
```

### 6.5 Systemd Service

```ini
# /etc/systemd/system/mbongo.service
[Unit]
Description=Mbongo Chain Node
After=network.target

[Service]
Type=simple
User=mbongo
ExecStart=/usr/local/bin/mbongo node start -c /etc/mbongo/config.toml
ExecStop=/usr/local/bin/mbongo node stop
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl enable mbongo
sudo systemctl start mbongo
sudo systemctl status mbongo
```

---

## 7. ASCII Diagrams

### 7.1 Node Startup Pipeline

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         NODE STARTUP PIPELINE                               │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   $ mbongo node start                                                       │
│         │                                                                   │
│         ▼                                                                   │
│   ┌───────────────┐    ┌───────────────┐    ┌───────────────┐              │
│   │ Load Config   │───▶│ Init Database │───▶│ Load Keystore │              │
│   └───────────────┘    └───────────────┘    └───────────────┘              │
│         │                                          │                        │
│         │              ┌───────────────────────────┘                        │
│         ▼              ▼                                                    │
│   ┌───────────────┐    ┌───────────────┐    ┌───────────────┐              │
│   │ Start Network │───▶│ Peer Discovery│───▶│ Begin Sync    │              │
│   └───────────────┘    └───────────────┘    └───────────────┘              │
│         │                                          │                        │
│         │              ┌───────────────────────────┘                        │
│         ▼              ▼                                                    │
│   ┌───────────────┐    ┌───────────────┐    ┌───────────────┐              │
│   │ Start RPC     │───▶│ Start Metrics │───▶│ Ready         │              │
│   └───────────────┘    └───────────────┘    └───────────────┘              │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 7.2 Peer Discovery & Sync

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         P2P SYNC FLOW                                       │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   LOCAL NODE                           PEER NETWORK                         │
│   ══════════                           ════════════                         │
│                                                                             │
│   ┌─────────────┐                      ┌─────────────┐                     │
│   │   Node      │ ──── Bootstrap ────▶ │  Bootnode   │                     │
│   └─────────────┘                      └─────────────┘                     │
│         │                                    │                              │
│         │ ◀──────── Peer List ───────────────┘                              │
│         │                                                                   │
│         ▼                                                                   │
│   ┌─────────────┐      Connect         ┌─────────────┐                     │
│   │   Node      │ ◀──────────────────▶ │   Peer 1    │                     │
│   └─────────────┘      Gossip          └─────────────┘                     │
│         │                                                                   │
│         │ ◀──────── Block Headers ──────────────────────                   │
│         │                                                                   │
│         │ ◀──────── Block Bodies ───────────────────────                   │
│         │                                                                   │
│         │ ◀──────── State Trie ─────────────────────────                   │
│         │                                                                   │
│         ▼                                                                   │
│   ┌─────────────┐                                                          │
│   │   SYNCED    │                                                          │
│   └─────────────┘                                                          │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 7.3 RPC Query Flow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         RPC QUERY FLOW                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   CLIENT                    RPC SERVER                  NODE CORE           │
│   ══════                    ══════════                  ═════════           │
│                                                                             │
│   ┌─────────┐               ┌─────────┐               ┌─────────┐          │
│   │ curl /  │  JSON-RPC     │ HTTP    │    Query      │ State   │          │
│   │ SDK     │ ────────────▶ │ Server  │ ────────────▶ │ Manager │          │
│   └─────────┘               └─────────┘               └─────────┘          │
│                                   │                         │               │
│                                   │                         ▼               │
│                                   │                   ┌─────────┐          │
│                                   │                   │ Database│          │
│                                   │                   └─────────┘          │
│                                   │                         │               │
│   ┌─────────┐               ┌─────────┐               ┌─────────┐          │
│   │ Response│ ◀──────────── │ Format  │ ◀──────────── │ Result  │          │
│   │ (JSON)  │               │ Handler │               │         │          │
│   └─────────┘               └─────────┘               └─────────┘          │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Quick Reference

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         NODE COMMANDS QUICK REFERENCE                       │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   LIFECYCLE                          STATUS                                 │
│   ─────────                          ──────                                 │
│   mbongo node start                  mbongo node status                     │
│   mbongo node stop                   mbongo node info                       │
│   mbongo node restart                mbongo node sync-status                │
│                                      mbongo node peers                      │
│                                                                             │
│   DATA                               SERVICES                               │
│   ────                               ────────                               │
│   mbongo node purge-data             mbongo node rpc status                 │
│   mbongo node snapshot export        mbongo node rpc enable                 │
│   mbongo node snapshot import        mbongo node metrics                    │
│   mbongo node logs                                                          │
│                                                                             │
│   COMMON FLAGS                                                              │
│   ────────────                                                              │
│   -c, --config <FILE>     Config file path                                 │
│   -d, --data-dir <DIR>    Data directory                                   │
│   --validator             Enable validator mode                            │
│   --compute-provider      Enable GPU provider mode                         │
│   --rpc-addr <ADDR>       RPC bind address                                 │
│   --rpc-port <PORT>       RPC port                                         │
│   -o, --output <FORMAT>   Output format (json, table, raw)                 │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Related Documentation

| Document | Description |
|----------|-------------|
| [cli_overview.md](./cli_overview.md) | CLI overview and conventions |
| [cli_validator.md](./cli_validator.md) | Validator commands |
| [cli_network.md](./cli_network.md) | Network commands |
| [cli_config.md](./cli_config.md) | Configuration reference |

---

*This document provides the complete reference for `mbongo node` commands. For general CLI information, see [cli_overview.md](./cli_overview.md).*

