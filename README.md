# 🛡️ WP-AIBOM Sentinel

> The missing AI governance layer for WordPress.

WordPress powers 43% of the web. AI coding tools are writing more code than ever. But no one can prove which code was written by AI. Enterprises using WordPress VIP are demanding governance. Automattic has AI Guidelines but no enforcement mechanism.

WP-AIBOM Sentinel scans any WordPress plugin or theme, detects AI-generated code, and produces a machine-readable AI Bill of Materials for compliance audits.

## ✨ Features

- 🔍 AI Code Detection — Heuristic-based detection of AI-generated code patterns
- 📜 AIBOM Generation — Machine-readable JSON output
- ⚠️ Risk Scoring — Context-aware risk assessment (medium/high/critical)
- 🎯 WP-CLI Command — Native WordPress integration (wp aibom scan <file>)
- ⚙️ GitHub Action — Automatically scans PRs for AI code and comments the report
- 📊 Zero Dependencies — Built with Go and Python, no cloud required

## 🚀 Quick Start (WP-CLI)

Scan any WordPress plugin file:
wp aibom scan wp-content/plugins/hello.php

Output:
📄 File: /wp-content/plugins/hello.php
🤖 AI Probability: 27%
⚠️ Risk Level: medium

📋 Evidence Found:
   - Line 7 [generic_names]: $data...
   - Line 8 [generic_names]: $result...
   - Line 6 [obvious_comments]: // Initialize ...

## ⚙️ GitHub Action (CI/CD)

Add this to your .github/workflows/aibom-scan.yml and every PR will be automatically scanned.

## 📦 Installation

### CLI
go build -o aibom ./cmd/aibom
./aibom scan <path_to_file>

### WP-CLI Plugin
1. Copy aibom-cli.php to wp-content/plugins/wp-aibom-sentinel/
2. Activate the plugin in WordPress admin
3. Run wp aibom scan <file>

## 🧠 How It Works

1. Go CLI wraps the core engine
2. Python Engine scans files for known AI patterns (verbose comments, generic names, obvious comments)
3. WP-CLI Plugin bridges WordPress to the Go binary
4. GitHub Action triggers the scan on PRs and posts the AIBOM report

## 🎯 Why This Matters

- SOC2 Compliance — Prove AI code is audited
- GDPR Compliance — Track AI code handling PII
- Supply Chain Security — Know what's in your dependencies
- Enterprise Trust — Give WordPress VIP clients confidence

## 👤 Author

Built by [Abhisharydv90](https://github.com/Abhisharydv90). 
