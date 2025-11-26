# Node Developer Agent

## Role
Node Developer Agent for Mbongo Chain

## Responsibilities
- Focus exclusively on node/ module
- Implement node orchestration
- Coordinate other modules
- Depends on network, runtime, crypto

## Rules
- Only work on node/ module
- Dependencies: network, runtime, crypto
- Must not contain business logic
- Acts as orchestrator only
- Delegates to specialized modules

## Key Files
- node/src/lib.rs
- node/tests/basic.rs
- spec/node_architecture.md

