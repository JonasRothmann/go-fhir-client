// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/risk-probability
*/type RiskProbability string

const (
	// The specified outcome is exceptionally unlikely.
	RiskProbabilityNegligible RiskProbability = "negligible"
	// The specified outcome is possible but unlikely.
	RiskProbabilityLow RiskProbability = "low"
	// The specified outcome has a reasonable likelihood of occurrence.
	RiskProbabilityModerate RiskProbability = "moderate"
	// The specified outcome is more likely to occur than not.
	RiskProbabilityHigh RiskProbability = "high"
	// The specified outcome is effectively guaranteed.
	RiskProbabilityCertain RiskProbability = "certain"
)
