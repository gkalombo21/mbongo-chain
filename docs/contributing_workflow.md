# Mbongo Chain — Contributing Workflow

Official contribution workflow for all Mbongo Chain contributors.

---

## 1. Overview

### Contribution Principles

All contributions to Mbongo Chain follow a structured workflow designed for:

| Principle | Description |
|-----------|-------------|
| **Quality** | Every change passes automated checks |
| **Traceability** | Clean commit history for auditing |
| **Collaboration** | Peer review before merge |
| **Consistency** | Standardized process for all contributors |

### Workflow Summary

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     CONTRIBUTION WORKFLOW                                   │
└─────────────────────────────────────────────────────────────────────────────┘

   Branch ──▶ Develop ──▶ Validate ──▶ Sync ──▶ PR ──▶ Review ──▶ Merge

   ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐
   │ Create   │  │ Write    │  │ Format   │  │ Sync     │  │ Create   │
   │ Branch   │──│ Code     │──│ & Lint   │──│ & Push   │──│ PR       │
   └──────────┘  └──────────┘  └──────────┘  └──────────┘  └──────────┘
                                                                │
   ┌──────────┐  ┌──────────┐  ┌──────────┐                    │
   │ Merge    │◀─│ Approval │◀─│ Review   │◀───────────────────┘
   └──────────┘  └──────────┘  └──────────┘
```

### Core Requirements

- **Branching**: Always branch from `main`
- **Commits**: Follow Conventional Commits format
- **Reviews**: All PRs require peer review
- **Sync**: Keep branch up-to-date with `main`
- **Checks**: All CI checks must pass

---

## 2. Branching Model

### Branch Types

| Branch | Pattern | Purpose | Merge Target |
|--------|---------|---------|--------------|
| `main` | `main` | Production-ready, protected | — |
| `feature/*` | `feature/<name>` | New functionality | `main` |
| `fix/*` | `fix/<name>` | Bug fixes | `main` |
| `docs/*` | `docs/<name>` | Documentation updates | `main` |
| `chore/*` | `chore/<name>` | Maintenance, tooling | `main` |

### Branch Diagram

```
main ─────────────────────────────────────────────────────────────────▶
       │                              │                        ▲
       │ branch                       │ branch                 │ merge
       ▼                              ▼                        │
feature/mempool-priority ────●────●────●─────────────────────▶│
                                                               │
fix/header-validation ────●────●──────────────────────────────▶│
                                                               │
docs/consensus-overview ────●─────────────────────────────────▶│
```

### When to Create Each Branch Type

| Type | Create When... | Examples |
|------|----------------|----------|
| `feature/*` | Adding new functionality | `feature/mempool-eviction`, `feature/sync-pipeline` |
| `fix/*` | Correcting bugs or errors | `fix/block-validation`, `fix/gas-overflow` |
| `docs/*` | Updating documentation only | `docs/runtime-guide`, `docs/api-reference` |
| `chore/*` | Maintenance without logic changes | `chore/update-deps`, `chore/ci-config` |

### Branch Naming Rules

```
✓ feature/add-tx-priority       # Lowercase, hyphenated
✓ fix/runtime-gas-calculation   # Descriptive, includes module
✓ docs/update-architecture      # Clear purpose

✗ Feature/AddTxPriority         # No PascalCase
✗ fix_runtime_gas               # No underscores
✗ update                        # Too vague
```

---

## 3. Development Flow

### Complete Workflow (13 Steps)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     DEVELOPMENT FLOW                                        │
└─────────────────────────────────────────────────────────────────────────────┘

  STEP 1: CREATE BRANCH
  ──────────────────────
  git checkout main
  git pull origin main
  git checkout -b feature/my-feature

  STEP 2: WRITE CODE OR DOCS
  ──────────────────────────
  • Implement feature or fix
  • Update documentation
  • Add tests if applicable

  STEP 3: FORMAT USING RUSTFMT
  ────────────────────────────
  cargo fmt --all

  STEP 4: LINT USING CLIPPY
  ─────────────────────────
  cargo clippy --workspace --all-targets -- -D warnings

  STEP 5: RUN MODULE-SPECIFIC VALIDATION
  ──────────────────────────────────────
  cargo check -p <module>
  cargo test -p <module>

  STEP 6: RUN WORKSPACE BUILD
  ───────────────────────────
  cargo build --workspace

  STEP 7: SYNC WITH MAIN
  ───────────────────────
  git fetch origin
  git rebase origin/main

  STEP 8: RESOLVE CONFLICTS
  ─────────────────────────
  • Open conflicting files
  • Resolve manually or with editor
  • git add <resolved-files>
  • git rebase --continue

  STEP 9: CREATE PULL REQUEST
  ───────────────────────────
  git push -u origin feature/my-feature
  • Open GitHub → New Pull Request
  • Fill PR template

  STEP 10: ADD REVIEWERS
  ──────────────────────
  • Assign module owners
  • Request review from team

  STEP 11: APPLY REQUESTED CHANGES
  ────────────────────────────────
  • Address feedback
  • Push additional commits
  • Reply to comments

  STEP 12: SQUASH COMMITS (IF NEEDED)
  ───────────────────────────────────
  git rebase -i origin/main
  • Mark commits as 'squash' or 'fixup'
  • Force push: git push --force-with-lease

  STEP 13: MERGE
  ──────────────
  • Ensure all checks pass
  • Reviewer approves
  • Squash and merge (preferred)
```

### Quick Reference Commands

```powershell
# Step 1: Branch
git checkout main && git pull origin main
git checkout -b feature/my-feature

# Steps 3-6: Validate
cargo fmt --all
cargo clippy --workspace --all-targets -- -D warnings
cargo build --workspace
cargo test --workspace

# Step 7: Sync
git fetch origin
git rebase origin/main

# Step 9: Push
git push -u origin feature/my-feature

# Step 12: Squash (interactive)
git rebase -i origin/main
git push --force-with-lease
```

---

## 4. PR Requirements

### Mandatory Requirements

| Requirement | Command/Check | Status |
|-------------|---------------|--------|
| CI passes | GitHub Actions | Required |
| No Clippy warnings | `cargo clippy -- -D warnings` | Required |
| Rustfmt compliant | `cargo fmt -- --check` | Required |
| Tests pass | `cargo test --workspace` | Required |
| Build succeeds | `cargo build --workspace` | Required |
| Clean history | Squashed or logical commits | Required |
| Conventional title | `feat:`, `fix:`, `docs:`, etc. | Required |

### PR Title Format

```
<type>(<scope>): <description>

Examples:
feat(runtime): add gas metering for state transitions
fix(network): resolve peer discovery timeout
docs(consensus): update PoUW validation guide
refactor(node): simplify block processing pipeline
chore(ci): add coverage reporting
```

### PR Checklist

Every PR must include this checklist in the description:

```markdown
## Checklist

- [ ] Code compiles without errors
- [ ] Clippy passes with no warnings
- [ ] Rustfmt applied to all files
- [ ] Tests pass (if applicable)
- [ ] Documentation updated (if applicable)
- [ ] No TODO comments left in code
- [ ] Commit messages follow Conventional Commits
- [ ] Branch is synced with main
```

### Documentation Requirements

| Change Type | Documentation Required |
|-------------|----------------------|
| New feature | API docs + guide update |
| Bug fix | Changelog entry |
| Breaking change | Migration guide |
| New module | Architecture doc update |

---

## 5. Review Guidelines

### For Reviewers

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     REVIEW CHECKLIST                                        │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  CORRECTNESS                                                                │
│  • Does the code do what it claims?                                        │
│  • Are edge cases handled?                                                 │
│  • Are error conditions handled properly?                                  │
│                                                                             │
│  STYLE                                                                      │
│  • Does it follow Rust idioms?                                             │
│  • Is it readable and maintainable?                                        │
│  • Are variable names descriptive?                                         │
│                                                                             │
│  ARCHITECTURE                                                               │
│  • Does it fit the existing design?                                        │
│  • Are module boundaries respected?                                        │
│  • Is there unnecessary coupling?                                          │
│                                                                             │
│  SECURITY                                                                   │
│  • Any potential vulnerabilities?                                          │
│  • Input validation present?                                               │
│  • No panics in production paths?                                          │
│                                                                             │
│  DOCUMENTATION                                                              │
│  • Are public APIs documented?                                             │
│  • Are complex sections commented?                                         │
│  • Is the PR description clear?                                            │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### For Authors

**Responding to Reviews:**

| Feedback Type | How to Respond |
|---------------|----------------|
| **Question** | Answer clearly, update code if needed |
| **Suggestion** | Implement or explain why not |
| **Required change** | Must address before merge |
| **Nitpick** | Address or acknowledge |

**Request Clarification:**

```markdown
> Reviewer: "This might cause issues with concurrent access"

Author: "Could you clarify the scenario you're concerned about?
        The current design assumes single-threaded access because..."
```

**Handle Requested Changes:**

1. Read feedback carefully
2. Ask for clarification if unclear
3. Implement changes
4. Push new commits
5. Reply to each comment
6. Re-request review

### Approval Criteria

| Status | Meaning | Action |
|--------|---------|--------|
| **Approved** | Ready to merge | Merge when CI passes |
| **Changes requested** | Issues must be fixed | Address feedback |
| **Commented** | Questions/suggestions | Respond and discuss |

### Rejection Reasons

- Fails CI checks
- Unresolved conflicts
- Missing documentation
- Security concerns
- Architectural issues
- Scope too large

---

## 6. Conflict Resolution

### Step-by-Step Commands

```powershell
# 1. Fetch latest changes
git fetch origin

# 2. Start rebase
git rebase origin/main

# 3. If conflicts occur, Git will pause
# Output: CONFLICT (content): Merge conflict in runtime/src/lib.rs

# 4. Check which files have conflicts
git status
# Shows: both modified: runtime/src/lib.rs

# 5. Open file and find conflict markers
# <<<<<<< HEAD
# your changes
# =======
# incoming changes
# >>>>>>> origin/main

# 6. Edit file to resolve (remove markers, keep correct code)

# 7. Stage resolved file
git add runtime/src/lib.rs

# 8. Continue rebase
git rebase --continue

# 9. If more conflicts, repeat steps 5-8

# 10. Push updated branch
git push --force-with-lease
```

### Conflict Resolution in Editors

**Cursor:**
1. Open conflicting file
2. Click "Accept Current", "Accept Incoming", or "Accept Both"
3. Save file
4. Stage and continue rebase

**VS Code:**
1. Conflict highlighting appears automatically
2. Click inline resolution buttons
3. Save and stage

### Best Practices to Avoid Conflicts

| Practice | Benefit |
|----------|---------|
| Sync frequently | Fewer accumulated changes |
| Small PRs | Less overlap with others |
| Coordinate on shared files | Avoid simultaneous edits |
| Rebase before push | Catch conflicts early |

### Pre-Merge Sync

Always sync before merging:

```powershell
# Ensure branch is up-to-date
git fetch origin
git rebase origin/main

# Verify no conflicts
git status

# Run validation
cargo fmt --all
cargo clippy --workspace --all-targets -- -D warnings
cargo build --workspace

# Push if needed
git push --force-with-lease
```

---

## 7. Documentation Rules

### Core Requirements

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     DOCUMENTATION REQUIREMENTS                              │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  EVERY CHANGE MUST BE DOCUMENTED                                            │
│  • Code changes → Update relevant docs                                     │
│  • New features → Add to architecture docs                                 │
│  • API changes → Update API documentation                                  │
│  • Bug fixes → Add to changelog                                            │
│                                                                             │
│  EVERY PR MUST UPDATE RELATED DOCS                                          │
│  • Modify runtime/ → Update runtime_architecture.md                        │
│  • Modify network/ → Update networking_overview.md                         │
│  • Modify consensus → Update consensus_overview.md                         │
│  • Add new module → Update architecture_overview.md                        │
│                                                                             │
│  ALL DIAGRAMS MUST BE ASCII-COMPATIBLE                                      │
│  • Use box-drawing characters                                              │
│  • Render correctly in GitHub                                              │
│  • Include in markdown code blocks                                         │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Documentation Mapping

| Code Change | Documentation to Update |
|-------------|------------------------|
| `runtime/` | `runtime_architecture.md`, `state_machine_validation.md` |
| `node/` | `node_architecture.md`, `block_validation_pipeline.md` |
| `network/` | `networking_overview.md`, `sync_validation.md` |
| `pow/` | `consensus_overview.md`, `consensus_validation.md` |
| `crypto/` | `architecture_overview.md` |
| `cli/` | `getting_started.md`, `developer_introduction.md` |
| `mempool/` | `mempool_overview.md` |

### ASCII Diagram Standards

```
GOOD:
┌─────────────────┐
│   Component A   │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│   Component B   │
└─────────────────┘

BAD (image-based, not renderable):
[Component A] → [Component B]
```

---

## 8. Cursor Integration

### Safe Local Development

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     CURSOR BEST PRACTICES                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ✓ DO                                                                       │
│  ────                                                                       │
│  • Use module-scoped prompts: "In runtime/src/lib.rs, fix..."              │
│  • Run validation after AI changes                                         │
│  • Review all generated code before committing                             │
│  • Use agents for specific modules only                                    │
│                                                                             │
│  ✗ DON'T                                                                    │
│  ──────                                                                     │
│  • Request workspace-wide refactors                                        │
│  • Accept changes without review                                           │
│  • Let AI modify unrelated files                                           │
│  • Skip validation steps                                                   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Running Validation Scripts

```powershell
# After any AI-generated changes, always run:

# 1. Format check
cargo fmt --all -- --check

# 2. Lint check
cargo clippy --workspace --all-targets -- -D warnings

# 3. Build check
cargo build --workspace

# 4. Test check
cargo test --workspace
```

### Workspace Linting

```powershell
# Full workspace validation
cargo fmt --all
cargo clippy --workspace --all-targets -- -D warnings

# Module-specific validation
cargo clippy -p runtime -- -D warnings
cargo clippy -p node -- -D warnings
cargo clippy -p network -- -D warnings
```

### Preventing Accidental Changes

| Risk | Prevention |
|------|------------|
| Unintended file edits | Always specify file in prompt |
| Mass refactoring | Never request "fix everything" |
| Breaking changes | Run tests after each change |
| Uncommitted changes | Use `git status` frequently |

### Safe Prompts

```
✓ "In runtime/src/lib.rs line 42, fix the gas calculation"
✓ "Add documentation to the Transaction struct in crypto/src/lib.rs"
✓ "Implement validate_header() in node/src/lib.rs following the spec"

✗ "Refactor all modules to use the new pattern"
✗ "Fix all the bugs in the codebase"
✗ "Update everything to latest Rust style"
```

---

## 9. Final Checklist (Before PR)

### Pre-PR Validation

Run this checklist before opening any PR:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     PRE-PR CHECKLIST                                        │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ☐ RUSTFMT OK                                                               │
│    cargo fmt --all -- --check                                              │
│    └── No formatting issues                                                │
│                                                                             │
│  ☐ CLIPPY OK                                                                │
│    cargo clippy --workspace --all-targets -- -D warnings                   │
│    └── No warnings or errors                                               │
│                                                                             │
│  ☐ BUILD OK                                                                 │
│    cargo build --workspace                                                 │
│    └── Compiles without errors                                             │
│                                                                             │
│  ☐ TESTS OK                                                                 │
│    cargo test --workspace                                                  │
│    └── All tests pass                                                      │
│                                                                             │
│  ☐ DOCS OK                                                                  │
│    └── Related documentation updated                                       │
│    └── Public APIs documented                                              │
│                                                                             │
│  ☐ NO TODOs LEFT                                                            │
│    grep -r "TODO" --include="*.rs"                                         │
│    └── No unresolved TODOs in code                                         │
│                                                                             │
│  ☐ CLEAN COMMITS                                                            │
│    git log --oneline                                                       │
│    └── Logical, well-formatted commit messages                             │
│    └── Follows Conventional Commits                                        │
│                                                                             │
│  ☐ SYNCED WITH MAIN                                                         │
│    git fetch origin && git rebase origin/main                              │
│    └── Branch is up-to-date                                                │
│    └── No merge conflicts                                                  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Quick Validation Script

```powershell
# Save as scripts/pre-pr-check.ps1

Write-Host "=== Pre-PR Validation ===" -ForegroundColor Cyan

# Rustfmt
Write-Host "`n[1/5] Checking formatting..." -ForegroundColor Yellow
cargo fmt --all -- --check
if ($LASTEXITCODE -ne 0) { Write-Host "FAIL: Run 'cargo fmt --all'" -ForegroundColor Red; exit 1 }
Write-Host "OK" -ForegroundColor Green

# Clippy
Write-Host "`n[2/5] Running Clippy..." -ForegroundColor Yellow
cargo clippy --workspace --all-targets -- -D warnings
if ($LASTEXITCODE -ne 0) { Write-Host "FAIL: Fix Clippy warnings" -ForegroundColor Red; exit 1 }
Write-Host "OK" -ForegroundColor Green

# Build
Write-Host "`n[3/5] Building workspace..." -ForegroundColor Yellow
cargo build --workspace
if ($LASTEXITCODE -ne 0) { Write-Host "FAIL: Fix build errors" -ForegroundColor Red; exit 1 }
Write-Host "OK" -ForegroundColor Green

# Tests
Write-Host "`n[4/5] Running tests..." -ForegroundColor Yellow
cargo test --workspace
if ($LASTEXITCODE -ne 0) { Write-Host "FAIL: Fix failing tests" -ForegroundColor Red; exit 1 }
Write-Host "OK" -ForegroundColor Green

# Sync check
Write-Host "`n[5/5] Checking sync with main..." -ForegroundColor Yellow
git fetch origin
$behind = git rev-list --count HEAD..origin/main
if ($behind -gt 0) { Write-Host "WARNING: Branch is $behind commits behind main" -ForegroundColor Yellow }
else { Write-Host "OK" -ForegroundColor Green }

Write-Host "`n=== All checks passed! Ready for PR ===" -ForegroundColor Green
```

### Checklist Summary

| Check | Command | Required |
|-------|---------|----------|
| Rustfmt | `cargo fmt --all -- --check` | ✓ |
| Clippy | `cargo clippy --workspace -- -D warnings` | ✓ |
| Build | `cargo build --workspace` | ✓ |
| Tests | `cargo test --workspace` | ✓ |
| Docs | Manual review | ✓ |
| TODOs | `grep -r "TODO"` | ✓ |
| Commits | `git log --oneline` | ✓ |
| Sync | `git rebase origin/main` | ✓ |

---

**Mbongo Chain** — Compute-first blockchain infrastructure for the global future.

