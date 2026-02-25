// Package crypto provides quantum cryptography security analysis.
package crypto

import (
	"regexp"
	"time"
)

// AlgorithmStatus represents quantum resistance status.
type AlgorithmStatus string

const (
	StatusSecure      AlgorithmStatus = "QUANTUM_SECURE"
	StatusVulnerable  AlgorithmStatus = "QUANTUM_VULNERABLE"
	StatusUnknown     AlgorithmStatus = "QUANTUM_UNKNOWN"
	StatusDeprecated  AlgorithmStatus = "QUANTUM_DEPRECATED"
)

// CryptoAlgorithm represents a cryptographic algorithm.
type CryptoAlgorithm struct {
	Name           string
	Type           string // asymmetric, symmetric, hash
	KeySize        int
	Status         AlgorithmStatus
	QuantumThreat  string
	Recommendation string
	Standard       string
}

// SecurityAssessment represents quantum security assessment.
type SecurityAssessment struct {
	AssessmentDate   time.Time
	Algorithms       []CryptoAlgorithm
	OverallRisk      string
	ComplianceStatus string
	Recommendations  []string
}

// Analyzer analyzes cryptographic algorithms for quantum threats.
type Analyzer struct {
	knownAlgorithms map[string]*CryptoAlgorithm
}

// NewAnalyzer creates a new quantum crypto analyzer.
func NewAnalyzer() *Analyzer {
	return &Analyzer{
		knownAlgorithms: make(map[string]*CryptoAlgorithm),
	}
}

// InitializeKnownAlgorithms initializes known algorithm database.
func (a *Analyzer) InitializeKnownAlgorithms() {
	// Asymmetric algorithms
	a.knownAlgorithms["RSA-1024"] = &CryptoAlgorithm{
		Name:           "RSA-1024",
		Type:           "asymmetric",
		KeySize:        1024,
		Status:         StatusVulnerable,
		QuantumThreat:  "Shor's algorithm can break in ~1000 qubits",
		Recommendation: "Upgrade to RSA-2048+ or use post-quantum crypto",
		Standard:       "NIST SP 800-57",
	}
	a.knownAlgorithms["RSA-2048"] = &CryptoAlgorithm{
		Name:           "RSA-2048",
		Type:           "asymmetric",
		KeySize:        2048,
		Status:         StatusVulnerable,
		QuantumThreat:  "Shor's algorithm can break in ~4000 qubits",
		Recommendation: "Plan migration to post-quantum algorithms",
		Standard:       "NIST SP 800-57",
	}
	a.knownAlgorithms["RSA-4096"] = &CryptoAlgorithm{
		Name:           "RSA-4096",
		Type:           "asymmetric",
		KeySize:        4096,
		Status:         StatusVulnerable,
		QuantumThreat:  "Shor's algorithm can break in ~8000 qubits",
		Recommendation: "Plan migration to post-quantum algorithms",
		Standard:       "NIST SP 800-57",
	}
	a.knownAlgorithms["ECC-256"] = &CryptoAlgorithm{
		Name:           "ECC-256",
		Type:           "asymmetric",
		KeySize:        256,
		Status:         StatusVulnerable,
		QuantumThreat:  "Shor's algorithm can break in ~2300 qubits",
		Recommendation: "Use post-quantum elliptic curve crypto",
		Standard:       "NIST FIPS 186-4",
	}
	a.knownAlgorithms["ECC-384"] = &CryptoAlgorithm{
		Name:           "ECC-384",
		Type:           "asymmetric",
		KeySize:        384,
		Status:         StatusVulnerable,
		QuantumThreat:  "Shor's algorithm can break in ~4000 qubits",
		Recommendation: "Use post-quantum elliptic curve crypto",
		Standard:       "NIST FIPS 186-4",
	}

	// Post-quantum algorithms
	a.knownAlgorithms["Kyber-512"] = &CryptoAlgorithm{
		Name:           "Kyber-512",
		Type:           "asymmetric",
		KeySize:        512,
		Status:         StatusSecure,
		QuantumThreat:  "Resistant to known quantum attacks",
		Recommendation: "Recommended for quantum-safe deployments",
		Standard:       "NIST CRYSTALS-Kyber",
	}
	a.knownAlgorithms["Kyber-768"] = &CryptoAlgorithm{
		Name:           "Kyber-768",
		Type:           "asymmetric",
		KeySize:        768,
		Status:         StatusSecure,
		QuantumThreat:  "Resistant to known quantum attacks",
		Recommendation: "Recommended for quantum-safe deployments",
		Standard:       "NIST CRYSTALS-Kyber",
	}
	a.knownAlgorithms["Dilithium-2"] = &CryptoAlgorithm{
		Name:           "Dilithium-2",
		Type:           "asymmetric",
		KeySize:        2,
		Status:         StatusSecure,
		QuantumThreat:  "Resistant to known quantum attacks",
		Recommendation: "Recommended for quantum-safe signatures",
		Standard:       "NIST CRYSTALS-Dilithium",
	}
	a.knownAlgorithms["FALCON-512"] = &CryptoAlgorithm{
		Name:           "FALCON-512",
		Type:           "asymmetric",
		KeySize:        512,
		Status:         StatusSecure,
		QuantumThreat:  "Resistant to known quantum attacks",
		Recommendation: "Recommended for compact signatures",
		Standard:       "NIST CRYSTALS-Dilithium",
	}

	// Symmetric algorithms
	a.knownAlgorithms["AES-128"] = &CryptoAlgorithm{
		Name:           "AES-128",
		Type:           "symmetric",
		KeySize:        128,
		Status:         StatusSecure,
		QuantumThreat:  "Grover's algorithm provides quadratic speedup",
		Recommendation: "Double key size to AES-256 for quantum security",
		Standard:       "NIST FIPS 197",
	}
	a.knownAlgorithms["AES-256"] = &CryptoAlgorithm{
		Name:           "AES-256",
		Type:           "symmetric",
		KeySize:        256,
		Status:         StatusSecure,
		QuantumThreat:  "Grover's algorithm provides quadratic speedup",
		Recommendation: "Quantum-safe with sufficient margin",
		Standard:       "NIST FIPS 197",
	}
	a.knownAlgorithms["ChaCha20"] = &CryptoAlgorithm{
		Name:           "ChaCha20",
		Type:           "symmetric",
		KeySize:        256,
		Status:         StatusSecure,
		QuantumThreat:  "Grover's algorithm provides quadratic speedup",
		Recommendation: "Quantum-safe with sufficient margin",
		Standard:       "RFC 8439",
	}

	// Hash algorithms
	a.knownAlgorithms["SHA-1"] = &CryptoAlgorithm{
		Name:           "SHA-1",
		Type:           "hash",
		KeySize:        160,
		Status:         StatusDeprecated,
		QuantumThreat:  "Collisions found, quantum attacks possible",
		Recommendation: "Migrate to SHA-256 or SHA-3",
		Standard:       "NIST FIPS 180-4",
	}
	a.knownAlgorithms["SHA-256"] = &CryptoAlgorithm{
		Name:           "SHA-256",
		Type:           "hash",
		KeySize:        256,
		Status:         StatusSecure,
		QuantumThreat:  "Grover's algorithm provides quadratic speedup",
		Recommendation: "Quantum-safe with sufficient margin",
		Standard:       "NIST FIPS 180-4",
	}
	a.knownAlgorithms["SHA-384"] = &CryptoAlgorithm{
		Name:           "SHA-384",
		Type:           "hash",
		KeySize:        384,
		Status:         StatusSecure,
		QuantumThreat:  "Grover's algorithm provides quadratic speedup",
		Recommendation: "Quantum-safe with sufficient margin",
		Standard:       "NIST FIPS 180-4",
	}
	a.knownAlgorithms["SHA-512"] = &CryptoAlgorithm{
		Name:           "SHA-512",
		Type:           "hash",
		KeySize:        512,
		Status:         StatusSecure,
		QuantumThreat:  "Grover's algorithm provides quadratic speedup",
		Recommendation: "Quantum-safe with sufficient margin",
		Standard:       "NIST FIPS 180-4",
	}
	a.knownAlgorithms["SHA3-256"] = &CryptoAlgorithm{
		Name:           "SHA3-256",
		Type:           "hash",
		KeySize:        256,
		Status:         StatusSecure,
		QuantumThreat:  "Resistant to quantum attacks",
		Recommendation: "Quantum-safe, recommended for new systems",
		Standard:       "NIST FIPS 202",
	}
	a.knownAlgorithms["SHA3-512"] = &CryptoAlgorithm{
		Name:           "SHA3-512",
		Type:           "hash",
		KeySize:        512,
		Status:         StatusSecure,
		QuantumThreat:  "Resistant to quantum attacks",
		Recommendation: "Quantum-safe, recommended for new systems",
		Standard:       "NIST FIPS 202",
	}
}

// AnalyzeAlgorithm analyzes an algorithm for quantum security.
func (a *Analyzer) AnalyzeAlgorithm(algorithm string) *CryptoAlgorithm {
	if alg, exists := a.knownAlgorithms[algorithm]; exists {
		return alg
	}

	// Try to extract algorithm info from string
	re := regexp.MustCompile(`(RSA|ECC|AES|SHA|Kyber|Dilithium|FALCON|ChaCha)\s*(-?\d+)?`)
	matches := re.FindStringSubmatch(algorithm)

	if len(matches) >= 2 {
		algName := matches[1]
		keySize := 0
		if len(matches) >= 3 && matches[2] != "" {
			keySize = parseInt(matches[2])
		}

		return &CryptoAlgorithm{
			Name:           algorithm,
			Type:           "unknown",
			KeySize:        keySize,
			Status:         StatusUnknown,
			QuantumThreat:  "Not in known database",
			Recommendation: "Verify quantum security separately",
			Standard:       "Unknown",
		}
	}

	return &CryptoAlgorithm{
		Name:           algorithm,
		Type:           "unknown",
		KeySize:        0,
		Status:         StatusUnknown,
		QuantumThreat:  "Unable to identify algorithm",
		Recommendation: "Provide full algorithm name",
		Standard:       "Unknown",
	}
}

// AssessSecurity performs full security assessment.
func (a *Analyzer) AssessSecurity(algorithms []string) *SecurityAssessment {
	a.InitializeKnownAlgorithms()

	assessment := &SecurityAssessment{
		AssessmentDate:  time.Now(),
		Algorithms:      make([]CryptoAlgorithm, 0),
		OverallRisk:     "LOW",
		ComplianceStatus: "COMPLIANT",
		Recommendations: make([]string, 0),
	}

	vulnerableCount := 0
	secureCount := 0
	deprecatedCount := 0

	for _, algName := range algorithms {
		alg := a.AnalyzeAlgorithm(algName)
		assessment.Algorithms = append(assessment.Algorithms, *alg)

		switch alg.Status {
		case StatusVulnerable:
			vulnerableCount++
			assessment.Recommendations = append(assessment.Recommendations,
				"Migrate from "+algName+" to quantum-safe alternative")
		case StatusSecure:
			secureCount++
		case StatusDeprecated:
			deprecatedCount++
			assessment.Recommendations = append(assessment.Recommendations,
				"Immediately replace deprecated "+algName)
		}
	}

	// Determine overall risk
	total := len(algorithms)
	if total == 0 {
		assessment.OverallRisk = "UNKNOWN"
	} else {
		vulnerableRatio := float64(vulnerableCount+deprecatedCount) / float64(total)

		if vulnerableRatio > 0.5 {
			assessment.OverallRisk = "CRITICAL"
		} else if vulnerableRatio > 0.3 {
			assessment.OverallRisk = "HIGH"
		} else if vulnerableRatio > 0.1 {
			assessment.OverallRisk = "MEDIUM"
		} else {
			assessment.OverallRisk = "LOW"
		}
	}

	// Determine compliance
	if deprecatedCount > 0 {
		assessment.ComplianceStatus = "NON_COMPLIANT"
	} else if vulnerableCount > 0 {
		assessment.ComplianceStatus = "AT_RISK"
	} else {
		assessment.ComplianceStatus = "COMPLIANT"
	}

	return assessment
}

// CalculateQuantumRisk calculates quantum risk score.
func (a *Analyzer) CalculateQuantumRisk(assessment *SecurityAssessment) float64 {
	if len(assessment.Algorithms) == 0 {
		return 0.0
	}

	score := 0.0
	for _, alg := range assessment.Algorithms {
		switch alg.Status {
		case StatusSecure:
			score += 0
		case StatusVulnerable:
			score += 0.5
		case StatusDeprecated:
			score += 1.0
		default:
			score += 0.25
		}
	}

	return score / float64(len(assessment.Algorithms))
}

// GenerateMigrationPlan generates migration plan.
func (a *Analyzer) GenerateMigrationPlan(assessment *SecurityAssessment) map[string]string {
	plan := make(map[string]string)

	for _, rec := range assessment.Recommendations {
		plan[rec] = "Migration required"
	}

	if assessment.OverallRisk == "CRITICAL" {
		plan["IMMEDIATE_ACTION"] = "Priority 1 - Quantum migration critical"
	} else if assessment.OverallRisk == "HIGH" {
		plan["IMMEDIATE_ACTION"] = "Priority 2 - Plan migration within 30 days"
	} else if assessment.OverallRisk == "MEDIUM" {
		plan["IMMEDIATE_ACTION"] = "Priority 3 - Plan migration within 90 days"
	}

	return plan
}

// parseInt converts string to int.
func parseInt(s string) int {
	var result int
	for _, c := range s {
		if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
		}
	}
	return result
}

// GetAlgorithmStatus returns algorithm status.
func GetAlgorithmStatus(alg *CryptoAlgorithm) AlgorithmStatus {
	return alg.Status
}