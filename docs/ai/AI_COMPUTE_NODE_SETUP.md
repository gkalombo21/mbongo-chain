# 🧠 Mbongo Chain — AI Compute Node Setup Guide

## 1. Introduction

The **AI Compute Node** is a participant in the Mbongo Chain network that contributes GPU or CPU power to execute useful computations — such as AI model training, inference, or scientific simulations.

These nodes are rewarded in **MBG tokens** through the Proof of Useful Work (PoUW) consensus mechanism.

---

## 2. Requirements

| Component | Minimum Spec | Recommended |
|------------|---------------|--------------|
| CPU | Quad-core 2.5 GHz | Ryzen 7 / i7 or better |
| GPU | 6 GB VRAM (CUDA/OpenCL) | RTX 3070 / RX 6800+ |
| RAM | 8 GB | 16 GB+ |
| Disk | 50 GB free | SSD 250 GB |
| OS | Windows, Linux, macOS | Ubuntu 22.04 LTS |
| Go | v1.21+ | Latest |
| Network | Stable connection | 10 Mbps upload/download |

---

## 3. Installation

### Step 1 — Clone the repository

```bash
git clone https://github.com/gkalombo21/mbongo-chain.git
cd mbongo-chain
```

| Issue | Fix |

> |--------|-----|
> | Node not syncing | Check network/firewall settings |
> | Low performance | Update GPU drivers or reduce batch size |
> | Invalid proof | Ensure clock is synchronized (NTP) |

| Feature | Description | ETA |

> |----------|--------------|-----|
> | GPU Benchmark Tool | Auto-tune GPU settings for PoUW jobs | Q1 2026 |
> | Cross-platform Daemon | Lightweight background process | Q2 2026 |
> | Node Reputation System | AI-based scoring for reliability | Q3 2026 |
