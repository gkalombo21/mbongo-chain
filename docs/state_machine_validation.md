# Mbongo Chain — State Machine Validation

This document provides a comprehensive specification of the state machine in Mbongo Chain, covering state representation, transaction execution, validation rules, and security invariants.

---

## 1. Overview

### Purpose of the State Machine

The **state machine** is the core component that defines how Mbongo Chain transitions from one valid state to another. It processes transactions deterministically, ensuring all nodes arrive at identical states given identical inputs.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     STATE MACHINE PURPOSE                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Input:                                                                     │
│  • Current state (S_n)                                                     │
│  • Block transactions (T_1, T_2, ..., T_k)                                 │
│  • Block context (proposer, timestamp, height)                             │
│                                                                             │
│  Processing:                                                                │
│  • Validate each transaction                                               │
│  • Execute state transitions                                               │
│  • Generate receipts and logs                                              │
│                                                                             │
│  Output:                                                                    │
│  • New state (S_{n+1})                                                     │
│  • Execution receipts                                                      │
│  • State root hash                                                         │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Execution Determinism Requirements

| Requirement | Description |
|-------------|-------------|
| **Bit-for-bit reproducibility** | Same inputs always produce same outputs |
| **No external dependencies** | No system calls, network, or file I/O |
| **Fixed arithmetic** | No floating point; checked integer math |
| **Ordered execution** | Transactions processed in strict order |
| **Bounded computation** | Gas limits prevent infinite loops |

### Role in Consensus-Critical Logic

The state machine is consensus-critical:

- All validators must agree on state transitions
- Any divergence causes chain split
- State root is committed in block header
- Incorrect implementation breaks consensus

### High-Level State Transition

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     STATE TRANSITION FLOW                                   │
└─────────────────────────────────────────────────────────────────────────────┘

  ┌──────────────┐                                         ┌──────────────┐
  │              │                                         │              │
  │   STATE_N    │                                         │  STATE_N+1   │
  │              │                                         │              │
  │  root: 0xABC │                                         │  root: 0xDEF │
  │              │                                         │              │
  └──────┬───────┘                                         └──────▲───────┘
         │                                                        │
         │                                                        │
         ▼                                                        │
  ┌─────────────────────────────────────────────────────────────────────────┐
  │                                                                         │
  │                         EXECUTION ENGINE                                │
  │                                                                         │
  │   ┌─────────┐    ┌─────────┐    ┌─────────┐    ┌─────────┐            │
  │   │  Load   │───▶│Validate │───▶│ Execute │───▶│ Commit  │────────────┼──▶
  │   │  State  │    │   Tx    │    │   Tx    │    │ Changes │            │
  │   └─────────┘    └─────────┘    └─────────┘    └─────────┘            │
  │                                                                         │
  │                        For each transaction                             │
  │                                                                         │
  └─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    │
                                    ▼
                           ┌───────────────┐
                           │   RECEIPTS    │
                           │   + LOGS      │
                           └───────────────┘
```

---

## 2. State Representation

### Account Model Structure

Mbongo Chain uses an account-based model (similar to Ethereum):

```rust
/// Account state representation
#[derive(Clone, Debug, Default, PartialEq, Eq)]
pub struct Account {
    /// Transaction counter (prevents replay)
    pub nonce: u64,
    
    /// Token balance (in smallest unit)
    pub balance: u128,
    
    /// Storage trie root (for contract accounts)
    pub storage_root: Hash,
    
    /// Code hash (for contract accounts, EMPTY_CODE_HASH for EOA)
    pub code_hash: Hash,
}

impl Account {
    /// Check if account is externally owned (not a contract)
    pub fn is_eoa(&self) -> bool {
        self.code_hash == EMPTY_CODE_HASH
    }
    
    /// Check if account is a contract
    pub fn is_contract(&self) -> bool {
        self.code_hash != EMPTY_CODE_HASH
    }
    
    /// Check if account is empty (can be pruned)
    pub fn is_empty(&self) -> bool {
        self.nonce == 0 
            && self.balance == 0 
            && self.code_hash == EMPTY_CODE_HASH
    }
    
    /// Encode account for trie storage
    pub fn encode(&self) -> Vec<u8> {
        let mut buf = Vec::with_capacity(8 + 16 + 32 + 32);
        buf.extend_from_slice(&self.nonce.to_le_bytes());
        buf.extend_from_slice(&self.balance.to_le_bytes());
        buf.extend_from_slice(&self.storage_root);
        buf.extend_from_slice(&self.code_hash);
        buf
    }
    
    /// Decode account from trie storage
    pub fn decode(data: &[u8]) -> Result<Self, DecodeError> {
        if data.len() != 88 {
            return Err(DecodeError::InvalidLength);
        }
        Ok(Self {
            nonce: u64::from_le_bytes(data[0..8].try_into().unwrap()),
            balance: u128::from_le_bytes(data[8..24].try_into().unwrap()),
            storage_root: data[24..56].try_into().unwrap(),
            code_hash: data[56..88].try_into().unwrap(),
        })
    }
}
```

### Key Data Structures

```rust
/// Global state representation
pub struct State {
    /// Account trie (address -> account)
    accounts: MerkleTrie<Address, Account>,
    
    /// Contract storage tries (address -> storage trie)
    storage: HashMap<Address, MerkleTrie<Hash, Hash>>,
    
    /// Contract code (code_hash -> bytecode)
    code: HashMap<Hash, Vec<u8>>,
    
    /// Current block context
    block_context: BlockContext,
    
    /// Accumulated transaction receipts
    receipts: Vec<TransactionReceipt>,
    
    /// Pending state changes (for rollback)
    journal: StateJournal,
}

/// Block execution context
#[derive(Clone, Debug)]
pub struct BlockContext {
    /// Block proposer (receives fees)
    pub proposer: Address,
    
    /// Block timestamp (unix seconds)
    pub timestamp: u64,
    
    /// Block height
    pub height: u64,
    
    /// Block gas limit
    pub gas_limit: u64,
    
    /// Base fee (EIP-1559 style, future)
    pub base_fee: u128,
    
    /// Previous block hash (for BLOCKHASH opcode)
    pub parent_hash: Hash,
}

/// Transaction receipt
#[derive(Clone, Debug)]
pub struct TransactionReceipt {
    /// Transaction hash
    pub tx_hash: Hash,
    
    /// Execution status
    pub status: ExecutionStatus,
    
    /// Gas consumed
    pub gas_used: u64,
    
    /// Cumulative gas in block
    pub cumulative_gas: u64,
    
    /// Emitted logs
    pub logs: Vec<Log>,
    
    /// Return data (if any)
    pub output: Vec<u8>,
    
    /// Post-execution state root
    pub state_root: Hash,
}

/// Execution status
#[derive(Clone, Debug, PartialEq, Eq)]
pub enum ExecutionStatus {
    Success,
    Revert,
    OutOfGas,
    InvalidOpcode,
    StackOverflow,
    StackUnderflow,
}

/// Log entry
#[derive(Clone, Debug)]
pub struct Log {
    /// Contract that emitted the log
    pub address: Address,
    
    /// Indexed topics (max 4)
    pub topics: Vec<Hash>,
    
    /// Non-indexed data
    pub data: Vec<u8>,
}
```

### Trie Layout

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     STATE TRIE LAYOUT                                       │
└─────────────────────────────────────────────────────────────────────────────┘

  World State Trie
  ════════════════
  
                         ┌─────────────────┐
                         │   State Root    │
                         │    (32 bytes)   │
                         └────────┬────────┘
                                  │
            ┌─────────────────────┼─────────────────────┐
            │                     │                     │
            ▼                     ▼                     ▼
     ┌────────────┐        ┌────────────┐        ┌────────────┐
     │  Account   │        │  Account   │        │  Account   │
     │   0x123... │        │   0x456... │        │   0x789... │
     └─────┬──────┘        └─────┬──────┘        └────────────┘
           │                     │
           ▼                     ▼
    ┌─────────────┐       ┌─────────────┐
    │   Storage   │       │   Storage   │
    │    Trie     │       │    Trie     │
    └─────────────┘       └─────────────┘


  Account Entry (RLP encoded):
  ════════════════════════════
  
  ┌─────────┬─────────┬──────────────┬────────────┐
  │  nonce  │ balance │ storage_root │ code_hash  │
  │ 8 bytes │16 bytes │   32 bytes   │  32 bytes  │
  └─────────┴─────────┴──────────────┴────────────┘


  Storage Entry:
  ══════════════
  
  Key: keccak256(slot_number)  →  Value: 32-byte word
```

### State Interface

```rust
impl State {
    /// Create new state from root
    pub fn from_root(root: Hash, db: &Database) -> Result<Self, StateError> {
        let accounts = MerkleTrie::from_root(root, db)?;
        Ok(Self {
            accounts,
            storage: HashMap::new(),
            code: HashMap::new(),
            block_context: BlockContext::default(),
            receipts: Vec::new(),
            journal: StateJournal::new(),
        })
    }
    
    /// Get account (or default if not exists)
    pub fn get_account(&self, address: &Address) -> Account {
        self.accounts
            .get(address)
            .unwrap_or_default()
    }
    
    /// Set account
    pub fn set_account(&mut self, address: &Address, account: &Account) {
        // Record for rollback
        let old = self.get_account(address);
        self.journal.record_account_change(address, old);
        
        if account.is_empty() {
            self.accounts.delete(address);
        } else {
            self.accounts.insert(address, account.clone());
        }
    }
    
    /// Get balance
    pub fn get_balance(&self, address: &Address) -> u128 {
        self.get_account(address).balance
    }
    
    /// Get nonce
    pub fn get_nonce(&self, address: &Address) -> u64 {
        self.get_account(address).nonce
    }
    
    /// Get storage value
    pub fn get_storage(&self, address: &Address, key: &Hash) -> Hash {
        self.storage
            .get(address)
            .and_then(|trie| trie.get(key))
            .unwrap_or(ZERO_HASH)
    }
    
    /// Set storage value
    pub fn set_storage(&mut self, address: &Address, key: &Hash, value: Hash) {
        // Record for rollback
        let old = self.get_storage(address, key);
        self.journal.record_storage_change(address, key, old);
        
        let trie = self.storage
            .entry(*address)
            .or_insert_with(MerkleTrie::new);
        
        if value == ZERO_HASH {
            trie.delete(key);
        } else {
            trie.insert(key, value);
        }
    }
    
    /// Compute current state root
    pub fn root(&self) -> Hash {
        self.accounts.root()
    }
    
    /// Create checkpoint for rollback
    pub fn checkpoint(&mut self) -> CheckpointId {
        self.journal.checkpoint()
    }
    
    /// Rollback to checkpoint
    pub fn rollback(&mut self, checkpoint: CheckpointId) {
        for change in self.journal.changes_since(checkpoint).rev() {
            match change {
                JournalEntry::AccountChange { address, old_value } => {
                    if old_value.is_empty() {
                        self.accounts.delete(&address);
                    } else {
                        self.accounts.insert(&address, old_value);
                    }
                }
                JournalEntry::StorageChange { address, key, old_value } => {
                    let trie = self.storage.get_mut(&address).unwrap();
                    if old_value == ZERO_HASH {
                        trie.delete(&key);
                    } else {
                        trie.insert(&key, old_value);
                    }
                }
            }
        }
        self.journal.revert_to(checkpoint);
    }
    
    /// Commit checkpoint (discard rollback data)
    pub fn commit(&mut self, checkpoint: CheckpointId) {
        self.journal.commit(checkpoint);
    }
}
```

---

## 3. Transaction Execution Flow

### Execution Pipeline

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     TRANSACTION EXECUTION PIPELINE                          │
└─────────────────────────────────────────────────────────────────────────────┘

  Input: Transaction + State
  ══════════════════════════

  ┌─────────────────┐
  │  Raw Transaction│
  │                 │
  │  • from (sig)   │
  │  • to           │
  │  • value        │
  │  • data         │
  │  • gas_limit    │
  │  • gas_price    │
  │  • nonce        │
  └────────┬────────┘
           │
           ▼
  ┌─────────────────────────────────────────────────────────────────────────┐
  │  PHASE 1: PRE-VALIDATION                                                │
  │                                                                         │
  │  ┌─────────────┐   ┌─────────────┐   ┌─────────────┐   ┌─────────────┐ │
  │  │  Decode &   │──▶│  Recover    │──▶│   Check     │──▶│   Check     │ │
  │  │  Format     │   │   Sender    │   │   Nonce     │   │  Balance    │ │
  │  └─────────────┘   └─────────────┘   └─────────────┘   └─────────────┘ │
  │                                                                         │
  │  Failure: Invalid format, bad signature, wrong nonce, insufficient bal │
  └────────────────────────────────────────┬────────────────────────────────┘
                                           │
                                           ▼
  ┌─────────────────────────────────────────────────────────────────────────┐
  │  PHASE 2: GAS RESERVATION                                               │
  │                                                                         │
  │  ┌─────────────┐   ┌─────────────┐   ┌─────────────┐                   │
  │  │  Calculate  │──▶│   Deduct    │──▶│  Increment  │                   │
  │  │ Intrinsic   │   │  Max Cost   │   │   Nonce     │                   │
  │  │    Gas      │   │ (upfront)   │   │             │                   │
  │  └─────────────┘   └─────────────┘   └─────────────┘                   │
  │                                                                         │
  │  max_cost = gas_limit * gas_price + value                              │
  └────────────────────────────────────────┬────────────────────────────────┘
                                           │
                                           ▼
  ┌─────────────────────────────────────────────────────────────────────────┐
  │  PHASE 3: EXECUTION                                                     │
  │                                                                         │
  │  ┌─────────────┐   ┌─────────────┐   ┌─────────────┐                   │
  │  │   Create    │──▶│   Execute   │──▶│   Collect   │                   │
  │  │  Context    │   │  Operation  │   │    Logs     │                   │
  │  └─────────────┘   └─────────────┘   └─────────────┘                   │
  │                                                                         │
  │  Transfer: sender → recipient (value)                                  │
  │  Contract: execute bytecode with gas metering                          │
  └────────────────────────────────────────┬────────────────────────────────┘
                                           │
                                           ▼
  ┌─────────────────────────────────────────────────────────────────────────┐
  │  PHASE 4: FINALIZATION                                                  │
  │                                                                         │
  │  ┌─────────────┐   ┌─────────────┐   ┌─────────────┐   ┌─────────────┐ │
  │  │   Refund    │──▶│    Pay      │──▶│   Create    │──▶│   Update    │ │
  │  │ Unused Gas  │   │  Proposer   │   │   Receipt   │   │   State     │ │
  │  └─────────────┘   └─────────────┘   └─────────────┘   └─────────────┘ │
  │                                                                         │
  │  refund = (gas_limit - gas_used) * gas_price                           │
  │  fee = gas_used * gas_price                                            │
  └────────────────────────────────────────┬────────────────────────────────┘
                                           │
                                           ▼
  Output: Receipt + Updated State
  ═══════════════════════════════
```

### Signature Verification Logic

```rust
/// Verify transaction signature and recover sender
pub fn verify_and_recover_sender(tx: &Transaction) -> Result<Address, SignatureError> {
    // Compute message hash (EIP-155 compatible)
    let message_hash = if let Some(chain_id) = tx.chain_id {
        // EIP-155: include chain_id in signing hash
        compute_eip155_hash(tx, chain_id)
    } else {
        // Legacy: sign over raw transaction fields
        compute_legacy_hash(tx)
    };
    
    // Extract signature components
    let (r, s, v) = decode_signature(&tx.signature)?;
    
    // Validate signature format
    if s > SECP256K1_N_DIV_2 {
        return Err(SignatureError::MalleableSignature);
    }
    
    // Recover public key
    let recovery_id = compute_recovery_id(v, tx.chain_id)?;
    let public_key = secp256k1_recover(&message_hash, &r, &s, recovery_id)?;
    
    // Derive address from public key
    let address = public_key_to_address(&public_key);
    
    Ok(address)
}

/// Compute EIP-155 signing hash
fn compute_eip155_hash(tx: &Transaction, chain_id: u64) -> Hash {
    let mut hasher = Keccak256::new();
    
    hasher.update(&rlp_encode_u64(tx.nonce));
    hasher.update(&rlp_encode_u128(tx.gas_price));
    hasher.update(&rlp_encode_u64(tx.gas_limit));
    hasher.update(&rlp_encode_address(&tx.to));
    hasher.update(&rlp_encode_u128(tx.value));
    hasher.update(&rlp_encode_bytes(&tx.data));
    hasher.update(&rlp_encode_u64(chain_id));
    hasher.update(&[0x80]);  // empty r
    hasher.update(&[0x80]);  // empty s
    
    hasher.finalize()
}
```

### Gas Metering Rules

```rust
/// Gas costs for operations
pub mod gas_costs {
    pub const TX_BASE: u64 = 21_000;
    pub const TX_DATA_ZERO: u64 = 4;
    pub const TX_DATA_NONZERO: u64 = 16;
    pub const TX_CREATE: u64 = 32_000;
    
    pub const SLOAD_COLD: u64 = 2_100;
    pub const SLOAD_WARM: u64 = 100;
    pub const SSTORE_SET: u64 = 20_000;
    pub const SSTORE_RESET: u64 = 2_900;
    pub const SSTORE_CLEAR_REFUND: u64 = 4_800;
    
    pub const CALL_BASE: u64 = 100;
    pub const CALL_VALUE: u64 = 9_000;
    pub const CALL_NEW_ACCOUNT: u64 = 25_000;
    
    pub const BALANCE: u64 = 2_600;
    pub const EXTCODESIZE: u64 = 2_600;
    pub const EXTCODECOPY_BASE: u64 = 2_600;
    
    pub const LOG_BASE: u64 = 375;
    pub const LOG_TOPIC: u64 = 375;
    pub const LOG_DATA: u64 = 8;
}

/// Gas metering context
pub struct GasMeter {
    /// Gas limit for transaction
    limit: u64,
    /// Gas used so far
    used: u64,
    /// Accumulated refunds
    refund: u64,
}

impl GasMeter {
    pub fn new(limit: u64) -> Self {
        Self {
            limit,
            used: 0,
            refund: 0,
        }
    }
    
    /// Consume gas (fails if insufficient)
    pub fn consume(&mut self, amount: u64) -> Result<(), OutOfGasError> {
        let new_used = self.used.checked_add(amount)
            .ok_or(OutOfGasError::Overflow)?;
        
        if new_used > self.limit {
            return Err(OutOfGasError::Exceeded {
                limit: self.limit,
                used: new_used,
            });
        }
        
        self.used = new_used;
        Ok(())
    }
    
    /// Add refund (capped at 50% of gas used)
    pub fn add_refund(&mut self, amount: u64) {
        self.refund = self.refund.saturating_add(amount);
    }
    
    /// Calculate gas used after refunds
    pub fn gas_used(&self) -> u64 {
        let max_refund = self.used / 2;  // Cap at 50%
        let actual_refund = self.refund.min(max_refund);
        self.used - actual_refund
    }
    
    /// Remaining gas
    pub fn gas_remaining(&self) -> u64 {
        self.limit.saturating_sub(self.used)
    }
}
```

### Nonce Validation

```rust
/// Validate transaction nonce
pub fn validate_nonce(
    tx: &Transaction,
    sender: &Address,
    state: &State,
) -> Result<(), NonceError> {
    let expected_nonce = state.get_nonce(sender);
    
    if tx.nonce < expected_nonce {
        return Err(NonceError::TooLow {
            expected: expected_nonce,
            got: tx.nonce,
        });
    }
    
    if tx.nonce > expected_nonce {
        return Err(NonceError::TooHigh {
            expected: expected_nonce,
            got: tx.nonce,
        });
    }
    
    Ok(())
}
```

### Balance Check Logic

```rust
/// Validate sender has sufficient balance
pub fn validate_balance(
    tx: &Transaction,
    sender: &Address,
    state: &State,
) -> Result<(), BalanceError> {
    let sender_balance = state.get_balance(sender);
    
    // Calculate maximum cost
    let gas_cost = tx.gas_limit
        .checked_mul(tx.gas_price as u64)
        .ok_or(BalanceError::Overflow)?;
    
    let total_cost = (gas_cost as u128)
        .checked_add(tx.value)
        .ok_or(BalanceError::Overflow)?;
    
    if sender_balance < total_cost {
        return Err(BalanceError::Insufficient {
            required: total_cost,
            available: sender_balance,
        });
    }
    
    Ok(())
}
```

### State Transition Phases

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     STATE TRANSITION PHASES                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  PRE-STATE (Before execution)                                               │
│  ═══════════════════════════                                                │
│  • Sender nonce: N                                                         │
│  • Sender balance: B                                                       │
│  • Recipient balance: R                                                    │
│  • State root: ROOT_PRE                                                    │
│                                                                             │
│                              │                                              │
│                              ▼                                              │
│                                                                             │
│  INTRA-STATE (During execution)                                            │
│  ═══════════════════════════════                                           │
│  • Checkpoint created                                                      │
│  • Sender balance: B - max_cost                                            │
│  • Sender nonce: N + 1                                                     │
│  • Execution proceeds with gas metering                                    │
│  • Storage reads/writes tracked                                            │
│  • Logs accumulated                                                        │
│                                                                             │
│                              │                                              │
│                   ┌──────────┴──────────┐                                  │
│                   │                     │                                  │
│                   ▼                     ▼                                  │
│                                                                             │
│  POST-STATE (Success)              POST-STATE (Revert)                     │
│  ════════════════════              ═══════════════════                     │
│  • Changes committed               • Changes rolled back                   │
│  • Sender: B - gas_used*price      • Sender: B - gas_used*price           │
│  • Recipient: R + value            • Recipient: R (unchanged)              │
│  • Proposer: +fee                  • Proposer: +fee                        │
│  • State root: ROOT_POST           • State root: ROOT_REVERT               │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Apply Transaction Implementation

```rust
/// Apply a single transaction to state
pub fn apply_transaction(
    state: &mut State,
    tx: &Transaction,
    block_context: &BlockContext,
    cumulative_gas: &mut u64,
) -> Result<TransactionReceipt, ExecutionError> {
    // ═══════════════════════════════════════════════════════════════════════
    // PHASE 1: PRE-VALIDATION
    // ═══════════════════════════════════════════════════════════════════════
    
    // Recover sender from signature
    let sender = verify_and_recover_sender(tx)
        .map_err(|e| ExecutionError::InvalidSignature(e))?;
    
    // Validate nonce
    validate_nonce(tx, &sender, state)
        .map_err(|e| ExecutionError::InvalidNonce(e))?;
    
    // Validate balance
    validate_balance(tx, &sender, state)
        .map_err(|e| ExecutionError::InsufficientBalance(e))?;
    
    // ═══════════════════════════════════════════════════════════════════════
    // PHASE 2: GAS RESERVATION
    // ═══════════════════════════════════════════════════════════════════════
    
    // Calculate intrinsic gas
    let intrinsic_gas = calculate_intrinsic_gas(tx);
    if tx.gas_limit < intrinsic_gas {
        return Err(ExecutionError::IntrinsicGasTooLow {
            required: intrinsic_gas,
            provided: tx.gas_limit,
        });
    }
    
    // Create gas meter
    let mut gas_meter = GasMeter::new(tx.gas_limit);
    gas_meter.consume(intrinsic_gas)?;
    
    // Create checkpoint for potential rollback
    let checkpoint = state.checkpoint();
    
    // Deduct maximum cost from sender
    let max_cost = tx.gas_limit as u128 * tx.gas_price + tx.value;
    let mut sender_account = state.get_account(&sender);
    sender_account.balance -= max_cost;
    sender_account.nonce += 1;
    state.set_account(&sender, &sender_account);
    
    // ═══════════════════════════════════════════════════════════════════════
    // PHASE 3: EXECUTION
    // ═══════════════════════════════════════════════════════════════════════
    
    let execution_result = match tx.to {
        Some(recipient) => {
            // Value transfer or contract call
            execute_call(
                state,
                &sender,
                &recipient,
                tx.value,
                &tx.data,
                &mut gas_meter,
                block_context,
            )
        }
        None => {
            // Contract creation
            execute_create(
                state,
                &sender,
                tx.value,
                &tx.data,
                &mut gas_meter,
                block_context,
            )
        }
    };
    
    // ═══════════════════════════════════════════════════════════════════════
    // PHASE 4: FINALIZATION
    // ═══════════════════════════════════════════════════════════════════════
    
    let (status, logs, output) = match execution_result {
        Ok(result) => {
            // Success: commit changes
            state.commit(checkpoint);
            (ExecutionStatus::Success, result.logs, result.output)
        }
        Err(ExecutionError::Revert { output, .. }) => {
            // Revert: rollback changes but keep nonce/gas
            state.rollback(checkpoint);
            
            // Re-apply nonce increment
            let mut sender_account = state.get_account(&sender);
            sender_account.nonce += 1;
            state.set_account(&sender, &sender_account);
            
            (ExecutionStatus::Revert, vec![], output)
        }
        Err(ExecutionError::OutOfGas) => {
            // Out of gas: rollback, consume all gas
            state.rollback(checkpoint);
            
            let mut sender_account = state.get_account(&sender);
            sender_account.nonce += 1;
            state.set_account(&sender, &sender_account);
            
            gas_meter.used = tx.gas_limit;  // Consume all gas
            (ExecutionStatus::OutOfGas, vec![], vec![])
        }
        Err(e) => {
            return Err(e);
        }
    };
    
    // Calculate actual gas used (after refunds)
    let gas_used = gas_meter.gas_used();
    
    // Refund unused gas to sender
    let gas_refund = (tx.gas_limit - gas_used) as u128 * tx.gas_price;
    let mut sender_account = state.get_account(&sender);
    sender_account.balance += gas_refund;
    state.set_account(&sender, &sender_account);
    
    // Pay fee to block proposer
    let fee = gas_used as u128 * tx.gas_price;
    let mut proposer_account = state.get_account(&block_context.proposer);
    proposer_account.balance += fee;
    state.set_account(&block_context.proposer, &proposer_account);
    
    // Update cumulative gas
    *cumulative_gas += gas_used;
    
    // Create receipt
    let receipt = TransactionReceipt {
        tx_hash: tx.hash(),
        status,
        gas_used,
        cumulative_gas: *cumulative_gas,
        logs,
        output,
        state_root: state.root(),
    };
    
    Ok(receipt)
}
```

---

## 4. Execution Rules

### Deterministic Function Ordering

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     DETERMINISTIC EXECUTION RULES                           │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  1. TRANSACTION ORDERING                                                    │
│     ─────────────────────                                                   │
│     • Transactions executed in block order (index 0, 1, 2, ...)            │
│     • No reordering during execution                                       │
│     • Parallel execution must preserve sequential semantics                 │
│                                                                             │
│  2. STATE ACCESS ORDERING                                                   │
│     ──────────────────────                                                  │
│     • Account lookups: deterministic trie traversal                        │
│     • Storage reads: deterministic key ordering                            │
│     • Map/Set iteration: use BTreeMap/BTreeSet (not HashMap/HashSet)       │
│                                                                             │
│  3. ARITHMETIC ORDERING                                                     │
│     ─────────────────────                                                   │
│     • Evaluation order: left-to-right                                      │
│     • Overflow checking: fail deterministically                            │
│     • Division by zero: fail deterministically                             │
│                                                                             │
│  4. LOG ORDERING                                                            │
│     ──────────────                                                          │
│     • Logs emitted in execution order                                      │
│     • Sub-call logs interleaved correctly                                  │
│     • Reverted logs discarded                                              │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Randomness Restrictions

```rust
/// Allowed sources of "randomness" (all deterministic)
pub mod allowed_randomness {
    use super::*;
    
    /// Derive random value from block context
    pub fn block_random(block: &BlockContext, domain: &[u8]) -> Hash {
        let mut hasher = Blake3::new();
        hasher.update(&block.parent_hash);
        hasher.update(&block.timestamp.to_le_bytes());
        hasher.update(&block.height.to_le_bytes());
        hasher.update(domain);
        hasher.finalize()
    }
    
    /// Derive random value from transaction
    pub fn tx_random(tx: &Transaction, block: &BlockContext, domain: &[u8]) -> Hash {
        let mut hasher = Blake3::new();
        hasher.update(&tx.hash());
        hasher.update(&block.parent_hash);
        hasher.update(domain);
        hasher.finalize()
    }
}

/// Forbidden randomness sources
pub mod forbidden_randomness {
    // ❌ std::time::SystemTime
    // ❌ std::time::Instant
    // ❌ rand::random()
    // ❌ /dev/urandom
    // ❌ RDRAND instruction
    // ❌ Thread ID
    // ❌ Process ID
    // ❌ Memory addresses
    // ❌ HashMap iteration order
}
```

### Memory Model Rules

```rust
/// Deterministic storage operations
pub struct DeterministicStorage {
    /// Ordered storage map (NOT HashMap)
    inner: BTreeMap<Hash, Hash>,
}

impl DeterministicStorage {
    /// Get value (deterministic)
    pub fn get(&self, key: &Hash) -> Hash {
        self.inner.get(key).copied().unwrap_or(ZERO_HASH)
    }
    
    /// Set value (deterministic)
    pub fn set(&mut self, key: Hash, value: Hash) {
        if value == ZERO_HASH {
            self.inner.remove(&key);
        } else {
            self.inner.insert(key, value);
        }
    }
    
    /// Iterate in deterministic order
    pub fn iter(&self) -> impl Iterator<Item = (&Hash, &Hash)> {
        self.inner.iter()  // BTreeMap guarantees sorted order
    }
    
    /// Compute storage root (deterministic)
    pub fn root(&self) -> Hash {
        let mut trie = MerkleTrie::new();
        for (key, value) in &self.inner {
            trie.insert(key, value);
        }
        trie.root()
    }
}
```

### Expected Execution Outputs

```rust
/// Execution output structure
pub struct ExecutionOutput {
    /// New state root after execution
    pub state_root: Hash,
    
    /// Transaction receipts
    pub receipts: Vec<TransactionReceipt>,
    
    /// Receipts merkle root
    pub receipts_root: Hash,
    
    /// All logs from all transactions
    pub logs: Vec<Log>,
    
    /// Logs bloom filter
    pub logs_bloom: Bloom,
    
    /// Total gas used in block
    pub gas_used: u64,
}

impl ExecutionOutput {
    /// Build from execution results
    pub fn build(
        state: &State,
        receipts: Vec<TransactionReceipt>,
    ) -> Self {
        let logs: Vec<Log> = receipts
            .iter()
            .flat_map(|r| r.logs.clone())
            .collect();
        
        let logs_bloom = compute_logs_bloom(&logs);
        let receipts_root = compute_receipts_root(&receipts);
        let gas_used = receipts.last().map(|r| r.cumulative_gas).unwrap_or(0);
        
        Self {
            state_root: state.root(),
            receipts,
            receipts_root,
            logs,
            logs_bloom,
            gas_used,
        }
    }
}
```

---

## 5. State Updates

### Update Rules

```rust
/// Nonce update rules
pub fn update_nonce(state: &mut State, address: &Address) {
    let mut account = state.get_account(address);
    
    // Nonce MUST increment by exactly 1
    account.nonce = account.nonce.checked_add(1)
        .expect("nonce overflow");
    
    state.set_account(address, &account);
}

/// Balance update rules
pub fn update_balance(
    state: &mut State,
    address: &Address,
    delta: i128,
) -> Result<(), BalanceError> {
    let mut account = state.get_account(address);
    
    if delta >= 0 {
        // Credit: checked addition
        account.balance = account.balance
            .checked_add(delta as u128)
            .ok_or(BalanceError::Overflow)?;
    } else {
        // Debit: checked subtraction
        let amount = (-delta) as u128;
        account.balance = account.balance
            .checked_sub(amount)
            .ok_or(BalanceError::Underflow)?;
    }
    
    state.set_account(address, &account);
    Ok(())
}

/// Storage write rules
pub fn write_storage(
    state: &mut State,
    address: &Address,
    key: &Hash,
    value: Hash,
    gas_meter: &mut GasMeter,
    access_tracker: &mut AccessTracker,
) -> Result<(), ExecutionError> {
    let current = state.get_storage(address, key);
    let original = access_tracker.get_original(address, key);
    
    // Calculate gas cost based on state change
    let gas_cost = calculate_sstore_gas(original, current, value, access_tracker);
    gas_meter.consume(gas_cost)?;
    
    // Calculate potential refund
    let refund = calculate_sstore_refund(original, current, value);
    gas_meter.add_refund(refund);
    
    // Apply storage change
    state.set_storage(address, key, value);
    
    Ok(())
}

/// Receipt creation rules
pub fn create_receipt(
    tx: &Transaction,
    status: ExecutionStatus,
    gas_used: u64,
    cumulative_gas: u64,
    logs: Vec<Log>,
    output: Vec<u8>,
    state: &State,
) -> TransactionReceipt {
    TransactionReceipt {
        tx_hash: tx.hash(),
        status,
        gas_used,
        cumulative_gas,
        logs,
        output,
        state_root: state.root(),
    }
}

/// Log emission rules
pub fn emit_log(
    logs: &mut Vec<Log>,
    address: Address,
    topics: Vec<Hash>,
    data: Vec<u8>,
    gas_meter: &mut GasMeter,
) -> Result<(), ExecutionError> {
    // Validate topic count
    if topics.len() > 4 {
        return Err(ExecutionError::TooManyTopics);
    }
    
    // Calculate gas cost
    let gas_cost = gas_costs::LOG_BASE
        + gas_costs::LOG_TOPIC * topics.len() as u64
        + gas_costs::LOG_DATA * data.len() as u64;
    
    gas_meter.consume(gas_cost)?;
    
    logs.push(Log { address, topics, data });
    
    Ok(())
}
```

### Failure Mode Table

| Category | Error | Cause | State Effect | Gas Consumed |
|----------|-------|-------|--------------|--------------|
| **Invalid Nonce** | `NonceTooLow` | Replayed transaction | None | 0 |
| | `NonceTooHigh` | Gap in sequence | None | 0 |
| **Insufficient Balance** | `InsufficientBalance` | Can't pay gas + value | None | 0 |
| | `TransferFailed` | Recipient rejection | Rollback | Partial |
| **Out of Gas** | `OutOfGas` | Gas limit exceeded | Rollback | All (limit) |
| | `IntrinsicGasTooLow` | Below minimum | None | 0 |
| **Overflow/Underflow** | `BalanceOverflow` | Credit exceeds u128 | Rollback | Partial |
| | `BalanceUnderflow` | Debit exceeds balance | Rollback | Partial |
| | `NonceOverflow` | Nonce exceeds u64 | None | 0 |
| **Unauthorized** | `InvalidSignature` | Bad signature | None | 0 |
| | `WrongChainId` | Cross-chain replay | None | 0 |
| **Execution** | `Revert` | REVERT opcode | Rollback | Partial |
| | `InvalidOpcode` | Unknown instruction | Rollback | All |
| | `StackOverflow` | Stack > 1024 | Rollback | All |

---

## 6. Commit Phase

### Commit Flow

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     COMMIT PHASE                                            │
└─────────────────────────────────────────────────────────────────────────────┘

  ┌─────────────────────────────────────────────────────────────────────────┐
  │                         EXECUTION COMPLETE                              │
  │                                                                         │
  │  Computed:                                                              │
  │  • state_root_computed                                                  │
  │  • receipts_root_computed                                               │
  │  • logs_bloom_computed                                                  │
  │  • gas_used_computed                                                    │
  └────────────────────────────────┬────────────────────────────────────────┘
                                   │
                                   ▼
  ┌─────────────────────────────────────────────────────────────────────────┐
  │                         VALIDATION                                      │
  │                                                                         │
  │  Compare against block header:                                          │
  │                                                                         │
  │  ┌─────────────────────┐    ┌─────────────────────┐                    │
  │  │ state_root_computed │ == │ header.state_root   │ ?                  │
  │  └─────────────────────┘    └─────────────────────┘                    │
  │                                                                         │
  │  ┌─────────────────────┐    ┌─────────────────────┐                    │
  │  │receipts_root_computed│ == │header.receipts_root │ ?                  │
  │  └─────────────────────┘    └─────────────────────┘                    │
  │                                                                         │
  │  ┌─────────────────────┐    ┌─────────────────────┐                    │
  │  │  gas_used_computed  │ == │  header.gas_used    │ ?                  │
  │  └─────────────────────┘    └─────────────────────┘                    │
  │                                                                         │
  │  If ANY mismatch: REJECT BLOCK (consensus failure)                     │
  └────────────────────────────────┬────────────────────────────────────────┘
                                   │
                                   │ All match
                                   ▼
  ┌─────────────────────────────────────────────────────────────────────────┐
  │                         COMMIT                                          │
  │                                                                         │
  │  1. Persist state trie changes to database                             │
  │  2. Persist block header                                               │
  │  3. Persist block body (transactions)                                  │
  │  4. Persist receipts                                                   │
  │  5. Update chain head pointer                                          │
  │  6. Update height index                                                │
  │                                                                         │
  │  All writes MUST be atomic (single batch)                              │
  └────────────────────────────────┬────────────────────────────────────────┘
                                   │
                                   ▼
  ┌─────────────────────────────────────────────────────────────────────────┐
  │                         BROADCAST                                       │
  │                                                                         │
  │  Announce commit to network peers                                       │
  └─────────────────────────────────────────────────────────────────────────┘
```

### Commit Logic Implementation

```rust
/// Final state root comparison and commit
pub fn validate_and_commit(
    state: &State,
    execution_output: &ExecutionOutput,
    block: &Block,
    storage: &mut Storage,
) -> Result<CommitResult, CommitError> {
    // ═══════════════════════════════════════════════════════════════════════
    // VALIDATION: Compare computed values against header
    // ═══════════════════════════════════════════════════════════════════════
    
    // State root comparison
    if execution_output.state_root != block.header.state_root {
        return Err(CommitError::StateRootMismatch {
            computed: execution_output.state_root,
            expected: block.header.state_root,
        });
    }
    
    // Receipts root comparison
    if execution_output.receipts_root != block.header.receipts_root {
        return Err(CommitError::ReceiptsRootMismatch {
            computed: execution_output.receipts_root,
            expected: block.header.receipts_root,
        });
    }
    
    // Gas used comparison
    if execution_output.gas_used != block.header.gas_used {
        return Err(CommitError::GasUsedMismatch {
            computed: execution_output.gas_used,
            expected: block.header.gas_used,
        });
    }
    
    // Logs bloom comparison (optional, can be recomputed)
    if execution_output.logs_bloom != block.header.logs_bloom {
        return Err(CommitError::LogsBloomMismatch {
            computed: execution_output.logs_bloom,
            expected: block.header.logs_bloom,
        });
    }
    
    // ═══════════════════════════════════════════════════════════════════════
    // COMMIT: Persist all data atomically
    // ═══════════════════════════════════════════════════════════════════════
    
    let mut batch = storage.begin_batch();
    
    // 1. Persist state trie nodes
    for (hash, node) in state.dirty_nodes() {
        batch.put(&trie_node_key(&hash), &node.encode());
    }
    
    // 2. Persist block header
    batch.put(
        &block_header_key(&block.header.hash()),
        &block.header.encode(),
    );
    
    // 3. Persist block body
    batch.put(
        &block_body_key(&block.header.hash()),
        &block.body.encode(),
    );
    
    // 4. Persist receipts
    for (i, receipt) in execution_output.receipts.iter().enumerate() {
        batch.put(
            &receipt_key(&block.header.hash(), i as u32),
            &receipt.encode(),
        );
    }
    
    // 5. Update chain head
    batch.put(HEAD_KEY, &block.header.hash());
    
    // 6. Update height index
    batch.put(
        &height_index_key(block.header.height),
        &block.header.hash(),
    );
    
    // Atomic commit
    batch.commit()?;
    
    Ok(CommitResult {
        block_hash: block.header.hash(),
        block_height: block.header.height,
        state_root: execution_output.state_root,
        tx_count: block.body.transactions.len(),
        gas_used: execution_output.gas_used,
    })
}
```

---

## 7. Security Invariants

### Determinism Checks

```rust
/// Determinism invariants that MUST hold
pub mod determinism_invariants {
    use super::*;
    
    /// Same inputs produce same outputs
    pub fn verify_determinism(
        state: &State,
        block: &Block,
    ) -> bool {
        // Execute twice
        let result1 = execute_block(state.clone(), block);
        let result2 = execute_block(state.clone(), block);
        
        // Must be identical
        result1.state_root == result2.state_root
            && result1.receipts_root == result2.receipts_root
            && result1.gas_used == result2.gas_used
    }
    
    /// No hidden state between transactions
    pub fn verify_no_hidden_state(
        state: &State,
        txs: &[Transaction],
    ) -> bool {
        // Execute all together
        let combined = execute_transactions(state.clone(), txs);
        
        // Execute one by one
        let mut sequential_state = state.clone();
        for tx in txs {
            execute_transaction(&mut sequential_state, tx);
        }
        
        // Must match
        combined.state_root == sequential_state.root()
    }
}
```

### State Isolation

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     STATE ISOLATION INVARIANTS                              │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  1. NO CROSS-TRANSACTION HIDDEN STATE                                       │
│     ────────────────────────────────                                        │
│     • Each transaction sees committed state from previous                  │
│     • No thread-local or global mutable state                              │
│     • No caching that affects behavior                                     │
│                                                                             │
│  2. ATOMIC TRANSACTION EXECUTION                                            │
│     ────────────────────────────────                                        │
│     • All changes committed or all rolled back                             │
│     • No partial state visible to other transactions                       │
│     • Checkpoint/rollback mechanism enforced                               │
│                                                                             │
│  3. ISOLATED STORAGE NAMESPACES                                             │
│     ──────────────────────────────                                          │
│     • Each account has separate storage trie                               │
│     • No direct access to other account's storage                          │
│     • Cross-account access only via CALL                                   │
│                                                                             │
│  4. BOUNDED EFFECTS                                                         │
│     ────────────────                                                        │
│     • Transaction can only modify:                                         │
│       - Sender account (nonce, balance)                                    │
│       - Recipient account (balance, storage, code)                         │
│       - Called contracts (storage)                                         │
│       - Block proposer (balance)                                           │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Gas Safety

```rust
/// Gas safety invariants
pub mod gas_safety {
    use super::*;
    
    /// Gas is always consumed or refunded (never lost)
    pub fn verify_gas_accounting(
        tx: &Transaction,
        receipt: &TransactionReceipt,
        sender_delta: i128,
        proposer_delta: i128,
    ) -> bool {
        let gas_cost = tx.gas_limit as u128 * tx.gas_price;
        let gas_used_cost = receipt.gas_used as u128 * tx.gas_price;
        let refund = gas_cost - gas_used_cost;
        
        // Sender pays gas_used, gets refund
        // Proposer receives gas_used
        (-sender_delta as u128) == gas_used_cost
            && proposer_delta as u128 == gas_used_cost
    }
    
    /// Gas limit bounds execution
    pub fn verify_gas_bounded(
        tx: &Transaction,
        receipt: &TransactionReceipt,
    ) -> bool {
        receipt.gas_used <= tx.gas_limit
    }
    
    /// Out of gas causes full consumption
    pub fn verify_oog_full_consumption(
        tx: &Transaction,
        receipt: &TransactionReceipt,
    ) -> bool {
        if receipt.status == ExecutionStatus::OutOfGas {
            receipt.gas_used == tx.gas_limit
        } else {
            true
        }
    }
}
```

### Replay Protection

```rust
/// Replay protection invariants
pub mod replay_protection {
    use super::*;
    
    /// Transaction can only execute once per nonce
    pub fn verify_nonce_uniqueness(
        txs: &[Transaction],
    ) -> bool {
        let mut seen: HashMap<(Address, u64), Hash> = HashMap::new();
        
        for tx in txs {
            let sender = recover_sender(tx).unwrap();
            let key = (sender, tx.nonce);
            
            if let Some(prev_hash) = seen.get(&key) {
                if *prev_hash != tx.hash() {
                    return false;  // Different tx with same nonce
                }
            }
            
            seen.insert(key, tx.hash());
        }
        
        true
    }
    
    /// Chain ID prevents cross-chain replay
    pub fn verify_chain_id(
        tx: &Transaction,
        expected_chain_id: u64,
    ) -> bool {
        tx.chain_id == Some(expected_chain_id)
    }
}
```

### No Partial State Invariant

```rust
/// No partial state invariant
pub fn verify_no_partial_state(
    pre_state: &State,
    post_state: &State,
    receipts: &[TransactionReceipt],
) -> bool {
    // For each transaction, either:
    // 1. All changes are applied (Success)
    // 2. Only nonce+gas changes are applied (Revert/OOG)
    
    for receipt in receipts {
        match receipt.status {
            ExecutionStatus::Success => {
                // All effects should be present
                // (verified by state root match)
            }
            ExecutionStatus::Revert | ExecutionStatus::OutOfGas => {
                // Only sender nonce and gas effects
                // Value transfer should NOT be present
            }
            _ => {
                // Other statuses should not produce partial state
            }
        }
    }
    
    true
}
```

---

## 8. GPU Offload Opportunities

### Parallelizable Workloads

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     GPU OFFLOAD OPPORTUNITIES                               │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  PARALLELIZABLE (GPU-suitable):                                             │
│  ═══════════════════════════════                                            │
│                                                                             │
│  1. Signature Verification (batch)                                          │
│     • ECDSA recovery is independent per transaction                        │
│     • Batch all signatures for parallel verification                       │
│     • GPU: ~100x speedup for large batches                                 │
│                                                                             │
│  2. Merkle Tree Computation                                                 │
│     • Hash tree levels can be parallelized                                 │
│     • Leaf hashing is independent                                          │
│     • GPU: ~50x speedup for large trees                                    │
│                                                                             │
│  3. State Root Computation                                                  │
│     • Trie hashing at each level                                           │
│     • Independent branch computation                                       │
│     • GPU: ~30x speedup for large state                                    │
│                                                                             │
│  4. Bloom Filter Computation                                                │
│     • Log bloom is independent per log                                     │
│     • OR operation is parallelizable                                       │
│     • GPU: ~20x speedup                                                    │
│                                                                             │
│  SEQUENTIAL (not GPU-suitable):                                             │
│  ═══════════════════════════════                                            │
│                                                                             │
│  • Transaction execution (state dependencies)                              │
│  • Nonce validation (sequential by sender)                                 │
│  • Balance tracking (order-dependent)                                      │
│  • Storage writes (conflict resolution)                                    │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### PoUW-Suitable Operations

```rust
/// Operations suitable for PoUW compute proofs
pub mod pouw_suitable {
    /// Cryptographic operations
    pub mod crypto {
        // SNARK proof generation
        // STARK proof generation
        // BLS signature aggregation
        // Hash preimage search (constrained)
    }
    
    /// Data operations
    pub mod data {
        // Large dataset merkle proofs
        // Data availability sampling
        // Erasure coding
    }
    
    /// AI/ML operations (future)
    pub mod ml {
        // Model inference
        // Gradient computation
        // Matrix multiplication
    }
}
```

### Future VM Acceleration

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     FUTURE VM ACCELERATION                                  │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  WASM VM with GPU:                                                          │
│  ─────────────────                                                          │
│  • WebGPU bindings for compute shaders                                     │
│  • Offload heavy computations to GPU                                       │
│  • Maintain determinism via fixed-function GPU ops                         │
│                                                                             │
│  RISC-V VM with GPU:                                                        │
│  ───────────────────                                                        │
│  • RISC-V vector extensions                                                │
│  • GPU coprocessor instructions                                            │
│  • ZK-friendly for proof generation                                        │
│                                                                             │
│  Acceleration Points:                                                       │
│  ────────────────────                                                       │
│  • Precompiled contracts (ECADD, ECMUL, etc.)                              │
│  • Hash functions (SHA256, Keccak, Blake3)                                 │
│  • Signature operations (ECDSA, BLS)                                       │
│  • Merkle proof verification                                               │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 9. Future Extensions

### ZK-Execution Proofs

*Status: Research*

```rust
/// ZK execution proof structure (future)
pub struct ZkExecutionProof {
    /// Proof bytes (SNARK/STARK)
    pub proof: Vec<u8>,
    
    /// Public inputs
    pub public_inputs: ZkPublicInputs,
    
    /// Proof system identifier
    pub system: ProofSystem,
}

pub struct ZkPublicInputs {
    /// Pre-state root
    pub pre_state_root: Hash,
    
    /// Post-state root
    pub post_state_root: Hash,
    
    /// Transactions hash
    pub transactions_hash: Hash,
    
    /// Block context hash
    pub block_context_hash: Hash,
}

pub enum ProofSystem {
    Groth16,
    Plonk,
    Stark,
    Halo2,
}
```

### Parallel VM Execution

*Status: Planned*

```rust
/// Parallel execution configuration
pub struct ParallelExecutionConfig {
    /// Number of worker threads
    pub workers: usize,
    
    /// Enable speculative execution
    pub speculative: bool,
    
    /// Conflict detection strategy
    pub conflict_strategy: ConflictStrategy,
}

pub enum ConflictStrategy {
    /// Pessimistic: lock accounts before execution
    Pessimistic,
    
    /// Optimistic: detect and retry on conflict
    Optimistic,
    
    /// Hybrid: predict conflicts, lock likely conflicts
    Hybrid,
}

/// Parallel execution engine (future)
pub async fn execute_parallel(
    state: &State,
    transactions: &[Transaction],
    config: &ParallelExecutionConfig,
) -> ExecutionOutput {
    // 1. Analyze transaction dependencies
    let deps = analyze_dependencies(transactions);
    
    // 2. Group independent transactions
    let groups = deps.independent_groups();
    
    // 3. Execute groups in parallel
    let results = futures::future::join_all(
        groups.iter().map(|group| {
            execute_group(state.clone(), group)
        })
    ).await;
    
    // 4. Merge results
    merge_results(results)
}
```

### Deterministic WASM Sandbox

*Status: Planned*

```rust
/// WASM execution sandbox (future)
pub struct WasmSandbox {
    /// WASM runtime (wasmer/wasmtime)
    runtime: WasmRuntime,
    
    /// Gas metering
    gas_meter: GasMeter,
    
    /// Memory limits
    memory_limit: usize,
    
    /// Stack limits
    stack_limit: usize,
}

impl WasmSandbox {
    /// Execute WASM contract
    pub fn execute(
        &mut self,
        code: &[u8],
        input: &[u8],
        state: &mut State,
    ) -> Result<ExecutionResult, WasmError> {
        // 1. Validate WASM module
        let module = self.runtime.validate(code)?;
        
        // 2. Instantiate with metering
        let instance = self.runtime.instantiate(
            &module,
            self.gas_meter.clone(),
        )?;
        
        // 3. Execute with sandbox
        let output = instance.call("main", input)?;
        
        // 4. Return result
        Ok(ExecutionResult {
            output,
            gas_used: self.gas_meter.gas_used(),
            logs: instance.logs(),
        })
    }
}
```

### Smart Contract Engine Roadmap

| Phase | Feature | Timeline |
|-------|---------|----------|
| 1 | Native transfers | Current |
| 2 | WASM VM integration | Q2 2026 |
| 3 | Standard library | Q3 2026 |
| 4 | Cross-contract calls | Q4 2026 |
| 5 | Upgradeable contracts | Q1 2027 |
| 6 | ZK-compatible VM | Q2 2027 |

---

## Summary

The Mbongo Chain state machine provides a deterministic, secure foundation for consensus-critical execution. All nodes must produce identical state transitions given identical inputs, ensuring network-wide agreement on the canonical state.

For block validation details, see [Block Validation Pipeline](block_validation_pipeline.md).

For consensus details, see [Consensus Validation](consensus_validation.md).

---

**Mbongo Chain** — Compute-first blockchain infrastructure for the global future.

