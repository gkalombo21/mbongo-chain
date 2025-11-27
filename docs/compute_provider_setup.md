# Mbongo Chain — Compute Provider Setup Guide

> **Document Version:** 1.0.0  
> **Last Updated:** November 2025  
> **Target Audience:** Compute Providers, GPU Operators, Data Center Administrators

---

## Table of Contents

1. [Introduction](#1-introduction)
2. [Hardware & Driver Requirements](#2-hardware--driver-requirements)
3. [Software Requirements](#3-software-requirements)
4. [Environment Setup](#4-environment-setup)
5. [Registering as a Compute Provider](#5-registering-as-a-compute-provider)
6. [Running the Provider Node](#6-running-the-provider-node)
7. [Security & Fraud Prevention](#7-security--fraud-prevention)
8. [Earnings & Reward Mechanics](#8-earnings--reward-mechanics)
9. [Troubleshooting](#9-troubleshooting)
10. [Cross-References](#10-cross-references)

---

## 1. Introduction

### What is a Compute Provider?

In Mbongo Chain's **Proof of Useful Work (PoUW)** consensus mechanism, **Compute Providers** are network participants who contribute computational resources to perform real-world useful work. Unlike traditional Proof of Work (PoW) systems that waste energy on meaningless hash puzzles, PoUW directs computational power toward productive tasks.

```
┌─────────────────────────────────────────────────────────────────────┐
│                    PROOF OF USEFUL WORK (PoUW)                      │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│   ┌─────────────────┐         ┌─────────────────┐                  │
│   │  Task Submitter │         │   Compute       │                  │
│   │  (User/DApp)    │────────▶│   Provider      │                  │
│   └─────────────────┘         └────────┬────────┘                  │
│           │                            │                            │
│           │  Submit Task               │  Execute Task              │
│           ▼                            ▼                            │
│   ┌─────────────────────────────────────────────────────────────┐  │
│   │                      MBONGO NETWORK                          │  │
│   │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐          │  │
│   │  │ Task Queue  │──│ Scheduler   │──│ Verifier    │          │  │
│   │  └─────────────┘  └─────────────┘  └─────────────┘          │  │
│   └─────────────────────────────────────────────────────────────┘  │
│                                │                                    │
│                                ▼                                    │
│   ┌─────────────────────────────────────────────────────────────┐  │
│   │              REWARDS DISTRIBUTED TO PROVIDER                 │  │
│   │                    (MBO Tokens)                              │  │
│   └─────────────────────────────────────────────────────────────┘  │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Supported Hardware

Mbongo Chain supports a diverse range of computational hardware:

| Hardware Type | Description | Use Cases |
|---------------|-------------|-----------|
| **GPU** | Graphics Processing Units | AI inference, ML training, 3D rendering |
| **TPU** | Tensor Processing Units | Large-scale ML training, inference |
| **CPU** | Central Processing Units | General computation, scientific simulation |
| **FPGA** | Field-Programmable Gate Arrays | Custom acceleration, cryptography |
| **ASIC** | Application-Specific Integrated Circuits | Specialized high-throughput tasks |

### Supported Workloads

```
┌─────────────────────────────────────────────────────────────────────┐
│                    SUPPORTED WORKLOAD TYPES                         │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  🤖 AI INFERENCE                                             │   │
│  │  • Large Language Model (LLM) inference                      │   │
│  │  • Image classification & object detection                   │   │
│  │  • Natural language processing                               │   │
│  │  • Speech recognition & synthesis                            │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                                                                     │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  🧠 MACHINE LEARNING TRAINING                                │   │
│  │  • Neural network training                                   │   │
│  │  • Federated learning tasks                                  │   │
│  │  • Model fine-tuning                                         │   │
│  │  • Hyperparameter optimization                               │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                                                                     │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  🎨 3D RENDERING                                             │   │
│  │  • Ray tracing & path tracing                                │   │
│  │  • Animation rendering                                       │   │
│  │  • Visual effects processing                                 │   │
│  │  • Real-time graphics pipelines                              │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                                                                     │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  🔬 SCIENTIFIC SIMULATION                                    │   │
│  │  • Molecular dynamics                                        │   │
│  │  • Climate modeling                                          │   │
│  │  • Fluid dynamics (CFD)                                      │   │
│  │  • Genome sequencing                                         │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                                                                     │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  🔐 CRYPTOGRAPHY                                             │   │
│  │  • Zero-knowledge proof generation                           │   │
│  │  • Homomorphic encryption                                    │   │
│  │  • Multi-party computation                                   │   │
│  │  • Signature aggregation                                     │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Provider Responsibilities

As a Compute Provider, you are responsible for:

| Responsibility | Description | Penalty for Failure |
|----------------|-------------|---------------------|
| **Correctness** | Produce accurate, verifiable computation results | Slashing (1,000 MBO per invalid proof) |
| **Uptime** | Maintain high availability (>95% recommended) | Reduced task allocation |
| **Valid Receipts** | Submit cryptographically valid compute receipts | Task rejection, potential slashing |
| **Timely Completion** | Complete tasks within allocated time windows | Timeout penalties |
| **Resource Honesty** | Accurately report hardware capabilities | Registration rejection |

### Reward Model

Compute Providers earn **MBO tokens** for successfully completing valid compute tasks:

```
┌─────────────────────────────────────────────────────────────────────┐
│                      REWARD FLOW                                    │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌──────────┐    ┌──────────┐    ┌──────────┐    ┌──────────┐     │
│  │  Task    │───▶│  Execute │───▶│  Submit  │───▶│  Receive │     │
│  │ Assigned │    │  Compute │    │  Receipt │    │  Reward  │     │
│  └──────────┘    └──────────┘    └──────────┘    └──────────┘     │
│                                                                     │
│  Reward Calculation:                                                │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  MBO Earned = (Compute Units × Base Rate) × Quality Bonus    │   │
│  │                                                               │   │
│  │  Where:                                                       │   │
│  │  • Compute Units = Task complexity metric                    │   │
│  │  • Base Rate = Network-determined MBO per compute unit       │   │
│  │  • Quality Bonus = 1.0 - 1.25x based on performance          │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                                                                     │
│  Example:                                                           │
│  • AI Inference Task: 100 compute units × 0.01 MBO × 1.1 = 1.1 MBO │
│  • ML Training Task: 10,000 compute units × 0.01 MBO × 1.2 = 120 MBO│
│  • 3D Render Task: 500 compute units × 0.01 MBO × 1.0 = 5.0 MBO    │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

## 2. Hardware & Driver Requirements

### GPU Requirements

#### Minimum GPU Specifications

| Specification | Minimum | Recommended |
|---------------|---------|-------------|
| **VRAM** | 8 GB | 16+ GB |
| **Compute Capability** | 6.0 (Pascal) | 7.5+ (Turing/Ampere) |
| **Memory Bandwidth** | 200 GB/s | 400+ GB/s |
| **FP32 Performance** | 5 TFLOPS | 15+ TFLOPS |
| **PCIe** | Gen 3 x8 | Gen 4 x16 |

#### GPU Tiers by Workload Type

| Tier | AI Inference | ML Training | 3D Rendering | Example GPUs |
|------|--------------|-------------|--------------|--------------|
| **Entry** | ✅ Basic | ❌ | ✅ Basic | RTX 3060, RX 6700 XT |
| **Mid** | ✅ Good | ✅ Small models | ✅ Good | RTX 3080, RTX 4070, RX 7800 XT |
| **High** | ✅ Excellent | ✅ Medium models | ✅ Excellent | RTX 4090, A4000, RX 7900 XTX |
| **Enterprise** | ✅ Optimal | ✅ Large models | ✅ Optimal | A100, H100, MI250X, MI300X |
| **Data Center** | ✅ Maximum | ✅ Distributed | ✅ Maximum | H100 SXM, MI300X OAM |

```
┌─────────────────────────────────────────────────────────────────────┐
│                    GPU TIER COMPARISON                              │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ENTRY TIER (8-12 GB VRAM)                                         │
│  ├── RTX 3060 12GB          │████████░░░░│ 12.74 TFLOPS            │
│  ├── RTX 4060 Ti 16GB       │█████████░░░│ 22.06 TFLOPS            │
│  └── RX 6700 XT 12GB        │███████░░░░░│ 13.21 TFLOPS            │
│                                                                     │
│  MID TIER (12-16 GB VRAM)                                          │
│  ├── RTX 3080 12GB          │████████████│ 29.77 TFLOPS            │
│  ├── RTX 4070 Ti 12GB       │████████████│ 40.09 TFLOPS            │
│  └── RX 7800 XT 16GB        │██████████░░│ 37.32 TFLOPS            │
│                                                                     │
│  HIGH TIER (24+ GB VRAM)                                           │
│  ├── RTX 4090 24GB          │████████████████│ 82.58 TFLOPS        │
│  ├── RTX A6000 48GB         │████████████████│ 38.71 TFLOPS        │
│  └── RX 7900 XTX 24GB       │█████████████░░░│ 61.42 TFLOPS        │
│                                                                     │
│  ENTERPRISE TIER (40-80 GB VRAM)                                   │
│  ├── A100 80GB              │██████████████████│ 19.5 TFLOPS (FP32)│
│  │                          │████████████████████████│ 312 TFLOPS (Tensor)│
│  ├── H100 80GB              │█████████████████████│ 67 TFLOPS (FP32)│
│  │                          │████████████████████████████│ 1,979 TFLOPS (Tensor)│
│  └── MI300X 192GB           │████████████████████████████████│ 5,300 TFLOPS (Mixed)│
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

#### NVIDIA GPU Recommended Models

| Use Case | Budget | Mid-Range | Professional |
|----------|--------|-----------|--------------|
| **AI Inference** | RTX 4060 Ti | RTX 4080 | A4000/A5000 |
| **ML Training** | RTX 4070 Ti | RTX 4090 | A100/H100 |
| **3D Rendering** | RTX 3060 | RTX 4070 | RTX A4500 |

#### AMD GPU Recommended Models

| Use Case | Budget | Mid-Range | Professional |
|----------|--------|-----------|--------------|
| **AI Inference** | RX 7700 XT | RX 7800 XT | MI100 |
| **ML Training** | RX 7800 XT | RX 7900 XTX | MI250X/MI300X |
| **3D Rendering** | RX 6700 XT | RX 7900 XT | PRO W7800 |

### CPU-Only Tier

For providers without GPU hardware:

| Specification | Minimum | Recommended |
|---------------|---------|-------------|
| **Cores** | 8 | 16+ |
| **Threads** | 16 | 32+ |
| **Base Clock** | 3.0 GHz | 3.5+ GHz |
| **L3 Cache** | 16 MB | 32+ MB |
| **RAM** | 32 GB | 64+ GB |
| **AVX Support** | AVX2 | AVX-512 preferred |

> ⚠️ **Note**: CPU-only providers receive lower-priority task assignments and reduced rewards compared to GPU providers.

**Recommended CPU Models:**
- AMD EPYC 7003 series
- AMD Ryzen 9 7950X
- Intel Xeon Scalable (3rd/4th Gen)
- Intel Core i9-13900K/14900K

### TPU Support

| TPU Type | Support Level | Notes |
|----------|---------------|-------|
| **Google TPU v4** | ✅ Full | Requires GCP account |
| **Google TPU v5** | ✅ Full | Requires GCP account |
| **Edge TPU** | 🟡 Limited | Inference only |

**TPU Configuration Requirements:**
- Google Cloud Platform account with TPU quota
- TPU VM setup or TPU Node configuration
- Mbongo TPU runtime adapter installed

### FPGA Support

| FPGA Platform | Support Level | Driver Required |
|---------------|---------------|-----------------|
| **Xilinx Alveo U50** | ✅ Full | Vitis Runtime |
| **Xilinx Alveo U250** | ✅ Full | Vitis Runtime |
| **Intel Stratix 10** | ✅ Full | Intel FPGA SDK |
| **Intel Agilex** | 🟡 Beta | Intel FPGA SDK |

### ASIC Support

| ASIC Type | Support Level | Use Case |
|-----------|---------------|----------|
| **Custom ZK ASICs** | ✅ Full | Zero-knowledge proofs |
| **AI Inference ASICs** | 🟡 Limited | Specialized inference |

### Driver Requirements

#### NVIDIA CUDA

| Component | Minimum Version | Recommended |
|-----------|-----------------|-------------|
| **NVIDIA Driver** | 525.x | 535.x+ |
| **CUDA Toolkit** | 11.8 | 12.2+ |
| **cuDNN** | 8.6 | 8.9+ |
| **TensorRT** | 8.5 | 8.6+ |

#### AMD ROCm

| Component | Minimum Version | Recommended |
|-----------|-----------------|-------------|
| **ROCm** | 5.4 | 6.0+ |
| **ROCm Driver** | Included | Latest |
| **MIOpen** | 2.18 | 2.20+ |

#### Other Drivers

| Hardware | Driver/Runtime |
|----------|----------------|
| **TPU** | Google Cloud TPU Runtime |
| **FPGA (Xilinx)** | Xilinx Runtime (XRT) 2.14+ |
| **FPGA (Intel)** | Intel FPGA SDK 21.4+ |
| **CPU** | No special driver required |

---

## 3. Software Requirements

### Core Software Stack

| Software | Version | Purpose |
|----------|---------|---------|
| **Python** | 3.10+ | Runtime environment |
| **CUDA/ROCm** | See above | GPU acceleration |
| **OpenCL** | 3.0+ | Cross-platform compute |
| **Docker** | 24.0+ | Container isolation (optional) |
| **Mbongo Compute CLI** | Latest | Provider management |

### Ubuntu Dependencies

```bash
# System packages
sudo apt install -y \
    build-essential \
    cmake \
    git \
    curl \
    wget \
    pkg-config \
    libssl-dev \
    libffi-dev \
    python3-dev \
    python3-pip \
    python3-venv \
    ocl-icd-opencl-dev \
    opencl-headers \
    clinfo \
    htop \
    nvtop \
    jq

# Python packages
pip3 install --user \
    numpy \
    scipy \
    torch \
    tensorflow \
    onnxruntime-gpu \
    pycryptodome \
    aiohttp \
    websockets
```

### Windows Dependencies

```powershell
# Install via Chocolatey
choco install -y `
    python310 `
    git `
    cmake `
    visualstudio2022buildtools `
    cuda `
    wget `
    jq

# Python packages
pip install `
    numpy `
    scipy `
    torch `
    tensorflow `
    onnxruntime-gpu `
    pycryptodome `
    aiohttp `
    websockets
```

### Software Architecture

```
┌─────────────────────────────────────────────────────────────────────┐
│                    COMPUTE PROVIDER SOFTWARE STACK                  │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │                   USER SPACE                                 │   │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐          │   │
│  │  │ Mbongo CLI  │  │  Python     │  │  Docker     │          │   │
│  │  │ (Provider)  │  │  Runtime    │  │  (Optional) │          │   │
│  │  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘          │   │
│  │         │                │                │                  │   │
│  │         └────────────────┼────────────────┘                  │   │
│  │                          │                                   │   │
│  │  ┌───────────────────────▼───────────────────────────────┐  │   │
│  │  │              COMPUTE FRAMEWORKS                        │  │   │
│  │  │  ┌────────┐  ┌────────┐  ┌────────┐  ┌────────┐       │  │   │
│  │  │  │PyTorch │  │TensorFlow│ │ ONNX   │  │OpenCL  │       │  │   │
│  │  │  └────────┘  └────────┘  └────────┘  └────────┘       │  │   │
│  │  └───────────────────────┬───────────────────────────────┘  │   │
│  └──────────────────────────┼──────────────────────────────────┘   │
│                             │                                       │
│  ┌──────────────────────────▼──────────────────────────────────┐   │
│  │                   DRIVER LAYER                               │   │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐          │   │
│  │  │ CUDA/cuDNN  │  │   ROCm      │  │   OpenCL    │          │   │
│  │  │ (NVIDIA)    │  │   (AMD)     │  │   (Generic) │          │   │
│  │  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘          │   │
│  └─────────┼────────────────┼────────────────┼─────────────────┘   │
│            │                │                │                      │
│  ┌─────────▼────────────────▼────────────────▼─────────────────┐   │
│  │                   KERNEL SPACE                               │   │
│  │              GPU/TPU/FPGA/ASIC Hardware                      │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

## 4. Environment Setup

### Installation Pipeline

```
┌─────────────────────────────────────────────────────────────────────┐
│                 INSTALLATION PIPELINE                               │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  STEP 1              STEP 2              STEP 3              STEP 4 │
│  ┌────────┐          ┌────────┐          ┌────────┐          ┌────────┐
│  │ System │  ───────▶│  GPU   │  ───────▶│ Client │  ───────▶│Provider│
│  │ Setup  │          │Drivers │          │ Setup  │          │Register│
│  └────────┘          └────────┘          └────────┘          └────────┘
│      │                   │                   │                   │
│      ▼                   ▼                   ▼                   ▼
│  ┌────────┐          ┌────────┐          ┌────────┐          ┌────────┐
│  │• Update│          │• CUDA  │          │• Mbongo│          │• Keys  │
│  │• Deps  │          │• ROCm  │          │  CLI   │          │• Config│
│  │• Python│          │• OpenCL│          │• Python│          │• RPC   │
│  └────────┘          └────────┘          └────────┘          └────────┘
│                                                                     │
│  Timeline: ~30 minutes (varies by driver download speed)            │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Ubuntu 22.04 LTS Setup

#### Step 1: System Update & Base Packages

```bash
#!/bin/bash
# Mbongo Compute Provider Setup - Ubuntu

set -e

echo "=== Step 1: System Update ==="
sudo apt update && sudo apt upgrade -y

echo "=== Installing Base Packages ==="
sudo apt install -y \
    build-essential \
    cmake \
    git \
    curl \
    wget \
    pkg-config \
    libssl-dev \
    libffi-dev \
    python3.10 \
    python3.10-venv \
    python3-pip \
    software-properties-common

echo "=== Setting Python 3.10 as Default ==="
sudo update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.10 1

echo "=== Verifying Python ==="
python3 --version
pip3 --version
```

#### Step 2: GPU Driver Installation (NVIDIA)

```bash
#!/bin/bash
# NVIDIA CUDA Installation

echo "=== Installing NVIDIA Drivers ==="

# Add NVIDIA repository
wget https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2204/x86_64/cuda-keyring_1.1-1_all.deb
sudo dpkg -i cuda-keyring_1.1-1_all.deb
sudo apt update

# Install CUDA toolkit and drivers
sudo apt install -y cuda-toolkit-12-2 nvidia-driver-535

# Add CUDA to PATH
echo 'export PATH=/usr/local/cuda-12.2/bin:$PATH' >> ~/.bashrc
echo 'export LD_LIBRARY_PATH=/usr/local/cuda-12.2/lib64:$LD_LIBRARY_PATH' >> ~/.bashrc
source ~/.bashrc

# Install cuDNN
sudo apt install -y libcudnn8 libcudnn8-dev

# Verify installation
nvidia-smi
nvcc --version
```

#### Step 2 (Alternative): GPU Driver Installation (AMD ROCm)

```bash
#!/bin/bash
# AMD ROCm Installation

echo "=== Installing AMD ROCm ==="

# Add ROCm repository
sudo apt install -y ./amdgpu-install_6.0.60000-1_all.deb

# Install ROCm
sudo amdgpu-install --usecase=rocm,hip --no-dkms

# Add user to render and video groups
sudo usermod -a -G render,video $USER

# Add ROCm to PATH
echo 'export PATH=/opt/rocm/bin:$PATH' >> ~/.bashrc
echo 'export LD_LIBRARY_PATH=/opt/rocm/lib:$LD_LIBRARY_PATH' >> ~/.bashrc
source ~/.bashrc

# Verify installation
rocm-smi
hipcc --version
```

#### Step 3: Mbongo Compute CLI Installation

```bash
#!/bin/bash
# Mbongo Compute CLI Installation

echo "=== Installing Mbongo Compute CLI ==="

# Create provider directories
mkdir -p ~/.mbongo/{bin,config,data,logs,compute}
chmod 700 ~/.mbongo

# Download Mbongo Compute CLI
MBONGO_VERSION="v1.0.0"
wget -O mbongo-compute.tar.gz \
    "https://github.com/mbongo-chain/mbongo-chain/releases/download/${MBONGO_VERSION}/mbongo-compute-linux-amd64.tar.gz"

# Verify checksum
wget -O checksums.txt \
    "https://github.com/mbongo-chain/mbongo-chain/releases/download/${MBONGO_VERSION}/checksums.txt"
sha256sum -c checksums.txt --ignore-missing

# Extract and install
tar -xzf mbongo-compute.tar.gz
mv mbongo-compute ~/.mbongo/bin/
chmod +x ~/.mbongo/bin/mbongo-compute

# Add to PATH
echo 'export PATH="$HOME/.mbongo/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# Verify installation
mbongo-compute --version

# Install Python dependencies for compute tasks
pip3 install --user \
    numpy \
    scipy \
    torch \
    tensorflow \
    onnxruntime-gpu \
    pycryptodome \
    aiohttp \
    websockets

echo "=== Installation Complete ==="
```

### Windows PowerShell Setup

#### Step 1: System Setup

```powershell
# Run as Administrator
# Mbongo Compute Provider Setup - Windows

Write-Host "=== Step 1: Installing Base Packages ===" -ForegroundColor Cyan

# Install Chocolatey if not present
if (!(Get-Command choco -ErrorAction SilentlyContinue)) {
    Set-ExecutionPolicy Bypass -Scope Process -Force
    [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072
    iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
}

# Install required packages
choco install -y `
    python310 `
    git `
    cmake `
    visualstudio2022buildtools `
    wget `
    jq `
    7zip

# Refresh environment
refreshenv

Write-Host "=== Verifying Python ===" -ForegroundColor Cyan
python --version
pip --version
```

#### Step 2: GPU Driver Installation (NVIDIA - Windows)

```powershell
# NVIDIA CUDA Installation for Windows

Write-Host "=== Installing NVIDIA CUDA ===" -ForegroundColor Cyan

# Install CUDA via Chocolatey
choco install -y cuda --version=12.2.0

# Or download manually from NVIDIA website:
# https://developer.nvidia.com/cuda-12-2-0-download-archive

# Add CUDA to PATH
$cudaPath = "C:\Program Files\NVIDIA GPU Computing Toolkit\CUDA\v12.2"
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";$cudaPath\bin;$cudaPath\libnvvp", "User")
[Environment]::SetEnvironmentVariable("CUDA_PATH", $cudaPath, "User")

# Refresh PATH
$env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User")

# Verify installation
nvidia-smi
nvcc --version
```

#### Step 3: Mbongo Compute CLI Installation (Windows)

```powershell
# Mbongo Compute CLI Installation

Write-Host "=== Installing Mbongo Compute CLI ===" -ForegroundColor Cyan

# Create directories
$mbongoPath = "$env:USERPROFILE\.mbongo"
New-Item -ItemType Directory -Force -Path "$mbongoPath\bin"
New-Item -ItemType Directory -Force -Path "$mbongoPath\config"
New-Item -ItemType Directory -Force -Path "$mbongoPath\data"
New-Item -ItemType Directory -Force -Path "$mbongoPath\logs"
New-Item -ItemType Directory -Force -Path "$mbongoPath\compute"

# Download Mbongo Compute CLI
$MBONGO_VERSION = "v1.0.0"
$downloadUrl = "https://github.com/mbongo-chain/mbongo-chain/releases/download/$MBONGO_VERSION/mbongo-compute-windows-amd64.zip"
$zipPath = "$env:TEMP\mbongo-compute.zip"

Invoke-WebRequest -Uri $downloadUrl -OutFile $zipPath

# Extract
Expand-Archive -Path $zipPath -DestinationPath "$mbongoPath\bin" -Force

# Add to PATH
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($currentPath -notlike "*$mbongoPath\bin*") {
    [Environment]::SetEnvironmentVariable("Path", "$currentPath;$mbongoPath\bin", "User")
}

# Refresh PATH
$env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User")

# Install Python dependencies
pip install `
    numpy `
    scipy `
    torch `
    tensorflow `
    onnxruntime-gpu `
    pycryptodome `
    aiohttp `
    websockets

Write-Host "=== Verifying Installation ===" -ForegroundColor Cyan
mbongo-compute --version

Write-Host "=== Installation Complete ===" -ForegroundColor Green
```

### Verification Commands

```bash
# Verify GPU detection
mbongo-compute hardware detect

# Expected output:
# ┌────────────────────────────────────────────────────────────┐
# │ Hardware Detection Results                                 │
# ├────────────────────────────────────────────────────────────┤
# │ GPU 0: NVIDIA RTX 4090                                     │
# │   VRAM: 24 GB                                              │
# │   Compute Capability: 8.9                                  │
# │   Driver: 535.104.05                                       │
# │   CUDA: 12.2                                               │
# │   Status: ✅ Compatible                                    │
# └────────────────────────────────────────────────────────────┘

# Run hardware benchmark
mbongo-compute benchmark --quick

# Check OpenCL devices
clinfo --list
```

---

## 5. Registering as a Compute Provider

### Step 1: Generate Provider Keys

```bash
# Generate provider keypair
mbongo-compute provider keygen \
    --output-dir ~/.mbongo/compute/keys

# Output:
# ✅ Provider signing key generated
# ✅ Provider public key: 0xabc123...
# ✅ Keys saved to ~/.mbongo/compute/keys/
```

### Step 2: Create compute_provider.toml

```bash
# Create the configuration file
cat > ~/.mbongo/config/compute_provider.toml << 'EOF'
# Mbongo Compute Provider Configuration
# Version: 1.0.0

[provider]
# Provider identifier (auto-generated or custom)
name = "MyComputeProvider"

# Provider public key (from keygen)
pubkey = "0xYOUR_PROVIDER_PUBLIC_KEY"

# Provider type: "gpu", "cpu", "tpu", "fpga", "asic"
type = "gpu"

# Provider signing key path
keystore_path = "~/.mbongo/compute/keys/provider_keystore.json"

[hardware]
# Hardware type
type = "nvidia_gpu"

# GPU model (auto-detected or manual)
model = "NVIDIA RTX 4090"

# VRAM in GB
vram_gb = 24

# Number of GPUs
gpu_count = 1

# Compute capability (NVIDIA) or GCN version (AMD)
compute_capability = "8.9"

[capabilities]
# Supported workload types
workloads = [
    "ai_inference",
    "ml_training",
    "3d_rendering",
    "scientific_simulation",
    "cryptography"
]

# Maximum model size (in parameters)
max_model_params = 70_000_000_000  # 70B parameters

# Supported precision
precision = ["fp32", "fp16", "bf16", "int8"]

# Supported frameworks
frameworks = ["pytorch", "tensorflow", "onnx", "tensorrt"]

[performance]
# Maximum parallel tasks
max_parallel_tasks = 4

# Maximum concurrent model loads
max_concurrent_models = 2

# Task timeout (seconds)
default_timeout = 3600

# Memory allocation limit (percentage of VRAM)
max_vram_allocation = 0.95

[pricing]
# Compute unit pricing (MBO per compute unit)
price_per_compute_unit = 0.01

# Minimum task price (MBO)
min_task_price = 0.001

# Priority task multiplier
priority_multiplier = 1.5

[network]
# Validator RPC endpoint
validator_rpc = "http://127.0.0.1:8545"

# P2P listen address
listen_addr = "0.0.0.0"

# P2P port
p2p_port = 40303

# Maximum peer connections
max_peers = 50

[receipt]
# Receipt submission endpoint
submission_endpoint = "http://127.0.0.1:8545"

# Retry attempts for failed submissions
retry_attempts = 3

# Retry delay (seconds)
retry_delay = 5

[logging]
# Log level
level = "info"

# Log file path
file = "~/.mbongo/logs/compute_provider.log"

# Enable structured JSON logging
json_format = false

[monitoring]
# Enable Prometheus metrics
metrics_enabled = true

# Metrics port
metrics_port = 9091

# Health check endpoint
health_port = 8080
EOF
```

### Configuration Options by Provider Type

#### GPU Provider Configuration

```toml
[hardware]
type = "nvidia_gpu"  # or "amd_gpu"
model = "NVIDIA RTX 4090"
vram_gb = 24
gpu_count = 1
compute_capability = "8.9"

# Multi-GPU configuration
[hardware.multi_gpu]
enabled = true
gpu_ids = [0, 1, 2, 3]  # GPU indices
distribution = "round_robin"  # or "load_balanced"
```

#### CPU Provider Configuration

```toml
[hardware]
type = "cpu"
model = "AMD EPYC 7763"
cores = 64
threads = 128
ram_gb = 256

[capabilities]
workloads = ["scientific_simulation", "cryptography"]
max_parallel_tasks = 32
```

#### TPU Provider Configuration

```toml
[hardware]
type = "tpu"
model = "Google TPU v4"
tpu_chips = 4
hbm_gb = 32

[tpu]
gcp_project = "your-gcp-project"
tpu_zone = "us-central1-a"
tpu_name = "your-tpu-node"
```

#### FPGA Provider Configuration

```toml
[hardware]
type = "fpga"
model = "Xilinx Alveo U250"
vendor = "xilinx"

[fpga]
bitstream_path = "/opt/xilinx/bitstreams/"
xrt_version = "2.14"
```

#### ASIC Provider Configuration

```toml
[hardware]
type = "asic"
model = "Custom ZK ASIC v1"

[asic]
driver_path = "/opt/asic/drivers/"
firmware_version = "1.2.3"
```

### Step 3: Register Provider on Network

```bash
# Register as a compute provider
mbongo-compute provider register \
    --config ~/.mbongo/config/compute_provider.toml \
    --stake 1000  # Minimum stake in MBO

# Output:
# ┌────────────────────────────────────────────────────────────┐
# │ Provider Registration                                      │
# ├────────────────────────────────────────────────────────────┤
# │ Provider ID: 0xabc123...                                   │
# │ Type: GPU (NVIDIA RTX 4090)                                │
# │ Capabilities: AI, ML, 3D, Science, Crypto                  │
# │ Stake: 1,000 MBO                                           │
# │ Status: ✅ Registered                                      │
# │ Activation: ~10 minutes (pending verification)             │
# └────────────────────────────────────────────────────────────┘

# Verify registration
mbongo-compute provider status \
    --config ~/.mbongo/config/compute_provider.toml
```

### Step 4: Connect to Validator RPC

```bash
# Test RPC connection
mbongo-compute provider test-connection \
    --rpc http://127.0.0.1:8545

# Expected output:
# ✅ RPC connection successful
# ✅ Block height: 1,234,567
# ✅ Provider recognized
# ✅ Task queue accessible
```

---

## 6. Running the Provider Node

### Starting the Provider

#### Ubuntu (Direct)

```bash
# Start provider in foreground (for testing)
mbongo-compute provider start \
    --config ~/.mbongo/config/compute_provider.toml

# Start with verbose logging
mbongo-compute provider start \
    --config ~/.mbongo/config/compute_provider.toml \
    --log-level debug
```

#### Ubuntu (systemd Service)

```bash
# Create systemd service
sudo tee /etc/systemd/system/mbongo-compute.service << 'EOF'
[Unit]
Description=Mbongo Compute Provider
After=network-online.target nvidia-persistenced.service
Wants=network-online.target
Requires=nvidia-persistenced.service

[Service]
Type=simple
User=YOUR_USERNAME
Group=YOUR_USERNAME
ExecStart=/home/YOUR_USERNAME/.mbongo/bin/mbongo-compute provider start \
    --config /home/YOUR_USERNAME/.mbongo/config/compute_provider.toml
Restart=always
RestartSec=10
LimitNOFILE=65535
Nice=-10

# GPU access
Environment="CUDA_VISIBLE_DEVICES=0"
Environment="NVIDIA_VISIBLE_DEVICES=all"
Environment="NVIDIA_DRIVER_CAPABILITIES=compute,utility"

# Memory limits
MemoryMax=28G
MemoryHigh=24G

# Security
NoNewPrivileges=true
ProtectSystem=strict
ProtectHome=read-only
ReadWritePaths=/home/YOUR_USERNAME/.mbongo

[Install]
WantedBy=multi-user.target
EOF

# Replace YOUR_USERNAME
sudo sed -i "s/YOUR_USERNAME/$USER/g" /etc/systemd/system/mbongo-compute.service

# Reload and start
sudo systemctl daemon-reload
sudo systemctl enable mbongo-compute
sudo systemctl start mbongo-compute

# Check status
sudo systemctl status mbongo-compute
```

#### Windows (NSSM Service)

```powershell
# Install NSSM
choco install -y nssm

# Create service
$serviceName = "MbongoCompute"
$binaryPath = "$env:USERPROFILE\.mbongo\bin\mbongo-compute.exe"
$configPath = "$env:USERPROFILE\.mbongo\config\compute_provider.toml"
$logPath = "$env:USERPROFILE\.mbongo\logs"

nssm install $serviceName $binaryPath
nssm set $serviceName AppParameters "provider start --config `"$configPath`""
nssm set $serviceName AppDirectory "$env:USERPROFILE\.mbongo"
nssm set $serviceName AppStdout "$logPath\compute-stdout.log"
nssm set $serviceName AppStderr "$logPath\compute-stderr.log"
nssm set $serviceName AppRotateFiles 1
nssm set $serviceName AppRotateBytes 52428800
nssm set $serviceName Start SERVICE_AUTO_START

# Set environment variables for CUDA
nssm set $serviceName AppEnvironmentExtra "CUDA_VISIBLE_DEVICES=0"

# Start service
nssm start $serviceName

# Check status
Get-Service $serviceName
```

### Log Monitoring

```bash
# View real-time logs (Ubuntu)
sudo journalctl -u mbongo-compute -f --no-hostname

# Or direct log file
tail -f ~/.mbongo/logs/compute_provider.log

# Search for specific events
grep -E "task_completed|task_failed|receipt_submitted" ~/.mbongo/logs/compute_provider.log

# Monitor GPU usage alongside logs
watch -n 1 'nvidia-smi; echo "---"; tail -5 ~/.mbongo/logs/compute_provider.log'
```

### Receipt Submission

```
┌─────────────────────────────────────────────────────────────────────┐
│                    RECEIPT SUBMISSION FLOW                          │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌──────────┐    ┌──────────┐    ┌──────────┐    ┌──────────┐     │
│  │  Task    │───▶│  Execute │───▶│  Generate│───▶│  Submit  │     │
│  │ Received │    │  Compute │    │  Receipt │    │  Receipt │     │
│  └──────────┘    └──────────┘    └──────────┘    └──────────┘     │
│       │               │               │               │            │
│       ▼               ▼               ▼               ▼            │
│  ┌──────────────────────────────────────────────────────────────┐ │
│  │  Receipt Contents:                                            │ │
│  │  • Task ID                                                    │ │
│  │  • Computation hash                                           │ │
│  │  • Output hash                                                │ │
│  │  • Execution time                                             │ │
│  │  • Provider signature                                         │ │
│  │  • Hardware attestation                                       │ │
│  └──────────────────────────────────────────────────────────────┘ │
│                                                                     │
│  Receipt Submission Status:                                        │
│  • ✅ Accepted → Reward queued                                     │
│  • ⚠️ Pending verification → Wait for confirmation                 │
│  • ❌ Rejected → Check error, possible slashing                    │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Handling Failed Tasks

```bash
# View failed tasks
mbongo-compute provider tasks --status failed

# Retry specific failed task
mbongo-compute provider task-retry --task-id TASK_ID

# View task failure details
mbongo-compute provider task-info --task-id TASK_ID --verbose

# Clear failed task queue (with caution)
mbongo-compute provider tasks-clear --status failed --confirm
```

### Payment Cycles

```
┌─────────────────────────────────────────────────────────────────────┐
│                    PAYMENT CYCLE                                    │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  EPOCH N (6.4 minutes)                                       │   │
│  │  ┌───────┐ ┌───────┐ ┌───────┐ ┌───────┐ ┌───────┐         │   │
│  │  │Task 1 │ │Task 2 │ │Task 3 │ │Task 4 │ │Task 5 │         │   │
│  │  │ ✅    │ │ ✅    │ │ ✅    │ │ ❌    │ │ ✅    │         │   │
│  │  └───────┘ └───────┘ └───────┘ └───────┘ └───────┘         │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                              │                                      │
│                              ▼                                      │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  END OF EPOCH: Receipt Aggregation                           │   │
│  │  • Total valid tasks: 4                                      │   │
│  │  • Total compute units: 1,250                                │   │
│  │  • Failed tasks: 1 (no penalty if < threshold)               │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                              │                                      │
│                              ▼                                      │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  REWARD DISTRIBUTION (every 256 epochs ≈ 27 hours)          │   │
│  │  • Provider reward: 12.5 MBO                                 │   │
│  │  • Network fee: 0.625 MBO (5%)                               │   │
│  │  • Net payout: 11.875 MBO                                    │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Health Monitoring Commands

```bash
# Check provider health
mbongo-compute provider health

# View current tasks
mbongo-compute provider tasks --status active

# View earnings summary
mbongo-compute provider earnings --period 24h

# GPU monitoring
nvidia-smi --query-gpu=utilization.gpu,utilization.memory,temperature.gpu,power.draw --format=csv -l 5

# System resource monitoring
htop
nvtop  # GPU-specific monitor
```

---

## 7. Security & Fraud Prevention

### Invalid Compute Slashing

> ⚠️ **CRITICAL**: Invalid compute proofs result in **1,000 MBO slashing per offense!**

```
┌─────────────────────────────────────────────────────────────────────┐
│                    SLASHING SCENARIOS                               │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ❌ INVALID PROOF (1,000 MBO Penalty)                               │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  • Incorrect computation result                              │   │
│  │  • Tampered output hash                                      │   │
│  │  • Forged hardware attestation                               │   │
│  │  • Mismatched execution fingerprint                          │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                                                                     │
│  ⚠️ TIMEOUT PENALTY (Variable)                                      │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  • First timeout: Warning                                    │   │
│  │  • Repeated timeouts: 10-100 MBO penalty                     │   │
│  │  • Chronic timeouts: Provider suspension                     │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                                                                     │
│  ❌ FRAUD DETECTION                                                 │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  • Submitting pre-computed results: PERMANENT BAN            │   │
│  │  • Hardware spoofing: PERMANENT BAN                          │   │
│  │  • Collusion with task submitters: FULL STAKE SLASH          │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Verification Pipeline

```
┌─────────────────────────────────────────────────────────────────────┐
│                    VERIFICATION PIPELINE                            │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌──────────┐                                                       │
│  │ Provider │                                                       │
│  │ Submits  │                                                       │
│  │ Receipt  │                                                       │
│  └────┬─────┘                                                       │
│       │                                                             │
│       ▼                                                             │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  STAGE 1: Format Validation                                  │   │
│  │  • Receipt structure valid?                                  │   │
│  │  • Signature valid?                                          │   │
│  │  • Task ID exists?                                           │   │
│  └─────────────────────────────────────────────────────────────┘   │
│       │                                                             │
│       ▼                                                             │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  STAGE 2: Cryptographic Verification                         │   │
│  │  • Output hash matches expected?                             │   │
│  │  • Execution proof valid?                                    │   │
│  │  • Hardware attestation authentic?                           │   │
│  └─────────────────────────────────────────────────────────────┘   │
│       │                                                             │
│       ▼                                                             │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  STAGE 3: Spot-Check Verification (Random)                   │   │
│  │  • 5% of tasks re-executed by verifier network               │   │
│  │  • Results compared against provider submission              │   │
│  │  • Discrepancy triggers investigation                        │   │
│  └─────────────────────────────────────────────────────────────┘   │
│       │                                                             │
│       ▼                                                             │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │  STAGE 4: Consensus Finalization                             │   │
│  │  • Multiple validators confirm receipt                       │   │
│  │  • Reward or penalty applied                                 │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### GPU Isolation (NVIDIA MIG)

For enterprise GPUs (A100, H100), use Multi-Instance GPU (MIG) for isolation:

```bash
# Enable MIG mode (requires reboot)
sudo nvidia-smi -mig 1

# Create GPU instances
sudo nvidia-smi mig -cgi 9,9,9,9,9,9,9 -C

# List instances
nvidia-smi mig -lgi
nvidia-smi mig -lci

# Configure provider for MIG
# In compute_provider.toml:
# [hardware.mig]
# enabled = true
# instances = ["MIG-GPU-xxxxx-0", "MIG-GPU-xxxxx-1"]
```

### Sandbox Environment Best Practices

```bash
# Use Docker for task isolation
docker run --gpus all \
    --memory=16g \
    --memory-swap=16g \
    --cpus=4 \
    --network=none \
    --read-only \
    --security-opt=no-new-privileges \
    --cap-drop=ALL \
    mbongo/compute-sandbox:latest \
    execute-task --task-id TASK_ID

# Alternative: Use NVIDIA Container Toolkit
docker run --runtime=nvidia \
    --gpus '"device=0"' \
    --isolation=process \
    mbongo/compute-sandbox:latest
```

### Security Checklist

- [ ] GPU drivers updated to latest stable version
- [ ] CUDA/ROCm version matches provider configuration
- [ ] Provider keys secured with strong encryption
- [ ] Firewall configured (only P2P and health ports open)
- [ ] Task execution sandboxed (Docker/VM)
- [ ] Regular security updates applied
- [ ] Monitoring and alerting configured
- [ ] Backup of provider keys maintained
- [ ] No unauthorized access to GPU hardware

---

## 8. Earnings & Reward Mechanics

### Compute Unit Pricing

| Workload Type | Compute Units/Task | Base Rate (MBO) | Typical Payout |
|---------------|-------------------|-----------------|----------------|
| **AI Inference (Small)** | 10-50 | 0.01 | 0.1-0.5 MBO |
| **AI Inference (Large)** | 100-500 | 0.01 | 1-5 MBO |
| **ML Training (Epoch)** | 1,000-10,000 | 0.01 | 10-100 MBO |
| **3D Render (Frame)** | 50-200 | 0.01 | 0.5-2 MBO |
| **Scientific Sim (Hour)** | 500-2,000 | 0.01 | 5-20 MBO |
| **ZK Proof Generation** | 200-1,000 | 0.01 | 2-10 MBO |

### Reward Epochs

```
┌─────────────────────────────────────────────────────────────────────┐
│                    REWARD EPOCH STRUCTURE                           │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │  EPOCH (6.4 minutes = 32 slots × 12 seconds)                  │  │
│  │                                                                │  │
│  │  Slot 0    Slot 8    Slot 16   Slot 24   Slot 31             │  │
│  │    │         │         │         │         │                  │  │
│  │    ▼         ▼         ▼         ▼         ▼                  │  │
│  │  ┌───┐     ┌───┐     ┌───┐     ┌───┐     ┌───┐              │  │
│  │  │T1 │     │T2 │     │T3 │     │T4 │     │T5 │              │  │
│  │  └───┘     └───┘     └───┘     └───┘     └───┘              │  │
│  │                                                                │  │
│  │  End of Epoch: Receipts aggregated                            │  │
│  └──────────────────────────────────────────────────────────────┘  │
│                              │                                      │
│                              ▼                                      │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │  REWARD EPOCH (256 epochs ≈ 27.3 hours)                       │  │
│  │                                                                │  │
│  │  At end of reward epoch:                                      │  │
│  │  • All valid receipts tallied                                 │  │
│  │  • Provider scores calculated                                 │  │
│  │  • MBO rewards distributed to providers                       │  │
│  │  • Performance metrics updated                                │  │
│  └──────────────────────────────────────────────────────────────┘  │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Sample Earning Scenarios

#### Scenario 1: Single GPU Provider (Normal Demand)

```
┌─────────────────────────────────────────────────────────────────────┐
│  SCENARIO: Single RTX 4090 - Normal Network Demand                  │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Hardware: 1× NVIDIA RTX 4090 (24GB VRAM)                          │
│  Uptime: 95%                                                        │
│  Task Types: AI Inference, 3D Rendering                            │
│                                                                     │
│  Daily Activity:                                                    │
│  ┌────────────────────────────────────────────────────────────┐    │
│  │ AI Inference Tasks:    150 tasks × 50 CU × 0.01 = 75 MBO   │    │
│  │ 3D Rendering Tasks:     30 tasks × 100 CU × 0.01 = 30 MBO  │    │
│  │ Quality Bonus (5%):                              = 5.25 MBO │    │
│  │ Network Fee (5%):                               = -5.51 MBO │    │
│  │ ──────────────────────────────────────────────────────────  │    │
│  │ NET DAILY EARNINGS:                             ≈ 105 MBO   │    │
│  └────────────────────────────────────────────────────────────┘    │
│                                                                     │
│  Monthly Projection: ~3,150 MBO                                    │
│  (Assumes consistent demand and uptime)                            │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

#### Scenario 2: Single GPU Provider (High Demand)

```
┌─────────────────────────────────────────────────────────────────────┐
│  SCENARIO: Single RTX 4090 - High Network Demand                    │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Hardware: 1× NVIDIA RTX 4090 (24GB VRAM)                          │
│  Uptime: 98%                                                        │
│  Task Types: AI Inference (LLM), ML Training                       │
│                                                                     │
│  Daily Activity:                                                    │
│  ┌────────────────────────────────────────────────────────────┐    │
│  │ LLM Inference Tasks:   200 tasks × 100 CU × 0.015 = 300 MBO│    │
│  │ ML Training Tasks:       5 tasks × 2000 CU × 0.015 = 150 MBO│   │
│  │ Priority Bonus (50%):                            = 225 MBO │    │
│  │ Quality Bonus (10%):                             = 67.5 MBO │    │
│  │ Network Fee (5%):                               = -37.1 MBO │    │
│  │ ──────────────────────────────────────────────────────────  │    │
│  │ NET DAILY EARNINGS:                             ≈ 705 MBO   │    │
│  └────────────────────────────────────────────────────────────┘    │
│                                                                     │
│  Monthly Projection: ~21,150 MBO                                   │
│  (High demand periods with priority pricing)                       │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

#### Scenario 3: GPU Cluster (Enterprise)

```
┌─────────────────────────────────────────────────────────────────────┐
│  SCENARIO: Enterprise GPU Cluster - 8× H100                         │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Hardware: 8× NVIDIA H100 80GB (NVLink interconnect)               │
│  Uptime: 99.5%                                                      │
│  Task Types: Large ML Training, LLM Fine-tuning                    │
│                                                                     │
│  Daily Activity:                                                    │
│  ┌────────────────────────────────────────────────────────────┐    │
│  │ Large Model Training:   2 jobs × 50000 CU × 0.02 = 2000 MBO│    │
│  │ LLM Fine-tuning:       10 jobs × 10000 CU × 0.02 = 2000 MBO│    │
│  │ Distributed Inference: 500 tasks × 200 CU × 0.02 = 2000 MBO│    │
│  │ Enterprise Tier Bonus (25%):                   = 1500 MBO  │    │
│  │ Quality Bonus (15%):                           = 1125 MBO  │    │
│  │ Network Fee (5%):                             = -431.25 MBO │    │
│  │ ──────────────────────────────────────────────────────────  │    │
│  │ NET DAILY EARNINGS:                           ≈ 8,194 MBO   │    │
│  └────────────────────────────────────────────────────────────┘    │
│                                                                     │
│  Monthly Projection: ~245,820 MBO                                  │
│  (Enterprise SLA with guaranteed task allocation)                  │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Earnings Dashboard

```bash
# View current earnings
mbongo-compute provider earnings --period 24h

# Expected output:
# ┌────────────────────────────────────────────────────────────────┐
# │ Earnings Summary (Last 24 Hours)                               │
# ├────────────────────────────────────────────────────────────────┤
# │ Tasks Completed:           187                                 │
# │ Tasks Failed:              3                                   │
# │ Success Rate:              98.4%                               │
# │                                                                │
# │ Compute Units Processed:   12,450                              │
# │ Gross Earnings:            124.50 MBO                          │
# │ Quality Bonus:             6.23 MBO                            │
# │ Network Fee:               -6.54 MBO                           │
# │ ────────────────────────────────────────────────────────────── │
# │ Net Earnings:              124.19 MBO                          │
# │                                                                │
# │ Pending Payout:            248.38 MBO (next epoch)             │
# │ Total Lifetime Earnings:   15,432.67 MBO                       │
# └────────────────────────────────────────────────────────────────┘

# View detailed breakdown
mbongo-compute provider earnings --detailed --period 7d

# Export earnings report
mbongo-compute provider earnings --export csv --output earnings_report.csv
```

---

## 9. Troubleshooting

### Problem 1: GPU Not Detected

**Symptoms:**
- `mbongo-compute hardware detect` shows no GPU
- nvidia-smi fails
- Tasks not being assigned

**Solutions:**

```bash
# Check if GPU is visible to system
lspci | grep -i nvidia

# Check driver status
nvidia-smi

# If nvidia-smi fails, reinstall driver
sudo apt purge nvidia-* 
sudo apt autoremove
sudo apt install nvidia-driver-535

# Reboot after driver installation
sudo reboot

# Verify after reboot
nvidia-smi

# Check CUDA visibility
echo $CUDA_VISIBLE_DEVICES

# Ensure GPU not blocked
cat /sys/bus/pci/devices/*/enable | head -1
```

---

### Problem 2: CUDA Missing

**Symptoms:**
- "CUDA not found" errors
- PyTorch/TensorFlow can't find CUDA
- Tasks fail with CUDA errors

**Solutions:**

```bash
# Check CUDA installation
nvcc --version

# If missing, install CUDA toolkit
wget https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2204/x86_64/cuda-keyring_1.1-1_all.deb
sudo dpkg -i cuda-keyring_1.1-1_all.deb
sudo apt update
sudo apt install cuda-toolkit-12-2

# Set environment variables
export PATH=/usr/local/cuda-12.2/bin:$PATH
export LD_LIBRARY_PATH=/usr/local/cuda-12.2/lib64:$LD_LIBRARY_PATH

# Add to .bashrc
echo 'export PATH=/usr/local/cuda-12.2/bin:$PATH' >> ~/.bashrc
echo 'export LD_LIBRARY_PATH=/usr/local/cuda-12.2/lib64:$LD_LIBRARY_PATH' >> ~/.bashrc

# Verify PyTorch CUDA
python3 -c "import torch; print(torch.cuda.is_available())"
```

---

### Problem 3: ROCm Mismatch

**Symptoms:**
- AMD GPU detected but tasks fail
- ROCm version conflicts
- hipcc compilation errors

**Solutions:**

```bash
# Check ROCm version
rocm-smi --showversion

# Check installed ROCm packages
apt list --installed | grep rocm

# Remove conflicting versions
sudo apt purge rocm-* hip-*
sudo apt autoremove

# Install specific ROCm version
sudo amdgpu-install --usecase=rocm --rocmrelease=6.0.0

# Verify installation
rocminfo
hipcc --version

# Test with PyTorch
python3 -c "import torch; print(torch.cuda.is_available())"  # Uses HIP on AMD
```

---

### Problem 4: Task Timeout

**Symptoms:**
- Tasks complete but timeout before receipt submission
- "Task exceeded time limit" errors
- Partial results lost

**Solutions:**

```bash
# Check current timeout settings
grep timeout ~/.mbongo/config/compute_provider.toml

# Increase timeout for large tasks
# In compute_provider.toml:
# default_timeout = 7200  # 2 hours

# Monitor task execution time
mbongo-compute provider tasks --show-duration

# Check for resource bottlenecks during execution
nvidia-smi dmon -s u  # GPU utilization monitoring

# Optimize task execution
# - Ensure GPU memory not fragmented
# - Close other GPU-using applications
# - Consider reducing max_parallel_tasks
```

---

### Problem 5: Invalid Proof

**Symptoms:**
- Receipt rejected with "invalid proof"
- Slashing warning received
- Verification failures

**Solutions:**

```bash
# Check task execution logs
grep "proof_generation" ~/.mbongo/logs/compute_provider.log

# Verify computation determinism
mbongo-compute debug verify-task --task-id TASK_ID --recompute

# Check for floating-point inconsistencies
# In compute_provider.toml:
# [execution]
# deterministic_mode = true
# fp_precision = "high"

# Clear task cache and retry
mbongo-compute provider cache-clear
mbongo-compute provider task-retry --task-id TASK_ID

# Report potential network issue
mbongo-compute report invalid-verification --task-id TASK_ID
```

---

### Problem 6: Receipt Submission Failed

**Symptoms:**
- "Receipt submission failed" errors
- Tasks complete but no reward
- Network connectivity issues

**Solutions:**

```bash
# Test RPC connectivity
curl -s http://127.0.0.1:8545/health

# Check submission endpoint
mbongo-compute provider test-connection

# View pending receipts
mbongo-compute provider receipts --status pending

# Manual receipt resubmission
mbongo-compute provider receipt-submit --task-id TASK_ID

# Check for rate limiting
grep "rate_limit" ~/.mbongo/logs/compute_provider.log

# Increase retry settings
# In compute_provider.toml:
# [receipt]
# retry_attempts = 5
# retry_delay = 10
```

---

### Problem 7: Network/RPC Issues

**Symptoms:**
- Cannot connect to validator RPC
- Frequent disconnections
- Task assignment failures

**Solutions:**

```bash
# Check validator node status
curl -s http://127.0.0.1:8545/health | jq .

# Test network connectivity
ping -c 4 127.0.0.1
nc -vz 127.0.0.1 8545

# Check firewall
sudo ufw status
sudo iptables -L -n | grep 8545

# If running own validator, check its logs
sudo journalctl -u mbongo-validator -f

# Try alternative RPC endpoints
# In compute_provider.toml:
# validator_rpc = "http://public-rpc.mbongo.network:8545"

# Check DNS resolution
nslookup public-rpc.mbongo.network
```

---

### Problem 8: Windows Driver Conflict

**Symptoms:**
- GPU not detected on Windows
- BSOD during compute tasks
- Driver crashes

**Solutions:**

```powershell
# Check current driver
nvidia-smi

# Use DDU (Display Driver Uninstaller) for clean removal
# Download from: https://www.guru3d.com/files-details/display-driver-uninstaller-download.html

# Boot into Safe Mode and run DDU
# Then install fresh NVIDIA driver

# Check for Windows driver conflicts
Get-WmiObject Win32_VideoController | Select-Object Name, DriverVersion, Status

# Disable Windows Update driver management
# Group Policy: Computer Configuration > Administrative Templates > 
# Windows Components > Windows Update > Do not include drivers with Windows Updates

# Install specific driver version
choco install nvidia-display-driver --version=535.104.05

# Verify installation
nvidia-smi
```

---

### Problem 9: Performance Issues

**Symptoms:**
- Tasks taking longer than expected
- Low GPU utilization
- Overheating warnings

**Solutions:**

```bash
# Monitor GPU metrics
nvidia-smi dmon -s pucvmet -d 1

# Check thermal throttling
nvidia-smi -q -d PERFORMANCE

# Optimize GPU settings
sudo nvidia-smi -pm 1  # Persistence mode
sudo nvidia-smi -pl 350  # Set power limit (adjust for your GPU)

# Check PCIe bandwidth
nvidia-smi -q | grep -A5 "PCI"

# Reduce parallel tasks if overloaded
# In compute_provider.toml:
# max_parallel_tasks = 2

# Clear GPU memory fragmentation
mbongo-compute provider gpu-reset

# Optimize CUDA settings
export CUDA_LAUNCH_BLOCKING=0
export CUDA_DEVICE_ORDER=PCI_BUS_ID
```

---

### Problem 10: Outdated Client

**Symptoms:**
- Protocol version mismatch
- Tasks rejected
- Feature unavailable errors

**Solutions:**

```bash
# Check current version
mbongo-compute --version

# Check for updates
mbongo-compute update --check

# Download latest version
LATEST=$(curl -s https://api.github.com/repos/mbongo-chain/mbongo-chain/releases/latest | jq -r '.tag_name')
echo "Latest version: $LATEST"

# Stop service
sudo systemctl stop mbongo-compute

# Download and install
wget -O mbongo-compute.tar.gz \
    "https://github.com/mbongo-chain/mbongo-chain/releases/download/${LATEST}/mbongo-compute-linux-amd64.tar.gz"
tar -xzf mbongo-compute.tar.gz
mv mbongo-compute ~/.mbongo/bin/

# Verify new version
mbongo-compute --version

# Restart service
sudo systemctl start mbongo-compute

# Check compatibility
mbongo-compute provider test-connection
```

---

## 10. Cross-References

### Related Documentation

| Document | Description | Link |
|----------|-------------|------|
| **Reward Mechanics** | Detailed reward calculation | [reward_mechanics.md](./reward_mechanics.md) |
| **Incentive Design** | Economic incentive structure | [incentive_design.md](./incentive_design.md) |
| **Oracle Model** | Data verification oracles | [oracle_model.md](./oracle_model.md) |
| **CLI Compute Provider** | CLI command reference | [cli_compute_provider.md](./cli_compute_provider.md) |
| **Architecture Overview** | System architecture | [architecture_master_overview.md](./architecture_master_overview.md) |
| **Compute Engine Overview** | PoUW compute engine details | [compute_engine_overview.md](./compute_engine_overview.md) |
| **Economic Security** | Security economics | [economic_security.md](./economic_security.md) |

### External Resources

- **Mbongo Chain GitHub**: `https://github.com/mbongo-chain/mbongo-chain`
- **Provider Dashboard**: `https://providers.mbongo.network`
- **Task Marketplace**: `https://tasks.mbongo.network`
- **Community Discord**: `https://discord.gg/mbongo`
- **Documentation Portal**: `https://docs.mbongo.network`

### Quick Command Reference

```bash
# Provider Management
mbongo-compute provider status           # Check provider status
mbongo-compute provider register         # Register as provider
mbongo-compute provider start            # Start provider node
mbongo-compute provider stop             # Stop provider node

# Hardware Management
mbongo-compute hardware detect           # Detect hardware
mbongo-compute hardware benchmark        # Run benchmarks
mbongo-compute hardware status           # Hardware health

# Task Management
mbongo-compute provider tasks            # List tasks
mbongo-compute provider task-info        # Task details
mbongo-compute provider task-retry       # Retry failed task

# Earnings & Receipts
mbongo-compute provider earnings         # View earnings
mbongo-compute provider receipts         # View receipts
mbongo-compute provider withdraw         # Withdraw earnings

# Diagnostics
mbongo-compute debug logs                # View logs
mbongo-compute debug metrics             # View metrics
mbongo-compute debug config              # Validate config
```

---

## Appendix: Quick Start Checklist

Use this checklist to ensure your compute provider is properly set up:

### Pre-Installation
- [ ] Hardware meets minimum requirements
- [ ] GPU compatible with CUDA/ROCm
- [ ] Stable internet connection
- [ ] Ubuntu 22.04 LTS or Windows 10/11

### Installation
- [ ] System packages installed
- [ ] GPU drivers installed and verified
- [ ] CUDA/ROCm toolkit installed
- [ ] Python 3.10+ with dependencies
- [ ] Mbongo Compute CLI installed

### Configuration
- [ ] Provider keys generated
- [ ] compute_provider.toml configured
- [ ] Hardware type and capabilities set
- [ ] Pricing and performance limits set
- [ ] RPC endpoint configured

### Registration
- [ ] Provider registered on network
- [ ] Initial stake deposited
- [ ] Connection to validator verified
- [ ] Hardware benchmark completed

### Security
- [ ] Firewall configured
- [ ] Provider keys backed up
- [ ] Task isolation (Docker) configured
- [ ] Monitoring enabled

### Operations
- [ ] Service configured (systemd/NSSM)
- [ ] Log rotation enabled
- [ ] Health monitoring active
- [ ] Alerts configured

### Verification
- [ ] Provider receiving tasks
- [ ] Receipts being submitted
- [ ] Rewards being earned
- [ ] No errors in logs

---

> **🎉 Congratulations!** Your Mbongo Chain Compute Provider is now ready to earn MBO tokens by contributing computational resources to the network.
>
> **Questions?** Join our Discord community for support.

---

*Document maintained by the Mbongo Chain Core Team*  
*Last reviewed: November 2025*

