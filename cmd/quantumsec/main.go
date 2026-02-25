package main

import (
	"fmt"
	"os"

	"github.com/hallucinaut/quantumsec/pkg/crypto"
	"github.com/hallucinaut/quantumsec/pkg/analyze"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "assess":
		if len(os.Args) < 3 {
			fmt.Println("Error: algorithm list required")
			printUsage()
			return
		}
		assessSecurity(os.Args[2:])
	case "analyze":
		if len(os.Args) < 3 {
			fmt.Println("Error: algorithm list required")
			printUsage()
			return
		}
		analyzeAlgorithms(os.Args[2:])
	case "timeline":
		showTimeline()
	case "check":
		checkAlgorithm(os.Args[2])
	case "version":
		fmt.Printf("quantumsec version %s\n", version)
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
	}
}

func printUsage() {
	fmt.Printf(`quantumsec - Quantum Cryptography Security Analyzer

Usage:
  quantumsec <command> [options]

Commands:
  assess <alg1,alg2,...>   Assess quantum security of algorithms
  analyze <alg1,alg2,...>  Analyze algorithms for quantum threats
  timeline                 Show quantum threat timeline
  check <algorithm>        Check single algorithm
  version                  Show version information
  help                     Show this help message

Examples:
  quantumsec assess "RSA-2048,AES-256,SHA-256"
  quantumsec analyze "Kyber-768,Dilithium-2"
  quantumsec check RSA-2048
`, "quantumsec")
}

func assessSecurity(algs []string) {
	fmt.Printf("Assessing quantum security of: %v\n", algs)
	fmt.Println()

	analyzer := crypto.NewAnalyzer()
	assessment := analyzer.AssessSecurity(algs)

	fmt.Printf("Assessment Date: %s\n", assessment.AssessmentDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("Overall Risk: %s\n", assessment.OverallRisk)
	fmt.Printf("Compliance: %s\n\n", assessment.ComplianceStatus)

	fmt.Println("Algorithm Analysis:")
	for i, alg := range assessment.Algorithms {
		status := "✓"
		if alg.Status == crypto.StatusVulnerable {
			status = "⚠"
		} else if alg.Status == crypto.StatusDeprecated {
			status = "✗"
		}
		fmt.Printf("  [%d] %s %s (%s)\n", i+1, status, alg.Name, alg.Status)
		fmt.Printf("      Threat: %s\n", alg.QuantumThreat)
		fmt.Printf("      Recommendation: %s\n\n", alg.Recommendation)
	}

	if len(assessment.Recommendations) > 0 {
		fmt.Println("Recommendations:")
		for _, rec := range assessment.Recommendations {
			fmt.Printf("  - %s\n", rec)
		}
	}
}

func analyzeAlgorithms(algs []string) {
	fmt.Printf("Analyzing algorithms: %v\n", algs)
	fmt.Println()

	analyzer := crypto.NewAnalyzer()
	assessment := analyzer.AssessSecurity(algs)

	riskScore := analyzer.CalculateQuantumRisk(assessment)
	riskLevel := analyze.GetRiskLevel(riskScore * 100)

	result := &analyze.AnalysisResult{
		AnalysisDate:   assessment.AssessmentDate,
		RiskScore:      riskScore * 100,
		RiskLevel:      analyze.RiskLevel(riskLevel),
		QuantumReadiness: "RISK_" + string(riskLevel),
		Timeline:       make([]analyze.Timeline, 0),
		Actions:        assessment.Recommendations,
	}

	fmt.Println(analyze.GenerateReport(result))
}

func showTimeline() {
	fmt.Println("Quantum Threat Timeline")
	fmt.Println("=======================")
	fmt.Println()

	analyzer := analyze.NewAnalyzer()
	for _, timeline := range analyzer.quantumTimeline {
		fmt.Printf("%ds:\n", timeline.Year)
		fmt.Printf("  Threat Level: %s\n", timeline.ThreatLevel)
		fmt.Printf("  %s\n", timeline.Description)
		fmt.Printf("  Action: %s\n\n", timeline.Recommendation)
	}
}

func checkAlgorithm(alg string) {
	fmt.Printf("Checking algorithm: %s\n", alg)
	fmt.Println()

	analyzer := crypto.NewAnalyzer()
	result := analyzer.AnalyzeAlgorithm(alg)

	fmt.Printf("Algorithm: %s\n", result.Name)
	fmt.Printf("Type: %s\n", result.Type)
	fmt.Printf("Status: %s\n", result.Status)
	fmt.Printf("Threat: %s\n", result.QuantumThreat)
	fmt.Printf("Recommendation: %s\n", result.Recommendation)
	fmt.Printf("Standard: %s\n", result.Standard)
}