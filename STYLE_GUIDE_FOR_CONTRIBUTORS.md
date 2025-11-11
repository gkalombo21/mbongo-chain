# Mbongo Chain - Contributor Style Guide

## 1. Purpose
This guide establishes clear conventions for all contributors. Consistency across code, documentation, and commit history improves readability, accelerates reviews, and maintains a professional presentation for the Mbongo Chain project.

## 2. Language Rules
- English is the mandatory project language.
- All comments, commit messages, documentation, and file names must be in English.
- Cursor AI will automatically flag or translate non-English content.

## 3. File Structure and Organization
- Follow the modular Go project architecture:
  - `/cmd/` for entry points  
  - `/internal/` for main modules (bank, blockchain, AI, etc.)  
  - `/pkg/` for reusable packages  
- Keep file names lowercase with underscores if necessary (no spaces).

## 4. Markdown Formatting
- Each Markdown file must begin with a title (H1).
- Use H2 and H3 for sub-sections.
- Avoid paragraphs longer than 8 lines.
- Use bullet points for lists, and prefer tables for comparisons.

## 5. Go Code Standards
- Use standard Go formatting (`gofmt`).
- Each exported function must include a GoDoc comment.
- Imports should be grouped: standard → external → internal.
- Variables and functions should be named descriptively and in English.

## 6. Commit Message Guidelines
- Use present tense and short, meaningful descriptions.
- Example: `Add validation layer to bank module`
- Avoid non-English text or emojis.

## 7. File Validation Rules (via Cursor)
- Cursor automatically checks:
  - English-only content  
  - Markdown structure and consistency  
  - Go code formatting and imports  
- Contributors should fix all red warnings before committing.

## 8. Code Review Process
- Each pull request must include:
  - A short summary of changes.  
  - Confirmation that `go test ./...` passes without errors.  
- Reviews focus on clarity, efficiency, and maintainability.

## 9. Security and Ethics
- Never commit sensitive data (API keys, credentials).
- Respect the open-source MIT license and contributor code of conduct.

## 10. Contact
For any questions, contributors can open an issue or contact the maintainers at:  
`team@mbongo.io`
