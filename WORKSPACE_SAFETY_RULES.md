# Mbongo-Chain — Unified Workspace Safety Rules

# Purpose: Prevent file loss, avoid duplicates, keep a clean repo, and automate periodic backups.

---

rule: Workspace Snapshotter (Daily & Manual)
scope: repository
triggers:
  - manual: "@daily-snapshot"
  - schedule: "0 18 * * *"   # every day at 18:00 local time
  - on-branch-switch: true
actions:
  - name: ensure-backup-folders
    run: |
      # Create folders if missing (idempotent)
      mkdir -p backups/daily
      mkdir -p backups/_meta
  - name: prune-old-snapshots
    run: |
      # Keep the last 14 daily snapshots only
      ls -1t backups/daily | tail -n +15 | xargs -I {} sh -c 'rm -rf "backups/daily/{}"' || true
  - name: write-manifest
    run: |
      TS="$(date +%Y%m%d-%H%M%S)"
      echo "$TS" > backups/_meta/LAST_SNAPSHOT_TS.txt
      git status --porcelain=v1 > "backups/_meta/status-$TS.txt" || true
  - name: create-zip-snapshot
    run: |
      TS="$(cat backups/_meta/LAST_SNAPSHOT_TS.txt)"
      OUT="backups/daily/mbongo-chain-$TS.zip"
      # Exclude heavy/derived files and local env
      zip -r "$OUT" . \
        -x "node_modules/**" \
        -x ".git/**" \
        -x "backups/**" \
        -x ".cursor/**" \
        -x ".DS_Store" \
        -x "*.exe" \
        -x "*.log" \
        -x "*.tmp" \
        -x "*.cache/**" \
        -x "bin/**" \
        -x "build/**" \
        -x "dist/**" \
        -x "*.ipynb_checkpoints/**" \
        -x ".vscode/**" \
        -x ".idea/**"
      echo "✅ Snapshot created at $OUT"

safety:
  - never-overwrite-existing: true
  - never-delete-outside: ["backups/", "backups/daily/", "backups/_meta/"]

notes:
  - "Manual trigger: type @daily-snapshot in Cursor to run at any time."
  - "Restores are manual: unzip a snapshot into a temp directory and copy files back intentionally."

---

rule: Safety Net (No Duplicates / No Clobber)
scope: repository
triggers:
  - on-file-create: true
  - on-file-move: true
  - on-rename: true
guards:
  - name: prevent-duplicate-configs
    check:
      # If a file already exists, prefer to patch/append instead of creating duplicates.
      when-creating:
        if-exists:
          then: "abort-and-suggest-merge"
          message: |
            A file with the same path already exists. Do not create a duplicate.
            Propose a patch to update the existing file instead (minimal diff).
  - name: idempotent-writes
    check:
      when-writing:
        normalize-line-endings: true
        preserve-encoding: true
  - name: protected-paths
    check:
      deny-write:
        - ".git/**"
        - "backups/**"
        - ".cursor/**"
        - "**/*.exe"
        - "**/*.zip"
  - name: stable-structure
    check:
      # Do not move top-level directories unless explicitly requested by the user.
      deny-move:
        - "cmd/**"
        - "internal/**"
        - "config/**"
        - "docs/**"
        - "README.md"
        - "CONTRIBUTING.md"
        - "SECURITY.md"
        - ".github/**"
  - name: go-module-integrity
    check:
      # If go.mod or go.sum is modified, run 'go mod tidy' afterwards (only if Go is installed)
      post-action:
        run: |
          if command -v go >/dev/null 2>&1; then
            go mod tidy || true
          fi

notes:
  - "If a target file exists, the agent must propose a minimal patch instead of overwriting."
  - "If the user asked to replace entirely, confirm explicitly before proceeding."

---

rule: Periodic Guard (Auto Snapshot Reminder)
scope: repository
triggers:
  - schedule: "0 */4 * * *"   # every 4 hours
  - on-large-diff: true        # >100 changed lines since last commit
actions:
  - name: check-last-commit
    run: |
      # Get last commit timestamp if repo exists
      if [ -d ".git" ]; then
        LAST=$(git log -1 --format=%ct 2>/dev/null || echo "0")
      else
        LAST="0"
      fi
      NOW=$(date +%s)
      DIFF=$(( NOW - LAST ))
      echo "$DIFF" > backups/_meta/LAST_COMMIT_AGE.txt
  - name: diff-size-check
    run: |
      mkdir -p backups/_meta
      # Count changed lines (best-effort)
      if [ -d ".git" ]; then
        git diff --stat | awk '{s+=$1} END {print (s==""?0:s)}' > backups/_meta/CHANGED_LINES.txt || echo "0" > backups/_meta/CHANGED_LINES.txt
      else
        echo "0" > backups/_meta/CHANGED_LINES.txt
      fi
  - name: maybe-prompt-reminder
    run: |
      AGE=$(cat backups/_meta/LAST_COMMIT_AGE.txt 2>/dev/null || echo "999999")
      CHG=$(cat backups/_meta/CHANGED_LINES.txt 2>/dev/null || echo "0")
      if [ "$AGE" -gt 14400 ] || [ "$CHG" -gt 100 ]; then
        echo "::PROMPT:: 🕓 It’s been a while since your last commit/snapshot or a large diff was detected.
        Options:
        [1] Create snapshot now
        [2] Commit changes (local)
        [3] Skip"
      else
        echo "No reminder necessary."
      fi
  - name: act-on-choice
    run: |
      # Cursor should pass the user's selection as $CHOICE (1/2/3).
      case "$CHOICE" in
        1)
          echo "Creating snapshot…"
          echo "@daily-snapshot"
          ;;
        2)
          if [ -d ".git" ]; then
            git add -A
            git commit -m "chore: periodic checkpoint $(date +%Y-%m-%dT%H:%M:%S)"
            echo "✅ Local commit created."
          else
            echo "No Git repository initialized."
          fi
          ;;
        *)
          echo "Reminder skipped."
          ;;
      esac

safety:
  - never-auto-push: true

notes:
  - "Default choice is [1] Create snapshot."
  - "You can also trigger: @guard check | @guard enable | @guard disable"
