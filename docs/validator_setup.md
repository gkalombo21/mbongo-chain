# Mbongo Chain — Validator Setup Guide

> **Document Type:** Setup Tutorial  
> **Last Updated:** November 2025  
> **Status:** Official Reference  
> **Audience:** Validator Operators, Infrastructure Engineers

---

## Table of Contents

1. [Introduction](#1-introduction)
2. [Hardware Requirements](#2-hardware-requirements)
3. [Software Prerequisites](#3-software-prerequisites)
4. [Key Management](#4-key-management)
5. [Installation Steps](#5-installation-steps)
6. [Security & Slashing Prevention](#6-security--slashing-prevention)
7. [Testnet vs Mainnet Differences](#7-testnet-vs-mainnet-differences)
8. [Troubleshooting](#8-troubleshooting)
9. [Cross-References](#9-cross-references)

---

## 1. Introduction

### 1.1 Role of Validators in PoS

Validators are the backbone of Mbongo Chain's Proof-of-Stake (PoS) consensus. They secure the network by staking MBO tokens as collateral and participating in block production and attestation.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         VALIDATOR ROLE IN CONSENSUS                         │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐  │
│   │                                                                      │  │
│   │                    MBONGO CHAIN CONSENSUS                            │  │
│   │                                                                      │  │
│   │   ┌───────────────────┐         ┌───────────────────┐               │  │
│   │   │    PoS (50%)      │         │   PoUW (50%)      │               │  │
│   │   │                   │         │                   │               │  │
│   │   │  ┌─────────────┐  │         │  ┌─────────────┐  │               │  │
│   │   │  │ VALIDATORS  │  │         │  │  COMPUTE    │  │               │  │
│   │   │  │             │  │         │  │  PROVIDERS  │  │               │  │
│   │   │  │ • Stake MBO │  │         │  │             │  │               │  │
│   │   │  │ • Propose   │  │         │  │ • Execute   │  │               │  │
│   │   │  │ • Attest    │  │         │  │ • Prove     │  │               │  │
│   │   │  │ • Earn 50%  │  │         │  │ • Earn 50%  │  │               │  │
│   │   │  └─────────────┘  │         │  └─────────────┘  │               │  │
│   │   │                   │         │                   │               │  │
│   │   └───────────────────┘         └───────────────────┘               │  │
│   │                                                                      │  │
│   └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│   Validators earn 50% of all block rewards + priority fees                 │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 1.2 Validator Responsibilities

| Responsibility | Description | Frequency |
|----------------|-------------|-----------|
| **Block Proposal** | Create new blocks when selected as leader | Per-slot (when selected) |
| **Attestation** | Vote on block validity | Every slot |
| **Signing** | Cryptographically sign proposals and votes | Continuous |
| **Uptime** | Maintain 24/7 availability | Always |
| **Sync** | Stay synchronized with the chain | Always |

### 1.3 Slashing Risks

> ⚠️ **WARNING: Slashing can result in permanent loss of staked MBO**

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         SLASHING PENALTIES                                  │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   OFFENSE                 │ PENALTY        │ RECOVERY                      │
│   ────────────────────────┼────────────────┼───────────────────────────────│
│   Double-signing          │ 5% of stake    │ None (permanent damage)       │
│   Surround vote           │ 5% of stake    │ None (permanent damage)       │
│   Extended downtime       │ 0.5% of stake  │ Automatic after recovery      │
│   (>500 consecutive slots)│ per day        │                               │
│                                                                             │
│   SLASHING SEVERITY LADDER                                                  │
│   ════════════════════════                                                  │
│                                                                             │
│   ┌───────────────────────────────────────────────────────────────────┐    │
│   │  MINOR                    MODERATE                  SEVERE        │    │
│   │    │                         │                         │          │    │
│   │    ▼                         ▼                         ▼          │    │
│   │  Downtime              Invalid attestation      Double-signing    │    │
│   │  (0.5%/day)            (1% stake)               (5% stake)        │    │
│   │                                                                   │    │
│   │  Recoverable           Recoverable              PERMANENT         │    │
│   └───────────────────────────────────────────────────────────────────┘    │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 1.4 Required Skills

Before proceeding, ensure you have:

| Skill | Level | Description |
|-------|-------|-------------|
| **Command Line** | Basic | Navigate filesystem, run commands |
| **System Administration** | Basic | Manage services, logs, updates |
| **Networking** | Basic | Firewall rules, port forwarding |
| **Security Awareness** | Important | Key management, backup procedures |
| **Blockchain Concepts** | Helpful | Understanding of PoS, consensus |

### 1.5 Supported Operating Systems

| OS | Support Level | Notes |
|----|---------------|-------|
| **Ubuntu 22.04 LTS** | ✓ Full support | Recommended |
| **Ubuntu 20.04 LTS** | ✓ Full support | Supported |
| **Debian 11/12** | ✓ Full support | Supported |
| **Windows Server 2022** | ○ Experimental | PowerShell required |
| **Windows 11** | ○ Experimental | For development only |
| **macOS** | ○ Experimental | For development only |

---

## 2. Hardware Requirements

### 2.1 Validator Node Specifications

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    VALIDATOR HARDWARE REQUIREMENTS                          │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   COMPONENT       │ MINIMUM            │ RECOMMENDED         │ OPTIMAL     │
│   ────────────────┼────────────────────┼─────────────────────┼─────────────│
│   CPU             │ 4 cores @ 2.5 GHz  │ 8 cores @ 3.0 GHz   │ 16 cores    │
│   RAM             │ 8 GB               │ 16 GB               │ 32 GB       │
│   Storage         │ 512 GB NVMe        │ 1 TB NVMe           │ 2 TB NVMe   │
│   Network         │ 50 Mbps            │ 100 Mbps            │ 200+ Mbps   │
│   Latency         │ <100ms to peers    │ <50ms to peers      │ <20ms       │
│                                                                             │
│   ADDITIONAL RECOMMENDATIONS                                                │
│   ══════════════════════════                                                │
│   • Dedicated server (not shared hosting)                                  │
│   • SSD with >50,000 IOPS                                                  │
│   • Redundant internet connection (optional but recommended)               │
│   • UPS for power protection                                               │
│   • ECC memory (for enterprise deployments)                                │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

| Component | Minimum | Recommended | Notes |
|-----------|---------|-------------|-------|
| **CPU** | 4 cores | 8 cores | Modern x86_64 (Intel Xeon, AMD EPYC) |
| **RAM** | 8 GB | 16 GB | DDR4/DDR5, ECC preferred |
| **Storage** | 512 GB NVMe | 1 TB NVMe | High IOPS, low latency |
| **Network** | 50 Mbps | 100-200 Mbps | Stable, low latency |
| **Redundancy** | Optional | Recommended | Dual ISP, UPS |

### 2.2 Cloud Provider Recommendations

| Provider | Instance Type | Monthly Cost (est.) |
|----------|---------------|---------------------|
| **AWS** | m6i.xlarge | ~$150 |
| **GCP** | n2-standard-4 | ~$140 |
| **Azure** | Standard_D4s_v5 | ~$145 |
| **Hetzner** | AX41-NVMe | ~$50 |
| **OVH** | Rise-1 | ~$70 |

---

## 3. Software Prerequisites

### 3.1 Required Software

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    SOFTWARE REQUIREMENTS                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   CORE COMPONENTS                                                           │
│   ═══════════════                                                           │
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐  │
│   │  1. Rust Toolchain                                                  │  │
│   │     • Version: 1.75+ (stable)                                       │  │
│   │     • Components: rustc, cargo                                      │  │
│   └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐  │
│   │  2. Mbongo CLI                                                      │  │
│   │     • Node management                                               │  │
│   │     • Wallet operations                                             │  │
│   │     • Validator commands                                            │  │
│   └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐  │
│   │  3. Mbongo Validator Client                                         │  │
│   │     • Consensus participation                                       │  │
│   │     • Block signing                                                 │  │
│   │     • Attestation generation                                        │  │
│   └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐  │
│   │  4. Mbongo Node                                                     │  │
│   │     • Blockchain sync                                               │  │
│   │     • P2P networking                                                │  │
│   │     • RPC server                                                    │  │
│   └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 3.2 Ubuntu System Packages

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install required packages
sudo apt install -y \
    build-essential \
    pkg-config \
    libssl-dev \
    libclang-dev \
    cmake \
    git \
    curl \
    wget \
    jq \
    htop \
    tmux \
    ufw

# Install Rust toolchain
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
source ~/.cargo/env

# Verify Rust installation
rustc --version
cargo --version
```

### 3.3 Windows PowerShell Prerequisites

```powershell
# Run PowerShell as Administrator

# Install Chocolatey (package manager)
Set-ExecutionPolicy Bypass -Scope Process -Force
[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072
iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

# Install required packages
choco install -y `
    git `
    cmake `
    visualstudio2022-workload-vctools `
    openssl `
    rust-ms

# Verify installations
rustc --version
cargo --version
git --version
```

### 3.4 Version Requirements

| Software | Minimum Version | Recommended Version |
|----------|-----------------|---------------------|
| **Rust** | 1.70.0 | 1.75.0+ |
| **Mbongo CLI** | 0.1.0 | Latest |
| **Mbongo Node** | 0.1.0 | Latest |
| **Mbongo Validator** | 0.1.0 | Latest |

---

## 4. Key Management

### 4.1 Validator Key Architecture

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    VALIDATOR KEY ARCHITECTURE                               │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   VALIDATOR IDENTITY                                                        │
│   ══════════════════                                                        │
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐  │
│   │                                                                      │  │
│   │                    ┌─────────────────────┐                           │  │
│   │                    │   MASTER SEED       │                           │  │
│   │                    │   (BIP-39 Mnemonic) │                           │  │
│   │                    └──────────┬──────────┘                           │  │
│   │                               │                                      │  │
│   │              ┌────────────────┼────────────────┐                     │  │
│   │              │                │                │                     │  │
│   │              ▼                ▼                ▼                     │  │
│   │   ┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐       │  │
│   │   │  SIGNING KEY    │ │ WITHDRAWAL KEY  │ │   NODE KEY      │       │  │
│   │   │                 │ │                 │ │                 │       │  │
│   │   │ • Signs blocks  │ │ • Controls MBO  │ │ • P2P identity  │       │  │
│   │   │ • Signs votes   │ │ • Exit stake    │ │ • Network auth  │       │  │
│   │   │ • Hot key       │ │ • Cold key      │ │ • Hot key       │       │  │
│   │   │ • On server     │ │ • OFFLINE       │ │ • On server     │       │  │
│   │   └─────────────────┘ └─────────────────┘ └─────────────────┘       │  │
│   │                                                                      │  │
│   └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│   KEY SECURITY LEVELS                                                       │
│   ═══════════════════                                                       │
│                                                                             │
│   ┌───────────────────────────────────────────────────────────────────┐    │
│   │  SIGNING KEY         │ WITHDRAWAL KEY      │ NODE KEY            │    │
│   │  ─────────────────── │ ─────────────────── │ ─────────────────── │    │
│   │  Security: HIGH      │ Security: CRITICAL  │ Security: MEDIUM    │    │
│   │  Location: Server    │ Location: OFFLINE   │ Location: Server    │    │
│   │  Encrypted: YES      │ Encrypted: YES      │ Encrypted: Optional │    │
│   │  Backup: REQUIRED    │ Backup: ESSENTIAL   │ Backup: Recommended │    │
│   └───────────────────────────────────────────────────────────────────┘    │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 4.2 Key Generation

> ⚠️ **CRITICAL: Generate keys on a secure, air-gapped machine when possible**

#### Generate Validator Keys

```bash
# Create secure directory
mkdir -p ~/.mbongo/validator-keys
chmod 700 ~/.mbongo/validator-keys
cd ~/.mbongo/validator-keys

# Generate new validator keys
mbongo wallet create \
    --keystore ./signing-key.json \
    --type validator

# You will be prompted:
# Enter password for keystore: [enter strong password]
# Confirm password: [confirm password]

# Output will show:
# ✓ Validator signing key generated
# ✓ Public key: 0x8a7b3c4d...
# ✓ Keystore saved to: ./signing-key.json
# ✓ IMPORTANT: Back up your keystore and password securely!
```

#### Generate Withdrawal Key (Separate, Cold Storage)

```bash
# On a SEPARATE, AIR-GAPPED machine:
mbongo wallet create \
    --keystore ./withdrawal-key.json \
    --type withdrawal

# Store this key OFFLINE in multiple secure locations
# This key controls your staked MBO - NEVER put it online
```

### 4.3 Keystore Encryption

```bash
# Keystore format (JSON)
# The keystore is encrypted with your password using scrypt

# Example keystore structure:
cat signing-key.json
```

```json
{
  "version": 3,
  "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
  "address": "8a7b3c4d5e6f7890abcdef1234567890abcdef12",
  "crypto": {
    "ciphertext": "...",
    "cipherparams": {
      "iv": "..."
    },
    "cipher": "aes-128-ctr",
    "kdf": "scrypt",
    "kdfparams": {
      "dklen": 32,
      "salt": "...",
      "n": 262144,
      "r": 8,
      "p": 1
    },
    "mac": "..."
  }
}
```

### 4.4 Backup Strategy (3-2-1 Rule)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    3-2-1 BACKUP STRATEGY                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   THE 3-2-1 RULE                                                            │
│   ══════════════                                                            │
│                                                                             │
│   3 │ Keep 3 copies of your keys                                           │
│   2 │ Store on 2 different media types                                     │
│   1 │ Keep 1 copy offsite                                                  │
│                                                                             │
│   RECOMMENDED BACKUP LOCATIONS                                              │
│   ════════════════════════════                                              │
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐  │
│   │                                                                      │  │
│   │   COPY 1: Primary (Server)                                          │  │
│   │   └── ~/.mbongo/validator-keys/signing-key.json                     │  │
│   │                                                                      │  │
│   │   COPY 2: Local Backup (USB Drive, encrypted)                       │  │
│   │   └── Store in safe or locked cabinet                               │  │
│   │                                                                      │  │
│   │   COPY 3: Offsite Backup (Safety deposit box)                       │  │
│   │   └── Paper backup of mnemonic in sealed envelope                   │  │
│   │                                                                      │  │
│   └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│   WITHDRAWAL KEY (CRITICAL)                                                 │
│   ═════════════════════════                                                 │
│   • NEVER store on the validator server                                    │
│   • Store in cold storage only                                             │
│   • Consider hardware wallet or multi-sig                                  │
│   • Test recovery process before staking                                   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 4.5 Slashing Protection Database

```bash
# Initialize slashing protection database
mbongo validator init-slashing-protection \
    --db-path ~/.mbongo/slashing_protection.db

# The slashing protection DB prevents double-signing
# ALWAYS back up this file and migrate it when moving validators

# Export slashing protection data (for migration)
mbongo validator export-slashing-protection \
    --db-path ~/.mbongo/slashing_protection.db \
    --output ./slashing_protection_backup.json

# Import on new machine (BEFORE starting validator)
mbongo validator import-slashing-protection \
    --db-path ~/.mbongo/slashing_protection.db \
    --input ./slashing_protection_backup.json
```

---

## 5. Installation Steps

### Step 1 — Install Dependencies

#### Ubuntu 22.04 LTS

```bash
#!/bin/bash
# save as: install_dependencies.sh

echo "=== Mbongo Validator: Installing Dependencies ==="

# Update system
sudo apt update && sudo apt upgrade -y

# Install system packages
sudo apt install -y \
    build-essential \
    pkg-config \
    libssl-dev \
    libclang-dev \
    cmake \
    git \
    curl \
    wget \
    jq \
    htop \
    tmux \
    ufw \
    fail2ban

# Install Rust
if ! command -v rustc &> /dev/null; then
    echo "Installing Rust..."
    curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
    source ~/.cargo/env
fi

# Verify installations
echo "=== Verification ==="
rustc --version
cargo --version

echo "=== Dependencies installed successfully ==="
```

#### Windows (PowerShell as Administrator)

```powershell
# save as: install_dependencies.ps1

Write-Host "=== Mbongo Validator: Installing Dependencies ===" -ForegroundColor Green

# Install Chocolatey if not present
if (!(Get-Command choco -ErrorAction SilentlyContinue)) {
    Set-ExecutionPolicy Bypass -Scope Process -Force
    [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072
    iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
}

# Install required packages
choco install -y git cmake visualstudio2022-workload-vctools openssl rust-ms

# Refresh environment
refreshenv

# Verify
Write-Host "=== Verification ===" -ForegroundColor Green
rustc --version
cargo --version
git --version

Write-Host "=== Dependencies installed successfully ===" -ForegroundColor Green
```

### Step 2 — Download Mbongo Validator Client

#### Option A: Install via Cargo (Recommended)

```bash
# Install Mbongo CLI and Node
cargo install mbongo-cli mbongo-node mbongo-validator

# Verify installation
mbongo --version
mbongo-node --version
mbongo-validator --version
```

#### Option B: Download Pre-built Binaries

```bash
# Download latest release (placeholder URL)
MBONGO_VERSION="v0.1.0"
wget https://github.com/mbongo-chain/mbongo/releases/download/${MBONGO_VERSION}/mbongo-linux-amd64.tar.gz

# Extract
tar -xzf mbongo-linux-amd64.tar.gz

# Move to system path
sudo mv mbongo* /usr/local/bin/

# Verify
mbongo --version
```

#### Add to PATH (if needed)

```bash
# Ubuntu
echo 'export PATH="$HOME/.cargo/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# Windows (PowerShell)
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";C:\Users\$env:USERNAME\.cargo\bin", "User")
```

### Step 3 — Configure the Validator

#### Create Validator Directory Structure

```bash
# Create directory structure
mkdir -p ~/.mbongo/{config,data,keys,logs}
chmod 700 ~/.mbongo/keys

# Directory structure:
# ~/.mbongo/
# ├── config/
# │   ├── config.toml          # Main configuration
# │   ├── validator.toml       # Validator-specific config
# │   └── genesis.json         # Network genesis file
# ├── data/
# │   ├── chaindata/           # Blockchain data
# │   └── slashing_protection.db
# ├── keys/
# │   └── signing-key.json     # Encrypted validator key
# └── logs/
#     └── mbongo.log           # Log files
```

#### Create Configuration File

```bash
# Initialize configuration for mainnet
mbongo config init --network mainnet --validator

# Or for testnet
mbongo config init --network testnet --validator
```

#### Configuration File Layout

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    CONFIG FILE STRUCTURE                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   ~/.mbongo/config/config.toml                                             │
│   ════════════════════════════                                              │
│                                                                             │
│   [node]                                                                    │
│   ├── data_dir = "~/.mbongo/data"                                          │
│   ├── log_level = "info"                                                   │
│   └── log_file = "~/.mbongo/logs/mbongo.log"                               │
│                                                                             │
│   [network]                                                                 │
│   ├── chain_id = 1                    # 1 = mainnet, 11155111 = testnet    │
│   ├── listen_addr = "0.0.0.0:30303"                                        │
│   ├── bootnodes = [...]               # Network bootstrap nodes            │
│   └── max_peers = 50                                                       │
│                                                                             │
│   [rpc]                                                                     │
│   ├── enabled = true                                                       │
│   ├── addr = "127.0.0.1:8545"         # Local only for security            │
│   └── cors = []                                                            │
│                                                                             │
│   [validator]                                                               │
│   ├── enabled = true                                                       │
│   ├── keystore = "~/.mbongo/keys/signing-key.json"                         │
│   ├── slashing_db = "~/.mbongo/data/slashing_protection.db"                │
│   └── fee_recipient = "0xYOUR_ADDRESS_HERE"                                │
│                                                                             │
│   [metrics]                                                                 │
│   ├── enabled = true                                                       │
│   └── addr = "127.0.0.1:9090"                                              │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

#### Sample config.toml

```toml
# ~/.mbongo/config/config.toml

[node]
data_dir = "~/.mbongo/data"
log_level = "info"
log_file = "~/.mbongo/logs/mbongo.log"

[network]
chain_id = 1
listen_addr = "0.0.0.0:30303"
bootnodes = [
    "enode://abc123...@bootnode1.mbongo.io:30303",
    "enode://def456...@bootnode2.mbongo.io:30303",
]
max_peers = 50

[rpc]
enabled = true
addr = "127.0.0.1:8545"
ws_addr = "127.0.0.1:8546"
cors = []
rate_limit = 100

[validator]
enabled = true
keystore = "~/.mbongo/keys/signing-key.json"
slashing_db = "~/.mbongo/data/slashing_protection.db"
fee_recipient = "0xYOUR_ADDRESS_HERE"
graffiti = "MyValidator"

[metrics]
enabled = true
addr = "127.0.0.1:9090"
```

### Step 4 — Import Validator Keys

```bash
# Import validator keystore
mbongo validator import-key \
    --keystore ~/.mbongo/keys/signing-key.json \
    --password-file ~/.mbongo/keys/password.txt

# Or enter password interactively (more secure)
mbongo validator import-key \
    --keystore ~/.mbongo/keys/signing-key.json

# Verify key imported
mbongo validator list-keys

# Initialize slashing protection
mbongo validator init-slashing-protection \
    --db-path ~/.mbongo/data/slashing_protection.db

# Verify slashing protection enabled
mbongo validator verify-slashing-protection
```

### Step 5 — Start the Validator

#### Ubuntu: Create Systemd Service

```bash
# Create systemd service file
sudo tee /etc/systemd/system/mbongo-validator.service > /dev/null << 'EOF'
[Unit]
Description=Mbongo Validator Node
After=network.target
Wants=network-online.target

[Service]
Type=simple
User=mbongo
Group=mbongo
ExecStart=/usr/local/bin/mbongo validator start \
    --config /home/mbongo/.mbongo/config/config.toml
Restart=always
RestartSec=5
LimitNOFILE=65535

# Security hardening
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ProtectHome=read-only
ReadWritePaths=/home/mbongo/.mbongo

[Install]
WantedBy=multi-user.target
EOF

# Create mbongo user (if not exists)
sudo useradd -r -s /bin/false mbongo || true

# Set permissions
sudo chown -R mbongo:mbongo /home/mbongo/.mbongo

# Reload systemd
sudo systemctl daemon-reload

# Enable service (start on boot)
sudo systemctl enable mbongo-validator

# Start the validator
sudo systemctl start mbongo-validator

# Check status
sudo systemctl status mbongo-validator
```

#### Windows: Create PowerShell Service

```powershell
# Create service wrapper script
$serviceScript = @"
# Mbongo Validator Service Script
Set-Location "C:\Users\$env:USERNAME\.mbongo"
& mbongo validator start --config "C:\Users\$env:USERNAME\.mbongo\config\config.toml"
"@

$serviceScript | Out-File -FilePath "C:\mbongo\start-validator.ps1"

# Install as Windows Service using NSSM
# Download NSSM: https://nssm.cc/download
nssm install MbongoValidator "powershell.exe" "-ExecutionPolicy Bypass -File C:\mbongo\start-validator.ps1"
nssm set MbongoValidator AppDirectory "C:\Users\$env:USERNAME\.mbongo"
nssm set MbongoValidator DisplayName "Mbongo Validator"
nssm set MbongoValidator Start SERVICE_AUTO_START

# Start service
nssm start MbongoValidator

# Check status
nssm status MbongoValidator
```

#### Monitor Validator Status

```bash
# Check if validator is running
sudo systemctl status mbongo-validator

# View logs (real-time)
sudo journalctl -u mbongo-validator -f

# Or check log file
tail -f ~/.mbongo/logs/mbongo.log

# Check sync status
mbongo node sync-status

# Check validator status
mbongo validator status

# Check peer connections
mbongo node peers

# Check RPC health
curl -X POST http://127.0.0.1:8545 \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_syncing","params":[],"id":1}'

# Check metrics
curl http://127.0.0.1:9090/metrics
```

---

## 6. Security & Slashing Prevention

### 6.1 Critical Security Rules

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    VALIDATOR SECURITY RULES                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   ⛔ NEVER DO                                                               │
│   ═══════════                                                               │
│                                                                             │
│   ✗ Run the same validator key on multiple machines                        │
│   ✗ Share your signing key with anyone                                     │
│   ✗ Store withdrawal key on the validator server                           │
│   ✗ Expose RPC ports (8545, 8546) to the public internet                   │
│   ✗ Run without slashing protection database                               │
│   ✗ Delete slashing protection database                                    │
│   ✗ Start validator before full sync                                       │
│                                                                             │
│   ✓ ALWAYS DO                                                               │
│   ═══════════                                                               │
│                                                                             │
│   ✓ Use encrypted keystores                                                │
│   ✓ Back up keys using 3-2-1 rule                                          │
│   ✓ Keep slashing protection database backed up                            │
│   ✓ Monitor validator health 24/7                                          │
│   ✓ Set up alerts for downtime                                             │
│   ✓ Keep software updated                                                  │
│   ✓ Use firewall to restrict access                                        │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 6.2 Firewall Configuration

```bash
# Ubuntu UFW configuration
sudo ufw default deny incoming
sudo ufw default allow outgoing

# Allow SSH (change port if using non-standard)
sudo ufw allow 22/tcp

# Allow P2P
sudo ufw allow 30303/tcp
sudo ufw allow 30303/udp

# DO NOT allow RPC from internet (keep local only)
# sudo ufw allow 8545/tcp  # DON'T DO THIS

# Enable firewall
sudo ufw enable

# Verify rules
sudo ufw status verbose
```

### 6.3 Server Hardening

```bash
# Disable root login via SSH
sudo sed -i 's/PermitRootLogin yes/PermitRootLogin no/' /etc/ssh/sshd_config

# Use SSH keys only (disable password auth)
sudo sed -i 's/#PasswordAuthentication yes/PasswordAuthentication no/' /etc/ssh/sshd_config

# Restart SSH
sudo systemctl restart sshd

# Install and configure fail2ban
sudo apt install -y fail2ban
sudo systemctl enable fail2ban
sudo systemctl start fail2ban

# Enable automatic security updates
sudo apt install -y unattended-upgrades
sudo dpkg-reconfigure -plow unattended-upgrades
```

### 6.4 Slashing Prevention Checklist

| Check | Command | Expected Result |
|-------|---------|-----------------|
| Slashing DB exists | `ls ~/.mbongo/data/slashing_protection.db` | File exists |
| Only one validator instance | `pgrep -c mbongo-validator` | Returns 1 |
| Node is synced | `mbongo node sync-status` | `synced: true` |
| Key is imported | `mbongo validator list-keys` | Shows your key |
| Correct network | `mbongo node info` | Correct chain ID |

---

## 7. Testnet vs Mainnet Differences

### 7.1 Comparison Table

| Aspect | Testnet | Mainnet |
|--------|---------|---------|
| **Chain ID** | 11155111 | 1 |
| **Tokens** | Test MBO (no value) | Real MBO |
| **Faucet** | Available | N/A |
| **Stake** | Test tokens | Real stake required |
| **Slashing** | No real loss | Real economic loss |
| **Resets** | May occur | Never |
| **Bootnodes** | testnet-boot*.mbongo.io | boot*.mbongo.io |

### 7.2 Testnet Setup

```bash
# Get testnet tokens from faucet
# Visit: https://faucet.testnet.mbongo.io
# Or use CLI:
mbongo faucet request --address 0xYOUR_ADDRESS

# Initialize for testnet
mbongo config init --network testnet --validator

# Testnet-specific config values:
# chain_id = 11155111
# bootnodes = ["enode://...@testnet-boot1.mbongo.io:30303"]

# Start testnet validator
mbongo validator start --network testnet
```

### 7.3 Mainnet Configuration

```bash
# Initialize for mainnet
mbongo config init --network mainnet --validator

# Mainnet-specific config values:
# chain_id = 1
# bootnodes = ["enode://...@boot1.mbongo.io:30303"]

# IMPORTANT: Verify chain ID before staking!
mbongo node info | grep chain_id
# Should show: chain_id: 1

# Stake MBO (minimum 50,000 MBO)
mbongo validator stake \
    --amount 50000 \
    --withdrawal-address 0xYOUR_WITHDRAWAL_ADDRESS
```

### 7.4 Migration: Testnet to Mainnet

```bash
# 1. Stop testnet validator
sudo systemctl stop mbongo-validator

# 2. Back up testnet data (optional)
cp -r ~/.mbongo ~/.mbongo-testnet-backup

# 3. Clear testnet data
rm -rf ~/.mbongo/data/chaindata

# 4. Reconfigure for mainnet
mbongo config set network.chain_id 1
mbongo config set network.bootnodes '["enode://...@boot1.mbongo.io:30303"]'

# 5. Generate NEW keys for mainnet (don't reuse testnet keys)
mbongo wallet create --keystore ~/.mbongo/keys/mainnet-signing-key.json --type validator

# 6. Initialize new slashing protection
rm ~/.mbongo/data/slashing_protection.db
mbongo validator init-slashing-protection

# 7. Start mainnet validator
sudo systemctl start mbongo-validator
```

---

## 8. Troubleshooting

### 8.1 Common Problems and Solutions

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    TROUBLESHOOTING GUIDE                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   This section covers the 10 most common validator issues.                 │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

#### Problem 1: Node Not Syncing

**Symptoms:**
- Sync status shows 0% progress
- No new blocks being processed
- Peer count is 0

**Solutions:**
```bash
# Check peer connections
mbongo node peers

# If no peers, check firewall
sudo ufw status
# Ensure 30303/tcp and 30303/udp are allowed

# Check if bootnodes are reachable
ping boot1.mbongo.io

# Restart with verbose logging
mbongo node start --log-level debug

# Clear peers and reconnect
mbongo node peers clear
mbongo node start
```

---

#### Problem 2: RPC Timeout

**Symptoms:**
- CLI commands hang or timeout
- "Connection refused" errors
- RPC queries fail

**Solutions:**
```bash
# Check if RPC is enabled
grep -A5 "\[rpc\]" ~/.mbongo/config/config.toml

# Verify RPC is listening
netstat -tlnp | grep 8545

# Test RPC locally
curl -X POST http://127.0.0.1:8545 \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'

# If not responding, check logs
journalctl -u mbongo-validator | grep -i rpc

# Restart validator
sudo systemctl restart mbongo-validator
```

---

#### Problem 3: Slashing Protection Database Missing

**Symptoms:**
- "Slashing protection not found" error
- Validator refuses to start
- Warning about double-signing risk

**Solutions:**
```bash
# Initialize new slashing protection (ONLY if you've never run before)
mbongo validator init-slashing-protection \
    --db-path ~/.mbongo/data/slashing_protection.db

# If migrating from another machine, import backup
mbongo validator import-slashing-protection \
    --db-path ~/.mbongo/data/slashing_protection.db \
    --input ./slashing_protection_backup.json

# Verify database exists
ls -la ~/.mbongo/data/slashing_protection.db
```

---

#### Problem 4: Wrong Keystore Password

**Symptoms:**
- "Invalid password" or "Decryption failed"
- Validator can't unlock key
- Authentication errors

**Solutions:**
```bash
# Test password manually
mbongo wallet unlock \
    --keystore ~/.mbongo/keys/signing-key.json

# If password forgotten:
# 1. Recover from mnemonic (if you have it)
mbongo wallet recover \
    --mnemonic "word1 word2 word3..." \
    --keystore ~/.mbongo/keys/new-signing-key.json

# 2. Or restore from backup keystore
cp /backup/signing-key.json ~/.mbongo/keys/

# Update config to point to correct keystore
mbongo config set validator.keystore "~/.mbongo/keys/new-signing-key.json"
```

---

#### Problem 5: Peer Connection Issues

**Symptoms:**
- Low peer count (<5)
- Frequent disconnections
- Sync stalls

**Solutions:**
```bash
# Check current peer count
mbongo node peers | wc -l

# Add more bootnodes
mbongo config set network.bootnodes '["enode://...@boot1.mbongo.io:30303","enode://...@boot2.mbongo.io:30303"]'

# Increase max peers
mbongo config set network.max_peers 100

# Check NAT/firewall
# Ensure port 30303 is forwarded if behind NAT

# Clear bad peers
mbongo node peers clear

# Restart
sudo systemctl restart mbongo-validator
```

---

#### Problem 6: Corrupted Database

**Symptoms:**
- "Database corruption" errors
- Node crashes on startup
- Invalid block errors

**Solutions:**
```bash
# Stop validator
sudo systemctl stop mbongo-validator

# Back up corrupted data (for debugging)
mv ~/.mbongo/data/chaindata ~/.mbongo/data/chaindata-corrupted

# Resync from scratch
mbongo node start --sync-mode full

# Or restore from snapshot (if available)
wget https://snapshots.mbongo.io/latest.tar.gz
tar -xzf latest.tar.gz -C ~/.mbongo/data/

# Restart validator
sudo systemctl start mbongo-validator
```

---

#### Problem 7: Outdated Client Version

**Symptoms:**
- "Unsupported protocol version" errors
- Network forks
- Peer rejections

**Solutions:**
```bash
# Check current version
mbongo --version

# Check latest version
curl -s https://api.github.com/repos/mbongo-chain/mbongo/releases/latest | jq -r '.tag_name'

# Update via cargo
cargo install mbongo-cli mbongo-node mbongo-validator --force

# Or download binary
wget https://github.com/mbongo-chain/mbongo/releases/latest/download/mbongo-linux-amd64.tar.gz
tar -xzf mbongo-linux-amd64.tar.gz
sudo mv mbongo* /usr/local/bin/

# Restart
sudo systemctl restart mbongo-validator
```

---

#### Problem 8: Memory Spikes / OOM

**Symptoms:**
- Validator killed by OOM killer
- System becomes unresponsive
- High memory usage

**Solutions:**
```bash
# Check memory usage
free -h
htop

# Set memory limits in systemd
sudo systemctl edit mbongo-validator
# Add:
# [Service]
# MemoryMax=12G

# Reduce cache size in config
mbongo config set node.cache_size 1024

# Add swap (if not present)
sudo fallocate -l 8G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab

# Restart
sudo systemctl restart mbongo-validator
```

---

#### Problem 9: Firewall Blocking P2P

**Symptoms:**
- 0 peers despite correct config
- "Connection timed out" in logs
- Can't reach bootnodes

**Solutions:**
```bash
# Check UFW status
sudo ufw status

# Add P2P rules
sudo ufw allow 30303/tcp
sudo ufw allow 30303/udp

# If behind NAT, configure port forwarding on router
# Forward external:30303 -> internal:30303 (TCP+UDP)

# Check if port is open externally
# Visit: https://www.yougetsignal.com/tools/open-ports/
# Or use:
nc -zv your-server-ip 30303

# Check cloud provider firewall (AWS Security Groups, GCP Firewall, etc.)
```

---

#### Problem 10: Log Rotation Setup

**Symptoms:**
- Disk filling up with logs
- `/var/log` or `~/.mbongo/logs` consuming GB of space

**Solutions:**
```bash
# Create logrotate config
sudo tee /etc/logrotate.d/mbongo << 'EOF'
/home/mbongo/.mbongo/logs/*.log {
    daily
    rotate 7
    compress
    delaycompress
    missingok
    notifempty
    create 640 mbongo mbongo
    postrotate
        systemctl reload mbongo-validator > /dev/null 2>&1 || true
    endscript
}
EOF

# Test logrotate
sudo logrotate -d /etc/logrotate.d/mbongo

# Force rotation
sudo logrotate -f /etc/logrotate.d/mbongo

# Check disk usage
df -h
du -sh ~/.mbongo/logs/
```

---

## 9. Cross-References

### 9.1 Related Documentation

| Document | Description |
|----------|-------------|
| [node_setup_overview.md](./node_setup_overview.md) | General node setup overview |
| [compute_provider_setup.md](./compute_provider_setup.md) | Compute provider setup guide |
| [cli_wallet.md](./cli_wallet.md) | Wallet CLI commands |
| [cli_validator.md](./cli_validator.md) | Validator CLI commands |
| [cli_node.md](./cli_node.md) | Node CLI commands |
| [cli_config.md](./cli_config.md) | Configuration CLI commands |

### 9.2 SDK Documentation

| Document | Description |
|----------|-------------|
| [rust_sdk_overview.md](./rust_sdk_overview.md) | Rust SDK reference |
| [ts_sdk_overview.md](./ts_sdk_overview.md) | TypeScript SDK reference |
| [rpc_overview.md](./rpc_overview.md) | RPC API reference |

### 9.3 Architecture Documentation

| Document | Description |
|----------|-------------|
| [architecture_master_overview.md](./architecture_master_overview.md) | Full architecture overview |
| [consensus_validation.md](./consensus_validation.md) | Consensus mechanism details |
| [staking_model.md](./staking_model.md) | Staking economics |
| [economic_security.md](./economic_security.md) | Economic security model |

### 9.4 Quick Reference Card

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    VALIDATOR QUICK REFERENCE                                │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   INITIAL SETUP                                                             │
│   ─────────────                                                             │
│   cargo install mbongo-cli mbongo-node mbongo-validator                    │
│   mbongo config init --network mainnet --validator                         │
│   mbongo wallet create --keystore ./signing-key.json --type validator      │
│   mbongo validator init-slashing-protection                                │
│                                                                             │
│   START VALIDATOR                                                           │
│   ───────────────                                                           │
│   sudo systemctl enable mbongo-validator                                   │
│   sudo systemctl start mbongo-validator                                    │
│                                                                             │
│   MONITORING                                                                │
│   ──────────                                                                │
│   mbongo validator status                                                  │
│   mbongo node sync-status                                                  │
│   journalctl -u mbongo-validator -f                                        │
│                                                                             │
│   STAKE MANAGEMENT                                                          │
│   ────────────────                                                          │
│   mbongo validator stake --amount 50000                                    │
│   mbongo validator unstake --amount 10000                                  │
│   mbongo validator withdraw-rewards                                        │
│                                                                             │
│   KEY BACKUP                                                                │
│   ──────────                                                                │
│   mbongo wallet export --keystore ./signing-key.json --output ./backup/    │
│   mbongo validator export-slashing-protection --output ./slashing.json     │
│                                                                             │
│   PORTS                                                                     │
│   ─────                                                                     │
│   P2P:       30303 (TCP/UDP) - ALLOW                                       │
│   RPC:       8545 (TCP) - LOCAL ONLY                                       │
│   Metrics:   9090 (TCP) - LOCAL ONLY                                       │
│                                                                             │
│   MINIMUM REQUIREMENTS                                                      │
│   ────────────────────                                                      │
│   Stake:     50,000 MBO                                                    │
│   CPU:       4 cores                                                       │
│   RAM:       8 GB                                                          │
│   Storage:   512 GB NVMe                                                   │
│   Network:   50 Mbps                                                       │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

*This document provides complete instructions for setting up a Mbongo Chain validator. For questions or support, visit the official documentation or community channels.*

