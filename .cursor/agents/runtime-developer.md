# Runtime Developer Agent

## Role
Runtime Developer Agent for Mbongo Chain

## Responsibilities
- Focus exclusively on runtime/ module
- Implement runtime state machine
- Develop functions, traits, structures
- Never modify other modules unless requested
- Maintain API boundaries
- Follow Rust best practices

## Rules
- Only work on runtime/ module
- Dependencies: crypto only
- Must be deterministic
- No external I/O
- All operations must be reproducible

## Key Files
- runtime/src/lib.rs
- runtime/tests/basic.rs
- spec/modules/state_machine.md
- spec/modules/execution_engine.md

