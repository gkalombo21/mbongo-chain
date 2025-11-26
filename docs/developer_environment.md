# Mbongo Chain — Developer Environment Guide

Definitive setup guide for consistent, reproducible development across all contributors.

---

## 1. Purpose of This Document

### Why This Guide Exists

This document ensures that **every contributor** to Mbongo Chain has a consistent, reproducible development environment. Standardization prevents:

- "Works on my machine" bugs
- CI/CD failures due to local configuration differences
- Wasted time debugging environment issues
- Inconsistent code formatting and style

### Goals

| Goal | Description |
|------|-------------|
| **Consistency** | All developers use the same tools and versions |
| **Reproducibility** | Any developer can replicate another's environment |
| **Efficiency** | Minimize setup time for new contributors |
| **Quality** | Enforce code standards through tooling |

---

## 2. Recommended System Setup

### Operating System

| OS | Status | Notes |
|----|--------|-------|
| **Windows 10/11** | Primary | Full support, all commands tested |
| **Linux (Ubuntu 22.04+)** | Supported | Alternative for advanced users |
| **macOS (13+)** | Supported | Alternative for advanced users |

*This guide focuses on Windows. Linux/macOS users should adapt commands accordingly.*

### Hardware Requirements

| Component | Minimum | Recommended | Notes |
|-----------|---------|-------------|-------|
| **CPU** | 4 cores | 8+ cores | Faster builds with more cores |
| **RAM** | 8 GB | 16+ GB | Cargo builds are memory-intensive |
| **Storage** | 20 GB free | 50+ GB SSD | SSD significantly improves build times |
| **GPU** | None | NVIDIA RTX / AMD RX | Optional for PoUW experimentation |

### Required Global Tools

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     REQUIRED TOOLS                                          │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  CORE DEVELOPMENT                                                           │
│  ────────────────                                                           │
│  • Git 2.30+                                                               │
│  • Rust (stable) via rustup                                                │
│  • Cargo (bundled with Rust)                                               │
│  • Rustfmt (code formatter)                                                │
│  • Clippy (linter)                                                         │
│                                                                             │
│  SHELL & EDITOR                                                             │
│  ──────────────                                                             │
│  • PowerShell 7+ (Windows)                                                 │
│  • VS Code or Cursor (recommended)                                         │
│                                                                             │
│  OPTIONAL                                                                   │
│  ────────                                                                   │
│  • Docker (future CI/cluster tests)                                        │
│  • NVIDIA CUDA / AMD ROCm (PoUW experiments)                               │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 3. Required Tools

### Git

**Purpose:** Version control, collaboration, branching

**Version:** 2.30 or newer

**Download:** https://git-scm.com/download/win

**Verify:**
```powershell
git --version
# Expected: git version 2.XX.X.windows.X
```

---

### Rust Toolchain (via rustup)

**Purpose:** Compile Rust code, manage toolchains

**Channels:**
- `stable` (required) — Primary development
- `nightly` (optional) — Experimental features, advanced tooling

**Download:** https://rustup.rs

**Verify:**
```powershell
rustc --version
# Expected: rustc 1.XX.X (XXXXXXXX YYYY-MM-DD)

cargo --version
# Expected: cargo 1.XX.X (XXXXXXXX YYYY-MM-DD)

rustup --version
# Expected: rustup 1.XX.X (XXXXXXXX YYYY-MM-DD)
```

---

### Cargo

**Purpose:** Build system, package manager, dependency resolution

**Bundled with:** Rust toolchain

**Key Commands:**
- `cargo build` — Compile the project
- `cargo test` — Run tests
- `cargo run` — Execute binary
- `cargo check` — Fast syntax check

---

### Rustfmt

**Purpose:** Automatic code formatting for consistency

**Install:**
```powershell
rustup component add rustfmt
```

**Verify:**
```powershell
cargo fmt --version
# Expected: rustfmt X.X.X-stable (XXXXXXXX YYYY-MM-DD)
```

**Usage:**
```powershell
cargo fmt --all           # Format all code
cargo fmt --all -- --check # Check without modifying
```

---

### Clippy

**Purpose:** Lint Rust code for bugs, style issues, performance

**Install:**
```powershell
rustup component add clippy
```

**Verify:**
```powershell
cargo clippy --version
# Expected: clippy X.X.X (XXXXXXXX YYYY-MM-DD)
```

**Usage:**
```powershell
cargo clippy --workspace --all-targets -- -D warnings
```

---

### PowerShell 7+

**Purpose:** Modern shell with better scripting support

**Install:**
```powershell
winget install Microsoft.PowerShell
```

**Verify:**
```powershell
$PSVersionTable.PSVersion
# Expected: Major = 7, Minor = X
```

---

### VS Code or Cursor (Recommended)

**Purpose:** IDE with Rust support

| Editor | Download | Recommended Extensions |
|--------|----------|------------------------|
| **VS Code** | https://code.visualstudio.com | rust-analyzer, crates, Error Lens |
| **Cursor** | https://cursor.sh | Built-in AI assistance |

**rust-analyzer Settings (VS Code):**
```json
{
  "rust-analyzer.checkOnSave.command": "clippy",
  "rust-analyzer.cargo.allFeatures": true,
  "editor.formatOnSave": true
}
```

---

### Docker (Optional)

**Purpose:** Future CI/CD, local cluster testing, reproducible builds

**Download:** https://www.docker.com/products/docker-desktop

**Verify:**
```powershell
docker --version
# Expected: Docker version XX.X.X
```

*Note: Docker integration is planned for future releases.*

---

### GPU Drivers (Optional)

**Purpose:** PoUW computation experiments, GPU-accelerated workloads

| Vendor | Driver | Download |
|--------|--------|----------|
| NVIDIA | CUDA Toolkit 12+ | https://developer.nvidia.com/cuda-downloads |
| AMD | ROCm 5+ | https://rocm.docs.amd.com |

**Verify NVIDIA:**
```powershell
nvidia-smi
# Expected: Driver version and GPU info
```

*Note: GPU support is planned for future PoUW releases.*

---

## 4. Install Instructions (Windows Only)

### Step 1: Install Rust Toolchain

```powershell
# Download and run rustup installer
Invoke-WebRequest -Uri https://win.rustup.rs/x86_64 -OutFile rustup-init.exe
.\rustup-init.exe -y
Remove-Item rustup-init.exe

# Restart PowerShell
exit
```

After restarting PowerShell:

```powershell
# Verify installation
rustc --version
cargo --version
rustup --version
```

### Step 2: Add Required Components

```powershell
# Install code formatter
rustup component add rustfmt

# Install linter
rustup component add clippy

# Install Rust source (for IDE support)
rustup component add rust-src

# Verify components
rustup component list --installed
```

Expected output:
```
cargo-x86_64-pc-windows-msvc
clippy-x86_64-pc-windows-msvc
rust-docs-x86_64-pc-windows-msvc
rust-src
rust-std-x86_64-pc-windows-msvc
rustc-x86_64-pc-windows-msvc
rustfmt-x86_64-pc-windows-msvc
```

### Step 3: Install PowerShell 7

```powershell
# Using winget (Windows 10/11)
winget install Microsoft.PowerShell

# Verify
pwsh --version
# Expected: PowerShell 7.X.X
```

### Step 4: Validate All Installations

```powershell
# Create validation script
@"
Write-Host "=== Mbongo Chain Environment Validation ===" -ForegroundColor Cyan

# Rust
Write-Host "`n[Rust]" -ForegroundColor Yellow
rustc --version
cargo --version

# Components
Write-Host "`n[Components]" -ForegroundColor Yellow
cargo fmt --version
cargo clippy --version

# Git
Write-Host "`n[Git]" -ForegroundColor Yellow
git --version

# PowerShell
Write-Host "`n[PowerShell]" -ForegroundColor Yellow
Write-Host "PowerShell $($PSVersionTable.PSVersion)"

Write-Host "`n=== Validation Complete ===" -ForegroundColor Green
"@ | Out-File -FilePath validate-env.ps1

# Run validation
.\validate-env.ps1
```

---

## 5. Directory Structure Guide

```
mbongo-chain/
├── cli/           # Command-line interface
├── crypto/        # Cryptographic primitives
├── docs/          # Documentation
├── network/       # P2P networking
├── node/          # Node orchestration
├── pow/           # Proof of Useful Work
├── runtime/       # Execution engine
├── scripts/       # Build and CI scripts
├── spec/          # Protocol specifications
├── Cargo.toml     # Workspace manifest
├── Cargo.lock     # Dependency lockfile
├── README.md      # Project overview
└── CONTRIBUTING.md # Contribution guide
```

### Module Purposes

| Directory | Purpose | Type |
|-----------|---------|------|
| `/node` | Node lifecycle, orchestration, coordination | Library |
| `/runtime` | State machine, transaction execution, gas metering | Library |
| `/pow` | PoUW verification, compute task management | Library |
| `/network` | P2P networking, gossip, peer management | Library |
| `/crypto` | Hashing, signatures, Merkle trees | Library |
| `/cli` | User-facing commands, node control | Binary |
| `/docs` | Developer documentation, architecture guides | Docs |
| `/scripts` | Automation, CI helpers, setup scripts | Scripts |
| `/spec` | Protocol specifications, consensus rules | Specs |

### Dependency Flow

```
cli ──▶ node ──▶ runtime ──▶ crypto
              │
              ├──▶ network (standalone)
              │
              └──▶ pow ──▶ crypto
```

---

## 6. Git Workflow Requirements

### Branch Strategy

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     BRANCH WORKFLOW                                         │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  main                                                                       │
│  ────                                                                       │
│  • Production-ready code                                                   │
│  • Protected (no direct commits)                                           │
│  • Merges only via reviewed PRs                                            │
│                                                                             │
│  develop                                                                    │
│  ───────                                                                    │
│  • Active development                                                      │
│  • Integration branch                                                      │
│  • PRs target this branch                                                  │
│                                                                             │
│  feature/<name>                                                             │
│  ──────────────                                                             │
│  • New features                                                            │
│  • Example: feature/mempool-priority                                       │
│                                                                             │
│  fix/<name>                                                                 │
│  ──────────                                                                 │
│  • Bug fixes                                                               │
│  • Example: fix/header-validation                                          │
│                                                                             │
│  docs/<name>                                                                │
│  ───────────                                                                │
│  • Documentation updates                                                   │
│  • Example: docs/consensus-overview                                        │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Creating Feature Branches

```powershell
# Sync with develop
git checkout develop
git pull origin develop

# Create feature branch
git checkout -b feature/my-feature

# Work on your changes...
```

### Conventional Commits

Use the following prefixes:

| Prefix | Purpose | Example |
|--------|---------|---------|
| `feat:` | New feature | `feat: add mempool priority queue` |
| `fix:` | Bug fix | `fix: resolve block validation error` |
| `docs:` | Documentation | `docs: update consensus overview` |
| `refactor:` | Code restructuring | `refactor: simplify runtime interface` |
| `test:` | Test additions | `test: add crypto hash tests` |
| `chore:` | Maintenance | `chore: update dependencies` |

**Example:**
```powershell
git commit -m "feat: implement transaction signature verification"
```

### Pull Request Workflow

1. **Create branch** from `develop`
2. **Implement changes** following code standards
3. **Run pre-PR checks** (see below)
4. **Push branch** to remote
5. **Open PR** targeting `develop`
6. **Address review feedback**
7. **Merge** after approval

### Syncing with Main Branch

```powershell
# Fetch latest changes
git fetch origin

# Rebase your branch on develop
git checkout feature/my-feature
git rebase origin/develop

# Resolve conflicts if any, then continue
git rebase --continue

# Push (force if rebased)
git push --force-with-lease
```

### Required Checks Before PR

```powershell
# 1. Format code
cargo fmt --all

# 2. Check formatting
cargo fmt --all -- --check

# 3. Run Clippy
cargo clippy --workspace --all-targets -- -D warnings

# 4. Build workspace
cargo build --workspace

# 5. Run tests
cargo test --workspace
```

All checks must pass before opening a PR.

---

## 7. Common Commands

### Essential Commands Table

| Command | Purpose |
|---------|---------|
| `git clone https://github.com/gkalombo21/mbongo-chain.git` | Clone repository |
| `git checkout -b feature/my-feature` | Create new branch |
| `git status` | Check working directory status |
| `git add .` | Stage all changes |
| `git commit -m "feat: description"` | Commit with message |
| `git push origin feature/my-feature` | Push branch to remote |
| `git pull origin develop` | Sync with develop |
| `git rebase origin/develop` | Rebase on latest develop |

### Cargo Commands Table

| Command | Purpose |
|---------|---------|
| `cargo build --workspace` | Build all modules |
| `cargo build --workspace --release` | Build optimized release |
| `cargo check --workspace` | Fast syntax check |
| `cargo test --workspace` | Run all tests |
| `cargo fmt --all` | Format all code |
| `cargo fmt --all -- --check` | Check formatting |
| `cargo clippy --workspace --all-targets -- -D warnings` | Run linter |
| `cargo doc --workspace --open` | Generate and open docs |
| `cargo clean` | Remove build artifacts |
| `cargo update` | Update dependencies |

### Module-Specific Commands

| Command | Purpose |
|---------|---------|
| `cargo build -p crypto` | Build crypto module only |
| `cargo test -p runtime` | Test runtime module only |
| `cargo clippy -p node` | Lint node module only |
| `cargo run -p cli -- --help` | Run CLI with help |

---

## 8. Environment Validation Script (PowerShell)

### Placeholder Script

Save as `scripts/validate-environment.ps1`:

```powershell
#Requires -Version 7.0
<#
.SYNOPSIS
    Validates the Mbongo Chain development environment.
.DESCRIPTION
    Checks all required tools and configurations for development.
.NOTES
    Status: PLACEHOLDER - To be fully implemented
#>

param(
    [switch]$Verbose
)

$ErrorActionPreference = "Stop"
$Results = @()

function Test-Command {
    param([string]$Command, [string]$Name)
    
    try {
        $output = Invoke-Expression "$Command 2>&1"
        $Results += [PSCustomObject]@{
            Check = $Name
            Status = "PASS"
            Details = $output | Select-Object -First 1
        }
        return $true
    }
    catch {
        $Results += [PSCustomObject]@{
            Check = $Name
            Status = "FAIL"
            Details = $_.Exception.Message
        }
        return $false
    }
}

Write-Host "`n╔══════════════════════════════════════════════════════════════╗" -ForegroundColor Cyan
Write-Host "║       MBONGO CHAIN ENVIRONMENT VALIDATION                    ║" -ForegroundColor Cyan
Write-Host "╚══════════════════════════════════════════════════════════════╝`n" -ForegroundColor Cyan

# Check Rust Version
Write-Host "[1/6] Checking Rust..." -ForegroundColor Yellow
Test-Command "rustc --version" "Rust Compiler" | Out-Null
Test-Command "cargo --version" "Cargo" | Out-Null

# Check Git Version
Write-Host "[2/6] Checking Git..." -ForegroundColor Yellow
Test-Command "git --version" "Git" | Out-Null

# Check Components
Write-Host "[3/6] Checking Rust Components..." -ForegroundColor Yellow
Test-Command "cargo fmt --version" "Rustfmt" | Out-Null
Test-Command "cargo clippy --version" "Clippy" | Out-Null

# Check Workspace Build
Write-Host "[4/6] Checking Workspace Build..." -ForegroundColor Yellow
$buildResult = cargo check --workspace 2>&1
if ($LASTEXITCODE -eq 0) {
    $Results += [PSCustomObject]@{
        Check = "Workspace Build"
        Status = "PASS"
        Details = "All modules compile"
    }
} else {
    $Results += [PSCustomObject]@{
        Check = "Workspace Build"
        Status = "FAIL"
        Details = "Build errors detected"
    }
}

# Check Formatting
Write-Host "[5/6] Checking Code Formatting..." -ForegroundColor Yellow
$fmtResult = cargo fmt --all -- --check 2>&1
if ($LASTEXITCODE -eq 0) {
    $Results += [PSCustomObject]@{
        Check = "Code Formatting"
        Status = "PASS"
        Details = "All files formatted"
    }
} else {
    $Results += [PSCustomObject]@{
        Check = "Code Formatting"
        Status = "WARN"
        Details = "Some files need formatting"
    }
}

# Check Clippy
Write-Host "[6/6] Checking Clippy..." -ForegroundColor Yellow
$clippyResult = cargo clippy --workspace --all-targets -- -D warnings 2>&1
if ($LASTEXITCODE -eq 0) {
    $Results += [PSCustomObject]@{
        Check = "Clippy Lints"
        Status = "PASS"
        Details = "No warnings"
    }
} else {
    $Results += [PSCustomObject]@{
        Check = "Clippy Lints"
        Status = "WARN"
        Details = "Warnings detected"
    }
}

# Summary
Write-Host "`n╔══════════════════════════════════════════════════════════════╗" -ForegroundColor Cyan
Write-Host "║                        RESULTS                               ║" -ForegroundColor Cyan
Write-Host "╚══════════════════════════════════════════════════════════════╝`n" -ForegroundColor Cyan

$passCount = ($Results | Where-Object { $_.Status -eq "PASS" }).Count
$failCount = ($Results | Where-Object { $_.Status -eq "FAIL" }).Count
$warnCount = ($Results | Where-Object { $_.Status -eq "WARN" }).Count

foreach ($result in $Results) {
    $color = switch ($result.Status) {
        "PASS" { "Green" }
        "FAIL" { "Red" }
        "WARN" { "Yellow" }
    }
    Write-Host "[$($result.Status)] $($result.Check): $($result.Details)" -ForegroundColor $color
}

Write-Host "`n────────────────────────────────────────────────────────────────" -ForegroundColor Gray
Write-Host "PASS: $passCount | WARN: $warnCount | FAIL: $failCount" -ForegroundColor White

if ($failCount -gt 0) {
    Write-Host "`n❌ Environment validation FAILED" -ForegroundColor Red
    exit 1
} elseif ($warnCount -gt 0) {
    Write-Host "`n⚠️  Environment validation PASSED with warnings" -ForegroundColor Yellow
    exit 0
} else {
    Write-Host "`n✅ Environment validation PASSED" -ForegroundColor Green
    exit 0
}
```

### Usage

```powershell
# Run validation
.\scripts\validate-environment.ps1

# Run with verbose output
.\scripts\validate-environment.ps1 -Verbose
```

*Note: This script is a placeholder. Full implementation is planned for future releases.*

---

## 9. Troubleshooting

### Rust Toolchain Mismatch

**Error:**
```
error: package `some-crate` requires rustc 1.70.0 or newer
```

**Solution:**
```powershell
rustup update stable
rustup default stable
rustc --version
```

---

### PATH Issues

**Error:**
```
'cargo' is not recognized as an internal or external command
```

**Solution:**
```powershell
# Add Cargo to PATH (current session)
$env:PATH += ";$env:USERPROFILE\.cargo\bin"

# Verify
cargo --version

# Permanent fix: Add to system PATH
# System Properties → Environment Variables → User PATH
# Add: %USERPROFILE%\.cargo\bin
```

---

### Clippy Not Found

**Error:**
```
error: no such command: `clippy`
```

**Solution:**
```powershell
rustup component add clippy
cargo clippy --version
```

---

### Rustfmt Not Found

**Error:**
```
error: no such command: `fmt`
```

**Solution:**
```powershell
rustup component add rustfmt
cargo fmt --version
```

---

### Git Credential Issues (HTTPS vs SSH)

**Error (HTTPS):**
```
fatal: Authentication failed for 'https://github.com/...'
```

**Solution (Use SSH):**
```powershell
# Generate SSH key
ssh-keygen -t ed25519 -C "your_email@example.com"

# Start SSH agent
Get-Service ssh-agent | Set-Service -StartupType Automatic
Start-Service ssh-agent

# Add key
ssh-add $env:USERPROFILE\.ssh\id_ed25519

# Add public key to GitHub
Get-Content $env:USERPROFILE\.ssh\id_ed25519.pub | clip
# Paste in GitHub → Settings → SSH Keys

# Test connection
ssh -T git@github.com

# Update remote URL
git remote set-url origin git@github.com:gkalombo21/mbongo-chain.git
```

**Solution (HTTPS with Credential Manager):**
```powershell
git config --global credential.helper manager
git pull  # Enter credentials when prompted
```

---

### VS Code / Cursor Terminal Problems

**Issue:** Terminal not recognizing commands

**Solution:**
1. Open Settings (Ctrl+,)
2. Search for "terminal.integrated.defaultProfile.windows"
3. Set to "PowerShell"
4. Restart VS Code/Cursor

**rust-analyzer not working:**
1. Install rust-analyzer extension
2. Reload window (Ctrl+Shift+P → "Reload Window")
3. Check Output → rust-analyzer for errors

---

### Windows Permission Issues

**Error:**
```
error: Permission denied (os error 5)
```

**Solutions:**

1. **Run as Administrator:**
   - Right-click PowerShell → Run as Administrator

2. **Antivirus interference:**
   - Add exclusions for:
     - `%USERPROFILE%\.cargo`
     - `%USERPROFILE%\.rustup`
     - Your project directory

3. **Target directory locked:**
   ```powershell
   cargo clean
   cargo build --workspace
   ```

---

### Build Fails - Missing MSVC

**Error:**
```
error: linker `link.exe` not found
```

**Solution:**
1. Download Visual Studio Build Tools
2. Install "Desktop development with C++"
3. Restart PowerShell
4. Rebuild: `cargo build --workspace`

---

### Cargo.lock Conflicts

**Error:**
```
error: failed to select a version for `some-crate`
```

**Solution:**
```powershell
git checkout develop -- Cargo.lock
cargo update
cargo build --workspace
```

---

## 10. Next Steps for Developers

After setting up your environment, follow this path:

### Immediate Next Steps

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     DEVELOPER ONBOARDING PATH                               │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  1. VALIDATE ENVIRONMENT                                                    │
│     └─▶ Run: .\scripts\validate-environment.ps1                            │
│                                                                             │
│  2. READ GETTING STARTED                                                    │
│     └─▶ docs/getting_started.md                                            │
│                                                                             │
│  3. READ DEVELOPER INTRODUCTION                                             │
│     └─▶ docs/developer_introduction.md                                     │
│                                                                             │
│  4. STUDY ARCHITECTURE                                                      │
│     └─▶ docs/final_architecture_overview.md                                │
│                                                                             │
│  5. EXPLORE YOUR FOCUS AREA                                                 │
│     ├─▶ Execution: runtime/, docs/runtime_architecture.md                  │
│     ├─▶ Consensus: pow/, docs/consensus_validation.md                      │
│     ├─▶ Networking: network/, docs/networking_overview.md                  │
│     └─▶ Node: node/, docs/node_architecture.md                             │
│                                                                             │
│  6. START CONTRIBUTING                                                      │
│     └─▶ Pick an issue, create a branch, submit a PR                        │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Recommended Reading Order

| Order | Document | Purpose |
|-------|----------|---------|
| 1 | [getting_started.md](getting_started.md) | Quick setup validation |
| 2 | [developer_introduction.md](developer_introduction.md) | Comprehensive overview |
| 3 | [final_architecture_overview.md](final_architecture_overview.md) | System architecture |
| 4 | Module-specific docs | Deep dive into your area |

### Domain-Specific Focus Areas

| Interest | Start With | Then Read |
|----------|------------|-----------|
| **Execution** | `runtime/` source | runtime_architecture.md, state_machine_validation.md |
| **Consensus** | `pow/` source | consensus_validation.md, consensus_overview.md |
| **Networking** | `network/` source | networking_overview.md, sync_validation.md |
| **Node** | `node/` source | node_architecture.md, block_validation_pipeline.md |
| **CLI** | `cli/` source | getting_started.md |

### Quick Reference

| Action | Command |
|--------|---------|
| Validate environment | `.\scripts\validate-environment.ps1` |
| Build project | `cargo build --workspace` |
| Run tests | `cargo test --workspace` |
| Format code | `cargo fmt --all` |
| Lint code | `cargo clippy --workspace --all-targets -- -D warnings` |
| Open documentation | `cargo doc --workspace --open` |

---

**Welcome to Mbongo Chain!** Your environment is ready for development.

**Mbongo Chain** — Compute-first blockchain infrastructure for the global future.

