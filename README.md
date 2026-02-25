# quantumsec - Quantum Cryptography Security Analyzer

[![Go](https://img.shields.io/badge/Go-1.21-blue)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

**Assess cryptographic systems for quantum computing threats and plan migration to post-quantum cryptography.**

Prepare your systems for the quantum computing era before current cryptography becomes vulnerable.

## 🚀 Features

- **Quantum Threat Assessment**: Analyze algorithms against quantum computing threats
- **Algorithm Database**: Known algorithms with quantum resistance status
- **Migration Planning**: Generate actionable migration plans
- **Timeline Analysis**: Understand quantum threat timeline
- **Compliance Checking**: Check quantum security compliance
- **Risk Scoring**: Calculate quantum risk scores

## 📦 Installation

### Build from Source

```bash
git clone https://github.com/hallucinaut/quantumsec.git
cd quantumsec
go build -o quantumsec ./cmd/quantumsec
sudo mv quantumsec /usr/local/bin/
```

### Install via Go

```bash
go install github.com/hallucinaut/quantumsec/cmd/quantumsec@latest
```

## 🎯 Usage

### Assess Security

```bash
# Assess multiple algorithms
quantumsec assess "RSA-2048,AES-256,SHA-256"

# Full quantum security assessment
quantumsec assess "Kyber-768,Dilithium-2,SHA3-512"
```

### Analyze Algorithms

```bash
# Analyze for quantum threats
quantumsec analyze "RSA-2048,ECC-256"

# Check single algorithm
quantumsec check RSA-4096
```

### View Timeline

```bash
# Show quantum threat timeline
quantumsec timeline
```

### Programmatic Usage

```go
package main

import (
    "fmt"
    "github.com/hallucinaut/quantumsec/pkg/crypto"
)

func main() {
    analyzer := crypto.NewAnalyzer()
    
    // Assess algorithms
    assessment := analyzer.AssessSecurity([]string{
        "RSA-2048",
        "AES-256",
        "SHA-256",
    })
    
    fmt.Printf("Overall Risk: %s\n", assessment.OverallRisk)
    fmt.Printf("Compliance: %s\n", assessment.ComplianceStatus)
    
    // Calculate risk score
    riskScore := analyzer.CalculateQuantumRisk(assessment)
    fmt.Printf("Risk Score: %.0f%%\n", riskScore*100)
    
    // Generate migration plan
    plan := analyzer.GenerateMigrationPlan(assessment)
    for action, priority := range plan {
        fmt.Printf("%s: %s\n", action, priority)
    }
}
```

## 🔍 Quantum Threats

### Shor's Algorithm
Breaks asymmetric cryptography:
- RSA
- ECC
- Diffie-Hellman

### Grover's Algorithm
Provides quadratic speedup for:
- Symmetric encryption (AES, ChaCha20)
- Hash functions (SHA-2, SHA-3)

## 📊 Algorithm Status

### Quantum-Safe Algorithms

| Algorithm | Type | Key Size | Quantum Threat |
|-----------|------|----------|----------------|
| Kyber-512 | Asymmetric | 512 | Resistant |
| Kyber-768 | Asymmetric | 768 | Resistant |
| Dilithium-2 | Asymmetric | 2 | Resistant |
| AES-256 | Symmetric | 256 | Safe (Grover) |
| SHA3-512 | Hash | 512 | Resistant |

### Quantum-Vulnerable Algorithms

| Algorithm | Type | Key Size | Quantum Threat |
|-----------|------|----------|----------------|
| RSA-1024 | Asymmetric | 1024 | ~1000 qubits |
| RSA-2048 | Asymmetric | 2048 | ~4000 qubits |
| ECC-256 | Asymmetric | 256 | ~2300 qubits |
| SHA-1 | Hash | 160 | Collisions found |

## 📋 Risk Levels

| Score | Level | Action |
|-------|-------|--------|
| 0-10% | MINIMAL | Monitor |
| 10-30% | LOW | Plan migration |
| 30-50% | MEDIUM | Prioritize migration |
| 50-70% | HIGH | Urgent migration |
| 70-100% | CRITICAL | Immediate action |

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test -v ./pkg/crypto -run TestAssessSecurity
```

## 📋 Example Output

```
Assessing quantum security of: [RSA-2048 AES-256 SHA-256]

Assessment Date: 2024-02-25 15:30:00
Overall Risk: HIGH
Compliance: AT_RISK

Algorithm Analysis:
  [1] ⚠ RSA-2048 (QUANTUM_VULNERABLE)
      Threat: Shor's algorithm can break in ~4000 qubits
      Recommendation: Plan migration to post-quantum algorithms

  [2] ✓ AES-256 (QUANTUM_SECURE)
      Threat: Grover's algorithm provides quadratic speedup
      Recommendation: Quantum-safe with sufficient margin

  [3] ✓ SHA-256 (QUANTUM_SECURE)
      Threat: Grover's algorithm provides quadratic speedup
      Recommendation: Quantum-safe with sufficient margin

Recommendations:
  - Migrate from RSA-2048 to quantum-safe alternative
```

## 🔒 Security Use Cases

- **Cryptographic Inventory Assessment**: Audit all cryptographic algorithms
- **Quantum Migration Planning**: Plan transition to post-quantum crypto
- **Compliance Verification**: Ensure quantum security compliance
- **Risk Management**: Assess quantum risk to systems
- **Long-term Data Protection**: Protect data for future decryption

## 🛡️ Best Practices

1. **Inventory all cryptographic algorithms** in your systems
2. **Assess quantum risk** for each algorithm
3. **Prioritize migration** based on risk scores
4. **Test post-quantum algorithms** before deployment
5. **Monitor quantum computing progress**
6. **Plan for hybrid crypto** during transition
7. **Protect sensitive data** with quantum-safe encryption

## 🏗️ Architecture

```
quantumsec/
├── cmd/
│   └── quantumsec/
│       └── main.go          # CLI entry point
├── pkg/
│   ├── crypto/
│   │   ├── crypto.go       # Quantum crypto analysis
│   │   └── crypto_test.go  # Unit tests
│   └── analyze/
│       ├── analyze.go      # Security analysis
│       └── analyze_test.go # Unit tests
└── README.md
```

## 📄 License

MIT License

## 🙏 Acknowledgments

- NIST Post-Quantum Cryptography Project
- Quantum computing research community
- Cryptography security researchers

## 🔗 Resources

- [NIST Post-Quantum Cryptography](https://csrc.nist.gov/Projects/post-quantum-cryptography)
- [Quantum Computing Report](https://quantumcomputingreport.com/)
- [Crypto Quantum Threat Timeline](https://www.zhenyu-liao.com/post/quantum-crypto-timeline/)

---

**Built with ❤️ by [hallucinaut](https://github.com/hallucinaut)**