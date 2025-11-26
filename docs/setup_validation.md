# Mbongo Chain — Setup Validation

This document provides step-by-step validation procedures to ensure your development environment is correctly configured for Mbongo Chain development.

---

## 1. Environment Prerequisites

### Required Tools

| Tool | Minimum Version | Purpose | Installation |
|------|-----------------|---------|--------------|
| **Rust** | 1.70.0+ | Programming language | https://rustup.rs |
| **Cargo** | 1.70.0+ | Package manager | Bundled with Rust |
| **Git** | 2.30.0+ | Version control | https://git-scm.com |
| **Clippy** | Latest | Linter | `rustup component add clippy` |
| **Rustfmt** | Latest | Formatter | `rustup component add rustfmt` |

### Optional Tools

| Tool | Purpose | Installation |
|------|---------|--------------|
| **Make** | Build automation | System package manager |
| **PowerShell** | Windows scripting | Pre-installed on Windows |
| **Docker** | Containerization | https://docker.com |

### System Requirements

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     SYSTEM REQUIREMENTS                                     │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Minimum:                                                                   │
│  • CPU: 2 cores                                                            │
│  • RAM: 4 GB                                                               │
│  • Disk: 10 GB free                                                        │
│  • OS: Windows 10+, macOS 10.15+, Linux (glibc 2.17+)                     │
│                                                                             │
│  Recommended:                                                               │
│  • CPU: 4+ cores                                                           │
│  • RAM: 8+ GB                                                              │
│  • Disk: 20+ GB SSD                                                        │
│  • OS: Latest stable release                                               │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 2. Validate Rust Toolchain

### Check Rust Compiler

```shell
rustc --version
```

**Expected Output:**
```
rustc 1.XX.X (XXXXXXXX YYYY-MM-DD)
```

**Validation:**
- Version should be 1.70.0 or higher
- If missing: Install from https://rustup.rs

### Check Cargo

```shell
cargo --version
```

**Expected Output:**
```
cargo 1.XX.X (XXXXXXXX YYYY-MM-DD)
```

**Validation:**
- Version should match Rust version
- If missing: Reinstall Rust toolchain

### Check Toolchain Details

```shell
rustup show
```

**Expected Output:**
```
Default host: x86_64-unknown-linux-gnu (or your platform)

installed toolchains
--------------------
stable-x86_64-unknown-linux-gnu (default)

active toolchain
----------------
stable-x86_64-unknown-linux-gnu (default)
rustc 1.XX.X (XXXXXXXX YYYY-MM-DD)
```

**Validation:**
- Stable toolchain should be installed and default
- If not default: Run `rustup default stable`

### Check Components

```shell
rustup component list --installed
```

**Required Components:**
```
cargo-x86_64-unknown-linux-gnu
clippy-x86_64-unknown-linux-gnu
rust-docs-x86_64-unknown-linux-gnu
rust-std-x86_64-unknown-linux-gnu
rustc-x86_64-unknown-linux-gnu
rustfmt-x86_64-unknown-linux-gnu
```

**If Missing:**
```shell
rustup component add clippy rustfmt
```

---

## 3. Validate Workspace Build

### Run Cargo Check

```shell
cargo check --workspace
```

**Expected Output:**
```
    Checking crypto v0.1.0 (C:\Dev\mbongo-chain\crypto)
    Checking network v0.1.0 (C:\Dev\mbongo-chain\network)
    Checking runtime v0.1.0 (C:\Dev\mbongo-chain\runtime)
    Checking pow v0.1.0 (C:\Dev\mbongo-chain\pow)
    Checking node v0.1.0 (C:\Dev\mbongo-chain\node)
    Checking cli v0.1.0 (C:\Dev\mbongo-chain\cli)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in X.XXs
```

**Validation:**
- All modules should check successfully
- No errors should appear
- Warnings are acceptable but should be minimized

### Check Individual Modules

```shell
# Check each module independently
cargo check -p crypto
cargo check -p network
cargo check -p runtime
cargo check -p pow
cargo check -p node
cargo check -p cli
```

**Validation:**
- Each module should compile independently
- Dependency errors indicate missing workspace setup

---

## 4. Validate Clippy

### Run Clippy (Workspace)

```shell
cargo clippy --workspace --all-targets
```

**Expected Output:**
```
    Checking crypto v0.1.0
    Checking network v0.1.0
    ...
    Finished `dev` profile [unoptimized + debuginfo] target(s) in X.XXs
```

**Validation:**
- No errors should appear
- Warnings should be reviewed and addressed

### Run Clippy (Strict Mode)

```shell
cargo clippy --workspace --all-targets -- -D warnings
```

**Expected Output:**
```
    Finished `dev` profile [unoptimized + debuginfo] target(s) in X.XXs
```

**Validation:**
- Exit code should be 0
- Any warning will cause failure in strict mode
- CI uses strict mode — ensure local passes before pushing

### Run Clippy (No Dependencies)

```shell
cargo clippy --workspace --no-deps -- -D warnings
```

**Validation:**
- Only checks workspace code (faster)
- Use for quick validation during development

---

## 5. Validate Rustfmt

### Check Formatting

```shell
cargo fmt -- --check
```

**Expected Output (Formatted):**
```
(no output — all files formatted)
```

**Expected Output (Unformatted):**
```
Diff in /path/to/file.rs at line X:
 fn example() {
-    let x=1;
+    let x = 1;
 }
```

**Validation:**
- No output means all files are correctly formatted
- If diff shown: Run `cargo fmt` to fix

### Apply Formatting

```shell
cargo fmt --all
```

**Validation:**
- Applies formatting to all files
- Re-run `cargo fmt -- --check` to verify

### Check Specific File

```shell
cargo fmt -- --check path/to/file.rs
```

---

## 6. Validate Node Execution

### Check Node Help

```shell
cargo run -p node -- --help
```

**Expected Output:**
```
Mbongo Node vX.X.X

USAGE:
    node [OPTIONS] [SUBCOMMAND]

OPTIONS:
    -h, --help       Print help information
    -V, --version    Print version information
    ...
```

**Validation:**
- Node binary should compile and run
- Help output should display without errors

### Check CLI Help

```shell
cargo run -p cli -- --help
```

**Expected Output:**
```
Mbongo CLI vX.X.X

USAGE:
    cli [OPTIONS] [SUBCOMMAND]

OPTIONS:
    -h, --help       Print help information
    ...
```

### Check Version

```shell
cargo run -p cli -- version
```

**Validation:**
- Version should display correctly
- No runtime errors

---

## 7. Validate Documentation Build

### Build Documentation

```shell
cargo doc --workspace
```

**Expected Output:**
```
 Documenting crypto v0.1.0
 Documenting network v0.1.0
 Documenting runtime v0.1.0
 Documenting pow v0.1.0
 Documenting node v0.1.0
 Documenting cli v0.1.0
    Finished `dev` profile [unoptimized + debuginfo] target(s) in X.XXs
   Generated target/doc/node/index.html
```

**Validation:**
- All modules should generate documentation
- No documentation errors or warnings

### Build and Open Documentation

```shell
cargo doc --workspace --open
```

**Validation:**
- Documentation opens in default browser
- All module links work correctly

### Check for Documentation Warnings

```shell
RUSTDOCFLAGS="-D warnings" cargo doc --workspace --no-deps
```

**Validation:**
- Strict mode catches missing documentation
- All public items should be documented

---

## 8. Validate Git Configuration

### Check Remote

```shell
git remote -v
```

**Expected Output:**
```
origin  https://github.com/gkalombo21/mbongo-chain.git (fetch)
origin  https://github.com/gkalombo21/mbongo-chain.git (push)
```

**Validation:**
- Remote should point to correct repository
- Both fetch and push URLs should be configured

### Check Branch

```shell
git branch
```

**Expected Output:**
```
* main
  develop
  feature/my-feature
```

**Validation:**
- Current branch should be marked with `*`
- `main` or `develop` should exist

### Check Status

```shell
git status
```

**Expected Output (Clean):**
```
On branch main
Your branch is up to date with 'origin/main'.

nothing to commit, working tree clean
```

**Expected Output (Changes):**
```
On branch main
Your branch is up to date with 'origin/main'.

Changes not staged for commit:
  modified:   src/lib.rs

Untracked files:
  new_file.rs
```

**Validation:**
- Branch should be up to date with remote
- Uncommitted changes should be intentional

### Check Git Config

```shell
git config --list --local
```

**Validation:**
- `user.name` should be set
- `user.email` should be set

---

## 9. Common Errors and Fixes

### Build Errors

| Error | Cause | Fix |
|-------|-------|-----|
| `error: could not find Cargo.toml` | Wrong directory | `cd` to workspace root |
| `error[E0432]: unresolved import` | Missing dependency | Check `Cargo.toml` dependencies |
| `error[E0433]: failed to resolve` | Module not found | Check module path and `mod` declarations |
| `error: linker 'cc' not found` | Missing C compiler | Install build-essential (Linux) or Xcode (macOS) |
| `LINK : fatal error LNK1181` | Missing MSVC | Install Visual Studio Build Tools (Windows) |

### Toolchain Errors

| Error | Cause | Fix |
|-------|-------|-----|
| `rustc: command not found` | Rust not installed | Install from https://rustup.rs |
| `error: toolchain 'stable' is not installed` | Missing toolchain | `rustup install stable` |
| `error: component 'clippy' is not available` | Missing component | `rustup component add clippy` |
| `error: rustup could not choose a version` | No default toolchain | `rustup default stable` |

### Clippy Errors

| Error | Cause | Fix |
|-------|-------|-----|
| `error: could not compile due to previous error` | Code error | Fix the underlying code issue |
| `warning: unused variable` | Unused code | Remove or prefix with `_` |
| `warning: this function has too many arguments` | Design issue | Consider refactoring |

### Format Errors

| Error | Cause | Fix |
|-------|-------|-----|
| `Diff in file.rs` | Unformatted code | Run `cargo fmt` |
| `error: couldn't read file.rs` | File permissions | Check file permissions |
| `error: expected expression` | Syntax error | Fix syntax before formatting |

### Git Errors

| Error | Cause | Fix |
|-------|-------|-----|
| `fatal: not a git repository` | Not in repo | Clone the repository first |
| `error: failed to push some refs` | Out of sync | `git pull --rebase` then push |
| `error: Your local changes would be overwritten` | Uncommitted changes | Commit or stash changes |

---

## 10. Final Checklist

### Pre-Development Checklist

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     SETUP VALIDATION CHECKLIST                              │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Environment                                                                │
│  □ Rust installed (1.70.0+)                                                │
│  □ Cargo available                                                          │
│  □ Clippy installed                                                         │
│  □ Rustfmt installed                                                        │
│  □ Git installed and configured                                             │
│                                                                             │
│  Repository                                                                 │
│  □ Repository cloned                                                        │
│  □ Remote configured correctly                                              │
│  □ On correct branch                                                        │
│  □ Working tree clean (or changes intentional)                              │
│                                                                             │
│  Build                                                                      │
│  □ cargo check --workspace passes                                          │
│  □ All modules compile                                                      │
│  □ No unresolved dependencies                                               │
│                                                                             │
│  Quality                                                                    │
│  □ cargo clippy --workspace passes                                         │
│  □ cargo fmt -- --check passes                                             │
│  □ No warnings in strict mode                                               │
│                                                                             │
│  Execution                                                                  │
│  □ Node runs (--help)                                                      │
│  □ CLI runs (--help)                                                       │
│  □ Documentation builds                                                     │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Quick Validation Script

Run all validations in sequence:

**Linux/macOS:**
```shell
#!/bin/bash
set -e

echo "=== Rust Toolchain ==="
rustc --version
cargo --version

echo "=== Workspace Build ==="
cargo check --workspace

echo "=== Clippy ==="
cargo clippy --workspace --all-targets -- -D warnings

echo "=== Format ==="
cargo fmt -- --check

echo "=== Node ==="
cargo run -p cli -- --help > /dev/null

echo "=== Documentation ==="
cargo doc --workspace --no-deps

echo "=== Git ==="
git status

echo "=== ALL VALIDATIONS PASSED ==="
```

**Windows (PowerShell):**
```powershell
$ErrorActionPreference = "Stop"

Write-Host "=== Rust Toolchain ===" -ForegroundColor Cyan
rustc --version
cargo --version

Write-Host "=== Workspace Build ===" -ForegroundColor Cyan
cargo check --workspace

Write-Host "=== Clippy ===" -ForegroundColor Cyan
cargo clippy --workspace --all-targets -- -D warnings

Write-Host "=== Format ===" -ForegroundColor Cyan
cargo fmt -- --check

Write-Host "=== CLI ===" -ForegroundColor Cyan
cargo run -p cli -- --help | Out-Null

Write-Host "=== Documentation ===" -ForegroundColor Cyan
cargo doc --workspace --no-deps

Write-Host "=== Git ===" -ForegroundColor Cyan
git status

Write-Host "=== ALL VALIDATIONS PASSED ===" -ForegroundColor Green
```

### Before Committing

```shell
# Quick pre-commit check
cargo fmt --all
cargo clippy --workspace --all-targets -- -D warnings
cargo check --workspace
git status
```

### Before Pushing

```shell
# Full validation
cargo fmt --all
cargo clippy --workspace --all-targets -- -D warnings
cargo check --workspace
cargo test --workspace
cargo doc --workspace --no-deps
```

---

## Summary

A properly validated development environment ensures smooth collaboration and reduces CI failures. Run through this checklist when setting up a new machine or troubleshooting build issues.

For development workflow, see [Developer Guide](developer_guide.md).

For contribution guidelines, see [CONTRIBUTING.md](../CONTRIBUTING.md).

---

**Mbongo Chain** — Compute-first blockchain infrastructure for the global future.

