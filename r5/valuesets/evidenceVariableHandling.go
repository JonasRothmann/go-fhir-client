// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/variable-handling
*/type EvidenceVariableHandling string

const (
	// A continuous variable is one for which, within the limits the variable ranges, any value is possible (from STATO http://purl.obolibrary.org/obo/STATO_0000251).
	EvidenceVariableHandlingContinuous EvidenceVariableHandling = "continuous"
	// A dichotomous variable is a categorical variable which is defined to have only 2 categories or possible values (from STATO http://purl.obolibrary.org/obo/STATO_0000090).
	EvidenceVariableHandlingDichotomous EvidenceVariableHandling = "dichotomous"
	// An ordinal variable is a categorical variable where the discrete possible values are ordered or correspond to an implicit ranking (from STATO http://purl.obolibrary.org/obo/STATO_0000228).
	EvidenceVariableHandlingOrdinal EvidenceVariableHandling = "ordinal"
	// A polychotomous variable is a categorical variable which is defined to have minimally 2 categories or possible values. (from STATO  http://purl.obolibrary.org/obo/STATO_0000087).  Suggestion to limit code use to situations when neither dichotomous nor ordinal variables apply.
	EvidenceVariableHandlingPolychotomous EvidenceVariableHandling = "polychotomous"
)
