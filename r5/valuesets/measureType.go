// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/measure-type
*/type MeasureType string

const (
	// A measure which focuses on a process which leads to a certain outcome, meaning that a scientific basis exists for believing that the process, when executed well, will increase the probability of achieving a desired outcome.
	MeasureTypeProcess MeasureType = "process"
	// A measure that indicates the result of the performance (or non-performance) of a function or process.
	MeasureTypeOutcome MeasureType = "outcome"
	// A measure that focuses on a health care provider's capacity, systems, and processes to provide high-quality care.
	MeasureTypeStructure MeasureType = "structure"
	// A measure that focuses on patient-reported information such as patient engagement or patient experience measures.
	MeasureTypePatientReportedOutcome MeasureType = "patient-reported-outcome"
	// A measure that combines multiple component measures in to a single quality measure.
	MeasureTypeComposite MeasureType = "composite"
)
