![Build Status](https://github.com/gkalombo21/mbongo-chain/actions/workflows/main.yml/badge.svg)
![License](https://img.shields.io/github/license/gkalombo21/mbongo-chain)
![Contributors](https://img.shields.io/github/contributors/gkalombo21/mbongo-chain)
![Last Commit](https://img.shields.io/github/last-commit/gkalombo21/mbongo-chain)
![Stars](https://img.shields.io/github/stars/gkalombo21/mbongo-chain?style=social)
![Issues](https://img.shields.io/github/issues/gkalombo21/mbongo-chain)
![MBG Supply](https://img.shields.io/badge/🪙%20MBG%20Total%20Supply-1%20Billion-gold?style=for-the-badge)

_Community-driven project for decentralized AI and financial blockchain._
_Fixed supply of 1,000,000,000 MBG tokens — distributed between validators, compute nodes, and the community._

# Mbongo Chain
Decentralized AI and Financial Blockchain built with Go + Cosmos SDK

> Mbongo Chain is a hybrid Proof of Stake + Proof of Useful Work blockchain powering decentralized AI compute and financial systems.

Mbongo Chain is an open-source, Africa-first protocol that blends scalable blockchain infrastructure with AI-driven utility. By rewarding useful GPU computation and staking, the network provides sustainable economic incentives for validators, AI contributors, developers, and communities worldwide.

---

## Key Features
- **Hybrid Consensus (PoS + PoUW):** Combines Tendermint-style staking security with AI/GPU powered useful work.
- **Fast Blocks:** 5-minute block time keeps the network responsive for finance and AI workloads.
- **Predictable Supply:** Halving every 5 years with a 100-year emission schedule.
- **Go + Cosmos SDK:** Built with Go, inspired by Cosmos architecture for modular design and interoperability.
- **Developer-First:** Clean module boundaries (`cmd/`, `internal/`, `pkg/`), clear APIs, and detailed docs for rapid onboarding.

---

## Architecture Snapshot
```
Applications & Marketplaces (AI services, DeFi tools)
└── API & SDK Layer
    └── Hybrid Blockchain Core (PoS validators)
        └── AI Compute Layer (PoUW GPU nodes)
```
- **Validator Layer:** Tendermint-based BFT consensus with MBG staking.
- **Compute Layer:** Schedules AI workloads, verifies proofs, distributes rewards.
- **Governance:** On-chain DAO proposals with staking-weighted voting.

---

## Tokenomics Highlights
- **Symbol:** MBG (Mbongo)
- **Total Supply:** 10,000,000,000 MBG
- **Initial Circulation:** 2,000,000,000 MBG
- **Reward Split:** 50% PoS | 30% PoUW | 15% Dev Fund | 5% Foundation
- **Emission:** Halving every 5 years, approaching zero by year 100

Full details in [`TOKENOMICS.md`](TOKENOMICS.md) and [`docs/tokenomics.md`](docs/tokenomics.md).

---

## Getting Started
### Prerequisites
- Go 1.21+
- Git
- Podman or Docker (optional for services)

### Quickstart
```bash
git clone https://github.com/gkalombo21/mbongo-chain.git
cd mbongo-chain
go mod tidy
go run ./cmd/mbongo-chain
```

### Run Tokenomics Simulation
```bash
go run ./scripts/simulate_tokenomics.go
```

---

## Documentation
- [Architecture Overview](docs/architecture/MBONGO_SYSTEM_OVERVIEW.md)
- [Network Deployment Guide](docs/architecture/NETWORK_DEPLOYMENT.md)
- [AI Compute Guide](docs/ai/AI_COMPUTE_GUIDE.md)
- [Whitepaper](WHITEPAPER.md)

---

## Community and Contribution
- Contribution guide: see [`CONTRIBUTING.md`](CONTRIBUTING.md)
- Code of Conduct: see [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md)
- Security policy: see [`SECURITY.md`](SECURITY.md)

---

## Whitepaper and Documentation
- 📄 [Whitepaper](WHITEPAPER.md)
- 📘 Technical Docs (to be added)

---

## License
Mbongo Chain is open-source under the [MIT License](LICENSE).

---

## Contact
- Email: [info@mbongo.io](mailto:info@mbongo.io)
- Website (placeholder): [https://mbongo.io](https://mbongo.io)
- Twitter / GitHub links: *to be added*

---

## 📄 Documentation

- [📘 Read the Full Whitepaper](./WHITEPAPER.md)
- [Contributor Style Guide](STYLE_GUIDE_FOR_CONTRIBUTORS.md)
- [Code of Conduct](CODE_OF_CONDUCT.md)
- [Security Policy](SECURITY.md)

---

## Mission
Mbongo Chain is building a decentralized, Africa-first AI infrastructure where useful work is rewarded, innovation thrives, and global collaboration shapes a resilient digital future.
