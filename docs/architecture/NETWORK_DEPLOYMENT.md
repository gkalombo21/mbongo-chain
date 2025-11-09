# ⚙️ Mbongo Chain Network Deployment Guide

## 1. Introduction

This guide explains how to deploy **Mbongo Chain** for local testing, testnet participation, and cloud-based environments — without Docker.  
Instead, we use **Go commands**, **Postman**, and **.env configuration files**.

---

## 2. Environment Requirements

| Component | Minimum Version | Notes |
|------------|-----------------|-------|
| Go | 1.21+ | Required for compilation |
| Node.js | 18+ | For API testing (optional) |
| Postman | Latest | For endpoint testing |
| Git | 2.30+ | To clone and sync repositories |
| OS | Windows 10/11, macOS, or Linux | 64-bit recommended |

---

## 3. Environment Setup

1️⃣ Clone the repository:

```bash
git clone https://github.com/gkalombo21/mbongo-chain.git
cd mbongo-chain
```

| Method | Endpoint | Description |

> |---------|-----------|-------------|
> | `GET` | `http://localhost:8080/status` | Check node health |
> | `POST` | `http://localhost:8080/account/create` | Create new wallet |
> | `POST` | `http://localhost:8080/transaction/send` | Send MBG tokens |
> | `GET` | `http://localhost:8080/ledger` | View recent blocks |

| Issue | Fix |

> |--------|-----|
> | “port already in use” | Change port in `.env` |
> | “go.mod missing” | Run `go mod init github.com/gkalombo21/mbongo-chain` |
> | Node crash | Run `go clean -cache` then rebuild |

---

## Next Steps

- Configure multiple validator nodes for testnet simulation.
- Use Postman to validate API responses and ledger integrity.
- Deploy first AI Compute Jobs using `/ai/policy` specifications.

---

## Dashboard

Mbongo provides an optional web dashboard: `http://localhost:8081/status`

Displays:

- Node status (online/offline)
- Current AI job
- Rewards earned
- Peer connections

---

## Security Best Practices

- Keep your private keys offline.
- Use hardware encryption for `.env` credentials.
- Run node under limited user permissions.
- Update regularly: `git pull && go build -o mbongod.exe ./cmd/mbongo-chain`
