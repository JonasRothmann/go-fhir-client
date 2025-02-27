// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/variable-role
*/type EvidenceVariableRole string

const (
	// variable represents a population.
	EvidenceVariableRolePopulation EvidenceVariableRole = "population"
	// variable represents a subpopulation.
	EvidenceVariableRoleSubpopulation EvidenceVariableRole = "subpopulation"
	// variable represents an exposure.
	EvidenceVariableRoleExposure EvidenceVariableRole = "exposure"
	// variable represents a reference exposure.
	EvidenceVariableRoleReferenceExposure EvidenceVariableRole = "referenceExposure"
	// variable represents a measured variable.
	EvidenceVariableRoleMeasuredVariable EvidenceVariableRole = "measuredVariable"
	// variable represents a confounder.
	EvidenceVariableRoleConfounder EvidenceVariableRole = "confounder"
	// variable represents a first of two measured variables to be used in a measured variable that is a mathematical operation of two measured variables.
	EvidenceVariableRoleMeasuredVariableA EvidenceVariableRole = "measuredVariableA"
	// variable represents a second of two measured variables to be used in a measured variable that is a mathematical operation of two measured variables.
	EvidenceVariableRoleMeasuredVariableB EvidenceVariableRole = "measuredVariableB"
	// variable represents a result of a mathematical operation of two measured variables.
	EvidenceVariableRoleMeasuredVariableAb EvidenceVariableRole = "measuredVariableAB"
)
