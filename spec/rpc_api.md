# Mbongo Chain — RPC API Specification  
Status: Canonical  
Version: v1.0

The RPC API allows applications, wallets, compute providers, explorers, and indexers to interact with the Mbongo Chain network.

All interfaces follow:

- **JSON-RPC 2.0**
- **WebSocket**
- **gRPC**

This specification defines core RPC endpoints for:

- blockchain queries  
- accounts & balances  
- transactions  
- PoS staking  
- PoUW compute system  
- AIDA economic regulator  
- governance  
- network introspection  

---

# 1. Conventions

Base URL examples:

http://localhost:8545 (JSON-RPC)
ws://localhost:8546 (WebSocket)
grpc://localhost:50051 (gRPC)

css
Copier le code

JSON structure:

```json
{
  "jsonrpc": "2.0",
  "method": "method_name",
  "params": [...],
  "id": 1
}
2. Blockchain Endpoints
2.1 mbongo_latestBlock
Returns the latest block.

Request
json
Copier le code
{"method":"mbongo_latestBlock","params":[],"id":1}
Response
json
Copier le code
{
  "number": 102394,
  "hash": "0xabc...",
  "parent": "0x123...",
  "state_root": "0xdead...",
  "timestamp": 1712345678
}
2.2 mbongo_getBlockByNumber
json
Copier le code
{"method":"mbongo_getBlockByNumber","params":[102394],"id":1}
2.3 mbongo_getBlockByHash
json
Copier le code
{"method":"mbongo_getBlockByHash","params":["0xhash"],"id":1}
2.4 mbongo_getTransaction
Returns transaction details.

Example:
json
Copier le code
{"method":"mbongo_getTransaction","params":["0xtrxid"],"id":1}
2.5 mbongo_getReceipt
Returns execution receipt for a transaction.

3. Account Endpoints
3.1 mbongo_getBalance
json
Copier le code
{"method": "mbongo_getBalance", "params": ["MBO1xyz..."], "id": 1}
3.2 mbongo_getNonce
json
Copier le code
{"method":"mbongo_getNonce","params":["MBO1xyz..."],"id":1}
4. Transaction Submission
4.1 mbongo_sendTransaction
Submit a signed transaction.

json
Copier le code
{
  "method": "mbongo_sendTransaction",
  "params": [{
    "from": "MBO1xyz...",
    "to": "MBO1abc...",
    "amount": 100,
    "gas_limit": 50000,
    "gas_price": 1,
    "nonce": 42,
    "signature": "0xsig..."
  }],
  "id": 1
}
4.2 mbongo_estimateGas
Estimate execution gas.

json
Copier le code
{
  "method": "mbongo_estimateGas",
  "params": ["0xpayload..."],
  "id": 1
}
5. PoS Staking Endpoints
5.1 mbongo_getValidatorSet
Returns current validators.

5.2 mbongo_delegateStake
Delegates stake:

json
Copier le code
{
  "method": "mbongo_delegateStake",
  "params": ["MBO1delegator", "MBO1validator", 5000],
  "id": 1
}
5.3 mbongo_undelegateStake
5.4 mbongo_getStakingInfo
Returns validator performance, uptime, rewards.

6. PoUW (Proof-of-Useful-Work) Endpoints
These endpoints interact with the compute marketplace & verification system.

6.1 mbongo_submitComputeJob
Submit a job to PoUW.

Example
json
Copier le code
{
  "method": "mbongo_submitComputeJob",
  "params": [{
    "job_type": "ai_inference",
    "payload": "base64-model-input",
    "cgas_limit": 200000,
    "max_fee": 3000
  }],
  "id": 1
}
6.2 mbongo_getJobStatus
6.3 mbongo_listJobs
Active jobs in the marketplace.

6.4 mbongo_getComputeReceipt
Returns VWP (Validatable Work Proof).

6.5 mbongo_submitFraudProof
Challengers submit fraud proofs:

json
Copier le code
{
  "method":"mbongo_submitFraudProof",
  "params":["job_id","proof_data"],
  "id":1
}
6.6 mbongo_listComputeNodes
Returns active compute nodes + hardware capabilities.

6.7 mbongo_getPoUWMultiplier
(From AIDA Regulator)

7. AIDA Economic Regulator Endpoints
These are required for transparency & client tooling.

7.1 mbongo_getAIDAState
Returns global economic parameters.

Example Response
json
Copier le code
{
  "burn_rate": 0.12,
  "base_fee_multiplier": 1.4,
  "pouw_multiplier": 1.05,
  "risk_level": "low",
  "last_update_block": 102340
}
7.2 mbongo_getBurnRate
7.3 mbongo_getFeeMultipliers
Returns:

base fee multiplier

priority fee limits

PoUW multiplier

7.4 mbongo_getEconomicForecast
AIDA advisory predictions (off-chain → bounded on-chain signals).

Example:

json
Copier le code
{
  "method":"mbongo_getEconomicForecast",
  "params":[],
  "id":1
}
7.5 mbongo_getAIDAWarnings
Returns warnings such as:

mempool congestion

abnormal demand

compute imbalance

pricing anomalies

upcoming parameter adjustments

8. Governance Endpoints
8.1 mbongo_submitProposal
Submit DAO proposal.

8.2 mbongo_vote
8.3 mbongo_getProposal
Returns:

state

snapshot

votes

timeline

AIDA risk report

8.4 mbongo_getFounderCouncilState
Transparency endpoint:

json
Copier le code
{
  "active": true,
  "years_remaining": 9,
  "vetoes_used": 0
}
9. Network & Node Introspection
9.1 mbongo_peerCount
9.2 mbongo_syncState
9.3 mbongo_nodeInfo
Returns:

node version

chain ID

supported modules

AIDA version

PoUW capabilities

10. WebSocket Event Subscriptions
Available Streams
Event	Description
newBlock	On block creation
newTransaction	On tx add
computeReceipt	On PoUW completion
fraudProof	On fraud detection
aidaUpdate	On AIDA economic changes

Example:

json
Copier le code
{
  "method":"mbongo_subscribe",
  "params":["newBlock"],
  "id":1
}
11. gRPC Endpoints
gRPC interface mirrors JSON-RPC but supports:

faster streaming

compute job streaming

PoUW receipt streaming

telemetry / metrics streaming

12. Summary
This RPC API provides:

full blockchain interaction

compute marketplace control

AIDA economic transparency

staking & governance integration

developer-friendly interfaces

WebSocket & gRPC streaming support

compute-native primitives not found in traditional L1s

This document represents the canonical API surface developers must rely on when building on Mbongo Chain.