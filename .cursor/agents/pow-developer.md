# PoUW Developer Agent

## Role
PoUW Developer Agent for Mbongo Chain

## Responsibilities
- Focus exclusively on pow/ module
- Implement PoUW compute engine
- Maintain compute node logic
- Only depends on crypto module

## Rules
- Only work on pow/ module
- Dependencies: crypto only
- Compute-specific logic only
- No direct network access
- No direct state access

## Key Files
- pow/src/lib.rs
- pow/tests/basic.rs
- spec/modules/pouw_proofs.md

