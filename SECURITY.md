# 🔐 Mbongo-Chain Security Policy

## Reporting a Vulnerability
If you discover a security vulnerability in Mbongo-Chain, please **do not open a public issue**.
Instead, contact our security team directly via email:

📧 **security@mbongo-chain.io**

We aim to respond to all reports within **72 hours** and to provide a fix or mitigation within **7 business days**.

If your report is validated, you will be publicly credited in our **Security Hall of Fame** (or optionally remain anonymous).

---

## Supported Versions

| Version | Supported |
|----------|------------|
| main (latest) | ✅ Supported |
| older releases | ⚠️ Community-supported only |

---

## Coordinated Disclosure
We follow the principles of **Coordinated Vulnerability Disclosure (CVD)** as recommended by [OpenSSF](https://openssf.org) and [CERT/CC](https://www.cert.org).
Please avoid sharing details of potential vulnerabilities publicly until a fix is confirmed and released.

---

## Security Best Practices for Contributors
When submitting code or pull requests, please:
- Use **secure coding patterns** and validate all inputs.
- Avoid hardcoded secrets or API keys.
- Use Go’s built-in crypto libraries for encryption/hashing (no custom crypto).
- Follow the **OWASP Top 10** and **CWE/SANS** secure development guidelines.

---

## Recommended Tools
Developers are encouraged to use:
- **gosec** for Go security analysis
- **dependabot** for dependency scanning
- **CodeQL** for advanced code analysis

---

## Acknowledgements
We thank all contributors and auditors who help make Mbongo-Chain safer for everyone.
Special thanks to the open-source security community for continuous support.

