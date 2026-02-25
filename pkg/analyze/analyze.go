// Package analyze provides quantum security analysis.
package analyze

import (
	"fmt"
	"time"
)

// RiskLevel represents quantum security risk level.
type RiskLevel string

const (
	RiskCritical RiskLevel = "CRITICAL"
	RiskHigh     RiskLevel = "HIGH"
	RiskMedium   RiskLevel = "MEDIUM"
	RiskLow      RiskLevel = "LOW"
	RiskMinimal  RiskLevel = "MINIMAL"
)

// Timeline represents quantum threat timeline.
type Timeline struct {
	Year         int
	ThreatLevel  string
	Description  string
	Recommendation string
}

// AnalysisResult contains quantum security analysis results.
type AnalysisResult struct {
	AnalysisDate   time.Time
	RiskScore      float64
	RiskLevel      RiskLevel
	QuantumReadiness string
	Timeline       []Timeline
	Actions        []string
}

// Analyzer analyzes quantum security posture.
type Analyzer struct {
	quantumTimeline []Timeline
}

// NewAnalyzer creates a new quantum security analyzer.
func NewAnalyzer() *Analyzer {
	return &Analyzer{
		quantumTimeline: []Timeline{
			{
				Year:         2024,
				ThreatLevel:  "LOW",
				Description:  "Quantum computers still limited, but data harvested now could be decrypted later",
				Recommendation: "Start planning quantum migration",
			},
			{
				Year:         2028,
				ThreatLevel:  "MEDIUM",
				Description:  "Quantum computers may reach 1000-2000 qubits",
				Recommendation: "Begin testing post-quantum algorithms",
			},
			{
				Year:         2030,
				ThreatLevel:  "HIGH",
				Description:  "Quantum computers could break current asymmetric crypto",
				Recommendation: "Full quantum migration required",
			},
			{
				Year:         2035,
				ThreatLevel:  "CRITICAL",
				Description:  "Quantum computers could break RSA-2048 and ECC-256",
				Recommendation: "Post-quantum crypto mandatory",
			},
		},
	}
}

// AnalyzeQuantumRisk analyzes quantum security risk.
func (a *Analyzer) AnalyzeQuantumRisk(algorithmCount, vulnerableCount int) *AnalysisResult {
	result := &AnalysisResult{
		AnalysisDate:   time.Now(),
		RiskScore:      0.0,
		RiskLevel:      RiskMinimal,
		QuantumReadiness: "UNKNOWN",
		Timeline:       a.quantumTimeline,
		Actions:        make([]string, 0),
	}

	if algorithmCount == 0 {
		result.RiskLevel = RiskMinimal
		result.QuantumReadiness = "NOT_ASSESSED"
		return result
	}

	ratio := float64(vulnerableCount) / float64(algorithmCount)
	result.RiskScore = ratio * 100.0

	if ratio > 0.7 {
		result.RiskLevel = RiskCritical
		result.QuantumReadiness = "CRITICAL_RISK"
		result.Actions = append(result.Actions,
			"Immediately replace all vulnerable algorithms")
		result.Actions = append(result.Actions,
			"Implement hybrid crypto solutions")
		result.Actions = append(result.Actions,
			"Create urgent migration plan")
	} else if ratio > 0.5 {
		result.RiskLevel = RiskHigh
		result.QuantumReadiness = "HIGH_RISK"
		result.Actions = append(result.Actions,
			"Prioritize quantum migration")
		result.Actions = append(result.Actions,
			"Test post-quantum algorithms")
	} else if ratio > 0.3 {
		result.RiskLevel = RiskMedium
		result.QuantumReadiness = "MODERATE_RISK"
		result.Actions = append(result.Actions,
			"Plan quantum migration")
		result.Actions = append(result.Actions,
			"Evaluate post-quantum solutions")
	} else if ratio > 0.1 {
		result.RiskLevel = RiskLow
		result.QuantumReadiness = "LOW_RISK"
		result.Actions = append(result.Actions,
			"Monitor quantum developments")
		result.Actions = append(result.Actions,
			"Review cryptographic inventory")
	} else {
		result.RiskLevel = RiskMinimal
		result.QuantumReadiness = "QUANTUM_SAFE"
		result.Actions = append(result.Actions,
			"Maintain current security posture")
		result.Actions = append(result.Actions,
			"Stay informed about quantum advances")
	}

	return result
}

// GenerateReport generates quantum security report.
func GenerateReport(result *AnalysisResult) string {
	var report string

	report += "=== Quantum Security Analysis Report ===\n\n"
	report += "Analysis Date: " + result.AnalysisDate.Format("2006-01-02 15:04:05") + "\n"
	report += "Risk Score: " + fmt.Sprintf("%.0f%%", result.RiskScore) + "%\n"
	report += "Risk Level: " + string(result.RiskLevel) + "\n"
	report += "Quantum Readiness: " + result.QuantumReadiness + "\n\n"

	if len(result.Actions) > 0 {
		report += "Recommended Actions:\n"
		for i, action := range result.Actions {
			report += "[" + string(rune(i+49)) + "] " + action + "\n"
		}
		report += "\n"
	}

	report += "Quantum Threat Timeline:\n"
	for _, timeline := range result.Timeline {
		report += "  " + string(rune(timeline.Year-48)) + "0s: " + timeline.ThreatLevel + " - " + timeline.Description + "\n"
	}

	return report
}

// GetRiskLevel returns risk level from score.
func GetRiskLevel(score float64) RiskLevel {
	if score >= 70 {
		return RiskCritical
	} else if score >= 50 {
		return RiskHigh
	} else if score >= 30 {
		return RiskMedium
	} else if score >= 10 {
		return RiskLow
	}
	return RiskMinimal
}
// GetTimeline returns the quantum timeline for analysis.
func (a *Analyzer) GetTimeline() []Timeline {
	return a.quantumTimeline
}
