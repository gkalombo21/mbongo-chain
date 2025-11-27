# Mbongo Chain — CLI Wallet Commands

> **Document Type:** CLI Reference  
> **Last Updated:** November 2025  
> **Status:** Official Reference  
> **Parent:** [cli_overview.md](./cli_overview.md)

---

## Table of Contents

1. [Purpose of Wallet Commands](#1-purpose-of-wallet-commands)
2. [Wallet Command Structure](#2-wallet-command-structure)
3. [Detailed Command Documentation](#3-detailed-command-documentation)
4. [Transfer Logic](#4-transfer-logic)
5. [Security Rules](#5-security-rules)
6. [Wallet Lifecycle Diagrams](#6-wallet-lifecycle-diagrams)
7. [Cross-Links](#7-cross-links)

---

## 1. Purpose of Wallet Commands

### 1.1 What Wallets Are Used For

The `mbongo wallet` commands manage cryptographic keys and facilitate all economic interactions on Mbongo Chain.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         WALLET CAPABILITIES                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   MBO TRANSFERS                         GAS FEES                            │
│   ═════════════                         ════════                            │
│   • Send MBO to any address             • Pay transaction fees              │
│   • Receive MBO                         • Set priority fees                 │
│   • Batch transfers                     • Estimate gas costs                │
│                                                                             │
│   STAKING OPERATIONS                    PoUW COMPUTE TASKS                  │
│   ══════════════════                    ══════════════════                  │
│   • Deposit validator stake             • Pay for compute jobs              │
│   • Delegate to validators              • Receive compute rewards           │
│   • Withdraw rewards                    • Sign compute receipts             │
│   • Manage unbonding                                                        │
│                                                                             │
│   SIGNING & AUTHENTICATION              GOVERNANCE                          │
│   ════════════════════════              ══════════                          │
│   • Sign transactions                   • Vote on proposals                 │
│   • Sign arbitrary messages             • Delegate voting power             │
│   • Verify signatures                   • Create proposals                  │
│   • Multi-sig coordination                                                  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 1.2 Key Types

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         KEY HIERARCHY                                       │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   SPENDING KEY (Primary)                                                    │
│   ══════════════════════                                                    │
│   • Full control over funds                                                │
│   • Required for transfers, staking, governance                            │
│   • Derived from mnemonic (BIP-39/BIP-44)                                  │
│   • NEVER expose this key                                                  │
│                                                                             │
│   Derivation: m/44'/60'/0'/0/0 (Ethereum-compatible)                       │
│   Algorithm: secp256k1 ECDSA                                               │
│                                                                             │
│   ─────────────────────────────────────────────────────────────────────────│
│                                                                             │
│   VIEW KEY (Optional)                                                       │
│   ═══════════════════                                                       │
│   • Read-only access to balance and history                                │
│   • Cannot sign transactions                                               │
│   • Safe to share with auditors/monitoring                                 │
│   • Derived from spending key                                              │
│                                                                             │
│   Use case: Portfolio tracking, tax reporting, monitoring                  │
│                                                                             │
│   ─────────────────────────────────────────────────────────────────────────│
│                                                                             │
│   SESSION KEY (Optional)                                                    │
│   ══════════════════════                                                    │
│   • Limited-scope temporary key                                            │
│   • Time-bounded or action-bounded                                         │
│   • For DApp interactions                                                  │
│   • Revocable by spending key                                              │
│                                                                             │
│   Use case: Gaming, automated trading, DApp sessions                       │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 1.3 Security Warnings

```
╔═════════════════════════════════════════════════════════════════════════════╗
║                                                                             ║
║   ⚠️  CRITICAL SECURITY WARNINGS                                            ║
║                                                                             ║
║   ✗ NEVER share your private key or mnemonic phrase                        ║
║   ✗ NEVER take screenshots of recovery phrases                             ║
║   ✗ NEVER store keys in plain text files                                   ║
║   ✗ NEVER enter keys on websites or untrusted software                     ║
║   ✗ NEVER use the same wallet on mainnet and testnet                       ║
║                                                                             ║
║   ✓ Store mnemonic offline (paper, metal backup)                           ║
║   ✓ Use hardware wallets for large holdings                                ║
║   ✓ Verify addresses character-by-character before sending                 ║
║   ✓ Test with small amounts first                                          ║
║   ✓ Keep keystore files encrypted                                          ║
║                                                                             ║
╚═════════════════════════════════════════════════════════════════════════════╝
```

---

## 2. Wallet Command Structure

### 2.1 Syntax

```
mbongo wallet <command> [subcommand] [flags]
```

### 2.2 Subcommands

| Command | Description | Risk Level |
|---------|-------------|------------|
| `create` | Create new wallet | 🟡 Medium |
| `restore` | Restore from mnemonic | 🔴 High |
| `import` | Import from keystore | 🔴 High |
| `export` | Export keystore | 🔴 High |
| `address` | Show wallet address | 🟢 Low |
| `balance` | Check MBO balance | 🟢 Low |
| `transfer` | Send MBO | 🔴 High |
| `history` | Transaction history | 🟢 Low |
| `sign` | Sign message/transaction | 🟡 Medium |
| `verify` | Verify signature | 🟢 Low |
| `watch` | Add watch-only address | 🟢 Low |
| `keys` | Key management | 🟡 Medium |
| `mnemonic` | Display recovery phrase | 🔴 Critical |
| `delete` | Delete wallet | 🔴 Critical |

---

## 3. Detailed Command Documentation

### 3.1 `mbongo wallet create`

**Description:** Create a new wallet with a fresh keypair.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--name` | `-n` | No | `default` | Wallet name |
| `--password-file` | | No | (prompt) | Password file path |
| `--output-dir` | `-o` | No | `~/.mbongo/wallets` | Output directory |
| `--words` | | No | `24` | Mnemonic words (12, 15, 18, 21, 24) |

**Examples:**

```bash
# Interactive creation
mbongo wallet create

# Named wallet with password file
mbongo wallet create --name validator-1 --password-file ~/.secrets/wallet.pass

# 12-word mnemonic (less secure, not recommended)
mbongo wallet create --words 12
```

**Output:**

```
Creating new wallet...

⚠️  IMPORTANT: Write down your recovery phrase and store it safely.
    Anyone with this phrase can access your funds.

Recovery Phrase (24 words):
────────────────────────────────────────────────────────────────────
abandon ability able about above absent absorb abstract absurd abuse
access accident account accuse achieve acid acoustic acquire across act
action actor actress actual adapt
────────────────────────────────────────────────────────────────────

Wallet Created Successfully
────────────────────────────────────────────────────────────────────
  Name:     validator-1
  Address:  0x742d35Cc6634C0532925a3b844Bc9e7595f8bDe7
  Path:     ~/.mbongo/wallets/validator-1.json
────────────────────────────────────────────────────────────────────

Press ENTER to confirm you have saved your recovery phrase...
```

**JSON Output (`--output json`):**

```json
{
  "name": "validator-1",
  "address": "0x742d35Cc6634C0532925a3b844Bc9e7595f8bDe7",
  "path": "/home/user/.mbongo/wallets/validator-1.json",
  "created_at": "2025-11-27T10:30:00Z"
}
```

**Error Cases:**

| Error | Exit Code | Cause |
|-------|-----------|-------|
| `WalletExists` | 9 | Wallet with same name exists |
| `InvalidPassword` | 4 | Password too short (<8 chars) |
| `IOError` | 5 | Cannot write to output directory |

---

### 3.2 `mbongo wallet restore`

**Description:** Restore wallet from mnemonic phrase.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--name` | `-n` | No | `restored` | Wallet name |
| `--mnemonic-file` | | No | (prompt) | File containing mnemonic |
| `--password-file` | | No | (prompt) | Password file |
| `--derivation-path` | | No | `m/44'/60'/0'/0/0` | HD derivation path |

**Examples:**

```bash
# Interactive restore
mbongo wallet restore

# From file (more secure)
mbongo wallet restore --name my-wallet --mnemonic-file ~/.secrets/mnemonic.txt

# Custom derivation path
mbongo wallet restore --derivation-path "m/44'/60'/0'/0/1"
```

**Output:**

```
Enter your recovery phrase (24 words):
> [hidden input]

Verifying mnemonic...
Deriving keys...

Wallet Restored Successfully
────────────────────────────────────────────────────────────────────
  Name:     my-wallet
  Address:  0x742d35Cc6634C0532925a3b844Bc9e7595f8bDe7
  Path:     ~/.mbongo/wallets/my-wallet.json
────────────────────────────────────────────────────────────────────
```

**Error Cases:**

| Error | Exit Code | Cause |
|-------|-----------|-------|
| `InvalidMnemonic` | 1 | Invalid words or checksum |
| `WalletExists` | 9 | Wallet name already exists |

---

### 3.3 `mbongo wallet import`

**Description:** Import wallet from encrypted keystore file.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--file` | `-f` | Yes | — | Keystore file path |
| `--name` | `-n` | No | (from file) | Wallet name |
| `--password-file` | | No | (prompt) | Password for keystore |

**Examples:**

```bash
mbongo wallet import --file ./keystore.json
mbongo wallet import -f ./keystore.json --name imported-wallet
```

---

### 3.4 `mbongo wallet export`

**Description:** Export wallet to encrypted keystore file.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--name` | `-n` | No | `default` | Wallet to export |
| `--output` | `-o` | Yes | — | Output file path |
| `--password-file` | | No | (prompt) | Export password |

**Examples:**

```bash
mbongo wallet export --name validator-1 --output ./backup.json
```

**Output:**

```
⚠️  WARNING: Keystore Export
────────────────────────────────────────────────────────────────────
You are exporting an encrypted keystore file.
Anyone with this file AND the password can access your funds.

Wallet:  validator-1
Address: 0x742d35Cc...

Type 'EXPORT' to confirm: EXPORT

Enter export password: [hidden]
Confirm password: [hidden]

Keystore exported to: ./backup.json
```

---

### 3.5 `mbongo wallet address`

**Description:** Display wallet address(es).

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--name` | `-n` | No | `default` | Wallet name |
| `--all` | | No | `false` | Show all wallets |
| `--qr` | | No | `false` | Display QR code |

**Examples:**

```bash
mbongo wallet address
mbongo wallet address --name validator-1
mbongo wallet address --all
mbongo wallet address --qr
```

**Output:**

```
0x742d35Cc6634C0532925a3b844Bc9e7595f8bDe7
```

**Output (`--all`):**

```
Wallets
────────────────────────────────────────────────────────────────────
  default       │ 0x742d35Cc6634C0532925a3b844Bc9e7595f8bDe7
  validator-1   │ 0x8Ba1f109551bD432803012645Ac136ddd64DBA72
  watch-only    │ 0x1234567890abcdef1234567890abcdef12345678 (view)
────────────────────────────────────────────────────────────────────
```

---

### 3.6 `mbongo wallet balance`

**Description:** Check MBO balance.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--name` | `-n` | No | `default` | Wallet name |
| `--address` | `-a` | No | — | Check any address |
| `--output` | `-o` | No | `table` | Output format |
| `--rpc-url` | | No | (config) | RPC endpoint |

**Examples:**

```bash
mbongo wallet balance
mbongo wallet balance --name validator-1
mbongo wallet balance --address 0x1234...
mbongo wallet balance --output json
```

**Output (table):**

```
Balance for 0x742d35Cc6634C0532925a3b844Bc9e7595f8bDe7
────────────────────────────────────────────────────────────────────
  Available     │ 1,234.56789 MBO
  Staked        │ 50,000.00000 MBO
  Delegated     │ 5,000.00000 MBO
  Unbonding     │ 0.00000 MBO
  Pending       │ 12.34567 MBO (rewards)
  ────────────────────────────────────────────────────────────────
  Total         │ 56,246.91356 MBO
────────────────────────────────────────────────────────────────────
```

**Output (JSON):**

```json
{
  "address": "0x742d35Cc6634C0532925a3b844Bc9e7595f8bDe7",
  "available": "1234567890000000000000",
  "available_formatted": "1234.56789 MBO",
  "staked": "50000000000000000000000",
  "delegated": "5000000000000000000000",
  "unbonding": "0",
  "pending_rewards": "12345670000000000000",
  "total": "56246913560000000000000"
}
```

**Error Cases:**

| Error | Exit Code | Cause |
|-------|-----------|-------|
| `ConnectionError` | 3 | Cannot connect to RPC |
| `InvalidAddress` | 1 | Malformed address |

---

### 3.7 `mbongo wallet transfer`

**Description:** Send MBO to another address.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--from` | `-f` | No | `default` | Source wallet |
| `--to` | `-t` | Yes | — | Recipient address |
| `--amount` | `-a` | Yes | — | Amount in MBO |
| `--gas-price` | | No | `auto` | Gas price (gwei) |
| `--priority-fee` | | No | `1` | Priority fee (gwei) |
| `--nonce` | | No | `auto` | Transaction nonce |
| `--dry-run` | `-n` | No | `false` | Simulate only |
| `--yes` | `-y` | No | `false` | Skip confirmation |

**Examples:**

```bash
# Basic transfer
mbongo wallet transfer --to 0x5678... --amount 100

# With custom gas
mbongo wallet transfer --to 0x5678... --amount 100 --gas-price 20 --priority-fee 2

# Dry run (simulation)
mbongo wallet transfer --to 0x5678... --amount 100 --dry-run

# Non-interactive
mbongo wallet transfer --to 0x5678... --amount 100 --yes
```

**Output:**

```
Transfer Confirmation
────────────────────────────────────────────────────────────────────
  From:          0x742d35Cc6634C0532925a3b844Bc9e7595f8bDe7
  To:            0x5678901234abcdef5678901234abcdef56789012
  Amount:        100.00000000 MBO
  Gas Limit:     21000
  Gas Price:     15 gwei
  Priority Fee:  1 gwei
  Max Fee:       0.000336 MBO
  ────────────────────────────────────────────────────────────────
  Total:         100.000336 MBO
────────────────────────────────────────────────────────────────────

Confirm transfer? [y/N]: y

Enter wallet password: [hidden]

Signing transaction...
Broadcasting...

Transaction Submitted
────────────────────────────────────────────────────────────────────
  Tx Hash:   0xabc123...def456
  Status:    Pending
  Nonce:     42
  Block:     (awaiting confirmation)
────────────────────────────────────────────────────────────────────

Waiting for confirmation...
✓ Confirmed in block #12,345,678 (2 confirmations)
```

**JSON Output:**

```json
{
  "tx_hash": "0xabc123def456...",
  "from": "0x742d35Cc...",
  "to": "0x5678901234...",
  "amount": "100000000000000000000",
  "gas_used": 21000,
  "gas_price": "15000000000",
  "nonce": 42,
  "block_number": 12345678,
  "status": "confirmed"
}
```

**Error Cases:**

| Error | Exit Code | Cause |
|-------|-----------|-------|
| `InsufficientBalance` | 8 | Not enough MBO |
| `InvalidRecipient` | 1 | Malformed address |
| `NonceTooLow` | 5 | Nonce already used |
| `GasTooLow` | 5 | Gas price below minimum |

---

### 3.8 `mbongo wallet history`

**Description:** View transaction history.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--name` | `-n` | No | `default` | Wallet name |
| `--limit` | `-l` | No | `25` | Max transactions |
| `--type` | | No | `all` | Filter (send, receive, stake) |
| `--output` | `-o` | No | `table` | Output format |

**Examples:**

```bash
mbongo wallet history
mbongo wallet history --limit 100 --type send
mbongo wallet history --output json
```

---

### 3.9 `mbongo wallet sign`

**Description:** Sign a message or transaction.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--name` | `-n` | No | `default` | Wallet name |
| `--message` | `-m` | No | — | Message to sign |
| `--file` | `-f` | No | — | File to sign |
| `--hex` | | No | `false` | Output hex signature |

**Examples:**

```bash
# Sign message
mbongo wallet sign --message "Hello, Mbongo!"

# Sign file
mbongo wallet sign --file ./document.txt

# Hex output
mbongo wallet sign --message "test" --hex
```

**Output:**

```
Signature
────────────────────────────────────────────────────────────────────
  Message:   Hello, Mbongo!
  Signer:    0x742d35Cc6634C0532925a3b844Bc9e7595f8bDe7
  Signature: 0x1234567890abcdef...
────────────────────────────────────────────────────────────────────
```

---

### 3.10 `mbongo wallet verify`

**Description:** Verify a signature.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--message` | `-m` | Yes | — | Original message |
| `--signature` | `-s` | Yes | — | Signature to verify |
| `--address` | `-a` | Yes | — | Expected signer |

**Examples:**

```bash
mbongo wallet verify \
  --message "Hello, Mbongo!" \
  --signature 0x1234... \
  --address 0x742d35Cc...
```

**Output:**

```
✓ Signature is VALID
  Signer: 0x742d35Cc6634C0532925a3b844Bc9e7595f8bDe7
```

---

### 3.11 `mbongo wallet watch`

**Description:** Add a watch-only address (no private key).

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--address` | `-a` | Yes | — | Address to watch |
| `--name` | `-n` | No | `watch-X` | Label |

**Examples:**

```bash
mbongo wallet watch --address 0x1234... --name treasury
```

---

### 3.12 `mbongo wallet keys`

**Description:** Key management operations.

**Subcommands:**

| Subcommand | Description |
|------------|-------------|
| `list` | List all keys |
| `rotate` | Generate new session key |
| `revoke` | Revoke session key |

**Examples:**

```bash
mbongo wallet keys list
mbongo wallet keys rotate --name validator-1
mbongo wallet keys revoke --key-id session-123
```

---

### 3.13 `mbongo wallet mnemonic`

**Description:** Display recovery phrase (DANGEROUS).

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--name` | `-n` | No | `default` | Wallet name |

**Examples:**

```bash
mbongo wallet mnemonic --name validator-1
```

**Output:**

```
⚠️  CRITICAL SECURITY WARNING
────────────────────────────────────────────────────────────────────
You are about to display your recovery phrase.
Anyone who sees this can STEAL ALL YOUR FUNDS.

• Ensure no one is watching your screen
• Do not take screenshots
• Do not copy to clipboard

Type 'SHOW MNEMONIC' to continue: SHOW MNEMONIC
Enter wallet password: [hidden]

Recovery Phrase for 'validator-1':
────────────────────────────────────────────────────────────────────
abandon ability able about above absent absorb abstract absurd abuse
access accident account accuse achieve acid acoustic acquire across act
action actor actress actual adapt
────────────────────────────────────────────────────────────────────

This message will clear in 30 seconds...
```

---

### 3.14 `mbongo wallet delete`

**Description:** Permanently delete a wallet.

**Flags:**

| Flag | Short | Required | Default | Description |
|------|-------|----------|---------|-------------|
| `--name` | `-n` | Yes | — | Wallet to delete |
| `--force` | `-f` | No | `false` | Skip balance check |

**Examples:**

```bash
mbongo wallet delete --name old-wallet
```

**Output:**

```
⚠️  PERMANENT DELETION WARNING
────────────────────────────────────────────────────────────────────
You are about to permanently delete this wallet.
This action CANNOT be undone.

Wallet:  old-wallet
Address: 0x742d35Cc6634C0532925a3b844Bc9e7595f8bDe7
Balance: 0.00000 MBO

Ensure you have:
✓ Backed up your recovery phrase
✓ Transferred all funds
✓ No pending transactions

Type the wallet name to confirm deletion: old-wallet
Enter wallet password: [hidden]

Wallet 'old-wallet' deleted.
```

---

## 4. Transfer Logic

### 4.1 MBO as Gas Token

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         GAS PAYMENT MODEL                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   MBO is the ONLY token accepted for gas fees on Mbongo Chain.             │
│                                                                             │
│   TRANSACTION COST                                                          │
│   ════════════════                                                          │
│                                                                             │
│   total_fee = gas_used × (base_fee + priority_fee)                         │
│                                                                             │
│   Where:                                                                    │
│   • gas_used: Actual computation units consumed                            │
│   • base_fee: Protocol-determined (burned)                                 │
│   • priority_fee: User-specified (to validator)                            │
│                                                                             │
│   EXAMPLE                                                                   │
│   ═══════                                                                   │
│   Simple transfer: 21,000 gas                                              │
│   Base fee: 10 gwei                                                        │
│   Priority fee: 1 gwei                                                     │
│   Total: 21,000 × 11 gwei = 231,000 gwei = 0.000231 MBO                    │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 4.2 Fee Computation

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         FEE BREAKDOWN                                       │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐  │
│   │                                                                     │  │
│   │   BASE FEE                                                          │  │
│   │   ════════                                                          │  │
│   │   • Algorithmically determined by network                          │  │
│   │   • Adjusts based on block utilization                             │  │
│   │   • 100% BURNED (removed from circulation)                         │  │
│   │   • Creates deflationary pressure                                  │  │
│   │                                                                     │  │
│   │   Query current base fee:                                          │  │
│   │   $ mbongo tools gas-price                                         │  │
│   │                                                                     │  │
│   └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐  │
│   │                                                                     │  │
│   │   PRIORITY FEE (Tip)                                                │  │
│   │   ══════════════════                                                │  │
│   │   • User-specified incentive for faster inclusion                  │  │
│   │   • Paid to block proposer (validator)                             │  │
│   │   • Higher tip = higher priority                                   │  │
│   │   • Minimum: 1 gwei (recommended)                                  │  │
│   │                                                                     │  │
│   │   Set priority fee:                                                │  │
│   │   $ mbongo wallet transfer ... --priority-fee 2                    │  │
│   │                                                                     │  │
│   └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 4.3 Signing Process

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         TRANSACTION SIGNING                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   SIGNING PIPELINE                                                          │
│   ════════════════                                                          │
│                                                                             │
│   1. BUILD TRANSACTION                                                      │
│      • Set recipient, amount, gas parameters                               │
│      • Query current nonce from network                                    │
│      • Estimate gas if not specified                                       │
│                                                                             │
│   2. SERIALIZE (RLP Encoding)                                               │
│      • Encode all fields to canonical format                               │
│      • Include chain ID for replay protection                              │
│                                                                             │
│   3. HASH                                                                   │
│      • Keccak256 hash of serialized transaction                            │
│      • Produces 32-byte digest                                             │
│                                                                             │
│   4. SIGN                                                                   │
│      • ECDSA signature with private key                                    │
│      • Produces (v, r, s) signature components                             │
│                                                                             │
│   5. BROADCAST                                                              │
│      • Submit signed transaction to network                                │
│      • Await inclusion in block                                            │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 4.4 Nonce Management

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         NONCE SYSTEM                                        │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   WHAT IS A NONCE?                                                          │
│   ════════════════                                                          │
│   • Sequential counter for each account                                    │
│   • Starts at 0 for new accounts                                           │
│   • Increments by 1 for each confirmed transaction                         │
│   • Ensures transaction ordering                                           │
│                                                                             │
│   AUTOMATIC NONCE                                                           │
│   ════════════════                                                          │
│   By default, CLI queries the current nonce from the network:              │
│   $ mbongo wallet transfer --to 0x... --amount 100                         │
│   (nonce automatically determined)                                         │
│                                                                             │
│   MANUAL NONCE                                                              │
│   ════════════                                                              │
│   For advanced use (batch transactions, replacement):                      │
│   $ mbongo wallet transfer --to 0x... --amount 100 --nonce 42              │
│                                                                             │
│   COMMON ISSUES                                                             │
│   ═════════════                                                             │
│   • "Nonce too low": Transaction with this nonce already confirmed        │
│   • "Nonce gap": Missing transaction with lower nonce                     │
│   • Solution: Use --nonce to specify correct value                        │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 4.5 Replay Protection

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         REPLAY PROTECTION                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   CHAIN ID                                                                  │
│   ════════                                                                  │
│   • Unique identifier for each network                                     │
│   • Included in transaction signature                                      │
│   • Prevents cross-chain replay attacks                                    │
│                                                                             │
│   Mbongo Chain IDs:                                                        │
│   • Mainnet: 1                                                             │
│   • Testnet: 5                                                             │
│   • Devnet: 1337                                                           │
│                                                                             │
│   PROTECTION MECHANISMS                                                     │
│   ══════════════════════                                                    │
│   1. Chain ID in signature (EIP-155)                                       │
│   2. Nonce prevents double-spend                                           │
│   3. Transaction hash uniqueness                                           │
│                                                                             │
│   ⚠️  WARNING: Never use the same wallet on multiple networks              │
│       without understanding replay risks.                                  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 5. Security Rules

### 5.1 Encrypted Keystore

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         KEYSTORE ENCRYPTION                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   FILE FORMAT                                                               │
│   ═══════════                                                               │
│   • Standard: Ethereum Web3 Secret Storage (v3)                            │
│   • Cipher: AES-128-CTR                                                    │
│   • KDF: scrypt (N=262144, r=8, p=1)                                       │
│   • Checksum: Keccak256                                                    │
│                                                                             │
│   STRUCTURE                                                                 │
│   ═════════                                                                 │
│   {                                                                        │
│     "version": 3,                                                          │
│     "id": "uuid",                                                          │
│     "address": "0x...",                                                    │
│     "crypto": {                                                            │
│       "ciphertext": "...",                                                 │
│       "cipherparams": { "iv": "..." },                                     │
│       "cipher": "aes-128-ctr",                                             │
│       "kdf": "scrypt",                                                     │
│       "kdfparams": { ... },                                                │
│       "mac": "..."                                                         │
│     }                                                                      │
│   }                                                                        │
│                                                                             │
│   The private key is NEVER stored in plain text.                           │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 5.2 Password Requirements

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         PASSWORD POLICY                                     │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   MINIMUM REQUIREMENTS                                                      │
│   ════════════════════                                                      │
│   • Length: 8+ characters (12+ recommended)                                │
│   • Complexity: Mix of upper, lower, numbers, symbols                      │
│   • Uniqueness: Never reuse passwords                                      │
│                                                                             │
│   BEST PRACTICES                                                            │
│   ══════════════                                                            │
│   ✓ Use a password manager                                                 │
│   ✓ Generate random passwords                                              │
│   ✓ Store password securely (not in plain text)                            │
│   ✓ Use --password-file for automation                                     │
│                                                                             │
│   PASSWORD FILE                                                             │
│   ═════════════                                                             │
│   # Create secure password file                                            │
│   $ echo "your-secure-password" > ~/.secrets/wallet.pass                   │
│   $ chmod 600 ~/.secrets/wallet.pass                                       │
│                                                                             │
│   # Use in commands                                                        │
│   $ mbongo wallet transfer ... --password-file ~/.secrets/wallet.pass      │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 5.3 Mnemonic Rules

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         RECOVERY PHRASE SECURITY                            │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   YOUR MNEMONIC IS YOUR WALLET                                              │
│   ════════════════════════════                                              │
│   Anyone with your 24 words has FULL ACCESS to your funds.                 │
│   There is NO way to recover funds if mnemonic is lost.                    │
│                                                                             │
│   STORAGE GUIDELINES                                                        │
│   ══════════════════                                                        │
│   ✓ Write on paper (multiple copies)                                       │
│   ✓ Store in fireproof safe                                                │
│   ✓ Consider metal backup (fire/water resistant)                           │
│   ✓ Geographic distribution (different locations)                          │
│   ✓ Consider Shamir's Secret Sharing for large holdings                    │
│                                                                             │
│   NEVER                                                                     │
│   ═════                                                                     │
│   ✗ Store digitally (computer, phone, cloud)                               │
│   ✗ Take photos or screenshots                                             │
│   ✗ Email or message to anyone                                             │
│   ✗ Enter on any website                                                   │
│   ✗ Share with "support" (there is no support)                             │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 5.4 What to NEVER Do

```
╔═════════════════════════════════════════════════════════════════════════════╗
║                                                                             ║
║   🚫 CRITICAL DON'Ts                                                        ║
║                                                                             ║
║   1. LEAK KEYS                                                              ║
║      • Never paste private key or mnemonic anywhere                        ║
║      • Never share screen while viewing keys                               ║
║      • Never use clipboard for sensitive data                              ║
║                                                                             ║
║   2. MIX NETWORKS                                                           ║
║      • Never use same wallet on mainnet AND testnet                        ║
║      • Create separate wallets for each network                            ║
║      • Label wallets clearly (mainnet-validator, testnet-dev)              ║
║                                                                             ║
║   3. TRUST UNVERIFIED SOFTWARE                                              ║
║      • Only use official mbongo CLI                                        ║
║      • Verify checksums of downloaded binaries                             ║
║      • Never install "wallet recovery" tools                               ║
║                                                                             ║
║   4. IGNORE WARNINGS                                                        ║
║      • CLI warnings exist for a reason                                     ║
║      • Double-check addresses before sending                               ║
║      • Test with small amounts first                                       ║
║                                                                             ║
║   5. SKIP BACKUPS                                                           ║
║      • Always backup mnemonic BEFORE funding                               ║
║      • Test recovery process with empty wallet                             ║
║      • Keep backups updated                                                ║
║                                                                             ║
╚═════════════════════════════════════════════════════════════════════════════╝
```

---

## 6. Wallet Lifecycle Diagrams

### 6.1 Wallet Creation → Funding → First Transaction

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         WALLET LIFECYCLE                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   PHASE 1: CREATION                                                         │
│   ═════════════════                                                         │
│                                                                             │
│   $ mbongo wallet create                                                   │
│         │                                                                   │
│         ▼                                                                   │
│   ┌───────────────┐    ┌───────────────┐    ┌───────────────┐              │
│   │ Generate      │───▶│ Derive Keys   │───▶│ Encrypt &     │              │
│   │ Entropy       │    │ (BIP-44)      │    │ Save          │              │
│   └───────────────┘    └───────────────┘    └───────────────┘              │
│         │                                          │                        │
│         ▼                                          ▼                        │
│   ┌───────────────┐                        ┌───────────────┐              │
│   │ Display       │                        │ ~/.mbongo/    │              │
│   │ Mnemonic      │                        │ wallets/*.json│              │
│   └───────────────┘                        └───────────────┘              │
│         │                                                                   │
│         ▼                                                                   │
│   ┌───────────────────────────────────────────────────────────┐           │
│   │ ⚠️  User MUST backup mnemonic offline before proceeding   │           │
│   └───────────────────────────────────────────────────────────┘           │
│                                                                             │
│   PHASE 2: FUNDING                                                          │
│   ════════════════                                                          │
│                                                                             │
│   $ mbongo wallet address                                                  │
│         │                                                                   │
│         ▼                                                                   │
│   ┌───────────────┐    ┌───────────────┐    ┌───────────────┐              │
│   │ Get Address   │───▶│ Send MBO from │───▶│ Confirm on    │              │
│   │ 0x742d35...   │    │ Exchange/Peer │    │ Explorer      │              │
│   └───────────────┘    └───────────────┘    └───────────────┘              │
│                                                                             │
│   PHASE 3: FIRST TRANSACTION                                                │
│   ══════════════════════════                                                │
│                                                                             │
│   $ mbongo wallet transfer --to 0x... --amount 10                          │
│         │                                                                   │
│         ▼                                                                   │
│   ┌───────────────┐    ┌───────────────┐    ┌───────────────┐              │
│   │ Build Tx      │───▶│ Sign with     │───▶│ Broadcast to  │              │
│   │ (nonce, gas)  │    │ Private Key   │    │ Network       │              │
│   └───────────────┘    └───────────────┘    └───────────────┘              │
│                                                   │                         │
│                                                   ▼                         │
│                                          ┌───────────────┐                 │
│                                          │ Confirmed!    │                 │
│                                          │ Block #12345  │                 │
│                                          └───────────────┘                 │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 6.2 Signing & Verification Pipeline

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         SIGNING PIPELINE                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   SIGNING                                                                   │
│   ═══════                                                                   │
│                                                                             │
│   Message/Tx ──▶ Hash ──▶ Sign ──▶ Signature                               │
│       │           │        │          │                                     │
│       │           │        │          │                                     │
│       ▼           ▼        ▼          ▼                                     │
│   ┌───────┐   ┌───────┐ ┌───────┐ ┌───────────┐                           │
│   │ "Hi"  │──▶│Keccak │─▶│ECDSA  │─▶│ (v, r, s) │                         │
│   │       │   │256    │ │Sign   │ │           │                           │
│   └───────┘   └───────┘ └───────┘ └───────────┘                           │
│                            ▲                                               │
│                            │                                               │
│                     Private Key                                            │
│                    (from keystore)                                         │
│                                                                             │
│   ─────────────────────────────────────────────────────────────────────────│
│                                                                             │
│   VERIFICATION                                                              │
│   ════════════                                                              │
│                                                                             │
│   Message + Signature ──▶ Recover ──▶ Compare ──▶ Valid?                   │
│       │          │           │           │          │                       │
│       │          │           │           │          │                       │
│       ▼          ▼           ▼           ▼          ▼                       │
│   ┌───────┐  ┌───────┐   ┌───────┐   ┌───────┐  ┌───────┐                 │
│   │ "Hi"  │  │(v,r,s)│──▶│ECDSA  │──▶│Equals │──▶│ ✓ or ✗│                │
│   │       │  │       │   │Recover│   │ addr? │  │       │                 │
│   └───────┘  └───────┘   └───────┘   └───────┘  └───────┘                 │
│                              │           ▲                                 │
│                              ▼           │                                 │
│                        ┌───────────┐     │                                 │
│                        │ Recovered │─────┘                                 │
│                        │ Address   │                                       │
│                        └───────────┘                                       │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 6.3 Key Rotation Flow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         KEY ROTATION FLOW                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   WHY ROTATE?                                                               │
│   • Suspected compromise                                                   │
│   • Security best practice                                                 │
│   • Session key expiration                                                 │
│                                                                             │
│   ROTATION PROCESS                                                          │
│   ════════════════                                                          │
│                                                                             │
│   ┌─────────────┐                                                          │
│   │ OLD WALLET  │                                                          │
│   │ 0x742d...   │                                                          │
│   └──────┬──────┘                                                          │
│          │                                                                  │
│          │  1. Create new wallet                                           │
│          │     $ mbongo wallet create --name new-wallet                    │
│          │                                                                  │
│          ▼                                                                  │
│   ┌─────────────┐                                                          │
│   │ NEW WALLET  │ ◀──── Backup mnemonic FIRST!                             │
│   │ 0x8Ba1...   │                                                          │
│   └──────┬──────┘                                                          │
│          │                                                                  │
│          │  2. Transfer all funds                                          │
│          │     $ mbongo wallet transfer --from old --to 0x8Ba1... --all    │
│          │                                                                  │
│   ┌──────┴──────┐                                                          │
│   │             │                                                          │
│   ▼             ▼                                                          │
│   ┌─────────────┐    ┌─────────────┐                                       │
│   │ Update      │    │ Update      │                                       │
│   │ Validator   │    │ Delegations │                                       │
│   │ Key         │    │             │                                       │
│   └─────────────┘    └─────────────┘                                       │
│          │                  │                                               │
│          └────────┬─────────┘                                               │
│                   │                                                         │
│                   │  3. Verify all moved                                   │
│                   │     $ mbongo wallet balance --name old-wallet          │
│                   │     (should be 0)                                      │
│                   │                                                         │
│                   │  4. Delete old wallet (optional)                       │
│                   │     $ mbongo wallet delete --name old-wallet           │
│                   ▼                                                         │
│          ┌─────────────┐                                                   │
│          │ ROTATION    │                                                   │
│          │ COMPLETE    │                                                   │
│          └─────────────┘                                                   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 7. Cross-Links

### Related Documentation

| Document | Description |
|----------|-------------|
| [cli_overview.md](./cli_overview.md) | CLI overview and conventions |
| [fee_model.md](./fee_model.md) | Gas and fee structure |
| [staking_model.md](./staking_model.md) | Staking operations |
| [compute_value.md](./compute_value.md) | PoUW compute payments |
| [economic_security.md](./economic_security.md) | Security model |

### Quick Reference

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         WALLET COMMANDS QUICK REFERENCE                     │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│   CREATION & RECOVERY              QUERIES                                  │
│   ────────────────────             ───────                                  │
│   mbongo wallet create             mbongo wallet address                    │
│   mbongo wallet restore            mbongo wallet balance                    │
│   mbongo wallet import             mbongo wallet history                    │
│   mbongo wallet export                                                      │
│                                                                             │
│   TRANSACTIONS                     SIGNING                                  │
│   ────────────                     ───────                                  │
│   mbongo wallet transfer           mbongo wallet sign                       │
│   --to <ADDR>                      mbongo wallet verify                     │
│   --amount <MBO>                                                            │
│   --gas-price <GWEI>                                                        │
│                                                                             │
│   MANAGEMENT                       DANGER ZONE                              │
│   ──────────                       ───────────                              │
│   mbongo wallet watch              mbongo wallet mnemonic                   │
│   mbongo wallet keys               mbongo wallet delete                     │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

*This document provides the complete reference for `mbongo wallet` commands. For general CLI information, see [cli_overview.md](./cli_overview.md).*

